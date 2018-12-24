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

func toIP(domain string) string {
	cmd := exec.Command("dig", "+short", domain)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
		os.Exit(1)
	}

	s := out.String()
	ip := strings.Split(s, "\n")[0]

	return ip
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("[ERROR] No argument supplied. To get the help message, run with `help', eg. wgb help.")
		os.Exit(1)
	}
	arg := os.Args[1]

	if arg == "help" {
		fmt.Printf("wasgubata takes one command-line argument, either an IP address or a domainname.\n\n")
		fmt.Println("USAGE: wgb {ipaddr|domainname}")
		fmt.Println("EXAMPLE: wgb gov.vu || wgb 103.7.197.89")
		os.Exit(0)
	}

	var ip string
	i := strings.Split(arg, ".")
	if len(i) < 4 {
		ip = toIP(arg)
	}

	for _, v := range i {
		ii, err := strconv.Atoi(v)
		if err != nil {
			ip = toIP(arg)
		} else {
			if ii < 0 || ii > 255 {
				ip = toIP(arg)
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

	fmt.Printf("%s knows geolocations...\n\n", os.Args[0])
	fmt.Println("IP:          ", loc.IP)
	fmt.Println("City:        ", loc.City)
	fmt.Println("Region:      ", loc.Region)
	fmt.Println("Country:     ", loc.Country)
	fmt.Println("Coordinates: ", loc.Coordinates)
}
