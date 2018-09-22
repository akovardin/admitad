package admitad

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type AdvCampaigns struct {
	Results []AdvCampaign `json:"results"`
}

type AdvCampaign struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Gotolink string `json:"gotolink"`
}

func (c *Client) AdvCampaigns(params url.Values) (*AdvCampaigns, error) {
	response, err := c.request("advcampaigns", "GET", params)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	campaigns := &AdvCampaigns{}
	err = json.NewDecoder(response.Body).Decode(campaigns)
	return campaigns, err
}

func (c *Client) AdvCampaign(id int, params url.Values) (*AdvCampaign, error) {
	response, err := c.request("advcampaigns/"+strconv.Itoa(id), "GET", params)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	campaign := &AdvCampaign{}
	err = json.NewDecoder(response.Body).Decode(campaign)
	return campaign, err
}

func (c *Client) AdvCampaignsByWebsite(id int, params url.Values) (*AdvCampaigns, error) {
	response, err := c.request("advcampaigns/website/"+strconv.Itoa(id), "GET", params)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	campaigns := &AdvCampaigns{}
	err = json.NewDecoder(response.Body).Decode(campaigns)
	return campaigns, err
}
