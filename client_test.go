package admitad

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ClientSuite struct {
	suite.Suite
	server *httptest.Server
}

func (s *ClientSuite) SetupSuite() {
	s.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, tokenTestData)
	}))
}

func (s *ClientSuite) TearDownSuite() {
	if s.server != nil {
		s.server.Close()
	}
}

func (s *ClientSuite) TestToken() {
	client := NewClient(
		s.server.URL,
		"xxx",
		"168d7bd617162c4ff00ae33e17eb5e",
		[]string{"advcampaigns", "banners", "websites"},
	)

	t, err := client.Token()

	s.Assertions.NoError(err)
	s.Assertions.EqualValues("horechek", t.Username)
	s.Assertions.EqualValues("82335ebbc7797d533e11", t.AccessToken)
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(ClientSuite))
}

const (
	tokenTestData = `{
	"username": "horechek", 
	"first_name": "\u0410\u0440\u0435\u0442\u043c", 
	"last_name": "\u041a\u043e\u0432\u0430\u0440\u0434\u0438\u043d", 
	"language": "ru", 
	"access_token": "82335ebbc7797d533e11", 
	"expires_in": 604800, 
	"token_type": "bearer", 
	"scope": "advcampaigns banners websites", 
	"id": 616295, 
	"refresh_token": "d8daa7a042a6e80ae109"
}`
)
