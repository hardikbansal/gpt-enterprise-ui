package core

import (
	"encoding/json"
	"fmt"
	"github.com/hardikbansal/gpt-enterprise-ui/config"
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
	if userInfo.Email == "" || userInfo.HD != config.GetInstance().EmailDomain {
		logger.LogProblem("%s", string(userData))
		return GoogleUserInfo{}, err
	}

	return userInfo, nil

}

func (srv *ChatGptService) GenerateAccessToken(token string) (AccessTokenData, error) {

	var userInfo GoogleUserInfo

	if config.GetInstance().IsDebug {
		// testing purpose
		userInfo = GoogleUserInfo{Email: "sameple_email@email", Name: "Sample Name"}
	} else {
		uInfo, err := validateGoogleUserDetails(token)
		if err != nil {
			return AccessTokenData{}, err
		}
		userInfo = uInfo
	}

	user, err := srv.dbService.GetOrCreateUser(userInfo.Email, userInfo.Name)
	if err != nil {
		return AccessTokenData{}, err
	}

	jwtToken, err := srv.authService.GenerateAuthToken(user, time.Now().Add(time.Hour*24).Unix())
	if err != nil {
		return AccessTokenData{}, err
	}
	return AccessTokenData{Token: jwtToken}, nil

}
