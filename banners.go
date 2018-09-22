package admitad

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type Banners struct {
	Results []Banner `json:"results"`
}

type Banner struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	DirectLink string `json:"direct_link"`
}

func (c *Client) Banners(cid, wid int, params url.Values) (*Banners, error) {
	response, err := c.request("banners/"+strconv.Itoa(cid)+"/website/"+strconv.Itoa(wid), "GET", params)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	banners := &Banners{}
	err = json.NewDecoder(response.Body).Decode(banners)
	return banners, err
}
