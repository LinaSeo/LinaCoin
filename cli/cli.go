package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/LinaSeo/LinaCoin/explorer"
	"github.com/LinaSeo/LinaCoin/rest"
)

func usage() {
	fmt.Printf("Please use the following commands\n\n")
	fmt.Printf("port: Sets the port of the server\n")
	fmt.Printf("mode: Choose between 'html'and 'rest'\n")
}

func Start() {
	if len(os.Args) < 2 {
		usage()
	}

	// flag.Int(parse starting point, default, flag notice)
	port := flag.Int("port", 4000, "Sets the port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html'and 'rest'")

	flag.Parse()
	switch *mode {
	case "explorer":
		fmt.Println("Start Explorer")
		explorer.Start(*port)
	case "rest":
		fmt.Println("Start REST API")
		rest.Start(*port)
	default:
		usage()
	}
}
