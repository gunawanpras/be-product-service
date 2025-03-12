package config

type (
	Config struct {
		Server  ServerConfig `yaml:"server"`
		Postgre PostgreList  `yaml:"postgre"`
		Redis   RedisList    `yaml:"redis"`
	}

	ServerConfig struct {
		Name string `yaml:"name"`
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}

	PostgreList struct {
		Primary DatabaseConfig `yaml:"primary"`
	}

	DatabaseConfig struct {
		ConnString              string `yaml:"connString"`
		MigrationConnString     string `yaml:"migrateConnString"`
		MaxOpenConn             int    `yaml:"maxOpenConn"`
		MaxIdleConn             int    `yaml:"maxIdleConn"`
		MaxConnLifeTimeInSecond int    `yaml:"maxConnLifeTimeInSecond"`
	}

	RedisList struct {
		Primary RedisConfig `yaml:"primary"`
	}

	RedisConfig struct {
		Host         string `yaml:"host"`
		Port         int    `yaml:"port"`
		Password     string `yaml:"password"`
		Db           int    `yaml:"db"`
		Ttl          int    `yaml:"ttl"`
		DialTimeout  int    `yaml:"dial_timeout"`
		ReadTimeout  int    `yaml:"read_timeout"`
		WriteTimeout int    `yaml:"write_timeout"`
	}
)
