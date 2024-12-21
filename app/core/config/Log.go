package config

import (
	"encoding/json"
	"go_admin/app/helper"
)

type Log struct {
	Path          string `yaml:"Path"`
	MaxSize       int64  `yaml:"MaxSize"`
	MaxBufferSize int64  `yaml:"MaxBufferSize"`
}

func LogConfigLoad() *Log {
	var instance = &Log{}
	err, str := helper.ReadJsonFile("config/log.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(str, instance)
	return instance
}
