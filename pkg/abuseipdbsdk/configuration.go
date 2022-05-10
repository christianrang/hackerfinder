package abuseipdbsdk

type Configuration struct {
	ApiKey string `mapstructure:"api_key"`
}

func (c *Configuration) SetApiKey(value string) {
	c.ApiKey = value
}
