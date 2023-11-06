# IPFS Upload Go 
This repository demonstrates how to upload files to an IPFS storage using Go and [Web3Storage](https://web3.storage/docs/). 
A similar thing was done using NodeJs and Typescript. The reasons for doing this again with Go are as follows: 
- When using Web3Storage to upload large files, there's a need to keep the file uploading in the background and send an immediate response back to the user. Goroutines allows this to be done more efficiently than when trying to implement this in Node. 
- Trying to demonstrate how to create APIs using Go and [mux](https://github.com/gorilla/mux). 
- Trying to demonstrate how to write Unit and e2e tests using Go. 
- How to Dockerize and package a go application. 


## Requirements 
- Go >= 1.18.*
- Docker 
- A [Web3Storage Token](https://web3.storage/docs/#get-an-api-token)

## Initial Setup 