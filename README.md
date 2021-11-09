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

## Writing and Compiling .proto files into .go
1. Create your go file in a directory under `src`
2. Add the following line after your package declaration:
`option go_package = "src/{PATH_TO_YOUR_PROTO_PACKAGE}"`
3. Run the following command one level above `src` (`ls` should print `README.md go.mod go.sum src`)
`protoc --go_out=. {PATH_TO_YOUR_PROTO_PACKAGE}`