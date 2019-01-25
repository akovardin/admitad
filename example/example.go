package main

import (
	"fmt"
	"log"
	"net/url"
	"strconv"

	"github.com/horechek/admitad"
)

var (
	base64Header = "xxxxx=="
	clientId     = "xxx"
)

type Websites struct {
	Results []map[string]interface{} `json:"results"`
}

type Campaigns struct {
	Results []map[string]interface{} `json:"results"`
}

type Banners struct {
	Results []map[string]interface{} `json:"results"`
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	client := admitad.NewClient(
		"https://api.admitad.com",
		base64Header,
		clientId,
		[]string{"advcampaigns", "banners", "websites", "advcampaigns_for_website", "banners_for_website"},
	)

	token, err := client.Token()
	if err != nil {
		log.Fatal(err)
	}
	client.Init(token)

	websites := &Websites{}
	err = client.Call("websites", "GET", url.Values{}, websites)
	if err != nil {
		log.Fatal(err)
	}

	banners := []map[string]interface{}{}
	for _, w := range websites.Results {
		p := url.Values{}
		p.Add("limit", "4")
		campaigns := Campaigns{}
		err := client.Call("advcampaigns/website/"+strconv.Itoa(int(w["id"].(float64))), "GET", p, &campaigns)
		if err != nil {
			log.Println(err)
			continue
		}

		for _, c := range campaigns.Results {
			list := &Banners{}
			err := client.Call("banners/"+strconv.Itoa(int(c["id"].(float64)))+"/website/"+strconv.Itoa(int(w["id"].(float64))), "GET", url.Values{}, list)
			if err != nil {
				log.Println(err)
				continue
			}

			banners = append(banners, list.Results...)
		}
	}

	for _, b := range banners {
		fmt.Println("==============")
		for k, v := range b {
			fmt.Println(k+":", v)
		}
	}
}
