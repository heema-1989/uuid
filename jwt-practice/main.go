package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io"
	"log"
	"net/http"
	"time"
)

var secretKey = []byte("secret key")
var api_key = "1234"

type Message struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

func CheckError(err error, msg string) {
	log.Fatal(err, msg)
}
func Home(w http.ResponseWriter, r *http.Request) {
	//fprintf, err := fmt.Fprintf(w, "secret token key")
	//if err != nil {
	//	log.Fatal("Unable to write response: Reason ", err)
	//	return
	//}
	//fmt.Println("Response is: ", string(fprintf))

	var msg Message
	decodeErr := json.NewDecoder(r.Body).Decode(&msg)
	CheckError(decodeErr, "Error decoding the request body")
	encodeErr := json.NewEncoder(w).Encode(msg)
	CheckError(encodeErr, "Error encoding json")
}
func GenerateJwtToken() (string, error) {
	//this is for generating jwt token
	token := jwt.New(jwt.SigningMethodHS256)
	//second is the payload part--which contains info about the entity that is making the request
	claims := token.Claims.(jwt.MapClaims)
	claims["expiration"] = time.Now().Add(10 * time.Minute)
	claims["authorised"] = true
	claims["user"] = "username"
	//signing the jwt token
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error signing token")
		return "", err
	}
	return tokenString, nil
}

// ValidateJWT validating the jwt-->The verifyJWT function is a middleware that takes in the handler function for the request you want to verify.
// The handler function uses the token parameter from the request header to verify the request and respond based on the status.
func ValidateJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Header["Token"] != nil {
			token, parseErr := jwt.Parse(request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				//this part checks whether the token passed in the request header has same signing method which was used to generate the JWT
				_, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					writer.WriteHeader(http.StatusUnauthorized) //401
					_, err := writer.Write([]byte("Your are unauthorised"))
					if err != nil {
						return nil, err
					}
				}
				return secretKey, nil
			})
			//this part is for error encountered while parsing the token
			if parseErr != nil {
				writer.WriteHeader(http.StatusUnauthorized)
				_, err := writer.Write([]byte("Error parsing token. You are unauthorised"))
				if err != nil {
					return
				}
			}
			//this part checks whether the token is valid or not--if token is valid then this home function will be called and output will be displayed else error
			if token.Valid {
				endpointHandler(writer, request)
			} else {
				writer.WriteHeader(http.StatusUnauthorized)
				_, err := writer.Write([]byte("Token is not valid. You are unauthorised"))
				if err != nil {
					return
				}
			}
		} else {
			//case when there is no token in the header
			writer.WriteHeader(http.StatusUnauthorized)
			_, err := writer.Write([]byte("There is no token present in the header. You are unauthorised"))
			if err != nil {
				return
			}
		}
	})
}
func GetJWT(writer http.ResponseWriter, request *http.Request) {
	if request.Header["Access"] != nil {
		if request.Header["Access"][0] == api_key {
			token, err := GenerateJwtToken()
			if err != nil {
				writer.WriteHeader(http.StatusUnauthorized)
				_, err := writer.Write([]byte("You are not authorised due to wrong api_key"))
				if err != nil {
					return
				}
			}
			fprintf, err := fmt.Fprintf(writer, token)
			//we can also write like this
			writeString, err := io.WriteString(writer, "heema")
			if err != nil {
				return
			}
			log.Println(writeString)
			if err != nil {
				return
			}
			log.Println(fprintf)
		}
	}
}
func main() {
	http.Handle("/api", ValidateJWT(Home))
	http.HandleFunc("/jwt", GetJWT)
	fmt.Println(http.ListenAndServe(":3500", nil))
}
