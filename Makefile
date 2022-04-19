app:
	go mod tidy && \
	go run main.go

server:
	go mod tidy && \
	go run mock-website/*go