package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Profile  string
	AWS      *aws
	Postgres *postgres
}

type aws struct {
	Region       string
	ClientID     string
	ClientSecret string
}

type postgres struct {
	dbname   string
	user     string
	password string
	host     string
	port     string //TODO - should be int16
	/*
	 * disable - No SSL
	 * require - Always SSL (skip verification)
	 * verify-ca - Always SSL (verify that the certificate presented by the server was signed by a trusted CA)
	 * verify-full - Always SSL (verify that the certification presented by the server was signed by a trusted CA and the server host name matches the one in the certificate)
	 */
	sslmode string
	// fallback_application_name string
	// connect_timeout           string
	// sslcert                   string
	// sslkey                    string
	// sslrootcert               string
}

func NewConfig() *Config {
	return &Config{
		Profile: os.Getenv("PROFILE"),
		AWS: &aws{
			Region:       os.Getenv("AWS_REGION"),
			ClientID:     os.Getenv("AWS_CLIENT_ID"),
			ClientSecret: os.Getenv("AWS_CLIENT_SECRET"),
		},
		Postgres: &postgres{
			dbname:   os.Getenv("POSTGRES_DBNAME"),
			user:     os.Getenv("POSTGRES_USER"),
			password: os.Getenv("POSTGRES_PASS"),
			host:     os.Getenv("POSTGRES_HOST"),
			port:     os.Getenv("POSTGRES_PORT"),
		},
	}
}

func (c *Config) Local() bool {
	return len(c.Profile) == 0 || strings.ToLower(c.Profile) == "local"
}

func (c *Config) Cloud() bool {
	return strings.ToLower(c.Profile) == "aws"
}

func (p *postgres) DatasourceURL() string {
	return fmt.Sprintf("user=%s dbname=%s sslmode=%s", p.user, p.dbname, p.sslmode)
}
