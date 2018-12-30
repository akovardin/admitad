package admitad

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/suite"
)

type BannersSuite struct {
	suite.Suite
	server *httptest.Server
}

func (s *BannersSuite) SetupSuite() {
	s.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case "/banners/1/website/2/?":
			fmt.Fprint(w, bannersTestData)
		}
	}))
}

func (s *BannersSuite) TearDownSuite() {
	if s.server != nil {
		s.server.Close()
	}
}

func (s *BannersSuite) TestBanners() {
	client := NewClient(
		s.server.URL,
		"xxx",
		"168d7bd617162c4ff00ae33e17eb5e",
		[]string{"advcampaigns", "banners", "websites"},
	)
	client.Init(&Token{})

	type Banners struct {
		Results []map[string]interface{} `json:"results"`
	}

	campaigns := Banners{}
	err := client.Call("banners/1/website/2", "GET", url.Values{}, &campaigns)
	s.Assertions.NoError(err)
	s.Assertions.Len(campaigns.Results, 2)
}

func TestBannersSuite(t *testing.T) {
	suite.Run(t, new(BannersSuite))
}

const (
	bannersTestData = `{
    "results": [
        {
            "ecpc": "0.00",
            "name": "120x600_Литература для детей",
            "mobile_content": false,
            "size_width": 120,
            "direct_link": "http://ad.admitad.com/g/92e3447eb1b72ea2020f27ad1bb5fb/",
            "creation_date": "2013-04-30T14:41:12",
            "size_height": 600,
            "html_code": "<!-- admitad.banner: 92e3447eb1b72ea2020f27ad1bb5fb Буквоед -->\n<a target=\"_blank\" rel=\"nofollow\" href=\"http://ad.admitad.com/g/92e3447eb1b72ea2020f27ad1bb5fb/?i=4\"><img width=\"120\" height=\"600\" border=\"0\" src=\"http://ad.admitad.com/b/92e3447eb1b72ea2020f27ad1bb5fb/\" alt=\"Буквоед\"/></a>\n<!-- /admitad.banner -->",
            "banner_image_url": "https://cdn.admitad.com/bs/2013/04/30/4f7a7707810a3775ea0421b000ed8adb.jpg",
            "is_flash": false,
            "type": "jpeg",
            "id": 40160
        },
        {
            "ecpc": "0.00",
            "name": "160x600_Литература для детей",
            "mobile_content": false,
            "size_width": 160,
            "direct_link": "http://ad.admitad.com/g/516083f8b0b72ea2020f27ad1bb5fb/",
            "creation_date": "2013-04-30T14:41:18",
            "size_height": 600,
            "html_code": "<!-- admitad.banner: 516083f8b0b72ea2020f27ad1bb5fb Буквоед -->\n<a target=\"_blank\" rel=\"nofollow\" href=\"http://ad.admitad.com/g/516083f8b0b72ea2020f27ad1bb5fb/?i=4\"><img width=\"160\" height=\"600\" border=\"0\" src=\"http://ad.admitad.com/b/516083f8b0b72ea2020f27ad1bb5fb/\" alt=\"Буквоед\"/></a>\n<!-- /admitad.banner -->",
            "banner_image_url": "https://cdn.admitad.com/bs/2013/04/30/245465838198c725361f3cb170575036.jpg",
            "is_flash": false,
            "type": "jpeg",
            "id": 40161
        }
    ],
    "_meta": {
        "count": 72,
        "limit": 2,
        "offset": 0
    }
}`
)
