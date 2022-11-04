package config

import (
	"encoding/json"
	"net"
	"os"
	"strings"
)

const (
	dbPassEscSeq = "{password}"
	password     = "note-service-password"
)

type Config interface {
	GetDBConfig() (string, error)
	GetGRPCAddress() string
	GetHTTPAddress() string
}

type DB struct {
	DSN               string `json:"dsn"`
	MaxOpenConnection int32  `json:"max_open_connections"`
}

type GRPC struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type HTTP struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Configure struct {
	DB   *DB   `json:"db"`
	GRPC *GRPC `json:"grpc"`
	HTTP *HTTP `json:"http"`
}

func NewConfig(path string) (Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Configure{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Configure) GetGRPCAddress() string {
	return net.JoinHostPort(c.GRPC.Host, c.GRPC.Port)
}

func (c *Configure) GetHTTPAddress() string {
	return net.JoinHostPort(c.HTTP.Host, c.HTTP.Port)
}

func (c *Configure) GetDBConfig() (string, error) {
	dbDsn := strings.ReplaceAll(c.DB.DSN, dbPassEscSeq, password)

	return dbDsn, nil
}
