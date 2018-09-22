package admitad

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	clientId string
	auth     string
	url      string
	client   *http.Client
	token    *Token
	scope    string
}

func NewClient(url, auth, id string, scope []string) *Client {
	client := &Client{
		auth:     auth,
		url:      url,
		client:   &http.Client{},
		clientId: id,
		scope:    strings.Join(scope, " "),
	}

	return client
}

type ApiError struct {
	ErrorName        string `json:"error"`
	ErrorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`
}

func (e ApiError) Error() string {
	return e.ErrorName + ": " + e.ErrorDescription
}

func (c *Client) request(u, m string, params url.Values) (*http.Response, error) {
	u = c.url + "/" + u + "/?" + params.Encode()
	request, err := http.NewRequest(m, u, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", "Bearer "+c.token.AccessToken)
	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == 200 {
		return response, nil
	}

	defer response.Body.Close()

	e := ApiError{}
	if err := json.NewDecoder(response.Body).Decode(&e); err != nil {
		return nil, err
	}
	return nil, e
}

type Token struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Language     string `json:"language"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	Scope        string `json:"scope"`
}

func (c *Client) Token() (*Token, error) {
	params := url.Values{}
	params.Add("grant_type", "client_credentials")
	params.Add("client_id", c.clientId)
	params.Add("scope", c.scope)
	u := c.url + "/token/?" + params.Encode()

	request, err := http.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", "Basic "+c.auth)
	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	token := &Token{}
	if err := json.NewDecoder(response.Body).Decode(token); err != nil {
		return nil, err
	}

	return token, nil
}

func (c *Client) Init(token *Token) {
	c.token = token
}
