# Admitad API

A Go wrapper around Admitad API

## Install

```
go get github.com/horechek/admitad
```

## Method

There are 2 common methods to communicate with api. `Token` and `Call`:

```go
Token() (*Token, error)
Call(url, method string, params url.Values, result interface{}) error
```

## Example

Request access token:

```go
var (
    base64Header = "xxx"
    clientId = "xxx"
)

client := NewClient(
    "https://api.admitad.com",
    base64Header,
    clientId,
    []string{"advcampaigns", "banners", "websites", "advcampaigns_for_website", "banners_for_website"},
)

token, err := client.Token()
if err != nil {
    return nil, err
}
client.Init(token)
```

Fetch banners from API

```go
websites := &Websites{}
err = client.Call("websites", "GET", url.Values{}, websites)
if err != nil {
    return nil, err
}

banners := []Banner{}
for _, w := range websites.Results {
    p := url.Values{}
    p.Add("limit", "4")
    campaigns := AdvCampaigns{}
    err := client.Call("advcampaigns/website/" + strconv.Itoa(w.Id), "GET", p, &campaigns)
    if err != nil {
        return nil, err
    }

    for _, c := range campaigns.Results {
        list := &Banners{}
        err := client.Call("banners/"+strconv.Itoa(c.Id)+"/website/"+strconv.Itoa(w.Id), "GET", url.Values{}, list)
        if err != nil {
            return nil, err
        }

        banners = append(banners, list.Results...)
    }
}

fmt.Println(banners)
```

`Banner`, `Websites`, `AdvCampaigns`, `Banners` - it's types from current application, not from package

## Credentials

You can get `base64Header` and `clientId` here: https://www.admitad.com/ru/webmaster/account/settings/credentials/