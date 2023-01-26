package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strings"
	"time"
)

var domains []string
var uploadDomains []string

type Response struct {
	StatusCode int    `json:"status-code"`
	Status     string `json:"status"`
}

func main() {

	domainFlagValue := flag.String("gateway-domains", "https://shuttle-4-bs1.estuary.tech,https://shuttle-4-bs2.estuary.tech", "")
	uploadDomainFlagValue := flag.String("upload-domains", "https://upload.estuary.tech", "")

	domains = strings.Split(*domainFlagValue, ",")
	uploadDomains = strings.Split(*uploadDomainFlagValue, ",")

	http.HandleFunc("/health", HealthCheck)
	http.HandleFunc("/", CheckRequestUriAndForward)

	// start the http server
	fmt.Println("Started the forwarder...... at 8080")
	http.ListenAndServe(":8080", nil)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Status:     "ok",
		StatusCode: http.StatusOK,
	}

	// set the content type to json
	w.Header().Set("Content-Type", "application/json")
	// set the http status code
	w.WriteHeader(resp.StatusCode)
	// encode the struct to json and write to the response body
	json.NewEncoder(w).Encode(resp)
}

// function to check if the endpoint is /gw or anything else

// function to redirect. if /gw then redirect to the gateway domains and if not then redirect to the upload domains
func CheckRequestUriAndForward(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	if strings.HasPrefix(r.URL.Path, "/gw") {
		for _, url := range domains {
			var lastErr error
			client := &http.Client{
				Timeout: time.Second * 60,
			}

			req, err := http.NewRequest(r.Method, url+r.URL.Path, r.Body)
			if err != nil {
				lastErr = err
				continue
			}
			req.Header = r.Header
			resp, err := client.Do(req)
			if err != nil {
				lastErr = err
				continue
			}
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				http.Redirect(w, r, url+r.URL.Path, http.StatusTemporaryRedirect)
				return
			}
			fmt.Errorf("error %s", lastErr)
		}
	} else {
		for _, url := range uploadDomains {
			fmt.Println(url + r.URL.Path)
			var lastErr error
			client := &http.Client{
				Timeout: time.Second * 60,
			}

			req, err := http.NewRequest(r.Method, url+r.URL.Path, r.Body)
			if err != nil {
				lastErr = err
				continue
			}
			req.Header = r.Header
			resp, err := client.Do(req)
			if err != nil {
				lastErr = err
				continue
			}
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				http.Redirect(w, r, url+r.URL.Path, http.StatusTemporaryRedirect)
				return
			}
			fmt.Errorf("error %s", lastErr)
		}
	}
}
