package utils

//The data will be in json format. So to use it by controllers, we need to unmarshal it.

import (
	"encoding/json"
	"io"
	"os"
	"net/http"
)

func ParseBody(r *http.Request,x interface{}){
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
}