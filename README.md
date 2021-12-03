# Go-Whiteboard

## Description


## Getting it running:
Make sure you have golang installed
1. Clone the repo
2. Navigate to the `./src` folder
3. Execute `go run .`

## Deploying to appengine
1. Install [Cloud SDK](https://cloud.google.com/sdk/docs/install)
2. Install AppEngine extension for Go: `gcloud components install app-engine-go`
3. Obtain credentials: `gcloud auth login`
4. Set default project: `gcloud config set project go-whiteboard`
5. Verify settings (account, project): `gcloud config list`
6. To deploy the project cd into the folder with app.yaml and then `gcloud app
deploy`
7. To verify deployed project:
`curl https://go-whiteboard.appspot.com/helloWorld`

## Installing protoc compiler for golang
### Install generic protoc compiler
1. [Download](https://github.com/protocolbuffers/protobuf/releases/tag/v3.18.1) `protoc` for your operating system *NOTE: you do not need to download protobuf*
2. Unzip the downloaded folder to your computer
3. Add the path to the `protoc/bin` to your PATH
### Install protoc-gen-go
1. `$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
2. Make sure your path includes your `go/bin` folder (`$PATH:[go_env_path]/bin`)
   
Check that `protoc` was installed by running `protoc --version`

## Writing and Compiling .proto files into .go
1. Create your proto file in a directory under `src`
2. Add the following line after your package declaration:
`option go_package = "src/{PATH_TO_YOUR_PROTO_PACKAGE}"`
1. Run the following command to generate your go module `protoc --go_out=. --go_opts=paths=source_relative {PATH_TO_YOUR_PROTO_FILE}`