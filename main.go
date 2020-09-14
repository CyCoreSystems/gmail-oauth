package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var configFile = "credentials.json"

func main() {
	cfgJSON, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalln("failed to read config:", err)
	}

	config, err := google.ConfigFromJSON(cfgJSON, "https://mail.google.com/")
	if err != nil {
		log.Fatalln("failed to parse config:", err)
	}

	fmt.Println()
	fmt.Println("1. Ensure that you are logged in as your user in your browser.")
	fmt.Println()
	fmt.Println("2. Open the following link and authorise sendgmail:")
	fmt.Println(config.AuthCodeURL("state", oauth2.AccessTypeOffline))
	fmt.Println()
	fmt.Println("3. Enter the authorisation code:")
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Failed to read authorisation code: %v.", err)
	}
	fmt.Println()
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		log.Fatalf("Failed to exchange authorisation code for token: %v.", err)
	}

	if json.NewEncoder(os.Stdout).Encode(token); err != nil {
		log.Fatalln("failed to encode token:", err)
	}

	return
}
