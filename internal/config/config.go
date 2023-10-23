package config

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

const (
	formatJSON  = "json"
	redirectURL = "http://localhost:8080/callback"
)

type Config struct {
	Server struct {
		Host        string `envconfig:"SERVER_HOST" default:":9000"`
		MetricsBind string `envconfig:"BIND_METRICS" default:":9090"`
		HealthHost  string `envconfig:"BIND_HEALTH" default:":9091"`
	}

	Service struct {
		LogLevel  string `envconfig:"LOGGER_LEVEL" default:"debug"`
		LogFormat string `envconfig:"LOGGER_FORMAT" default:"console"`
	}

	DB struct {
		Address  string `envconfig:"DB_ADDRESS" default:"localhost"`
		Name     string `envconfig:"DB_NAME" default:"mydb"`
		User     string `envconfig:"DB_USER" default:"root"`
		Password string `envconfig:"DB_PASSWORD" default:"mydbpass"`
		Port     int    `envconfig:"DB_PORT" default:"5432"`
		MaxConn  int    `envconfig:"DB_MAX_CONN" default:"15"`
	}

	Auth struct {
		ClientID     string `envconfig:"AUTH_CLIENT_ID"`
		ClientSecret string `envconfig:"AUTH_CLIENT_SECRET"`
	}
}

func Parse() (*Config, error) {
	var cfg = &Config{}

	//мое--------------------------------------------
	err2 := godotenv.Load()
	if err2 != nil {
		log.Fatal("Error loading .env file")
	}

	err := envconfig.Process("", cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (cfg Config) Logger() (logger zerolog.Logger) {
	level := zerolog.InfoLevel
	if newLevel, err := zerolog.ParseLevel(cfg.Service.LogLevel); err == nil {
		level = newLevel
	}

	var out io.Writer = os.Stdout
	if cfg.Service.LogFormat != formatJSON {
		out = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.StampMicro}
	}
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	return zerolog.New(out).Level(level).With().Caller().Timestamp().Logger()
}

func (cfg Config) PgPoolConfig() (*pgxpool.Config, error) {
	poolCfg, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s port=%d dbname=%s sslmode=disable user=%s password=%s pool_max_conns=%d",
		cfg.DB.Address, cfg.DB.Port, cfg.DB.Name, cfg.DB.User, cfg.DB.Password, cfg.DB.MaxConn,
	))
	if err != nil {
		return nil, err
	}

	return poolCfg, nil
}

// Для Google аутентификации
func (cfg Config) SetupConfig() *oauth2.Config {
	conf := &oauth2.Config{
		RedirectURL:  redirectURL,
		ClientID:     cfg.Auth.ClientID,
		ClientSecret: cfg.Auth.ClientSecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}
