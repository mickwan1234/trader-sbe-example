package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/trader-sbe/trader"
)

func main() {
	http.HandleFunc("/data", handleSBEData)
	fmt.Println("Listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handleSBEData(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	var header trader.MessageHeader
	traderOrder := trader.TradeOrder{}
	bodyReader := bytes.NewReader(body)

	m := trader.NewSbeGoMarshaller()
	if err := header.Decode(m, bodyReader, header.Version); err != nil {
		http.Error(w, "Error decoding header", http.StatusInternalServerError)
		return
	}

	if err := traderOrder.Decode(m, bodyReader, traderOrder.SbeSchemaVersion(), traderOrder.SbeBlockLength(), false); err != nil {
		http.Error(w, "Error decoding body", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Received order: %+v\n", traderOrder)
}
