package admitad

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/suite"
)

type WebsitesSuite struct {
	suite.Suite
	server *httptest.Server
}

func (s *WebsitesSuite) SetupSuite() {
	s.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case "/websites/?":
			fmt.Fprint(w, websitesTestData)
		case "/websites/123?":
			fmt.Fprint(w, websiteTestData)
		}
	}))
}

func (s *WebsitesSuite) TearDownSuite() {
	if s.server != nil {
		s.server.Close()
	}
}

func (s *WebsitesSuite) TestWebsites() {
	client := NewClient(
		s.server.URL,
		"xxx",
		"168d7bd617162c4ff00ae33e17eb5e",
		[]string{"advcampaigns", "banners", "websites"},
	)
	client.Init(&Token{})

	type Websites struct {
		Results []map[string]interface{} `json:"results"`
	}

	result := Websites{}

	err := client.Call("websites", "GET", url.Values{}, &result)
	s.Assertions.NoError(err)
	s.Assertions.Len(result.Results, 2)
}

func (s *WebsitesSuite) Website() {
	client := NewClient(
		s.server.URL,
		"xxx",
		"168d7bd617162c4ff00ae33e17eb5e",
		[]string{"advcampaigns", "banners", "websites"},
	)
	client.Init(&Token{})

	type Website struct {
		Name string `json:"name"`
	}

	result := Website{}

	err := client.Call("websites/123", "GET", url.Values{}, &result)
	s.Assertions.NoError(err)
	s.Assertions.EqualValues("Язык программирования Go", result.Name)
}

func TestWebsitesSuite(t *testing.T) {
	suite.Run(t, new(WebsitesSuite))
}

const (
	websitesTestData = `{
    "results": [
        {
            "status": "active",
            "account_id": "",
            "kind": "website",
            "is_old": false,
            "name": "Сайт о языке программирования",
            "collecting_method": null,
            "verification_code": "64026f39b0",
            "mailing_targeting": false,
            "site_url": "http://4gophers.ru/",
            "regions": [
                {
                    "region": "RU",
                    "id": 1772826
                }
            ],
            "db_size": null,
            "adservice": null,
            "id": 710629,
            "creation_date": "2017-09-01T12:16:18",
            "validation_passed": true,
            "categories": [
                {
                    "language": "ru",
                    "id": 18,
                    "parent": null,
                    "name": "Интернет услуги"
                },
                {
                    "language": "ru",
                    "id": 21,
                    "parent": {
                        "language": "ru",
                        "id": 18,
                        "parent": null,
                        "name": "Интернет услуги"
                    },
                    "name": "Хостинги"
                },
                {
                    "language": "ru",
                    "id": 70,
                    "parent": {
                        "language": "ru",
                        "id": 62,
                        "parent": null,
                        "name": "Интернет-магазины"
                    },
                    "name": "Книги"
                },
                {
                    "language": "ru",
                    "id": 122,
                    "parent": {
                        "language": "ru",
                        "id": 18,
                        "parent": null,
                        "name": "Интернет услуги"
                    },
                    "name": "IT-решения"
                }
            ],
            "description": "Размещаю ссылки на своем сайте о языке программирования Go и в группе в vk"
        },
        {
            "status": "active",
            "account_id": "",
            "kind": "social_group",
            "is_old": false,
            "name": "Язык программирования Go",
            "collecting_method": null,
            "verification_code": "b72ea2020f",
            "mailing_targeting": false,
            "site_url": "https://vk.com/golang",
            "regions": [
                {
                    "region": "RU",
                    "id": 1772840
                }
            ],
            "db_size": null,
            "adservice": {
                "id": 8,
                "name": "ВКонтакте"
            },
            "id": 710635,
            "creation_date": "2017-09-01T12:35:44",
            "validation_passed": true,
            "categories": [
                {
                    "language": "ru",
                    "id": 21,
                    "parent": {
                        "language": "ru",
                        "id": 18,
                        "parent": null,
                        "name": "Интернет услуги"
                    },
                    "name": "Хостинги"
                },
                {
                    "language": "ru",
                    "id": 70,
                    "parent": {
                        "language": "ru",
                        "id": 62,
                        "parent": null,
                        "name": "Интернет-магазины"
                    },
                    "name": "Книги"
                },
                {
                    "language": "ru",
                    "id": 122,
                    "parent": {
                        "language": "ru",
                        "id": 18,
                        "parent": null,
                        "name": "Интернет услуги"
                    },
                    "name": "IT-решения"
                }
            ],
            "description": "Я размещаю партнерские ссылки в группе в социальной сети VK"
        }
    ],
    "_meta": {
        "count": 2,
        "limit": 2,
        "offset": 0
    }
}`
	websiteTestData = `{
    "status": "active",
    "account_id": "",
    "kind": "social_group",
    "is_old": false,
    "name": "Язык программирования Go",
    "collecting_method": null,
    "verification_code": "b72ea2020f",
    "mailing_targeting": false,
    "site_url": "https://vk.com/golang",
    "regions": [
        {
            "region": "RU",
            "id": 1772840
        }
    ],
    "db_size": null,
    "adservice": {
        "id": 8,
        "name": "ВКонтакте"
    },
    "id": 710635,
    "creation_date": "2017-09-01T12:35:44",
    "validation_passed": true,
    "categories": [
        {
            "language": "ru",
            "id": 21,
            "parent": {
                "language": "ru",
                "id": 18,
                "parent": null,
                "name": "Интернет услуги"
            },
            "name": "Хостинги"
        },
        {
            "language": "ru",
            "id": 70,
            "parent": {
                "language": "ru",
                "id": 62,
                "parent": null,
                "name": "Интернет-магазины"
            },
            "name": "Книги"
        },
        {
            "language": "ru",
            "id": 122,
            "parent": {
                "language": "ru",
                "id": 18,
                "parent": null,
                "name": "Интернет услуги"
            },
            "name": "IT-решения"
        }
    ],
    "description": "Я размещаю партнерские ссылки в группе в социальной сети VK"
}`
)
