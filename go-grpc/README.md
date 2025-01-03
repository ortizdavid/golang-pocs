<h1>What is it?</h1>
This is an ultra-simplified example of a gRPC client and server written in Go.

<h1>What will I need to have to make it work?</h1>
You will need to install the protobuf compiler and the go language follow these commands if you are in a Linux environment:

`sudo apt install protobuf-compiler `

`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

Export the PATH, so we can run the commands on the terminal

`export PATH="$PATH:$(go env GOPATH)/bin"`

For the last update your terminal

`source ~/.bashrc`

<h1>How can I run the project?</h1>

First, compile the protoc:

`protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb`

Now you are good to run the server and the client:

`go run cmd/client/client.go`

`go run cmd/server/server.go`

![image](https://user-images.githubusercontent.com/33813203/169661855-76d034ef-ac66-4389-95ea-3879a4a6ca6b.png)
