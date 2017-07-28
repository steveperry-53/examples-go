package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {
	data := []byte(`{"alg":"RS256","typ":"JWT"}`)
	str := base64.URLEncoding.EncodeToString(data)
	fmt.Println(str)

	data2 := []byte(`{"iss":"auth-experiment@steve-70.iam.gserviceaccount.com","aud":"https://accounts.google.com/o/oauth2/token","scope":"https://www.googleapis.com/auth/compute.readonly","iat": time.Now().Unix(),"exp": time.Now().Add(time.Second * 3600).Unix(),}`)

	fmt.Println()
	str2 := base64.URLEncoding.EncodeToString(data2)
	fmt.Println(str2)

	signature, err := ioutil.ReadFile("/home/steve/base64/file.txt.sha256")
	fmt.Println(err)

	fmt.Println()
	fmt.Println("Form str3")
    str3 := base64.URLEncoding.EncodeToString(signature)
	fmt.Println()
	fmt.Println(str3)
}
