package core

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	logger "github.com/hardikbansal/gpt-enterprise-ui/logger"
	"io/ioutil"
	"net/http"
	"time"
)

func validateGoogleUserDetails(token string) (GoogleUserInfo, error) {
	fmt.Println(token)
	client := &http.Client{}
	// use google api to get the user info
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	if err != nil {
		logger.LogProblem("Error creating request: ", err)
		return GoogleUserInfo{}, err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		logger.LogProblem("Error getting user info: ", err)
		return GoogleUserInfo{}, err
	}
	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.LogProblem("Error reading user")
		return GoogleUserInfo{}, err
	}
	err = resp.Body.Close()
	if err != nil {
		return GoogleUserInfo{}, err
	}
	var userInfo GoogleUserInfo
	err = json.Unmarshal(userData, &userInfo)
	if err != nil {
		return GoogleUserInfo{}, err
	}
	if userInfo.Email == "" {
		logger.LogProblem("%s", string(userData))
		return GoogleUserInfo{}, err
	}

	return userInfo, nil

}

func (srv *ChatGptService) GenerateAccessToken(token string) (AccessTokenData, error) {

	userInfo, err := validateGoogleUserDetails(token)
	if err != nil {
		return AccessTokenData{}, err
	}

	// testing purpose
	//userInfo := GoogleUserInfo{Email: "hardik@frnd.app", Name: "Hardik Bansal"}

	user, err := srv.dbService.GetOrCreateUser(userInfo.Email, userInfo.Name)
	if err != nil {
		return AccessTokenData{}, err
	}

	jwtToken, err := generateUserJwtToken(user, time.Now().Add(time.Hour*24).Unix())
	if err != nil {
		return AccessTokenData{}, err
	}
	return AccessTokenData{Token: jwtToken}, nil

}

func generateUserJwtToken(user User, expiry int64) (string, error) {
	jsonProfile, err := json.Marshal(user)
	if err != nil {
		return "", fmt.Errorf("unable to marshal user model")
	}

	// Generate a new JWT token
	claims := jwt.MapClaims{
		"profile": string(jsonProfile),
		"exp":     expiry,
	}
	intToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := []byte("YOUR_SECRET_KEY")
	jwtToken, err := intToken.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}
