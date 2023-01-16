
dev:
	go mod tidy && go run main.go

test:
	go test ./...

build:
	gcloud builds submit --tag gcr.io/stackpaper/penny-pincher

deploy:
	gcloud run deploy penny-pincher \
		--image gcr.io/stackpaper/penny-pincher \
		--platform managed

ship:
	make test && make build && make deploy