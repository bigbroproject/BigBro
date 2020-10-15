
install-dep:
	go get -d -u -v all ./...
	cd webserver/frontend && npm install
build:
	cd webserver/frontend && npm run build
	cd webserver && statik -src=www -f
	mkdir -p bin
	go build -o bin/bigbro cmd/main.go
install:
	mkdir -p /home/${USER}/bigbro
	cp bin/bigbro /home/${USER}/bigbro/
	cp -R config /home/${USER}/bigbro/config
	chmod +x /home/${USER}//bigbro
	cd /home/${USER}/bigbro/config && rm config.yml && touch config.yml
	ln -s /home/${USER}/bigbro /usr/local/bin/bigbro

