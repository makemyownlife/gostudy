package cms

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func GetHomeData() {
	cmsUrl := "http://cmsapi.biyao.com/homepage/getData?pageSize=10&curPageNumber=1";
	resp, err := http.Get(cmsUrl)
	if err != nil {
		// handle error
		fmt.Println("can't get data from cmsapi")
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
}
