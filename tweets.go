package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"	
	"github.com/dghubble/go-twitter/twitter"	
	"io/ioutil"
	"os"	
)

type Secrets struct {
	ApiKey string `json:"api_key"`
	ApiSecretKey string `json:"api_secret_key"`
}

func ReadSecrets() Secrets {
	jsonFile, _ := os.Open("secrets.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var secrets Secrets
	json.Unmarshal(byteValue, &secrets)
	return secrets
}

func GetLatestTweet() (string, string) {
	var secrets Secrets = ReadSecrets()

	config := &clientcredentials.Config{
		ClientID:     secrets.ApiKey,
		ClientSecret: secrets.ApiSecretKey,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	httpClient := config.Client(oauth2.NoContext)	

	client := twitter.NewClient(httpClient)

	user, resp, err := client.Users.Show(&twitter.UserShowParams{
		ScreenName: "realDonaldTrump",
	})

	if resp.StatusCode == 200 {
		return user.Status.Text, user.Status.CreatedAt
	} else {
		fmt.Println("could not get last tweet!")
		fmt.Println("response: " + resp.Status)
		fmt.Println("error: " + err.Error())
		return "", ""
	}

}
