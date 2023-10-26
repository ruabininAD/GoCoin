package main

import (
	"fmt"
	"goapp/GoApp/configs"
	. "goapp/GoApp/pkg/cli"
)
import . "goapp/GoApp/pkg/blockchain"

func main() {
	fmt.Printf("TargetBits: %x\n", configs.TargetBits)

	bc := NewBlockchain()
	defer bc.CloseDB()

	cli := CLI{Bc: bc}
	cli.Run()

}
