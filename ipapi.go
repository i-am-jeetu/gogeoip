//Package ipapi is an unofficial client for ip-api.com.
// Designed to be simple to use, without the need for any authentication.
// Do not use in production as the limit is very low i.e 45  requests/minute
package gogeoip

import (
	"bytes"
	"encoding/json"
	"errors"
	"net"
	"net/http"
)

const ipapiURL string = "http://ip-api.com/json/"
const fields string = "?fields=66846719"

// IPData struct holds the data acquired from the ip-api.com for a given Ip Address ...
type IPData struct {
	Query       string  `json:"query"`
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
}

// GetCountryCode returns the two-letter country code ISO 3166-1 alpha-2e for a given Ip.
func (ipdata *IPData) GetCountryCode() string {
	return ipdata.CountryCode
}

// GetRegion returns the region/state short code (FIPS or ISO).
func (ipdata *IPData) GetRegion() string {
	return ipdata.Region
}

// GetRegionName returns the Region/state for a given Ip
func (ipdata *IPData) GetRegionName() string {
	return ipdata.RegionName
}

// GetCity returns the CITY
func (ipdata *IPData) GetCity() string {
	return ipdata.City
}

// GetTimezone returns the Timezone
func (ipdata *IPData) GetTimezone() string {
	return ipdata.Timezone
}

// GetIsp gives the ISP.
func (ipdata *IPData) GetIsp() string {
	return ipdata.Isp
}

// GetOrg returns the organisation
func (ipdata *IPData) GetOrg() string {
	return ipdata.Org
}

// GetAs returns the AS number and organization, separated by space (RIR). Empty for IP blocks not being announced in BGP tables.
func (ipdata *IPData) GetAs() string {
	return ipdata.As
}

// GetLatitude gives Latitude
func (ipdata *IPData) GetLatitude() float64 {
	return ipdata.Lat
}

//GetLongitude gives Longitude
func (ipdata *IPData) GetLongitude() float64 {
	return ipdata.Lon
}

// GetZip gives Zip
func (ipdata *IPData) GetZip() string {
	return ipdata.Zip
}

// GetCountry is
func (ipdata *IPData) GetCountry() string {
	return ipdata.Country
}

// FromString creates a new Ipdata object from String.
func FromString(ip string) (*IPData, error) {
	resp, err := http.Get(ipapiURL + ip + fields)

	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	ipData := new(IPData)
	err = json.Unmarshal(buf.Bytes(), &ipData)

	if ipData.Status != "success" || err != nil {
		return nil, errors.New("IPAPI: Ip Address is Invalid")
	}

	return ipData, nil
}

// FromIP creates a new Ipdata object from IP Address.
func FromIP(ip *net.IP) (*IPData, error) {
	ipString := ip.String()
	ipdata, err := FromString(ipString)
	return ipdata, err
}
