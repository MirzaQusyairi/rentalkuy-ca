package geolocation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentalkuy-ca/business/geolocation"
)

type IpAPI struct {
	httpClient http.Client
}

func NewIpAPI() geolocation.Repository {
	return &IpAPI{
		httpClient: http.Client{},
	}
}

func (geo *IpAPI) GetLocationByIP(ip string) (geolocation.Domain, error) {
	ip = "175.158.36.2"
	req, _ := http.NewRequest("GET", "https://ipapi.co/"+ip+"/json/", nil)
	req.Header.Set("User-Agent", "ipapi.co/#go-v1.3")
	resp, err := geo.httpClient.Do(req)
	if err != nil {
		return geolocation.Domain{}, err
	}

	defer resp.Body.Close()

	data := Response{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return geolocation.Domain{}, err
	}
	fmt.Println(data)
	return data.toDomain(), nil
}
