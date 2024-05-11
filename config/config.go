package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	SystemConfig Configuration
)

// 패키지가 로드될 때 가장 먼저 호출되는 함수
func init() {
	configFileName := os.Getenv("CONFIG_FILE")
	if configFileName == "" {
		configFileName = ".\\config\\config.json"
	}
	SystemConfig = loadConfigration(configFileName)
}

// JSON struct
type Configuration struct {
	IssuerAddr           string `json:"issuer_addr"`
	VerifierAddr         string `json:"verifier_addr"`
	RegistrarAddr        string `json:"registrar_addr"`
	RegistrarGatewayAddr string `json:"registrar_gateway_addr"`
	ResolverAddr         string `json:"resolver_addr"`
	ResolverGatewayAddr  string `json:"resolver_gateway_addr"`
}

func loadConfigration(path string) Configuration {
	var configuration Configuration
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&configuration)
	return configuration
}
