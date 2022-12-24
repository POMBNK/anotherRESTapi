package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server struct {
		Port string `yml:"port"`
	}
	Db struct {
		Host     string `yml:"host"`
		Port     string `yml:"port"`
		Username string `yml:"username"`
		DbName   string `yml:"dbname"`
		SSLMode  string `yml:"sslmode"`
	}
}

var cfg *Config
var once sync.Once

func GetCfg() *Config {
	once.Do(func() {
		log.Println("Инициализация конфига...")
		cfg = &Config{}
		if err := cleanenv.ReadConfig("config.yml", cfg); err != nil {
			help, _ := cleanenv.GetDescription(cfg, nil)
			log.Println(help)
			log.Fatalln(err)
		}
	})
	return cfg
}
