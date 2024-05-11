// go\config\main.go

package main

import (
	"fmt"
	"go/config"
)

func main() {
	fmt.Println("Config Registrar address: ", config.SystemConfig.RegistrarAddr)
	fmt.Println("Config Resolver address: ", config.SystemConfig.ResolverAddr)
}
