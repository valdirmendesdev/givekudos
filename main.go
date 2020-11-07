package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"log"
	"net/http"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func healthStatus(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Write([]byte(`{"message":"ok"}`))
}

func users(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Write([]byte(`{"message":"users"}`))
}

func generateToken(w http.ResponseWriter, _ *http.Request) {

	claims := &Claims{
		Username: "valdir.mfjesus",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 2).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString([]byte("MINHA_API_KEY"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	resp := fmt.Sprintf(`{"token":%v}`, tokenString)
	w.Write([]byte(resp))
}

func checkToken(w http.ResponseWriter, r *http.Request) {
	token_string, err := request.AuthorizationHeaderExtractor.ExtractToken(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Erro no extractor")
		return
	}

	token, err := jwt.ParseWithClaims(token_string, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("MINHA_API_KEY"), nil
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Erro na conversão")
		return
	}

	fmt.Println(token)

	fmt.Println(token.Claims.(*Claims).Username)

	if err != nil {
		panic(err)
	}
	w.Write([]byte("top"))

}

func main() {
	server := http.NewServeMux()
	server.Handle("/", http.HandlerFunc(healthStatus))
	server.Handle("/users", http.HandlerFunc(users))
	server.Handle("/token", http.HandlerFunc(generateToken))
	server.Handle("/check", http.HandlerFunc(checkToken))

	fmt.Println("Server is running...")
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatal("Server error")
	}

}
