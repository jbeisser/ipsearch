package shodan

import (
	"encoding/json"
	"fmt"
	"io"
	"ipsearch/http"
)

type DnsResponse struct {
	Answers []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"answers"`
	FromLoc struct {
		City    string `json:"city"`
		Country string `json:"country"`
		Latlon  string `json:"latlon"`
	} `json:"from_loc"`
}

type GeoDnsResponse []DnsResponse

const (
	Dns    = geoNetUri + "/dns/"
	GeoDns = geoNetUri + "/geodns/"
)

func GeoDnsLookup(name string) (response *[]DnsResponse, err error) {
	resp, err := http.Get(GeoDns + name)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error: %s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &response)
	return
}
