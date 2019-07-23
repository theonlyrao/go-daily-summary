package main

//import "github.com/dghubble/go-twitter/twitter"
import (
	"encoding/json"
	"fmt"
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

func GetLatestTweet() string {
	var secrets Secrets = ReadSecrets()
	fmt.Println(secrets)
	return ""
}
