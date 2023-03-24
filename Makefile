dev:
	go get github.com/cespare/reflex
	reflex -r '\.go$$' -s go run .