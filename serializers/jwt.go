package serializers

import (
	"github.com/golang-jwt/jwt"
)

// TokenDetails Represents token object
type TokenDetails struct {
	Token        string
	TokenUUID    string
	TokenExpires int64
}

// RefreshClaims represents refresh token JWT claims
type RefreshClaims struct {
	RefreshTokenID string `json:"refreshTokenID"`
	ExternalID     string `json:"userID"`
	Role           string `json:"role"`
	jwt.StandardClaims
}

// AccessClaims represents access token JWT claims
type AccessClaims struct {
	AccessTokenID string `json:"accessTokenID"`
	ExternalID    string `json:"userID"`
	Role          string `json:"role"`
	jwt.StandardClaims
}
