package sylius

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"highlite-parser/internal/log"
	"highlite-parser/internal/sylius/transfer"
)

const (
	tokenRequestRetryCount int           = 10
	tokenRefreshInterval   time.Duration = 30 * time.Minute
	requestTimeout         time.Duration = 5 * time.Second

	methodGet   string = "get"
	methodPost  string = "post"
	methodPatch string = "patch"
)

var _ IClient = (*Client)(nil)

// ErrNotFound tells that http request returned 404 Status
var ErrNotFound = errors.New("not found")

// IClient is a Sylius Client interface
type IClient interface {
	GetTaxon(ctx context.Context, code string) (*transfer.Taxon, error)
	CreateTaxon(ctx context.Context, body transfer.TaxonNew) (*transfer.Taxon, error)
	GetProduct(ctx context.Context, code string) (*transfer.ProductEntire, error)
	CreateProduct(ctx context.Context, body transfer.Product) (*transfer.ProductEntire, error)
	UpdateProduct(ctx context.Context, body transfer.Product) error
	CreateProductVariant(ctx context.Context, product string, body transfer.ProductVariantNew) (*transfer.ProductVariant, error)
}

// NewClient is a Sylius Client constructor.
func NewClient(logger log.ILogger, endpoint string, auth Auth) *Client {
	c := &Client{
		endpoint:       endpoint,
		auth:           auth,
		logger:         logger,
		tokenChan:      make(chan *transfer.Token),
		requestTimeout: requestTimeout,
	}

	go c.tokenServer()

	// wait until the token is received
	c.getToken()

	return c
}

// Auth contains credentials to obtain Sylius API token.
type Auth struct {
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
}

// Client is a sylius api client
type Client struct {
	endpoint       string
	auth           Auth
	logger         log.ILogger
	requestTimeout time.Duration
	tokenChan      chan *transfer.Token
}

// SetRequestTimeout sets request timeout
func (c *Client) SetRequestTimeout(t time.Duration) {
	c.requestTimeout = t
}

// Returns full url using endpoint and resource paths.
func (c *Client) getURL(path string, args ...interface{}) string {
	if len(args) > 0 {
		path = fmt.Sprintf(path, args...)
	}

	return strings.TrimSuffix(c.endpoint, "/") + "/" + strings.TrimPrefix(path, "/")
}
