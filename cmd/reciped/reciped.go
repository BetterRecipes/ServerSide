package main


import "fmt"
import "github.com/BetterRecipe/ServerSide/internal/config"



func main() {
	fmt.Printf("%+v\n", config.ReadDatabaseConfig("configs/reciped.toml"))
	fmt.Printf("%+v\n", config.ReadLogConfig("configs/reciped.toml"))
	fmt.Printf("%+v\n", config.ReadServerConfig("configs/reciped.toml"))
}
