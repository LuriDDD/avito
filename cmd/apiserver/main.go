package main

// К сожалению, тут есть баги, потому что я решал все в последний день перед дедлайном. Я думал, что смогу сам выбрать, когда начать
// тестовое задание, а меня поставили перед фактом, пришло письмо, у тебя есть неделя.

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"

	"http-rest-api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
