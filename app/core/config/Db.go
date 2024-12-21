package config

type DbSingle struct {
	Host           string `yaml:"Host"`
	Port           int    `yaml:"Port"`
	Database       string `yaml:"Database"`
	Username       string `yaml:"Username"`
	Password       string `yaml:"Password"`
	Charset        string `yaml:"Charset"`
	TimeOut        int    `yaml:"TimeOut"`
	WriteTimeOut   int    `yaml:"WriteTimeOut"`
	ReadTimeOut    int    `yaml:"ReadTimeOut"`
	MaxIdleConnect int    `yaml:"MaxIdleConnect"`
	MaxOpenConnect int    `yaml:"MaxOpenConnect"`
}

//func DbConfigLoad() *map[string]DbGroup {
//	var instance = make(map[string]DbGroup)
//	err, str := helper.ReadJsonFile("config/db.json")
//	if err != nil {
//		panic(err)
//	}
//	err = json.Unmarshal(str, &instance)
//	return &instance
//}
