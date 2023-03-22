package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/hardikbansal/gpt-enterprise-ui/core"
	"time"
)

// Service implements core.AuthService
type Service struct {
}

func (a *Service) GenerateAuthToken(user core.User, expiry int64) (string, error) {
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

func (a *Service) VerifyAuthToken(token string) (core.User, error) {
	// Parse the token string into a JWT token
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Check that the signing method is HMAC with SHA-256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, core.InvalidAuthToken{Issue: fmt.Sprintf("unexpected signing method: %v", token.Header["alg"])}
		}
		// Return the secret key used to sign the token
		return []byte("YOUR_SECRET_KEY"), nil
	})
	if errors.Is(err, jwt.ValidationError{}) {
		return core.User{}, core.InvalidAuthToken{Issue: fmt.Sprintf("failed to parse token: %v", err)}
	}

	// Check that the token is valid
	if !jwtToken.Valid {
		return core.User{}, core.InvalidAuthToken{Issue: fmt.Sprintf("invalid token")}
	}

	// Extract the claims from the token
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return core.User{}, core.InvalidAuthToken{Issue: fmt.Sprintf("invalid claims")}
	}

	// Check that the token has not expired
	if exp, ok := claims["exp"].(float64); ok {
		if time.Unix(int64(exp), 0).Before(time.Now()) {
			return core.User{}, core.InvalidAuthToken{Issue: fmt.Sprintf("token has expired")}
		}
	} else {
		return core.User{}, core.InvalidAuthToken{fmt.Sprintf("missing or invalid exp claim")}
	}

	profile, ok := claims["profile"].(string)
	if !ok {
		return core.User{}, core.InvalidAuthToken{Issue: fmt.Sprintf("profile missing from token")}
	}

	var user core.User
	err = json.Unmarshal([]byte(profile), &user)
	if err != nil {
		return core.User{}, core.InvalidAuthToken{Issue: fmt.Sprintf("Invalid auth token")}
	}

	return user, nil
}
