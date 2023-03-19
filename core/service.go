package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	logger "github.com/hardikbansal/gpt-enterprise-ui/logger"
)

type UnauthorizedError struct {
}

func (e UnauthorizedError) Error() string {
	return "UnAuthorized"
}

type ChatGptService struct {
}

func GetNewService() *ChatGptService {
	return &ChatGptService{}
}

func (srv *ChatGptService) CallChatGPTAPI(query string) (string, error) {
	// create the request payload
	data := map[string]string{
		"prompt": query,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// create the HTTP client
	client := http.Client{}

	// create the request object
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/engines/davinci-codex/completions", bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	// set the API key in the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer YOUR_API_KEY_HERE")

	// send the request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// parse the response
	var respData struct {
		Choices []struct {
			Text string `json:"text"`
		} `json:"choices"`
	}
	err = json.Unmarshal(body, &respData)
	if err != nil {
		return "", err
	}

	// return the response text
	if len(respData.Choices) > 0 {
		return respData.Choices[0].Text, nil
	} else {
		return "", nil
	}
}

func (srv *ChatGptService) GenerateAccessToken(token string) (AccessTokendata, error) {
	fmt.Println(token)
	client := &http.Client{}
	// use google api to get the user info
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	if err != nil {
		logger.LogProblem("Error creating request: ", err)
		return GetEmptyAccessToken(), UnauthorizedError{}
	}
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		logger.LogProblem("Error getting user info: ", err)
		return GetEmptyAccessToken(), UnauthorizedError{}
	}
	userData, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		logger.LogProblem("Error reading user")
		return GetEmptyAccessToken(), UnauthorizedError{}
	}
	var userInfo GoogleUserInfo
	json.Unmarshal(userData, &userInfo)

	if userInfo.Email == "" {
		logger.LogProblem("%s", string(userData))
		return GetEmptyAccessToken(), UnauthorizedError{}
	}

	logger.LogMessage("%s", userInfo.Name)

	jsonProfile, err := json.Marshal(userInfo)
	if err != nil {
		return GetEmptyAccessToken(), UnauthorizedError{}
	}

	// Generate a new JWT token
	claims := jwt.MapClaims{
		"profile": string(jsonProfile),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	intToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := []byte("YOUR_SECRET_KEY")
	jwtToken, err := intToken.SignedString(jwtSecret)
	if err != nil {
		return GetEmptyAccessToken(), UnauthorizedError{}
	}
	return AccessTokendata{Token: jwtToken}, nil
}
