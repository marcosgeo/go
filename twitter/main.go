package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var keys struct {
		Key    string `json:"api_key"`
		Secret string `json:"api_secret_key"`
	}
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
	fmt.Printf("%+v\n", keys)
}
