package config

import (
	"github.com/tkanos/gonfig"
)

type MongoConfig struct {
	ConnectionUrl string
}

func (c *MongoConfig) Read() {
	if err := gonfig.GetConf("config/mongo.json", c); err != nil {
		panic(err)
	}
}
