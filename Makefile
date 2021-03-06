install-dep:
	@go get -u github.com/rakyll/statik
	@mkdir -p "webserver/www"
	@if [[ ! -e webserver/www/index.html ]]; then\
    	echo "BigBro - Need to do \"make build\" or similar build (for other platforms) before using GUI!" > webserver/www/index.html;\
	fi
	@cd webserver && statik -src=www -f 1>/dev/null
	@go get -u -v all
	@cd webserver/frontend && npm install
build:
	@echo -e "\e[96mBuilding for \e[95m${SUFFIX}\e[39m"
	@echo -e "\e[96mBuilding the \e[94muser interface\e[39m"
	@cd webserver/frontend && npm run build 1>/dev/null
	@echo -e "\e[96mPacking all \e[91massets\e[39m"
	@cd webserver && statik -src=www -f 1>/dev/null
	@mkdir -p bin
	@mkdir -p bin/config
	@echo -e "\e[96mBuilding \e[93mBigBro\e[39m"
	@go build -o bin/bigbro$(SUFFIX) cmd/main.go 1>/dev/null
	@cp -R config bin/
	@cd bin/config && rm config.yml && touch config.yml
	@echo -e "\e[96mPacking the \e[91mzip file\e[39m"
	@zip -u bin/bigbro$(SUFFIX).zip bin/bigbro$(SUFFIX) bin/config/* &>/dev/null
	@echo -e "\e[92mBuild Complete\e[39m"
build-embedded:
	@echo -e "\e[96mBuilding for \e[95m${GOARCH} (version ${GOARM})\e[39m"
	@echo -e "\e[96mBuilding the \e[94muser interface\e[39m"
	@cd webserver/frontend && npm run build 1>/dev/null
	@echo -e "\e[96mPacking all \e[91massets\e[39m"
	@cd webserver && statik -src=www -f 1>/dev/null
	@mkdir -p bin
	@mkdir -p bin/config
	@echo -e "\e[96mBuilding \e[93mBigBro\e[39m"
	@export GHW_DISABLE_WARNINGS=1
	@go build -o bin/bigbro_$(GOARCH)$(SUFFIX) cmd/main.go 1>/dev/null
	@cp -R config bin/
	@cd bin/config && rm config.yml && touch config.yml
	@echo -e "\e[96mPacking the \e[91mzip file\e[39m"
	@zip -u bin/bigbro_$(GOARCH)$(SUFFIX).zip bin/bigbro_$(GOARCH)$(SUFFIX) bin/config/* &>/dev/null
	@echo -e "\e[92mBuild Complete\e[39m"
build-x64:
	@make build GOOS=linux GOARCH=amd64 SUFFIX="_x64"
build-x86:
	@make build GOOS=linux GOARCH=386 SUFFIX="_x86"
build-arm:
	@make build-embedded GOOS=linux GOARCH=arm GOARM=7 SUFFIX="_7"
build-arm5:
	@make build-embedded GOOS=linux GOARCH=arm GOARM=5 SUFFIX="_5"
build-all:
	@make build-x64
	@make build-x86
	@make build-arm
	@make build-arm5
build-ui:
	@echo -e "\e[96mBuilding the \e[94muser interface\e[39m"
	@cd webserver/frontend && npm run build 1>/dev/null
	@echo -e "\e[96mPacking all \e[91massets\e[39m"
	@cd webserver && statik -src=www -f 1>/dev/null
	@echo -e "\e[92mUI Build Complete\e[39m"