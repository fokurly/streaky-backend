package utils

import (
	"fmt"
	"io"
	"os"

	"github.com/fokurly/streaky-backend/users_info_api/models"
	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

func ParseDatabaseConfigByKey(key string, fullPath bool) models.DatabaseConfig {
	if key == "" {
		logrus.Panicf("key for client init config is empty")
	}

	var filePath string
	if fullPath {
		filePath = key
	} else {
		filePath = fmt.Sprintf("./storage/postgre/config/%s.json", key)
	}

	file, err := os.Open(filePath)
	if err != nil {
		logrus.Panicf("could not open config file by <%s> key. error: %v", key, err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		logrus.Panicf("could not read info from file config by <%s> key. error: %v", key, err)
	}
	var config models.DatabaseConfig
	if err = jsoniter.Unmarshal(data, &config); err != nil {
		logrus.Panicf("could not correctly unmarshal config info from file with <%s> key. error: %v", key, err)
	}
	v := validator.New()
	if err = v.Struct(config); err != nil {
		logrus.Panicf("could not validate config by <%s> key. error: %v", key, err)
	}
	return config
}
