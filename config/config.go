package config

type Config struct {
	Application Application
	Database    DatabaseConfig
	Redis       RedisConfig
}

type Application struct {
	ServiceName string
	Port        string
}

type DatabaseConfig struct {
	Host                  string
	Port                  string
	User                  string
	Password              string
	Name                  string
	MaxIdleConnection     int
	MaxOpenConnection     int
	MaxLifetimeConnection int
	MaxIdletimeConnection int
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}
