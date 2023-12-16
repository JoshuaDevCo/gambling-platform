

go install google.golang.org/protobuf/cmd/protoc-gen-go

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative wallet/wallet.proto

docker build -t aigesh-test .

docker scout quickview

//INFO New version 1.2.2 available (installed version is 0.20.0)
//[2023-12-16T09:50:15.802626500Z][docker-credential-desktop.system][W] Windows version might not be up-to-date: The system cannot find the file specified.
//level=error msg="Status: login using Docker Desktop or 'docker login' command: no credential found for \"index.docker.io\", Code: 1"

docker run -p 8080:8080 aigesh-test



# gambling-platform TEST

I am looking forward to working with you.
I was not a senior. But I have passion. I will become a senior.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Addition](#addition)


## Installation

    Clone the repository.
   
    git clone https://github.com/trayanus1026/Kristina-Test.git

    cd Kristina Test

    npm install

## Usage

    npm run start1 for validExpressionCheck

    npm run start2 for nonFibonacciGen

    npm run start3 for lruCache

    npm run start4 for weightCounter

    npm run start5 for findConcat

    npm run test for test

## Addition

    My complexiy of the problem "Divide and rule" is O(K*S).