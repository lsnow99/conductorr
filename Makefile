
default:
	go build ./cmd/conductorr

run: default
	./conductorr
