package shodan

import (
	"encoding/json"
	"fmt"
	"io"
	"ipsearch/http"
)

type Record struct {
	Cpes      []interface{} `json:"cpes"`
	Hostnames []string      `json:"hostnames"`
	Ip        string        `json:"ip"`
	Ports     []int         `json:"ports"`
	Tags      []interface{} `json:"tags"`
	Vulns     []interface{} `json:"vulns"`
}

func InternetDBLookup(ip string) (payload Record, err error) {

	resp, err := http.Get(internetdbUri + ip)
	if err != nil {
		fmt.Printf("the horror: %v", err)
	}

	if resp.StatusCode != 200 {
		err := fmt.Errorf("%d: not found", resp.StatusCode)
		return payload, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("the horror: %v\n", err)
	}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		fmt.Printf("the horror: %v", err)
	}

	return
}
