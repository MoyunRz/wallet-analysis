package utils

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func HttpGet(urlStr string) (string, error) {
	resp, err := http.Get(urlStr)
	if err != nil {
		// handle error
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return "", err
	}
	return string(body), nil
}

func HttpPost(urlStr string, params *strings.Reader) (string, error) {
	client := &http.Client{}
	// url, strings.NewReader("name=cjb")
	req, err := http.NewRequest("POST", urlStr, params)
	if err != nil {
		// handle error
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")
	resp, err := client.Do(req)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return "", err
	}
	return string(body), nil
}


func HttpPostForm(urlStr string,values url.Values) (string, error) {
    resp, err := http.PostForm(urlStr, values)

    if err != nil {
        // handle error
		return "", err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
		return "", err
	}
	return string(body), nil

}