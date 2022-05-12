package vtsdk

type Configuration struct {
	ApiKey  string `mapstructure:"api_key"`
	Premium bool   `mapstructure:"premium"`
}

func (c *Configuration) SetApiKey(value string) {
	c.ApiKey = value
}
