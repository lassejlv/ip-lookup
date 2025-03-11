package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ApiResponse struct {
	Ip              string `json:"ip"`
	IpNumber        string `json:"ip_number"`
	IpVersion       int    `json:"ip_version"`
	CountryName     string `json:"country_name"`
	CountryCode2    string `json:"country_code2"`
	ISP             string `json:"isp"`
	ResponseCode    string `json:"response_code"`
	ResponseMessage string `json:"response_message"`
}

var API_URL string = "https://api.iplocation.net"

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Provide an IP please")
		os.Exit(1)
	}

	ip := os.Args[1]

	resp, err := http.Get(fmt.Sprintf("%s/?ip=%s", API_URL, ip))

	if err != nil {
		fmt.Println("No response sadly :(")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var result ApiResponse
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err)
	}

	fmt.Printf("Country: %s \n", result.CountryName)
	fmt.Printf("ISP: %s \n", result.ISP)
	fmt.Printf("IP-type: ipv%d \n", result.IpVersion)

}
