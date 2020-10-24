package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Users is struct used to collect username and password
var Users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type UsersDB struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username,omitempty"`
	Password string             `json:"password,omitempty"`
}

//JwtKey is JWT keythat use to create the signature (secret key)
var JwtKey = []byte("ThisIsSuperSecretKey")

//Credentials is a struct to recieve username and password from the post request
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//Claims is a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
