install-dep:
	go get -d -u -v all ./...
	cd webserver/frontend && npm install
build:
	cd webserver/frontend && npm run build
	cd ../../
