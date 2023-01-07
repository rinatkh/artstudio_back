package constants

const (
	JWTKeyUserID     = "user_id"
	JWTKeyExpiration = "exp"

	ViperJWTTTLKey    = "service.jwt_ttl"
	ViperJWTSecretKey = "service.jwt_secret"
	ViperSecretKey    = "service.secret"
)

const (
	HeaderKeyRequestID   = "X-Request-ID"
	CookieKeyAuthToken   = "Auth-Token"
	CookieKeySecretToken = "Secret-Token"
)
const CtxKeyUserID = "User-Id"
