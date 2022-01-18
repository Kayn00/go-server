BASEDIR      =  $(shell pwd)
GOPATH       =  $(shell go env GOPATH)

gowatch:
	@if [ ! -f "$(GOPATH)/bin/gowatch" ];then go get github.com/silenceper/gowatch; fi

dev:export GIN_MODE=debug
dev:export CONFIG_RUNMODE=dev
dev:export CONFIG_ROOTPATH=$(BASEDIR)
dev:
	gowatch

# run swag init
swagger:
	@if [ ! -f "$(GOPATH)/bin/swag" ];then go get github.com/swaggo/swag/cmd/swag; fi
	@swag init