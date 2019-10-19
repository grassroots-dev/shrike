OSFLAG 				:=
ifeq ($(OS),Windows_NT)
	OSFLAG += -D WIN32
	ifeq ($(PROCESSOR_ARCHITECTURE),AMD64)
		OSFLAG += -D AMD64
	endif
	ifeq ($(PROCESSOR_ARCHITECTURE),x86)
		OSFLAG += -D IA32
	endif
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		PROTOC_ZIP=protoc-3.10.0-rc-1-linux-x86_64.zip
		OSFLAG += -D LINUX
	endif
	ifeq ($(UNAME_S),Darwin)
		PROTOC_ZIP=protoc-3.10.0-rc-1-osx-x86_64.zip
		OSFLAG += -D OSX
	endif
		UNAME_P := $(shell uname -p)
	ifeq ($(UNAME_P),x86_64)
		OSFLAG += -D AMD64
	endif
		ifneq ($(filter %86,$(UNAME_P)),)
	OSFLAG += -D IA32
		endif
	ifneq ($(filter arm%,$(UNAME_P)),)
		OSFLAG += -D ARM
	endif
endif

deps:
	PROTOC_ZIP=$(PROTOC_ZIP) ./scripts/install-protoc.sh
	go get -u github.com/golang/protobuf/protoc-gen-go
	echo "Installing dependnencies.."
protogen:
	./scripts/protogen.sh
start:
	./scripts/start-docker.sh
startdb:
	docker-compose up -d db
dev: protogen startdb
	nodemon go run shrike.go
