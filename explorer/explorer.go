package explorer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/LinaSeo/LinaCoin/blockchain"
)

const (
	templateDir string = "explorer/templates/"
)

// set templates variable to load all templates before execute
var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"Explorer Home", blockchain.GetBlockchain().AllBlocks()}
	// execute various templates
	templates.ExecuteTemplate(rw, "home", data)
}

func add(rw http.ResponseWriter, r *http.Request) {
	// r.method POST || GET
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}

}

var port string

func Start(aPort int) {
	port = fmt.Sprintf(":%d", aPort)
	handler := http.NewServeMux()
	// template.Must() : handling error
	// template.ParseGlob() : loading more than one file
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	// Do use templates object instead template
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	handler.HandleFunc("/", home)
	handler.HandleFunc("/add", add)
	fmt.Printf("Listening on httl://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
