package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"github.com/dgrijalva/jwt-go"
)

func main() {
	// firstTry()
	secondTry()
}

func firstTry() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://www.googleapis.com/discovery/v1/apis/urlshortener/v1/rest", nil)
	req.Write(os.Stdout)
	resp, _ := client.Do(req)
	resp.Write(os.Stdout)
}


func secondTry() {
	rsaPrivateKey := `secret`

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":"auth-experiment@steve-70.iam.gserviceaccount.com",
  		"scope":"https://www.googleapis.com/auth/compute.readonly",	
    		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Second * 3600 * 24).Unix(),
	})

	jwtString, err := token.SignedString([]byte(rsaPrivateKey))
	fmt.Println()
	fmt.Println(err)
	fmt.Println(jwtString)
}
Contact GitHub API Training Shop Blog About
