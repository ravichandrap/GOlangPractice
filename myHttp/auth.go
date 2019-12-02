package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
)

func main() {
	username:="CoREDSUser"
	password:="IFeelHardCore_AlwaysHardCore"
	paramss := url.Values{}
	paramss.Set("username", "Q29SRURTVXNlcg==")
	paramss.Set("password", "SUZlZWxIYXJkQ29yZV9BbHdheXNIYXJkQ29yZQ==")
	fmt.Println(paramss.Encode())

	data := fmt.Sprintf(username+":"+password)
	encod := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encod)
	//fmt.Println("Q29SRURTVXNlcjpJRmVlbEhhcmRDb3JlX0Fsd2F5c0hhcmRDb3Jl")
	message := encod
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(message)))
	base64.StdEncoding.Decode(base64Text, []byte(message))
	fmt.Printf("base64: %s\n", base64Text)
	//return string(base64Text)

	l, _ := base64.StdEncoding.Decode(base64Text, []byte(message))
	log.Printf("base64: %s\n", base64Text[:l])

}
