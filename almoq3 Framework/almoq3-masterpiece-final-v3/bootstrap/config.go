package bootstrap

import (
	"fmt"
	"log"
	"os"

	"almoq3-masterpiece-final-v3/config"
	"github.com/fatih/color"
	"github.com/spf13/viper"
)

func InitializeConfig() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("%s .env file not found. Please create one based on .env.example", color.RedString("FATAL:"))
		} else {
			log.Fatalf("%s Error reading .env file: %v", color.RedString("FATAL:"), err)
		}
	}

	// Map to struct
	err := viper.Unmarshal(&config.Global.App)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	// Manual mapping for nested structs if needed, or use Viper's Unmarshal with tags
	config.Global.App.Name = viper.GetString("APP_NAME")
	config.Global.App.Env = viper.GetString("APP_ENV")
	config.Global.App.Port = viper.GetString("APP_PORT")
	config.Global.App.Debug = viper.GetBool("APP_DEBUG")
	config.Global.App.URL = viper.GetString("APP_URL")
	config.Global.App.Key = viper.GetString("APP_KEY")

	config.Global.Database.Driver = viper.GetString("DB_DRIVER")
	config.Global.Database.Host = viper.GetString("DB_HOST")
	config.Global.Database.Port = viper.GetString("DB_PORT")
	config.Global.Database.Database = viper.GetString("DB_DATABASE")
	config.Global.Database.Username = viper.GetString("DB_USERNAME")
	config.Global.Database.Password = viper.GetString("DB_PASSWORD")

	config.Global.Cache.Driver = viper.GetString("CACHE_DRIVER")
	config.Global.Cache.Host = viper.GetString("CACHE_HOST")
	config.Global.Cache.Port = viper.GetString("CACHE_PORT")
	config.Global.Cache.Password = viper.GetString("CACHE_PASSWORD")

	// Validation
	validateConfig()

	fmt.Printf("%s Configuration loaded successfully\n", color.CyanString("⚙️"))
}

func validateConfig() {
	if config.Global.App.Key == "" {
		fmt.Printf("%s APP_KEY is missing in .env. This is a critical security risk.\n", color.RedString("ARCHITECTURAL ERROR:"))
		os.Exit(1)
	}

	if config.Global.Database.Driver == "" {
		fmt.Printf("%s DB_DRIVER is not specified in .env.\n", color.RedString("ARCHITECTURAL ERROR:"))
		os.Exit(1)
	}
}
