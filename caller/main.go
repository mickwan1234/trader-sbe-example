package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/trader-sbe/trader"
)

func main() {
	client := &http.Client{}

	buf := new(bytes.Buffer)
	m := trader.NewSbeGoMarshaller()

	data := trader.TradeOrder{
		OrderId:        1234567890,
		TradeDate:      20230401, // Format YYYYMMDD
		OrderQty:       100,
		Price:          150.75,
		OrderType:      trader.OrderType.STOP,
		OrderSide:      trader.OrderSide.BUY,
		Symbol:         []uint8("AAPL"),
		ExecutionVenue: []uint8("NASDAQ"),
	}
	header := trader.MessageHeader{
		BlockLength: data.SbeBlockLength(),
		TemplateId:  data.SbeTemplateId(),
		SchemaId:    data.SbeSchemaId(),
		Version:     data.SbeSchemaVersion(),
	}
	header.Encode(m, buf)

	if err := data.Encode(m, buf, false); err != nil {
		fmt.Println("Error encoding data: ", err)
		panic(err)
	}

	req, err := http.NewRequest("POST", "http://localhost:8081/data", buf)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		panic(err)
	}
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := client.Do(req)
	fmt.Printf("Sending order: %+v\n", data)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response Status: ", resp.Status)
}
