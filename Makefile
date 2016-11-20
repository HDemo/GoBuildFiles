BINARY = DemoF/Demo

GO_FLAGS = 
SOURCE_DIR = ./DemoF

all: build

build: deps
	go build $(GO_FLAGS) -o $(BINARY) $(SOURCE_DIR)

deps:
	go get $(GO_FLAGS) -d $(SOURCE_DIR)

clean:
	go clean -i $(GO_FLAGS) $(SOURCE_DIR)
	rm -rf $(BINARY)
