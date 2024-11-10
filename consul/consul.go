package consul

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ConsulInterface interface {
	buildHttpRequest(hostgroup string) ([]Consul, error)
	GetListIPHostgroup(hostgroup string) ([]string, error)
}

type ConsulConfig struct {
	ConsulBaseURL string
}

type Consul struct {
	ID          string `json:"ID"`
	Node        string `json:"Node"`
	Address     string `json:"Address"`
	Datacenter  string `json:"Datacenter"`
	ServiceName string `json:"ServiceName"`
	ServiceTags []any  `json:"ServiceTags"`
}

func NewConsul(url string) ConsulInterface {
	return &ConsulConfig{
		ConsulBaseURL: url,
	}
}

func (c *ConsulConfig) buildHttpRequest(hostgroup string) ([]Consul, error) {

	var response []Consul
	url := fmt.Sprintf("%s/v1/catalog/service/%s", c.ConsulBaseURL, hostgroup)

	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error occurred while making request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error occurred while reading response body: %v", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println(err.Error(), err)
		return nil, err
	}

	return response, nil

}

// GetListIPHostgroup get lists of IP Address from consul hostgroup name
func (c *ConsulConfig) GetListIPHostgroup(hostgroup string) ([]string, error) {

	list, err := c.buildHttpRequest(hostgroup)
	if err != nil {
		return nil, err
	}

	var ipaddrs []string
	for _, v := range list {
		ipaddrs = append(ipaddrs, v.Address)
	}

	return ipaddrs, nil
}
