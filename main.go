package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/marcusrios/check-prom-query/types"
)

var (
	host               = flag.String("host", "172.16.32.115", "prometheus host")
	port               = flag.String("port", "30900", "prometheus port")
	query              = flag.String("query", "", "Query that be executed in prometheus")
	prometheusEndpoint = "http://" + *host + ":" + *port + "/api/v1/query?query="
	critical           = flag.Float64("critical", 0.0, "Critical if value is greater then")
	warning            = flag.Float64("warning", 0.0, "Warning if value is greater then")
)

func init() {
	flag.Parse()

	if *query == "" {
		fmt.Println("You need to pass the prometheus query that be executed")
		os.Exit(1)
	}
}

func main() {
	url := prometheusEndpoint + *query

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

	var response types.APIResponse

	json.Unmarshal(bytes, &response)

	switch value, _ := strconv.ParseFloat(response.FilterAPIResponseValue(), 64); {
	case value >= *critical:
		fmt.Printf("CRITICAL: critical value is %f and got %f\n", value, *critical)
		os.Exit(2)
	case value >= *warning:
		fmt.Printf("WARNING: warning value is %f and got %f\n", value, *critical)
		os.Exit(1)
	default:
		fmt.Println("OK")
		os.Exit(0)
	}
}
