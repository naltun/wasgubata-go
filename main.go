package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
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

func geoDomain(domain string) string {
	cmd := exec.Command("dig", "+short", domain)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
		os.Exit(1)
	}

	return strings.TrimSpace(out.String())
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("[ERROR] No argument supplied. Exiting.")
		os.Exit(1)
	}
	arg := os.Args[1]

	var ip string
	i := strings.Split(arg, ".")
	if len(i) < 4 {
		ip = geoDomain(arg)
	}

	for _, v := range i {
		i, err := strconv.Atoi(v)
		if err != nil {
			ip = geoDomain(arg)
		} else {
			if i < 0 || i > 255 {
				ip = geoDomain(arg)
			} else {
				ip = arg
			}
		}
	}

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

	fmt.Printf("./wgb knows geolocations...\n\n")
	fmt.Println("IP:          ", loc.IP)
	fmt.Println("City:        ", loc.City)
	fmt.Println("Region:      ", loc.Region)
	fmt.Println("Country:     ", loc.Country)
	fmt.Println("Coordinates: ", loc.Coordinates)
}
