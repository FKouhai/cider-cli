package config
import (
	"os"
)
type Config struct {
	Token string
	Port string
	Host string
}
func (c *Config) SetConfig() *Config{
	c.Token = os.Getenv("CIDER_CLI_TOKEN")
	c.Port = os.Getenv("CIDER_CLI_PORT")
	c.Host = os.Getenv("CIDER_CLI_HOST")
	return c
}
