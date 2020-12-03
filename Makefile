SOURCES := $(shell find . -name *.go)

SUPPORTED_ARCHITECTURE := linux_386 linux_amd64 linux_arm64 darwin_amd64 windows_386 windows_amd64
COMMAND := expose
TARGETS := $(addprefix out/$(COMMAND)-,$(SUPPORTED_ARCHITECTURE))

TARGET_ARCHOS ?= $(word 2,$(subst -, ,$@))
TARGET_OS ?= $(word 1,$(subst _, ,$(TARGET_ARCHOS)))
TARGET_ARCH ?= $(word 2,$(subst _, ,$(TARGET_ARCHOS)))


.PHONY: clean all fmt

# special

all: $(TARGETS)

clean:
	rm -fr out

fmt:
	go fmt ./...

# build binaries

$(TARGETS): $(SOURCES)
	GOOS=$(TARGET_OS) GOARCH=$(TARGET_ARCH) CGO_ENABLED=0 go build -o $@ cmd/$(COMMAND)/main.go
