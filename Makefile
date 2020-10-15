
install-dep:
	go get -d -u -v all ./...
	cd webserver/frontend && npm install
build:
	cd webserver/frontend && npm run build
	cd webserver && statik -src=www -f
	mkdir -p bin
	mkdir -p bin/config
	go build -o bin/bigbro cmd/main.go
	cp -R config bin/
	cd bin/config && rm config.yml && touch config.yml
install:
	mkdir -p /home/${USER}/bigbro
	cp -R bin/ /home/${USER}/bigbro/
	cp -R config /home/${USER}/bigbro/config
	chmod +x /home/${USER}/bigbro
	#ln -s /home/${USER}/bigbro/bigbro /usr/local/bin/bigbro

