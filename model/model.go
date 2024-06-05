package model

type Setting struct {
	Port          string `mapstructure:"port"`
	RedisIp       string `mapstructure:"redis_ip"`
	RedisPort     string `mapstructure:"redis_port"`
	RedisPassword string `mapstructure:"redis_password"`
	RedisDB       int    `mapstructure:"redis_db"`
}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}