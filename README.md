# BaestaMap Cloud Function

```sh
# local execute
cd cmd
go run .
curl http://127.0.0.1:8080
# Google Cloud Functions execute
# In repository root directory
# Install gcloud -> https://cloud.google.com/sdk/docs/install?hl=ja
gcloud functions deploy baestamap --entry-point GcloudMain --runtime go116 --trigger-http
curl -X POST -H "Content-Type: application/json" -d '{"lat":35.615304235976,"lng":139.7175761816}' https://us-central1-baestamap.cloudfunctions.net/baestamap
```
