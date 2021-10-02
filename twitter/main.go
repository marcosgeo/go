package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/oauth2"
)

func main() {
	var keys struct {
		Key    string `json:"api_key"`
		Secret string `json:"api_secret_key"`
	}
	// get current path
	_path, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		panic(err)
	}
	path := strings.TrimSpace(string(_path))
	f, err := os.Open(fmt.Sprintf("%s/twitter/.keys.json", path))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	dec.Decode(&keys)
	// fmt.Printf("%+v\n", keys)

	req, err := http.NewRequest("POST", "https://api.twitter.com/oauth2/token",
		strings.NewReader("grant_type=client_credentials"))
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(keys.Key, keys.Secret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	var client http.Client
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var token oauth2.Token
	dec = json.NewDecoder(res.Body)
	err = dec.Decode(&token)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)

	var conf oauth2.Config
	tclient := conf.Client(context.Background(), &token)
	res2, err := tclient.Get("https://api.twitter.com/1.1/statuses/retweets/991053593250758658.json")
	if err != nil {
		panic(err)
	}
	defer res2.Body.Close()
	io.Copy(os.Stdout, res2.Body)
}
