package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	Get()
}

func Get()  {
	http.HandleFunc("/dataversion", getDataVersion)
	http.ListenAndServe(":8085", nil)
	}

func getDataVersion(w http.ResponseWriter, r *http.Request) {
	//client := http.Client{}
	var body []byte
	res, err := http.NewRequest("GET","https://core-qa.phantomworksngwcore.de/spiderApplications/v1", &body)

	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
