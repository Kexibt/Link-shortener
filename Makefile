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
<<<<<<< HEAD
	go run .
=======
	go run .
>>>>>>> e5e1916bc9f57b2dcdc0db0b81e29c76684ed804
