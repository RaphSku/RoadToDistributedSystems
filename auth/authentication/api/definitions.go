package api

import (
	"crypto"
	"crypto/ecdsa"

	"github.com/golang-jwt/jwt"
	"github.com/hashicorp/go-hclog"
)

// ECDSA Signing Method Definition
var ECDSA = jwt.SigningMethodECDSA{
	Name:      "ECDSASigningMethod",
	Hash:      crypto.SHA256,
	KeySize:   32,
	CurveBits: 256,
}

// Handler with logger and privatKey
type Handler struct {
	logger     hclog.Logger
	privateKey *ecdsa.PrivateKey
}

// Request body when the user tries to login
type UserLoginTrial struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Stored information on the user during login
type UserLogin struct {
	username     string
	passwordHash []byte
	salt         []byte
}

// User Sign Up
type UserSignUp struct {
	Username     string `json:"username"`
	PasswordHash []byte `json:"passwordhash"`
}

// Body of the Request body when a user signs up
type UserSignUpBody struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims that a user makes
type UserClaims struct {
	Username string `json:"name"`
	jwt.StandardClaims
}

// A message that is returned to the user if a request was being processed successfully
// swagger:response successMessage
type successMessage struct {
	// The success message
	// in: body
	Body struct {
		// The congratulations message
		Message string `json:"message"`
	}
}

// A message that is returned to the user if the login was successful, a JWT will be returned
// swagger:response successMessageWithJWT
type successMessageWithJWT struct {
	// The JWT is being provided to the user
	// in: body
	Body struct {
		// JWT
		Token string `json:"token"`
	}
}

// A message that is returned to the user if a valid JWT is provided, the username will be returned
// swagger:response successMessageOnInfo
type successMessageOnInfo struct {
	// The username is returned
	// in: body
	Body struct {
		// JWT
		User string `json:"user"`
	}
}

// BadRequest represents an error where the server does not know what to do with the incoming request
// swagger:response badRequest
type badRequest struct {
	// Empty body
	// in: body
	Body struct{}
}

// A message that is returned to the user when the server has an internal error
// swagger:response internalServerError
type internalServerError struct {
	// The failure message
	// in: body
	Body struct {
		Message string `json:"message"`
	}
}

// A message that is returned to the user when the server has an authorization error
// swagger:response unauthorizedError
type unauthorizedError struct {
	// Empty body
	// in: body
	Body struct{}
}
