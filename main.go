package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
)

func init() {
	if len(os.Args) < 2 {
		fmt.Println("You need te pass prometheus query")
		os.Exit(1)
	}
}

const (
	prometheusEndpoint = "http://172.16.32.115:30900/api/v1/query?query="
)

func main() {
	prometheusQuery := os.Args[1]
	url := prometheusEndpoint + prometheusQuery

	resp, err := http.Get(url)

	if resp.StatusCode != 200 {
		fmt.Printf("Failed to query prometheus. Got %d status code\n", resp.StatusCode)
		os.Exit(1)
	}

	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response interface{}

	json.Unmarshal(bytes, &response)

	fmt.Println(response.(map[string]interface{})["data"].(map[string]interface{})["result"].([]interface{})[0].(map[string]interface{})["value"].([]interface{})[0])
	fmt.Println(reflect.TypeOf(response.(map[string]interface{})["data"].(map[string]interface{})["result"].([]interface{})[0].(map[string]interface{})["value"].([]interface{})[0]))
}
