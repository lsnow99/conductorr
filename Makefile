
default:
	go build ./cmd/conductorr

run: default
	./conductorr

migrations:
	go build ./cmd/migrations

web:
	cd conductorr-web && npm run build
	rm -rf static/
	mv conductorr-web/dist static/

web-dependencies:
	cd conductorr-web && npm install

web-update: web-dependencies web
