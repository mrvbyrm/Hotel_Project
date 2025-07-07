package helpers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

// SignedDetails contains user information and JWT's standard claims
type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	UserID    string
	Role      string
	jwt.StandardClaims
}

// SecretKey is obtained from the environment variable
var secretKey = os.Getenv("SECRET_KEY")

func getSecretKey() string {
	if secretKey == "" {
		panic("SECRET_KEY environment variable not set")
	}
	return secretKey
}

// GenerateAllTokens creates access and refresh tokens
func GenerateAllTokens(email, firstName, lastName, userID, role string) (string, string, error) {
	claims := &SignedDetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		UserID:    userID,
		Role:      role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day expiration
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(getSecretKey()))
	if err != nil {
		return "", "", fmt.Errorf("failed to sign token: %w", err)
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 168).Unix(), // 7 days expiration
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefreshToken, err := refreshToken.SignedString([]byte(getSecretKey()))
	if err != nil {
		return "", "", fmt.Errorf("failed to sign refresh token: %w", err)
	}

	return signedToken, signedRefreshToken, nil
}
