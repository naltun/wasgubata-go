package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// GeoLocation holds the data returned after making a GET request
// to https://ipinfo.io/<ipAddr>/geo.
type GeoLocation struct {
	IP          string `json:"ip,omitempty"`
	City        string `json:"city,omitempty"`
	Region      string `json:"region,omitempty"`
	Country     string `json:"country,omitempty"`
	Coordinates string `json:"loc,omitempty"`
}

func geoDomain(s string) string {
}

func geoIP(s string) string {
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("[ERROR] No argument supplied. Exiting.")
		os.Exit(1)
	}
	arg := os.Args[1]

	// I need to check which function to call, eg. geoDomain or geoIP

	cmd := exec.Command("dig", "+short", arg)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
		os.Exit(1)
	}
	ip := strings.TrimSpace(out.String())

	res, err := http.Get("https://ipinfo.io/" + ip + "/geo")
	if err != nil {
		fmt.Printf("[ERROR]: %s\n", err)
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("[ERROR]: %s\n", err)
		os.Exit(1)
	}

	var loc GeoLocation
	err = json.Unmarshal(data, &loc)
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
		os.Exit(1)
	}
}
