package exampleconnector

import "fmt"

type Config struct {
	AttributeName string `mapstructure:"attributeName"`
}

func (c *Config) Validate() error {
	if c.AttributeName != "request.n" {
		return fmt.Errorf("the attribute for the data passed through the connector needs to be 'request.n' for this tutorial ")
	}
	return nil
}
