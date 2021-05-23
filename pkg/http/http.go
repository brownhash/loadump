package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/brownhash/golog"
)

func MakeRequest(method, url string, header map[string]string, body interface{}) (string, []byte, error) {
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(body)

	client := &http.Client{}
	request, err := http.NewRequest(method, url, payload)

	if err != nil {
		golog.Error(err)
	}

	if len(header) > 0 {
		for key, value := range header {
			request.Header.Set(key, value)
		}
	}

	response, err := client.Do(request)

	if err != nil {
		return "", nil, err
	}

	respBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", nil, err
	}

	return response.Status, respBody, err
}
