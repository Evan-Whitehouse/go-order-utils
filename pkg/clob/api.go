package clob

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Evan-Whitehouse/go-order-utils/pkg/utils"
)

// ApiCreds represents the API credentials structure
type ApiCreds struct {
	ApiKey        string `json:"apiKey"`
	ApiSecret     string `json:"secret"`
	ApiPassphrase string `json:"passphrase"`
}

// RequestArgs represents the arguments for a request
type RequestArgs struct {
	Method      string
	RequestPath string
}

// SignerType represents the signer interface
type SignerType interface {
	// Define methods required for signing
}

// Client represents the client structure
type Client struct {
	Host   string
	Signer SignerType
	Creds  *ApiCreds
	Logger *log.Logger
	Mode   string
}

// NewClient creates a new client instance
func NewClient(signer SignerType) *Client {
	return &Client{
		Host:   ClobBaseUrl,
		Signer: signer,
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

// CreateAPIKey creates a new CLOB API key
func (c *Client) CreateAPIKey(nonce *int) (*ApiCreds, error) {
	if err := c.assertLevel1Auth(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s%s", c.Host, CREATE_API_KEY)
	headers := createLevel1Headers(c.Signer, nonce)

	var credsRaw map[string]string
	if err := utils.Post(endpoint, headers, nil, &credsRaw); err != nil {
		c.Logger.Println("Couldn't parse created CLOB creds")
		return nil, err
	}

	creds := &ApiCreds{
		ApiKey:        credsRaw["apiKey"],
		ApiSecret:     credsRaw["secret"],
		ApiPassphrase: credsRaw["passphrase"],
	}

	return creds, nil
}

// DeriveAPIKey derives an existing CLOB API key
func (c *Client) DeriveAPIKey(nonce *int) (*ApiCreds, error) {
	if err := c.assertLevel1Auth(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s%s", c.Host, DERIVE_API_KEY)
	headers := createLevel1Headers(c.Signer, nonce)

	var credsRaw map[string]string
	if err := utils.Get(endpoint, headers, &credsRaw); err != nil {
		c.Logger.Println("Couldn't parse derived CLOB creds")
		return nil, err
	}

	creds := &ApiCreds{
		ApiKey:        credsRaw["apiKey"],
		ApiSecret:     credsRaw["secret"],
		ApiPassphrase: credsRaw["passphrase"],
	}

	return creds, nil
}

// CreateOrDeriveAPICreds creates or derives API credentials
func (c *Client) CreateOrDeriveAPICreds(nonce *int) (*ApiCreds, error) {
	creds, err := c.CreateAPIKey(nonce)
	if err != nil {
		creds, err = c.DeriveAPIKey(nonce)
		if err != nil {
			return nil, err
		}
	}
	return creds, nil
}

// SetAPICreds sets the client API credentials
func (c *Client) SetAPICreds(creds *ApiCreds) {
	c.Creds = creds
	c.Mode = c.getClientMode()
}

// GetAPIKeys retrieves available API keys for the address
func (c *Client) GetAPIKeys() (interface{}, error) {
	if err := c.assertLevel2Auth(); err != nil {
		return nil, err
	}

	requestArgs := RequestArgs{
		Method:      "GET",
		RequestPath: GET_API_KEYS,
	}
	headers := createLevel2Headers(c.Signer, c.Creds, requestArgs)

	var response interface{}
	if err := utils.Get(fmt.Sprintf("%s%s", c.Host, GET_API_KEYS), headers, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// DeleteAPIKey deletes an API key
func (c *Client) DeleteAPIKey() (interface{}, error) {
	if err := c.assertLevel2Auth(); err != nil {
		return nil, err
	}

	requestArgs := RequestArgs{
		Method:      "DELETE",
		RequestPath: DELETE_API_KEY,
	}
	headers := createLevel2Headers(c.Signer, c.Creds, requestArgs)

	var response interface{}
	if err := utils.DeleteRequest(fmt.Sprintf("%s%s", c.Host, DELETE_API_KEY), headers, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// Helper methods

func (c *Client) assertLevel1Auth() error {
	if c.Signer == nil {
		return errors.New("level 1 authentication not initialized")
	}
	return nil
}

func (c *Client) assertLevel2Auth() error {
	if c.Creds == nil {
		return errors.New("level 2 authentication not initialized")
	}
	return c.assertLevel1Auth()
}

func (c *Client) getClientMode() string {
	// Implement logic to determine the client mode
	return "default_mode"
}

func createLevel1Headers(signer SignerType, nonce *int) map[string]string {
	headers := make(map[string]string)
	// Add required headers
	return headers
}

func createLevel2Headers(signer SignerType, creds *ApiCreds, args RequestArgs) map[string]string {
	headers := make(map[string]string)
	// Add required headers
	return headers
}
