package main

import (
	"encoding/json"
	"fmt"
	"github.com/madoka/app/common"
	"io/ioutil"
	"net/http"
)

// CreateContainerRequests is Requests
func CreateContainerRequests() {
	client := &http.Client{}
	url := "http://localhost:8080/container/list/id"
	req, _ := http.NewRequest(
		"GET",
		url,
		nil,
	)

	
	req.Header.Set("Content-type", "application/json")
	res, err := client.Do(req)
	fmt.Println(err)
	data, _ := ioutil.ReadAll(res.Body)

	containerIDList := new(common.ResCotainerIDList)
	json.Unmarshal(data, &containerIDList)
	res.Body.Close()
	fmt.Println(*containerIDList)
}

func main() {
	CreateContainerRequests()
}
