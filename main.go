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

type Response struct {
	StatusCode int    `json:"status-code"`
	Status     string `json:"status"`
}

func main() {

	domainFlagValue := flag.String("domains", "https://shuttle-4-bs1.estuary.tech,https://shuttle-4-bs2.estuary.tech", "")

	domains = strings.Split(*domainFlagValue, ",")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
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
	})
	// create a custom http handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for _, url := range domains {
			var lastErr error
			client := &http.Client{
				Timeout: time.Second * 60,
			}
			fmt.Println(url + r.URL.Path)
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
	})

	// start the http server
	fmt.Println("Started the forwarder......")
	http.ListenAndServe(":8080", nil)
}

func handleHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
