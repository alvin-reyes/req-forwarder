package main

import (
	"flag"
	"fmt"
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
		for _, url := range domains {
			var lastErr error
			client := &http.Client{
				Timeout: time.Second * 60,
			}
			fmt.Println(url + r.URL.Path)
			req, err := http.NewRequest("GET", url+r.URL.Path, r.Body)
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
