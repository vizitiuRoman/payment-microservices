package bitcoin

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/vizitiuRoman/blockchain-api/pkg/blockchain"
)

type Client struct {
	client *http.Client

	config *Config
}

type Config struct {
	BaseURL  string
	Password string
	Username string
}

func NewClient(config *Config) *Client {
	return &Client{
		client: &http.Client{
			Timeout: time.Second * 30,
		},
		config: config,
	}
}

func (c *Client) CreateWallet() (*blockchain.Account, error) {
	res, err := c.call("createwallet", "")
	if err != nil {
		return nil, err
	}
	address := res.(string)

	privateKey, err := c.call("dumpprivkey", address)
	if err != nil {
		return nil, err
	}
	_, err = c.call("setlabel", address, address)
	if err != nil {
		return nil, err
	}

	return &blockchain.Account{
		Address:    address,
		PrivateKey: privateKey.(string),
	}, nil
}

func (c *Client) call(method string, params ...string) (interface{}, error) {
	u, err := url.Parse(c.config.BaseURL)
	if err != nil {
		return nil, err
	}
	u = u.ResolveReference(&url.URL{Path: method})

	body := &Body{
		ID:       1,
		JsonRPC:  "2.0",
		Method:   method,
		Params:   params,
		Username: c.config.Username,
		Password: c.config.Password,
	}

	buf := &bytes.Buffer{}
	err = json.NewEncoder(buf).Encode(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), buf)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
