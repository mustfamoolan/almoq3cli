package config

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Cache    CacheConfig
}

var Global Config
