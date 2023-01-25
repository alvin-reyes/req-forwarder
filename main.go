package main

import (
	"flag"
	"net/http"
	"strings"
	"time"
)

var domains []string

func main() {

	domainFlagValue := flag.String("domains", "https://shuttle-4-bs1.estuary.tech,https://shuttle-4-bs2.estuary.tech", "")

	domains = strings.Split(*domainFlagValue, ",")

	// create a custom http handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//var resp *http.Response

		for _, url := range domains {
			var _ error
			client := &http.Client{
				Timeout: time.Second * 10,
			}
			req, err := http.NewRequest(r.Method, url+r.URL.Path, r.Body)
			if err != nil {
				_ = err
				continue
			}
			req.Header = r.Header
			resp, err := client.Do(req)
			if err != nil {
				_ = err
				continue
			}
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				w.WriteHeader(http.StatusOK)
				return
			}
		}

		//defer resp.Body.Close()
	})

	// start the http server
	http.ListenAndServe(":8080", nil)
}
