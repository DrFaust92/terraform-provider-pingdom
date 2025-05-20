package pingdom

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/DrFaust92/go-pingdom/pingdom"
	"github.com/DrFaust92/go-pingdom/pingdomext"
	"github.com/hashicorp/go-retryablehttp"
)

// Config respresents the client configuration
type Config struct {
	APIToken           string `mapstructure:"api_token"`
	SolarwindsUser     string `mapstructure:"solarwinds_user"`
	SolarwindsPassword string `mapstructure:"solarwinds_passwd"`
	SolarwindsOrgID    string `mapstructure:"solarwinds_org_id"`
	MaxRetries         int    `mapstructure:"max_retries"`
	RetryWaitMin       int    `mapstructure:"retry_wait_min"`
	RetryWaitMax       int    `mapstructure:"retry_wait_max"`
}

type Clients struct {
	Pingdom    *pingdom.Client
	PingdomExt *pingdomext.Client
	// Solarwinds *solarwinds.Client
}

func (c *Config) Client() (*Clients, error) {
	pingdomClient, err := c.pingdomClient()
	if err != nil {
		return nil, err
	}
	if user := os.Getenv("SOLARWINDS_USER"); user != "" {
		if password := os.Getenv("SOLARWINDS_PASSWD"); password != "" {
			c.SolarwindsUser = user
			c.SolarwindsPassword = password
		} else {
			return nil, errors.New("user and password must be present together")
		}
	}
	if orgID := os.Getenv("SOLARWINDS_ORG_ID"); orgID != "" {
		c.SolarwindsOrgID = orgID
	}
	// solarwindsClient, err := solarwinds.NewClient(solarwinds.ClientConfig{
	// 	Username: c.SolarwindsUser,
	// 	Password: c.SolarwindsPassword,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// err = solarwindsClient.Init()
	// if err != nil {
	// 	return nil, err
	// }

	var pingdomClientExt *pingdomext.Client
	if c.SolarwindsUser != "" {
		pingdomClientExt, err = pingdomext.NewClientWithConfig(pingdomext.ClientConfig{
			Username: c.SolarwindsUser,
			Password: c.SolarwindsPassword,
			OrgID:    c.SolarwindsOrgID,
		})

		if err != nil {
			return nil, err
		}
	}

	return &Clients{
		Pingdom:    pingdomClient,
		PingdomExt: pingdomClientExt,
		// Solarwinds: solarwindsClient,
	}, nil
}

// Client returns a new client for accessing pingdom.
func (c *Config) pingdomClient() (*pingdom.Client, error) {
	if v := os.Getenv("PINGDOM_API_TOKEN"); v != "" {
		c.APIToken = v
	}

	// Set default retry configuration if not provided
	maxRetries := 3
	if c.MaxRetries > 0 {
		maxRetries = c.MaxRetries
	}

	retryWaitMin := 1
	if c.RetryWaitMin > 0 {
		retryWaitMin = c.RetryWaitMin
	}

	retryWaitMax := 30
	if c.RetryWaitMax > 0 {
		retryWaitMax = c.RetryWaitMax
	}

	// Create a retryable HTTP client
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = maxRetries
	retryClient.RetryWaitMin = time.Duration(retryWaitMin) * time.Second
	retryClient.RetryWaitMax = time.Duration(retryWaitMax) * time.Second

	// Add logging for retries
	retryClient.Logger = nil // Disable default logger
	retryClient.RequestLogHook = func(logger retryablehttp.Logger, req *http.Request, retry int) {
		if retry > 0 {
			log.Printf("[INFO] Retrying Pingdom API request to %s (attempt %d/%d)", req.URL, retry, maxRetries)
		}
	}

	// Get standard client with retry transport
	standardClient := retryClient.StandardClient()

	// Create Pingdom client with the retry-enabled HTTP client
	client, _ := pingdom.NewClientWithConfig(pingdom.ClientConfig{
		APIToken:   c.APIToken,
		HTTPClient: standardClient,
	})

	log.Printf("[INFO] Pingdom Client configured with retry capability (max attempts: %d).", maxRetries+1)

	return client, nil
}
