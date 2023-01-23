package services

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"

	cfg "app/configs"
	cnst "app/constants"
	mdl "app/models"
	slzr "app/serializers"
)

// IssueAccessToken generate access tokens used for auth
func IssueAccessToken(u mdl.User) (*slzr.TokenDetails, error) {
	expireTime := time.Now().Add(time.Hour * cnst.ACCESS_EXPR_HR) // hours
	tokenUUID := uuid.New().String()
	// Generate encoded token
	claims := slzr.AccessClaims{
		AccessTokenID: tokenUUID,
		ExternalID:    u.ExternalID,
		Role:          u.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    cfg.GetConfig().JWTIssuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tk, err := tokenClaims.SignedString([]byte(cfg.GetConfig().JWTAccessSecret))

	if err != nil {
		return nil, err
	}

	return &slzr.TokenDetails{
		Token:        tk,
		TokenUUID:    tokenUUID,
		TokenExpires: expireTime.Unix(),
	}, nil
}

// IssueRefreshToken generate refresh tokens used for auth
func IssueRefreshToken(u mdl.User) (*slzr.TokenDetails, error) {
	expireTime := time.Now().Add((24 * time.Hour) * cnst.REFRESH_EXPR_DAY) // in days
	tokenUUID := uuid.New().String()

	// Generate encoded token
	claims := slzr.RefreshClaims{
		RefreshTokenID: tokenUUID,
		ExternalID:     u.ExternalID,
		Role:           u.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    cfg.GetConfig().JWTIssuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tk, err := tokenClaims.SignedString([]byte(cfg.GetConfig().JWTRefreshSecret))

	if err != nil {
		return nil, err
	}

	return &slzr.TokenDetails{
		Token:        tk,
		TokenUUID:    uuid.New().String(),
		TokenExpires: expireTime.Unix(),
	}, nil
}
