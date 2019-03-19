package main

import (
	"fmt"
	"time"
	"golang.org/x/oauth2/jws"
	"encoding/base64"
	"encoding/json"
)

func main() {

	hdr := &jws.Header{
		Algorithm: "RS256",
		Typ: "JWT",
		KeyID: "38d02480d5eb1a7b8cec8880fd38b6ddbb9c3009",
	}

	var b []byte
	b, _ = json.Marshal(hdr)

	var hdr64 string
	hdr64 = base64.RawURLEncoding.EncodeToString(b)

	clSet := &jws.ClaimSet{
		Iss: "edward-02@steve-71-209916.iam.gserviceaccount.com",
		Scope: "https://www.googleapis.com/auth/bigquery",
		Aud: "https://oauth2.googleapis.com/token",
		Exp: time.Now().Unix() + 3500,
		Iat: time.Now().Unix(),
	}

	b, _ = json.Marshal(clSet)

	var clSet64 string
	clSet64 = base64.RawURLEncoding.EncodeToString(b)

	var hc string
	hc = hdr64 + "." + clSet64

	fmt.Println("\nhc:\n", hc)
}
