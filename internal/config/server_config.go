package config

type ServerConfig struct {
	Name     string `json:"name"                 mapstructure:"name"`
	Address  string `json:"address"                 mapstructure:"address"`
	ApiToken string `json:"api_token"                 mapstructure:"api_token"`
}
