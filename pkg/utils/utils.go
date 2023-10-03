package utils

import (
	"encoding/json"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(x); err != nil {
		return err
	}

	return nil
}

/*
package utils

import(
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}){
	if body, err := ioutil.ReadAll(r.Body);err == nil{
		if err := json.Unmarshal([]byte(body), x); err != nil{
			return
		}
	}
}
*/
