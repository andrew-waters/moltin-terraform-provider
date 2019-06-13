package moltin

import (
	"github.com/andrew-waters/gomo"
)

// config contains Moltin provider settings
type config struct {
	ClientID     string
	ClientSecret string
}

func (c *config) client() (*gomo.Client, error) {
	client := gomo.NewClient(
		gomo.NewClientCredentials(
			c.ClientID,
			c.ClientSecret,
		),
	)
	return &client, client.Authenticate()
}
