package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/LinaSeo/LinaCoin/blockchain"
	"github.com/LinaSeo/LinaCoin/blockchain/utils"
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
	// send json data to rw
	rw.Header().Add("Content-Type", "application/json")
	// encode data struct to json format
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlockInfo addBlockInfo
		utils.HanderErr(json.NewDecoder(r.Body).Decode(&addBlockInfo))
		fmt.Println(addBlockInfo.Data)
		blockchain.GetBlockchain().AddBlock(addBlockInfo.Data)
		rw.WriteHeader(http.StatusCreated)
	}
}

var port string

func Start(aPort int) {
	handler := http.NewServeMux()
	port = fmt.Sprintf(":%d", aPort)
	handler.HandleFunc("/", documentation)
	handler.HandleFunc("/blocks", blocks)
	fmt.Printf("http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
