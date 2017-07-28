package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

func main() {
	requrl := "http://127.0.0.1:8080/api/login"
	posturl := "http://127.0.0.1:8080/api/post"
	client := http.Client{}
	jar, err := cookiejar.New(nil)
	CheckErrorOnExit(err)
	client.Jar = jar

	bodyType := "application/json"
	user := make(map[string]interface{})
	user["username"] = "abd"
	user["password"] = "111111"
	bytesData, err := json.Marshal(user)
	CheckErrorOnExit(err)

	reader := bytes.NewReader(bytesData)
	resp, err := client.Post(requrl, bodyType, reader)
	CheckErrorOnExit(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	CheckErrorOnExit(err)
	fmt.Println("body:", string(body))
	for _, c := range resp.Cookies() {
		fmt.Println(c)
	}

	type loginResponse struct {
		Result  int    `json:"result"`
		Message string `json:"message"`
	}
	var lr loginResponse
	json.Unmarshal(body, &lr)
	if lr.Result != 1 {
		fmt.Println("failed to login")
		return
	}

	// post
	post := make(map[string]interface{})
	post["title"] = "post from go client"
	post["body"] = "test...test....test....."
	bytesData, err = json.Marshal(post)
	CheckErrorOnExit(err)

	reader = bytes.NewReader(bytesData)
	resp, err = client.Post(posturl, bodyType, reader)
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	CheckErrorOnExit(err)
	fmt.Println("post result:", string(body))
}

func CheckErrorOnExit(err error) {
	if err != nil {
		panic(err)
	}
}
