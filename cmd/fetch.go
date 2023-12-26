package cmd

import (
	"io"
	"log"
	"net/http"
)

// Fetch 发起请求
func Fetch() string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.jixs.cc/api/wenan-fkxqs/index.php", nil)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Http get error is ", err)
	}
	if res.StatusCode != http.StatusOK {
		log.Fatal("Http status code is ", res.StatusCode)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return ""
	}
	return string(bytes)
}
