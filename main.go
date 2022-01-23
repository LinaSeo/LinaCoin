package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/LinaSeo/LinaCoin/blockchain"
)

const (
	port        string = ":4000"
	templateDir        = "templates/"
)

// set templates variable to load all templates before execute
var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	// execute various templates
	templates.ExecuteTemplate(rw, "home", data)
}

func main() {
	// template.Must() : handling error
	// template.ParseGlob() : loading more than one file
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	// Do use templates object instead template
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", home)
	fmt.Printf("Listending on httl://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
