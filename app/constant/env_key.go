package constant

const (
	ServerPort  = "SERVER_PORT"
	Environment = "ENVIRONMENT"

	EnvJWT_SECRET_KEY = "JWT_SECRET_KEY"

	// socials
	EnvFacebookClientID     = "FACEBOOK_CLIENT_ID"
	EnvFacebookClientSecret = "FACEBOOK_CLIENT_SECRET"
	EnvFacebookRedirectURL  = "FACEBOOK_REDIRECT_URL"
	EnvGoogleClientID       = "GOOGLE_CLIENT_ID"
	EnvGoogleClientSecret   = "GOOGLE_CLIENT_SECRET"
	EnvGoogleRedirectURL    = "GOOGLE_REDIRECT_URL"

	// db
	DBHostKey   string = "DB_HOST"
	DBUserKey   string = "DB_USER"
	DBPassKey   string = "DB_PASS"
	DBPortKey   string = "DB_PORT"
	DBDBNameKey string = "DB_DB_NAME"
	DBSSLMode   string = "DB_SSL_MODE"
)
