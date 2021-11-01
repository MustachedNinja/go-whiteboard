# Go-Whiteboard

## Description


## Getting it running:
Make sure you have golang installed
1. Clone the repo
2. Navigate to the `./src` folder
3. Execute `go run .`

## Deploying to appengine
1. Install [Cloud SDK](https://cloud.google.com/sdk/docs/install)
2. Obtain credentials: `gcloud auth login`
3. Set default project: `gcloud config set project go-whiteboard`
4. Verify settings (account, project): `gcloud config list`
5. To deploy the project cd into the folder with app.yaml and then 'gcloud app deploy'