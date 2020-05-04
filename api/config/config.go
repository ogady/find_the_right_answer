package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

const confDir = "./config/env/" //　設定ファイルへの実行ファイルからの相対パスを指定

var DynamoDB DynamoDBConf

type Config struct { // toml内の名前を入れる
	DynamoDB DynamoDBConf `toml:"dynamodb"`
}

type DynamoDBConf struct {
	Region       string `toml:"region"`
	Endpoint     string `toml:"endpoint"`
	AccessKey    string `toml:"access_key"`
	SecretKey    string `toml:"secret_key"`
	SessionToken string `toml:"session_token"`
}

func InitConf() error {
	var err error

	env := os.Getenv("ENV")
	if env == "" {
		panic("failed to get application mode, check whether APP_MODE is set.")
	}

	// Get configuration
	err = newConfig(confDir, env) // 引数に渡す
	if err != nil {
		return err
	}
	return err
}

// NewConfig return configuration struct.
func newConfig(path string, env string) error {
	var conf Config

	confPath := path + env + ".toml" // tomlファイルを読み設定情報を取得
	if _, err := toml.DecodeFile(confPath, &conf); err != nil {
		return err
	}

	DynamoDB = conf.DynamoDB

	return nil
}
