package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserAgent struct {
	Product  string
	Version  string
	Comment  string
	RawValue string
}

type IfConfig struct {
	Ip         string
	IpDecimal  uint32
	Country    string
	CountryEu  bool
	CountryIso string
	City       string
	Latitude   float32
	Longitude  float32
	Asn        string
	AsnOrg     string
	UserAgent  UserAgent
}

func main() {
	resp, err := http.Get("http://ifconfig.co/json")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if_config := new(IfConfig)

	err = json.NewDecoder(resp.Body).Decode(&if_config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n%s", if_config.Ip, if_config.Country)
}
