.DEFAULT_GOAL := run

fmt:
	go fmt ./...

test:
	go test ./...

make-bin:
	GOOS=linux GOARCH=amd64 go build -o app app.go

make-leader:
	docker build -f Dockerfile-leader -t leader .

make-client:
	docker build -f Dockerfile-client -t client .

run: fmt test make-bin make-leader make-client
	docker run -v /var/run/docker.sock:/var/run/docker.sock -p 8080:8080 --rm --name leader leader /app
