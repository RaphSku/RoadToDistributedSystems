package main

import (
	"fmt"
	"os"
	"time"

	"github.com/RaphSku/RoadToDistributedSystems/tree/main/auth/jwt_authentication/pkg"
	"github.com/golang-jwt/jwt"
	"github.com/hashicorp/go-hclog"
)

type UserClaims struct {
	Username string `json:"name"`
	jwt.StandardClaims
}

func main() {
	// Create Logger
	logger := hclog.Default()

	// Get private and public key that are generated with ECSDA with P256
	logger.Info("Retrieve private and public key...")
	encodedPrivateKey, encodedPublicKey := pkg.GetPrivatePublicKey("./")
	privateKey, publicKey := pkg.Decode(encodedPrivateKey, encodedPublicKey)

	// Create JWT with Claims
	logger.Info("Creating a JWT...")
	token := jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			"exp":  time.Now().Add(time.Hour * 24).Unix(),
			"name": "Raphael",
		})
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		logger.Error("%v", err)
		os.Exit(1)
	}

	// Show signed JWT
	logger.Info("Showing signed token...")
	logger.Info(fmt.Sprintf("Signed Token: %v", signedToken))

	// Parsing token with custom claims
	logger.Info("Parsing JWT with custom claims...")
	token, err = jwt.ParseWithClaims(signedToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		logger.Error(fmt.Sprintf("Parsing failed due to %v", err))
		os.Exit(1)
	}

	// Retrieving claims from JWT
	logger.Info("Retrieving claims from JWT...")
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		fmt.Println("Claims could not be read!")
		os.Exit(1)
	}

	// Showing username and when the token expires in Unix time
	logger.Info(fmt.Sprintf("username: %v, expires at: %v", claims.Username, claims.ExpiresAt))
}
