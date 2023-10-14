package utils

import (
	"fmt"

	"github.com/beego/beego/httplib"
)

// for get api request
func Get_api_request(url string, channel chan string) {
	req := httplib.Get(url)
	res, err := req.String()
	if err != nil {
		fmt.Println(err)
	} else {
		channel <- res
	}
}
