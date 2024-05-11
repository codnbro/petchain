// examples/config/main.go
package main

import (
	"fmt"
	"petchain/config"
)

func main() {
	fmt.Println("Config Registrar address: ", config.SystemConfig.RegistrarAddr)
	fmt.Println("Config Resolver address: ", config.SystemConfig.ResolverAddr)
}

/*
Config Registrar address:  localhost:9000
Config Resolver address:  localhost:9001
*/
