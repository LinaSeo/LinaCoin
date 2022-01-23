package main

import (
	"github.com/LinaSeo/LinaCoin/explorer"
	"github.com/LinaSeo/LinaCoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
