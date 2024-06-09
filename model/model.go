package model

type Setting struct {
	RouterIp      string `mapstructure:"router_ip"`
	GinPort       string `mapstructure:"gin_port"`
	FastPort      string `mapstructure:"fast_port"`
	GzipPort      string `mapstructure:"gzip_port"`
	RedisIp       string `mapstructure:"redis_ip"`
	RedisPort     string `mapstructure:"redis_port"`
	RedisPassword string `mapstructure:"redis_password"`
	RedisDB       int    `mapstructure:"redis_db"`
}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
