package sarepay

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	// library version
	version = "0.1.0"

	// defaultHTTPTimeout is the default timeout on the http client
	defaultHTTPTimeout = 60 * time.Second

	// base URL for all Sarepay API requests
	baseURL = "https://staging-app.sarepay.com/api"
)

type service struct {
	client *Client
}

// Client manages communication with the Sarepay API
type Client struct {
	common  service      // Reuse a single struct instead of allocating one for each service on the heap.
	client  *http.Client // HTTP client used to communicate with the API.
	baseURL *url.URL
	// the API Key used to authenticate all Sarepay API requests
	publicKey string
	secretKey string

	//logger Logger
	// Services supported by the Sarepay API.
	// Miscellaneous actions are directly implemented on the Client object
	Transaction    *TransactionService
	VirtualAccount *VirtualAccountService
	Transfer       *TransferService
	// SubAccount  *SubAccountService
	// Charge      *ChargeService

	LoggingEnabled bool
	Log            Logger
}

// Logger interface for custom loggers
type Logger interface {
	Printf(format string, v ...interface{})
}

// Response represents arbitrary response data
type Response map[string]interface{}

// NewClient creates a new Sarepay API client with the given API key
func NewClient(publickey, secretkey string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: defaultHTTPTimeout}
	}

	u, _ := url.Parse(baseURL)
	c := &Client{
		client:         httpClient,
		publicKey:      publickey,
		secretKey:      secretkey,
		baseURL:        u,
		LoggingEnabled: true,
		Log:            log.New(os.Stderr, "", log.LstdFlags),
	}

	c.common.client = c

	c.Transaction = (*TransactionService)(&c.common)
	c.VirtualAccount = (*VirtualAccountService)(&c.common)
	c.Transfer = (*TransferService)(&c.common)

	return c
}
