package config

import (
	"encoding/json"
	"go_admin/app/helper"
)

type Http struct {
	Listen string `yaml:"Listen"`
	Host   string `yaml:"Host"`
}

func HttpConfigLoad() *Http {
	var instance = &Http{}
	err, str := helper.ReadJsonFile("config/http.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(str, instance)
	return instance
}
