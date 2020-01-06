package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	prisma "github.com/dmitrychurkin/hotelier/server/prisma-client"
)

const (
	maxLength                       = 1000
	reEmail                         = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	defaultTokenSecret              = "jwt_access_secret|123"
	defaultTokenCookieName          = "_u"
	defaultTokenLifetime            = 300
	defaultResendTokenTimespan      = 30
	defaultPasswordResetTokenLength = 32
	defaultPasswordHashRounds       = 14
)

/*
// User resolver
func User(ctx context.Context, p *prisma.Client) (*models.User, error) {
	// 1. get jwt token claims
	claims, err := parseAuthToken(ctx)
	if err != nil {
		return nil, err
	}
	if claims == nil {
		return nil, &gqlerror.Error{
			Message: "Unautorized",
			Extensions: map[string]interface{}{
				"code": http.StatusUnauthorized,
			},
		}
	}

	// 2. query user
	user, err := p.User(prisma.UserWhereUniqueInput{
		ID: &claims.Subject,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("We don't have records with associated credentials")
	}

	// 3. issue auth token
	// 4. set cookie
	if err := issueAuthToken(ctx, user); err != nil {
		return nil, err
	}

	return &models.User{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      models.UserRoles(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// Login resolver
func Login(ctx context.Context, p *prisma.Client, email string, password string) (*models.User, error) {
	// 1. validate input
	email, password = strings.TrimSpace(email), strings.TrimSpace(password)
	if emailLen, passwordLen := len(email), len(password); emailLen == 0 ||
		emailLen > 1000 ||
		passwordLen == 0 ||
		passwordLen > 1000 {
		return nil, errors.New("We don't have records with associated credentials")
	}

	matched, err := regexp.MatchString(reEmail, email)
	if err != nil {
		return nil, err
	}
	if !matched {
		return nil, errors.New("Invalid email pattern")
	}

	// 2. get user from DB
	user, err := p.User(prisma.UserWhereUniqueInput{
		Email: &email,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("We don't have records with associated credentials")
	}

	// 3. check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("We don't have records with associated credentials")
	}

	// 4. issue auth token
	// 5. set cookie
	if err := issueAuthToken(ctx, user); err != nil {
		return nil, err
	}

	return &models.User{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      models.UserRoles(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// Signup resolver
func Signup(ctx context.Context, p *prisma.Client, email string, firstName *string, lastName *string, password string, confirmPassword string) (*models.User, error) {
	// 1. validate input
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)
	confirmPassword = strings.TrimSpace(confirmPassword)
	fName := *firstName
	fName = strings.TrimSpace(fName)
	lName := *lastName
	lName = strings.TrimSpace(lName)
	if emailLen, passwordLen := len(email), len(password); emailLen == 0 ||
		emailLen > maxLength ||
		passwordLen == 0 ||
		passwordLen > maxLength ||
		password != confirmPassword {
		return nil, errors.New("Some of input was invalid")
	}

	matched, err := regexp.MatchString(reEmail, email)
	if err != nil {
		return nil, err
	}
	if !matched {
		return nil, errors.New("Invalid email pattern")
	}

	// 2. Hash password
	roundsHash := defaultPasswordHashRounds
	if r, err := strconv.Atoi(os.Getenv("PASSWORD_HASH_ROUNDS")); err == nil {
		roundsHash = r
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), roundsHash)
	if err != nil {
		return nil, err
	}

	// 3. save user into DB
	user, err := p.CreateUser(prisma.UserCreateInput{
		Email:     email,
		Password:  string(hash),
		FirstName: &fName,
		LastName:  &lName,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	// 4. issue auth token
	// 5. set cookie
	if err := issueAuthToken(ctx, user); err != nil {
		return nil, err
	}

	return &models.User{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      models.UserRoles(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// Logout resolver
func Logout(ctx context.Context) (*bool, error) {
	var (
		authTokenCookieName = os.Getenv("AUTH_TOKEN_COOKIE_NAME")
		result              = false
	)

	if len(authTokenCookieName) == 0 {
		authTokenCookieName = defaultTokenCookieName
	}

	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return &result, err
	}

	gc.SetCookie(authTokenCookieName, "", -1, "", "", false, true)
	result = true
	return &result, nil
}

// SendPasswordResetLink resolver
func SendPasswordResetLink(ctx context.Context, p *prisma.Client, email, path string) (*bool, error) {
	success := false
	// 1. validate input
	email = strings.TrimSpace(email)
	if emailLen := len(email); emailLen == 0 || emailLen > 1000 {
		return &success, errors.New("We don't have records with associated credentials")
	}

	matched, err := regexp.MatchString(reEmail, email)
	if err != nil {
		return &success, err
	}
	if !matched {
		return &success, errors.New("Invalid email pattern")
	}

	// 2. get user from DB
	user, err := p.User(prisma.UserWhereUniqueInput{
		Email: &email,
	}).Exec(ctx)

	if err != nil {
		return &success, err
	}

	if user == nil {
		return &success, errors.New("We don't have records with associated credentials")
	}

	resendTokenTimespan := defaultResendTokenTimespan
	if r, err := strconv.Atoi(os.Getenv("RESEND_TOKEN_TIMESPAN")); err == nil {
		resendTokenTimespan = r
	}

	// 2. Check if elapsed time since last email over resendTokenTimespan
	if user.PasswordResetToken != nil && len(*user.PasswordResetToken) > 0 {
		if t, err := time.Parse(time.RFC3339, *user.PasswordResetTokenCreatedAt); err == nil {
			if el := int(time.Since(t).Minutes()); el < resendTokenTimespan {
				return &success, &gqlerror.Error{
					Message: strconv.Itoa(el),
					Extensions: map[string]interface{}{
						"code": http.StatusForbidden,
					},
				}
			}
		}
	}

	passwordResetTokenLength := defaultPasswordResetTokenLength
	if r, err := strconv.Atoi(os.Getenv("PASSWORD_RESET_TOKEN_LENGTH")); err == nil {
		passwordResetTokenLength = r
	}

	// 3. Generate password reset token
	passwordResetToken, err := generateURLHash(passwordResetTokenLength)
	if err != nil {
		return &success, err
	}

	var (
		now                         = time.Now()
		passwordResetTokenCreatedAt = now.Format(time.RFC3339)
		passwordResetTokenExpiresAt = now.Add(time.Minute * 30).Format(time.RFC3339)
	)

	// 4. Store into db
	updatedUser, err := p.UpdateUser(prisma.UserUpdateParams{
		Where: prisma.UserWhereUniqueInput{
			ID: &user.ID,
		},
		Data: prisma.UserUpdateInput{
			PasswordResetToken:          &passwordResetToken,
			PasswordResetTokenCreatedAt: &passwordResetTokenCreatedAt,
			PasswordResetTokenExpiresAt: &passwordResetTokenExpiresAt,
		},
	}).Exec(ctx)

	if err != nil {
		return &success, err
	}

	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return &success, err
	}
	req := gc.Request
	// 5. get request url
	if path[len(path)-1:] != "/" {
		path += "/"
	}

	// 6. construct reset url
	link := path + *updatedUser.PasswordResetToken

	// 7. Send email
	host := req.Host
	if i := strings.Index(host, ":"); i != -1 {
		host = host[:i]
	}

	from := "no-reply@" + host
	if !govalidator.IsEmail(from) {
		from = ""
	}

	emailAgent := &mailer.ResetPasswordData{Link: link}
	success, err = emailAgent.Send(&mailer.MailRequest{
		From: from,
		To:   []string{user.Email},
	})
	return &success, err
}

// ResetPasswordCreds resolver
func ResetPasswordCreds(ctx context.Context, p *prisma.Client, passwordResetToken string) (*models.ResetPasswordCreds, error) {
	// 1. check if token exists
	user, err := p.User(prisma.UserWhereUniqueInput{
		PasswordResetToken: prisma.Str(strings.TrimSpace(passwordResetToken)),
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &models.ResetPasswordCreds{Email: user.Email}, err
}

// ResetPassword resolver
func ResetPassword(ctx context.Context, p *prisma.Client, email, password, confirmPassword, passwordResetToken string) (*models.User, error) {
	// 1. validate input
	email = strings.TrimSpace(email)
	passwordResetToken = strings.TrimSpace(passwordResetToken)
	password = strings.TrimSpace(password)
	confirmPassword = strings.TrimSpace(confirmPassword)

	if emailLen, passwordLength := len(email), len(password); emailLen == 0 ||
		emailLen > maxLength ||
		passwordLength == 0 ||
		passwordLength > maxLength ||
		password != confirmPassword {
		return nil, errors.New("Some of input was invalid")
	}

	matched, err := regexp.MatchString(reEmail, email)
	if err != nil {
		return nil, err
	}
	if !matched {
		return nil, errors.New("Invalid email pattern")
	}

	// 2. check if token expired
	user, err := p.User(prisma.UserWhereUniqueInput{
		Email: &email,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	expiresAt, err := time.Parse(time.RFC3339, *user.PasswordResetTokenExpiresAt)
	if err != nil {
		return nil, err
	}
	if expiresAt.Before(time.Now()) {
		return nil, errors.New("Password reset token expired")
	}

	// 3. Hash password
	roundsHash := defaultPasswordHashRounds
	if r, err := strconv.Atoi(os.Getenv("PASSWORD_HASH_ROUNDS")); err == nil {
		roundsHash = r
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), roundsHash)
	if err != nil {
		return nil, err
	}

	// 4. update user record
	user, err = p.UpdateUser(prisma.UserUpdateParams{
		Where: prisma.UserWhereUniqueInput{
			Email: &email,
		},
		Data: prisma.UserUpdateInput{
			Password:           prisma.Str(string(hash)),
			PasswordResetToken: prisma.Str(""),
		},
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	// 5. issue auth token
	// 6. set cookie
	if err := issueAuthToken(ctx, user); err != nil {
		return nil, err
	}

	return &models.User{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      models.UserRoles(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
*/
func generateURLHash(n int) (string, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	return base64.URLEncoding.EncodeToString(bytes), err
}

func parseAuthToken(ctx context.Context) (*jwt.StandardClaims, error) {
	var (
		authTokenSecret     = os.Getenv("AUTH_TOKEN_SECRET")
		authTokenCookieName = os.Getenv("AUTH_TOKEN_COOKIE_NAME")
	)

	if len(authTokenSecret) == 0 {
		authTokenSecret = defaultTokenSecret
	}
	if len(authTokenCookieName) == 0 {
		authTokenCookieName = defaultTokenCookieName
	}

	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	// 1. get token from req cookie
	t, err := gc.Cookie(authTokenCookieName)

	if err != nil {
		gc.SetCookie(authTokenCookieName, "", -1, "", "", false, true)
		return nil, err
	}

	if len(t) == 0 {
		gc.SetCookie(authTokenCookieName, "", -1, "", "", false, true)
		return nil, nil
	}

	// 2. parse token claims
	claims := &jwt.StandardClaims{}
	jwtToken, err := jwt.ParseWithClaims(t, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(authTokenSecret), nil
	})

	if err != nil {
		gc.SetCookie(authTokenCookieName, "", -1, "", "", false, true)
		return nil, err
	}

	if !jwtToken.Valid {
		gc.SetCookie(authTokenCookieName, "", -1, "", "", false, true)
		return nil, nil
	}

	return claims, err
}

func issueAuthToken(ctx context.Context, user *prisma.User) error {
	var (
		authTokenSecret     = os.Getenv("AUTH_TOKEN_SECRET")
		authTokenLifetime   = os.Getenv("AUTH_TOKEN_LIFETIME")
		authTokenCookieName = os.Getenv("AUTH_TOKEN_COOKIE_NAME")
		authTokenMin        = defaultTokenLifetime
	)

	if len(authTokenSecret) == 0 {
		authTokenSecret = defaultTokenSecret
	}

	d, err := strconv.Atoi(authTokenLifetime)
	if err == nil {
		authTokenMin = d
	}

	authTokenClaims := &jwt.StandardClaims{
		Subject:   user.ID,
		ExpiresAt: time.Now().Add(time.Duration(authTokenMin) * time.Minute).Unix(),
	}

	authToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, authTokenClaims).SignedString([]byte(authTokenSecret))

	if err != nil {
		return err
	}

	if len(authTokenCookieName) == 0 {
		authTokenCookieName = defaultTokenCookieName
	}

	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return err
	}

	gc.SetCookie(authTokenCookieName, authToken, authTokenMin*60, "", "", false, true)

	return nil
}
