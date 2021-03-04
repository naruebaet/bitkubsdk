package curl

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func HttpGet(url string, out interface{}) (statusCode int, err error) {
	resp, err := http.Get(url)

	// Error checking of the http.Get() request
	if err != nil {
		log.Fatal(err)
	}

	// Resource leak if response body isn't closed
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	// Error checking of the ioutil.ReadAll() request
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("!http status code : ", resp.StatusCode, string(bodyBytes))
		return resp.StatusCode, nil
	}

	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		log.Println("ticker error : ", err.Error())
		return resp.StatusCode, err
	}

	return resp.StatusCode, nil
}

func HttpGetRaw(url string) (raw []byte, statusCode int, err error) {
	resp, err := http.Get(url)

	// Error checking of the http.Get() request
	if err != nil {
		log.Fatal(err)
	}

	// Resource leak if response body isn't closed
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	// Error checking of the ioutil.ReadAll() request
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("!http status code : ", resp.StatusCode, string(bodyBytes))
		return []byte{}, resp.StatusCode, nil
	}

	return bodyBytes, resp.StatusCode, nil
}
