package main

import (
	"addons/addons"
	"fmt"
)

func main() {
	addons.StordbConnect()
	addons.MySqlConnect()
	cc, err := addons.IsBlackListed("791020", "792015")
	fmt.Println(cc, err)
	// api.RPCServer()
	// api.RPCClient()
	// api.ServerStart()

}
