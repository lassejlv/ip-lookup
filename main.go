package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/lassejlv/iplookup/utils"
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
var IP string

func main() {

	if len(os.Args) < 2 {
		IP = utils.FetchUsersIp()
		fmt.Println("👀 Using your networks ip since no ip was provided")
	} else {
		IP = os.Args[1]
	}

	full_url := fmt.Sprintf("%s/?ip=%s", API_URL, IP)

	resp, err := http.Get(full_url)

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
