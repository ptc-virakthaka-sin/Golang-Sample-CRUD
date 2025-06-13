package config

import (
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Host             string `mapstructure:"host"`
		Port             string `mapstructure:"port"`
		RateLimit        int    `mapstructure:"rate_limit"`
		IdleTimeout      int    `mapstructure:"idle_timeout"`
		FailedAttempts   int    `mapstructure:"failed_attempts"`
		EnablePrintRoute bool   `mapstructure:"enable_print_route"`

		App      `mapstructure:",squash"`
		Log      `mapstructure:",squash"`
		DB       `mapstructure:",squash"`
		Email    `mapstructure:",squash"`
		Redis    `mapstructure:",squash"`
		PermLogs `mapstructure:",squash"`
		JWT      `mapstructure:",squash"`
		AES      `mapstructure:",squash"`
	}

	App struct {
		Env     string `mapstructure:"app_env"`
		Name    string `mapstructure:"app_name"`
		Version string `mapstructure:"app_version"`
	}

	Log struct {
		Level string `mapstructure:"log_level"`
	}

	DB struct {
		Host        string `mapstructure:"db_host"`
		Port        string `mapstructure:"db_port"`
		User        string `mapstructure:"db_user"`
		Password    string `mapstructure:"db_password"`
		Name        string `mapstructure:"db_name"`
		UseSSL      string `mapstructure:"db_use_ssl"`
		AutoMigrate bool   `mapstructure:"db_auto_migrate"`
	}

	Email struct {
		Host     string `mapstructure:"email_host"`
		Port     string `mapstructure:"email_port"`
		Username string `mapstructure:"email_username"`
		Password string `mapstructure:"email_password"`
	}

	Redis struct {
		Addr     string `mapstructure:"redis_addr"`
		Port     string `mapstructure:"redis_port"`
		Password string `mapstructure:"redis_pwd"`
	}

	PermLogs struct {
		TabEnableText string `mapstructure:"tab_enable_text"`
		TableNames    []string
	}

	JWT struct {
		Secret string `mapstructure:"jwt_secret"`
	}

	AES struct {
		KEY string `mapstructure:"aes_key"`
	}
)

var Cfg *Config

func Init() error {
	// 1. load existing env vars are found into viper
	viper.AutomaticEnv()

	// 2. read .env file if exists
	var result map[string]interface{}

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("error reading config file, %s\n", err)
	}
	err := viper.Unmarshal(&result)
	if err != nil {
		fmt.Printf("unable to decode into map, %v\n", err)
	}

	// use WeakDecode to ignore typed error
	decErr := mapstructure.WeakDecode(result, &Cfg)
	if decErr != nil {
		fmt.Printf("error decode map to config, %v\n", decErr)
	}

	// 3. replace match any of the existing env vars used in flags
	pflag.String("host", "", "server host")
	pflag.Int("port", 8000, "server port")
	pflag.Int("rate_limit", 100, "rate limit")
	pflag.Int("idle_timeout", 5, "server idle timeout")
	pflag.Int("failed_attempts", 5, "login failed attempts")
	pflag.Bool("enable_print_route", true, "server enable print route")

	pflag.String("app_env", "local", "app env")
	pflag.String("app_name", "ptbank core api", "app name")
	pflag.String("app_version", "0.0.1", "app version")

	pflag.String("log_level", "info", "log level")

	pflag.String("db_host", "", "database host")
	pflag.String("db_port", "", "database port")
	pflag.String("db_user", "", "database user")
	pflag.String("db_password", "", "database password")
	pflag.String("db_name", "", "database name")
	pflag.String("db_use_ssl", "", "database use ssl")
	pflag.Bool("db_auto_migrate", false, "database auto migrate")

	pflag.String("email_host", "", "database host")
	pflag.String("email_port", "", "database port")
	pflag.String("email_username", "", "database username")
	pflag.String("email_password", "", "database password")

	pflag.String("redis_addr", "", "redis address")
	pflag.String("redis_port", "6379", "redis port")
	pflag.String("redis_pwd", "", "redis password")

	pflag.String("tab_enable_text", "", "tablenames before split")
	pflag.String("jwt_secret", "", "jwt secret")
	pflag.String("aes_key", "", "aes key")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	// fmt.Println("server port:", viper.GetInt64("port"))
	// fmt.Println("app env:", viper.GetString("app_env"))

	// 4. bind viper to config struct
	Cfg = &Config{
		Host:             viper.GetString("host"),
		Port:             viper.GetString("port"),
		RateLimit:        viper.GetInt("rate_limit"),
		IdleTimeout:      viper.GetInt("idle_timeout"),
		FailedAttempts:   viper.GetInt("failed_attempts"),
		EnablePrintRoute: viper.GetBool("enable_print_route"),
		App: App{
			Env:     viper.GetString("app_env"),
			Name:    viper.GetString("app_name"),
			Version: viper.GetString("app_version"),
		},
		Log: Log{
			Level: viper.GetString("log_level"),
		},
		DB: DB{
			Host:        viper.GetString("db_host"),
			Port:        viper.GetString("db_port"),
			User:        viper.GetString("db_user"),
			Password:    viper.GetString("db_password"),
			Name:        viper.GetString("db_name"),
			UseSSL:      viper.GetString("db_use_ssl"),
			AutoMigrate: viper.GetBool("db_auto_migrate"),
		},
		Email: Email{
			Host:     viper.GetString("email_host"),
			Port:     viper.GetString("email_port"),
			Username: viper.GetString("email_username"),
			Password: viper.GetString("email_password"),
		},
		Redis: Redis{
			Addr:     viper.GetString("redis_addr"),
			Port:     viper.GetString("redis_port"),
			Password: viper.GetString("redis_pwd"),
		},
		PermLogs: PermLogs{
			TabEnableText: viper.GetString("tab_enable_text"),
		},
		JWT: JWT{
			Secret: viper.GetString("jwt_secret"),
		},
		AES: AES{
			KEY: viper.GetString("aes_key"),
		},
	}

	Cfg.PermLogs.TableNames = strings.Split(Cfg.PermLogs.TabEnableText, ",")

	return nil
}

func (db *DB) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&tls=%s&parseTime=true",
		db.User, db.Password, db.Host, db.Port, db.Name, db.UseSSL)
}
