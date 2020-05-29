package main

import (
	//h "./internal"
	"fmt"
	"log"
	"net/http"
	"regexp"
	//"net/http/httptest"
	// "encoding/json"
	"io/ioutil"
	"net/url"
	"strings"
	"testing"
)

func TestIndexPage(t *testing.T) {
	formData := url.Values{
		"bannerSelectControl": {"standard.txt"},
		"inputText":           {"123asd;as"},
	}
	req, err := http.NewRequest("POST", "http://localhost:8181/", strings.NewReader(formData.Encode()))
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	if err != nil {
		log.Fatalln(err)
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	re := regexp.MustCompile(`<pre>(.|\n)*?</pre>`)

	fmt.Println(re.FindString(string(body)))
}
