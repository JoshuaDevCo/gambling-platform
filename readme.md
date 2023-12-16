# gambling-platform TEST

I am looking forward to working with you.
I was not a senior. But I have passion. I will become a senior.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Addition](#addition)


## Installation

    Clone the repository.
   
    git clone https://github.com/trayanus1026/gambling-platform.git

    cd gambling-platform

    go install google.golang.org/protobuf/cmd/protoc-gen-go

    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative wallet/wallet.proto
    
    docker build -t aigesh-test .

    docker scout quickview

    docker run -p 8080:8080 aigesh-test

## Usage
    
    You can test using postman.

    (POST)http://localhost:8080/api/wallet/deposit
        request: {
                    "user_id":"3",
                    "amount":100
                }

        response: {
                    "balance": 100
                }

    (POST)http://localhost:8080/api/wallet/withdraw
        request: {
                    "user_id":"3",
                    "amount":40
                }
        response: {
                    "balance": 60
                }
        
    (GET)http://localhost:8080/api/wallet/balance/3
        response: {
                    "balance": 60,
                    "user_id": "3"
                }

    You can test websocket using by console log. I logged.

    You can test grpc using by command "go test" (modify user_id in 25 line of main_test.go )
    
## Addition

    