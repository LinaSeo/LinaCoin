package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/LinaSeo/LinaCoin/blockchain"
	"github.com/LinaSeo/LinaCoin/blockchain/utils"
	"github.com/gorilla/mux"
)

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description,omitempty"`
	Payload     string `json:"payload,omitempty"`
}

type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type addBlockInfo struct {
	Data string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "Home",
		},
		{
			URL:         url("/blocks"),
			Method:      "GET",
			Description: "See All blocks",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add a block",
			Payload:     "data:string",
		},
	}
	// encode data struct to json format
	json.NewEncoder(rw).Encode(data)
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// send json data to rw
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlockInfo addBlockInfo
		utils.HanderErr(json.NewDecoder(r.Body).Decode(&addBlockInfo))
		fmt.Println(addBlockInfo.Data)
		blockchain.GetBlockchain().AddBlock(addBlockInfo.Data)
		rw.WriteHeader(http.StatusCreated)
	}
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["height"])
	utils.HanderErr(err)
	block, err := blockchain.GetBlockchain().GetBlock(id)
	encoder := json.NewEncoder(rw)
	if err == blockchain.ErrNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	} else {
		encoder.Encode(block)
	}
}

var port string

func Start(aPort int) {
	//multiplex
	router := mux.NewRouter()
	//handler := http.NewServeMux()
	router.Use(jsonContentTypeMiddleware)

	port = fmt.Sprintf(":%d", aPort)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{height:[0-9]+}", block).Methods("GET")

	fmt.Printf("http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
