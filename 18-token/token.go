package main

import (
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"context"
	"io/ioutil"
)

func main() {

	var jsonKey []byte
	jsonKey, _ = ioutil.ReadFile("/home/steve/Documents/keys/edward-02-key.json")

	var ctx context.Context
	ctx = context.Background()

	var creds *google.Credentials
	creds, _ = google.CredentialsFromJSON(ctx, jsonKey, "https://www.googleapis.com/auth/bigquery")

	var tknSrc oauth2.TokenSource
	tknSrc = creds.TokenSource
	
	var tkn *oauth2.Token
	tkn, _ = tknSrc.Token()

	fmt.Println("\n\ntkn:", tkn.AccessToken)
}
