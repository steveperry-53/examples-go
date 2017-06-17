package main

import (
	"fmt"
	"net/http"
	"os"
	"encoding/base64"
	"time"
)

func main() {
	firstTry()
	secondTry()
	tryTime()
}

func firstTry() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://www.googleapis.com/discovery/v1/apis/urlshortener/v1/rest", nil)
	req.Write(os.Stdout)
	resp, _ := client.Do(req)
	resp.Write(os.Stdout)
}

func secondTry() {
	fmt.Println("second")
	// client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://www.googleapis.com/discovery/v1/apis/urlshortener/v1/rest", nil)
	req.Write(os.Stdout)
	jwtHeader := `{"alg":"RS256","typ":"JWT"}`
	fmt.Println()
	fmt.Println(jwtHeader)
	encodedJwtHeader := base64.StdEncoding.EncodeToString([]byte(jwtHeader))
	fmt.Println(encodedJwtHeader)

	// iss: issuer
	// aud: audience
	// iat: issued at time, seconds since epoch
	// exp: expiration time, seconds since epoch

	jwtPayload := `{
  "iss":"auth-experiment@steve-70.iam.gserviceaccount.com",
  "scope":"https://www.googleapis.com/auth/compute.readonly",
  "aud":"https://www.googleapis.com/oauth2/v4/token",
  "exp":"1497719392",
  "iat":"1497892192"
}`
	fmt.Println()
	fmt.Println(jwtPayload)
	encodedJwtPayload := base64.StdEncoding.EncodeToString([]byte(jwtPayload))
	fmt.Println(encodedJwtPayload)
}

func tryTime() {
    now := time.Now()
    iat := now.Unix()
	exp := iat + 2*24*60*60
    fmt.Println(now)
    fmt.Println(iat)
	fmt.Println(exp)
}


// Client ID
// 642580684294-k3tsbg9iqg5qg756llbg1mckh7b40r13.apps.googleusercontent.com
// Client secret	
// 4PpFw_R9fwnZqFhaM1AwIalr

// req.Header.Add("Authentication", "Bearer xyzxyz")


// curl -d "client_id=642580684294-k3tsbg9iqg5qg756llbg1mckh7b40r13.apps.googleusercontent.com&scope=email profile" https://accounts.google.com/o/oauth2/device/code

// ID: 642580684294-88d2vmkjfi86pn4785ub4bta8jrpvveb.apps.googleusercontent.com
// Secret: lhDv4bbrJhXAkpXHNsvy-L7l


// POST /o/oauth2/device/code HTTP/1.1
// Host: accounts.google.com
// Content-Type: application/x-www-form-urlencoded

// client_id=client_id&
// scope=email%20profile

