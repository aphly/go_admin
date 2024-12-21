package config

type Config struct {
	AppKey   string                `yaml:"AppKey"`
	PageSize int                   `yaml:"PageSize"`
	Cors     Cors                  `yaml:"Cors"`
	Http     Http                  `yaml:"Http"`
	Db       map[string][]DbSingle `yaml:"Db"`
	Log      Log                   `yaml:"Log"`
	Redis    Redis                 `yaml:"Redis"`
	Oss      Oss                   `yaml:"Oss"`
}

//func NewConfig() *Config {
//	//var config = &Config{}
//	//config.Cors = CorsConfigLoad()
//	//config.Http = HttpConfigLoad()
//	//config.Db = DbConfigLoad()
//	//config.Log = LogConfigLoad()
//	//config.Redis = RedisConfigLoad()
//	////config.Oss = OssConfigLoad()
//	//return config
//}
