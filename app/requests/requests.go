package requests

import (
	"bytes"
	"net/http"
	"encoding/json"
	"github.com/madoka/app/common"
)

// CreateContainerRequests is Requests 
func CreateContainerRequests(c common.Container) {
	url := "/win"
	json, _ := json.Marshal(c)
	req, _ := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(json)),
	)
	req.Header.Set("Content-type", "application/json")
}
