<h1> Depedency </h1>
<h2> proto </h2>

1. `go get github.com/golang/protobuf/{proto,protoc-gen-go}`

2. `go get github.com/micro/protoc-gen-micro`


<h1> IDE </h1>

<h2>VS Code</h2>

1. Install go extension
2. Install Go laugage server: 
Open settings. Search and enable `go.useLanguageServer`.
3. set up go language server by following https://github.com/microsoft/vscode-go#go-language-server

<h1> Generate Proto </h1>
1. `protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. proto/package/pause.proto`

<h1> Example Rules </h1>

1. A client sends a pause request to a server.
2. The server validates the request based on rules below:
  * membership is valid for pause.
  * only 1 pause is allowed.
  * pause period is between 14 days and 31 days.
3. If pause request is valid, save it.
4. The server returns success or failure.