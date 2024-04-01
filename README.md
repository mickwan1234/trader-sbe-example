# SBE Encoder HTTP Server

This project is an example of creating an HTTP server that communicates using Simple Binary Encoding (SBE) instead of JSON. SBE is a high-performance encoding system for binary messages. This project is designed to demonstrate how to set up an HTTP server that can communicate using SBE-encoded messages.

## Project Structure

- `caller/main.go`: This file contains the main function for the caller service. It sends SBE-encoded messages to the server.
- `cmd/main.go`: This file contains the main function for the server. It listens for incoming SBE-encoded messages and processes them. It uses the `handleSBEData` function to handle incoming requests.
- `trader/`: This directory contains the SBE encoding logic.
  - `MessageHeader.go`: This file defines the structure of the message header in SBE format.
  - `OrderSide.go`: This file defines the OrderSide enumeration in SBE format.
  - `OrderType.go`: This file defines the OrderType enumeration in SBE format.
  - `SbeMarshalling.go`: This file contains functions for marshalling and unmarshalling SBE messages.
  - `TradeOrder.go`: This file defines the TradeOrder message in SBE format.
  - `VarAsciiEncoding.go`: This file contains functions for encoding and decoding ASCII strings in SBE format.

## Getting Started

1. Install Go. You can download it from the official website: https://golang.org/dl/. This project requires Go version 1.16 or later.
2. Clone this repository: `git clone https://github.com/mickwan1234/trader-sbe-example.git`
3. Navigate to the project directory: `cd trader-sbe-example`
4. Build the project: `go build`
5. Run the server: `go run cmd/main.go`
6. In a separate terminal, run the caller: `go run caller/main.go`

## Usage

To use the server, send a POST request to `http://localhost:8081/data` with an SBE-encoded TradeOrder message in the body. The server will respond with a status message. The server uses the `handleSBEData` function to read the request body, decode the header and the body of the SBE message, and print the received order.

## Contributing

Contributions are welcome! Please fork the repository and create a pull request with your changes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.