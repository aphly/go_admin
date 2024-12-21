package config

import (
	"encoding/json"
	"go_admin/app/helper"
)

// {
// 	"AccessKeyId":"",
// 	"AccessKeySecret":"",
// 	"Endpoint":"",
// 	"Bucket" : "",
// 	"IsCname":"false",
// 	"Url"  :""
// }

type Oss struct {
	AccessKeyId     string `yaml:"AccessKeyId"`
	AccessKeySecret string `yaml:"AccessKeySecret"`
	Endpoint        string `yaml:"Endpoint"`
	Bucket          string `yaml:"Bucket"`
	IsCname         string `yaml:"IsCname"`
	Url             string `yaml:"Url"`
}

func OssConfigLoad() *Oss {
	var instance = Oss{}
	err, str := helper.ReadJsonFile("config/oss.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(str, &instance)
	return &instance
}
