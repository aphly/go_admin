package config

import (
	"encoding/json"
	"go_admin/app/helper"
)

type Cors struct {
	Origin []string `yaml:"Origin"`
}

func CorsConfigLoad() *Cors {
	var instance = &Cors{}
	err, str := helper.ReadJsonFile("config/cors.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(str, instance)
	return instance
}
