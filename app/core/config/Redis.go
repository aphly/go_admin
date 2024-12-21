package config

type Redis struct {
	Single  []Single `yaml:"Single"`
	Cluster Cluster  `yaml:"Cluster"`
}

type Single struct {
	Addr     string `yaml:"Addr"`
	Password string `yaml:"Password"`
	PoolSize int    `yaml:"PoolSize"`
	Retries  int    `yaml:"Retries"`
	Db       int    `yaml:"Db"`
}

type Cluster struct {
	Addrs    []string `yaml:"Addrs"`
	Password string   `yaml:"Password"`
	PoolSize int      `yaml:"PoolSize"`
	Retries  int      `yaml:"Retries"`
}
