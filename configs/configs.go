package configs

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Server struct {
	Host string
	URL  string
	Port string
	TTL  time.Duration
}

type Redis struct {
	Addrs string
}

type Consul struct {
	Interval string
	Timeout  string
}

type Config struct {
	Server Server
	Redis  Redis
	Consul Consul
}

var C Config

func init() {
	env, ok := os.LookupEnv("ENV")
	if !ok {
		env = "development"
	}
	viper.SetConfigFile("configs/" + env + ".yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err = viper.Unmarshal(&C)
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}
