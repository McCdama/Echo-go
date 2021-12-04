package auth

import (
	"Echo-go/user"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

const (
	accessTokenCookieName = "access-token"

	jwtSecretKey = "'iFtc:E_HM3=b~?" // Todo: move it to .env
)

func GetJWTSecret() string {
	return jwtSecretKey
}

// Create a struct that will be encoded to a JWT.
// jwt.StandardClaims as an embedded type, to provide fields like expiry time.
type Claims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func GenerateTokenAndSetCookie(user *user.User, c echo.Context) error {
	accessToken, exp, err := generateAccessToken(user)
	if err != nil {
		return err
	}

	setTokenCookie(accessTokenCookieName, accessToken, exp, c)
	setUserCookie(user, exp, c)

	return nil
}

func generateAccessToken(user *user.User) (string, time.Time, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	return generateToken(user, expirationTime, []byte(GetJWTSecret()))
}

func generateToken(user *user.User, expirationTime time.Time, secret []byte) (string, time.Time, error) {
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	// JWT str
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}

func setTokenCookie(name, token string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.Path = "/"
	cookie.HttpOnly = true

	c.SetCookie(cookie)
}

func setUserCookie(user *user.User, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = user.Name
	cookie.Expires = expiration
	cookie.Path = "/"

	c.SetCookie(cookie)
}

// executed when user try to access a protected path.
func JWTErrorChecker(err error, c echo.Context) error {
	// Redirects to the signIn form.
	return c.Redirect(http.StatusMovedPermanently, c.Echo().Reverse("userSignInForm"))
}
