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
	"github.com/marcusrios/check-prom-query/utils"
)

var (
	host     = flag.String("host", "", "prometheus host")
	port     = flag.String("port", "30900", "prometheus port")
	query    = flag.String("query", "", "Query that be executed in prometheus")
	critical = flag.Float64("critical", 0.0, "Critical if value is greater than")
	warning  = flag.Float64("warning", 0.0, "Warning if value is greater than")
	lessThan = flag.Bool("lt", false, "Change whether value is less than check")
)

const (
	criticalStatus = 2
	warningStatus  = 1
	okStatus       = 0
)

func init() {
	flag.Parse()

	if *query == "" {
		fmt.Println("You need to pass the prometheus query that be executed")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *host == "" {
		fmt.Println("You need to pass the host")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	utils.RemoveHTTPPrefix(host)
	url := "http://" + *host + ":" + *port + "/api/v1/query?query=" + *query

	encodedURL := utils.EncodeURL(url)
	resp, err := http.Get(encodedURL)

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		fmt.Printf("Failed to query prometheus. Got %d status code\n", resp.StatusCode)
		os.Exit(1)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	var response types.APIResponse

	if err := json.Unmarshal(bytes, &response); err != nil {
		log.Fatal(err)
	}

	if *lessThan {
		switch value, _ := strconv.ParseFloat(response.FilterAPIResponseValue(), 64); {
		case value < *critical:
			fmt.Printf("CRITICAL: critical value is %.4f and got %.4f\n", *critical, value)
			os.Exit(criticalStatus)
		case value < *warning:
			fmt.Printf("WARNING: warning value is %.4f and got %.4f\n", *warning, value)
			os.Exit(warningStatus)
		default:
			fmt.Printf("OK - %.4f\n", value)
			os.Exit(okStatus)
		}
	} else {
		switch value, _ := strconv.ParseFloat(response.FilterAPIResponseValue(), 64); {
		case value >= *critical:
			fmt.Printf("CRITICAL: critical value is %.4f and got %.4f\n", *critical, value)
			os.Exit(criticalStatus)
		case value >= *warning:
			fmt.Printf("WARNING: warning value is %.4f and got %.4f\n", *warning, value)
			os.Exit(warningStatus)
		default:
			fmt.Printf("OK - %.4f\n", value)
			os.Exit(okStatus)
		}
	}
}
