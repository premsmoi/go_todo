package middleware

import (
	"Generalkhun/go-todo-server/models"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
)

//Signin is Handler function for route /signin
func Signin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var creds models.Credentials
		// Get the JSON body and decode into a Credentials type
		err := json.NewDecoder(c.Request.Body).Decode(&creds)
		if err != nil {
			c.Writer.WriteHeader(http.StatusBadRequest)
			return
		}
		// Get the expected password from our in memory map
		matchedUser, err := getUserPass(creds.Username, IntiateMongoConn())
		expectedPassword, ok := matchedUser["password"]

		// If a password exists for the given user
		// AND, if it is the same as the password we received, the we can move ahead
		// if NOT, then we return an "Unauthorized" status
		if !ok || creds.Password != expectedPassword {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Declare the expiration time of the token
		// here, we have kept it as 5 minutes
		expirationTime := time.Now().Add(10 * time.Minute)

		// Create the JWT claims, which includes the username and expiry time
		claims := models.Claims{
			Username: creds.Username,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expre	ssed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}
		// Declare the token with the algorithm used for signing, and the claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Create the JWT string
		tokenString, err := token.SignedString(models.JwtKey)
		if err != nil {
			c.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Finally, we set the client cookie for "token" as the JWT we just generated
		// we also set an expiry time which is the same as the token itself

		http.SetCookie(c.Writer, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
			//Domain:  "localhost:3000",
			Domain: "127.0.0.1:3000",
			Path:   "/",
		})
		c.Writer.WriteString("go thourgh this line")
		c.JSON(200, gin.H{"message": tokenString})
		//c.Redirect(http.StatusMovedPermanently, "../task/welcome")
	}
}

//Logout is handler function for /logout route, used to remove token from client's cookie
func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set the new token as the users `token` cookie
		http.SetCookie(c.Writer, &http.Cookie{
			Name:    "token",
			Value:   "",
			Domain:  "127.0.0.1:3000",
			Path:    "/",
			Expires: time.Now().Add(5 * time.Minute),
		})
		c.Redirect(http.StatusPermanentRedirect, "/")
	}
}

//Refresh is handler function for /refresh route, used to sent refresh token to client
func Refresh() gin.HandlerFunc {
	return func(c *gin.Context) {
		// (BEGIN) The code uptil this point is the same as the first part of the `Welcome` route
		cookie, err := c.Request.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				c.JSON(http.StatusUnauthorized, nil)
				return
			}
			c.AbortWithError(http.StatusBadRequest, err)
			c.Writer.WriteHeader(http.StatusBadRequest)
			return
		}
		tknStr := cookie.Value

		claims := &models.Claims{}
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return models.JwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {

				c.Writer.WriteHeader(http.StatusUnauthorized)
				return
			}
			c.AbortWithError(http.StatusBadRequest, err)
			c.Writer.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			c.Writer.WriteString(tknStr)
			c.Writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		// (END) The code up-till this point is the same as the first part of the `Welcome` route

		// Now, create a new token for the current use, with a renewed expiration time
		expirationTime := time.Now().Add(10 * time.Minute)
		claims.ExpiresAt = expirationTime.Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(models.JwtKey)
		if err != nil {
			c.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Set the new token as the users `token` cookie
		http.SetCookie(c.Writer, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
			Domain:  "127.0.0.1:3000",
			Path:    "/",
		})

	}

}

// AuthRequired is middleware function that used to specified users from a token they've got
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// We can obtain the session token from the requests cookies, which come with every request

		_, err := c.Request.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				c.Writer.WriteString("No cookie!!!")
				fmt.Println("no cookie")
				c.AbortWithError(http.StatusUnauthorized, err)
				return
			}
			c.Writer.WriteString("internal server error")
			c.AbortWithError(http.StatusInternalServerError, err)
			return

		}

		// // Get the JWT string from the cookie
		// tknStr := cookie.Value
		//tknStr := c.Request.Header.Get("Cookie")[6:]
		token, _ := c.Request.Cookie("token")
		tknStr := token.Value

		// Initialize a new instance of `Claims`
		claims := &models.Claims{}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return models.JwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.Writer.WriteString("Signature invalid")
				c.AbortWithError(http.StatusUnauthorized, err)
				return
			}
			c.Writer.WriteString("Bad request")
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if !tkn.Valid {
			c.Writer.WriteString("Token is invalid")
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		// If all conditions pass, store username inside gin.Context to further use on
		val := reflect.ValueOf(tkn.Claims).Elem()
		username := val.FieldByName("Username").Interface().(string)
		fmt.Println(username)
		c.Set("contextUsername", username)
	}
}
