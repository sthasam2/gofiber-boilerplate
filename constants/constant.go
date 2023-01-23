package constants

const (
	APPNAME = "APP"

	PROD = "PRODUCTION"
	DEV  = "DEVELOPMENT"

	PG    = "POSTGRES"
	REDIS = "REDIS"
	MONGO = "MONGODB"

	// Roles
	DEFAULT = "default"
	ADMIN   = "admin"
	MOD     = "moderator"

	// Expiry in HRS
	ACCESS_EXPR_HR   = 1
	REFRESH_EXPR_DAY = 7
)
