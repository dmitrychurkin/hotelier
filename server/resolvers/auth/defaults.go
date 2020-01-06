package auth

const (
	defaultPasswordMinLength = 8
	defaultPasswordMaxLength = 1000
	// maxLength                       = 1000
	// reEmail                         = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	defaultTokenSecret              = "jwt_access_secret|123"
	defaultTokenCookieName          = "_u"
	defaultTokenLifetime            = 300
	defaultResendTokenTimespan      = 30
	defaultPasswordResetTokenLength = 32
	defaultPasswordHashRounds       = 14
)
