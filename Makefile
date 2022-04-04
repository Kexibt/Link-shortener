.PHONY: test
test:
	go test ./url
	go test ./db

.PHONY: deps
deps:
	ls go.mod || go mod init github.com/Kexibt/Link-shortener/
	go get github.com/lib/pq
	go mod tidy

.PHONY: run-db
run-db:
	go run . -db

.PHONY: run
run:
	go run .
