
default:
	go build ./cmd/conductorr

run: default
	./conductorr

migrations:
	go build ./cmd/migrations