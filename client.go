package admitad

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type ApiError struct {
	ErrorName        string `json:"error"`
	ErrorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`
}

func (e ApiError) Error() string {
	return e.ErrorName + ": " + e.ErrorDescription
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

type Client struct {
	clientId     string
	base64Header string
	url          string
	token        *Token
	scope        string

	client *http.Client
}

func NewClient(url, base64Header, clientId string, scope []string) *Client {
	return &Client{
		base64Header: base64Header,
		url:          url,
		clientId:     clientId,
		scope:        strings.Join(scope, " "),
		client:       &http.Client{},
	}
}

func (c *Client) Call(url, method string, params url.Values, result interface{}) error {
	response, err := c.request(url, method, params)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(result); err != nil {
		return err
	}

	return err
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

	request.Header.Add("Authorization", "Basic "+c.base64Header)
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
