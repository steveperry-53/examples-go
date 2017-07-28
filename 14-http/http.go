package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

func main() {
	// firstTry()
	secondTry()
}

func firstTry() {

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss":"auth-experiment@steve-70.iam.gserviceaccount.com",
		"aud":"https://accounts.google.com/o/oauth2/token",
		"scope":"https://www.googleapis.com/auth/compute.readonly",	
    	"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Second * 3600).Unix(),
	})

	jwtString, err := token.SignedString([]byte("secret"))
	fmt.Println("\nSignedString Error: ", err)
	fmt.Println("\n", jwtString)
}


func secondTry() {

	tokenRequestBody := `grant_type=urn%3Aietf%3Aparams%3Aoauth%3Agrant-type%3Ajwt-bearer&assertion=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhdXRoLWV4cGVyaW1lbnRAc3RldmUtNzAuaWFtLmdzZXJ2aWNlYWNjb3VudC5jb20iLCJhdWQiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20vby9vYXV0aDIvdG9rZW4iLCJzY29wZSI6Imh0dHBzOi8vd3d3Lmdvb2dsZWFwaXMuY29tL2F1dGgvY29tcHV0ZS5yZWFkb25seSIsImlhdCI6IHRpbWUuTm93KCkuVW5peCgpLCJleHAiOiB0aW1lLk5vdygpLkFkZCh0aW1lLlNlY29uZCAqIDM2MDApLlVuaXgoKSx9.mhFjU7Pqy1-WJAYZrt04fi2b5aneqLOX7iJweqUrKVR8yMu2glYR_dKGOaZ5rH2uaEI2_dqTQCX7o5V1FV3-VMEN1BhIbKSDdhGtq7kCeijN4tkKU4AaSSJYQ2UYKIxPME5Oroikn7U3L4wN6A9AGOMtCQkUimnOWwN4XKGux5WX5nhIikxMTrDs2b05G1SpI0LTbfPqB8QgnbvqTIZOHjOSOIVvNdufOUh-qiiS3-MCvqPwV2emOxuwoC-QEP7cs8XgAUjrJgFas4GevQMcikydpcplqN-yHH0op_hA-dnKBflv59n2S3syMqu0xtM6n7bt4W1VKH7PP2WRULwM-g`

	req, err := http.NewRequest("POST", "https://accounts.google.com/o/oauth2/token", strings.NewReader(tokenRequestBody))
	fmt.Println("\nNewRequest error: ", err)
	req.Header.Add("Content-Type",  "application/x-www-form-urlencoded")
	// req.Write(os.Stdout)

	c := &http.Client{}
	resp, err := c.Do(req)
	fmt.Println("\nDo error: ", err)
	resp.Write(os.Stdout)
}
