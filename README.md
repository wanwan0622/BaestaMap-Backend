# BaestaMap Cloud Function

## Deploy & Execute

### local execute

```sh
cd cmd
go run .
curl http://127.0.0.1:8080
```

### Google Cloud Functions execute

In repository root directory and install gcloud from [here](https://cloud.google.com/sdk/docs/install?hl=ja)

* API Access Text

```sh
gcloud functions deploy hello --gen2 --runtime=go116 --region=us-central1 --source=. --entry-point=HelloCommand --trigger-http --allow-unauthenticated
curl https://hello-qpz6p6e7bq-uc.a.run.app
```

* Get Instagram Posts API (from Coordinates)

```sh
gcloud functions deploy baestamap --gen2 --runtime=go116 --region=us-central1 --source=. --entry-point=GcloudMain --trigger-http --allow-unauthenticated
curl -X POST -H "Content-Type: application/json" -d '{"lat":35.615304235976,"lng":139.7175761816}' https://baestamap-qpz6p6e7bq-uc.a.run.app
```

* Get Instagram Posts API (from Query)

```sh
gcloud functions deploy baestamap-query --gen2 --runtime=go116 --region=us-central1 --source=. --entry-point=GetPostFromQuery --trigger-http --allow-unauthenticated
curl -X POST -H "Content-Type: application/json" -d '{"query":"東京タワー"}' https://baestamap-query-qpz6p6e7bq-uc.a.run.app
```

* Get Location from Query API

```sh
gcloud functions deploy baestamap-location --gen2 --runtime=go116 --region=us-central1 --source=. --entry-point=GetLocationFromQuery --trigger-http --allow-unauthenticated
curl -X POST -H "Content-Type: application/json" -d '{"query":"東京タワー"}' https://baestamap-location-qpz6p6e7bq-uc.a.run.app
```
