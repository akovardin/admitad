package admitad

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type Websites struct {
	Results []Website `json:"results"`
}

type Website struct {
	Id      int    `json:"id"`
	Status  string `json:"status"`
	Name    string `json:"name"`
	SiteUrl string `json:"site_url"`
}

func (c *Client) Websites(params url.Values) (*Websites, error) {
	response, err := c.request("websites", "GET", params)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	websites := &Websites{}
	if err := json.NewDecoder(response.Body).Decode(websites); err != nil {
		return nil, err
	}

	return websites, err
}

func (c *Client) Website(id int, params url.Values) (*Website, error) {
	response, err := c.request("websites/"+strconv.Itoa(id), "GET", params)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	website := &Website{}
	if err := json.NewDecoder(response.Body).Decode(website); err != nil {
		return nil, err
	}

	return website, err
}
