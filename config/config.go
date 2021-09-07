package config

// @Created 06/09/2021
// @Updated
type Config struct {
	Postgres PostgresConfig
}

// @Created 06/09/2021
// @Updated
type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlDbName   string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlSSLMode  bool
}
