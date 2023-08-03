package cfg

import (
	"path"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Cfg struct {
	Port      string
	FileName  string
	BatchSize int
}

func LoadAndStoreConfig() Cfg {
	v := viper.New()
	v.SetEnvPrefix("MAKVES")

	v.SetDefault("PORT", "8080")
	v.SetDefault("FILENAME", path.Join("data", "ueba.csv"))
	v.SetDefault("BATCHSIZE", 4096)

	v.SetConfigFile(".env")
	_ = v.ReadInConfig()
	v.AutomaticEnv()

	var cfg Cfg
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Panic(err)
	}
	log.Println(cfg)

	return cfg
}
