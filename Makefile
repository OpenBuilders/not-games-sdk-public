.PHONY: gen_proto
gen_proto:
	buf generate

install_buf:
	sudo curl -sSL \
        "https://github.com/bufbuild/buf/releases/download/v1.50.0/buf-$(uname -s)-$(uname -m)" \
        -o "/usr/local/bin/buf" && \
		sudo chmod +x "/usr/local/bin/buf"
	if [[ "$(shell uname -s)" = "Linux" ]]; then \
		sudo apt-get install -y protobuf-compiler; \
	elif [[ "$(shell uname -s)" = "Darwin" ]]; then \
		brew install protobuf; \
	else \
		echo "Unsupported OS. Please install protobuf manually."; \
	fi
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	echo "$(go env GOPATH)/bin" >> $GITHUB_PATH
	sudo npm install -g protoc || true
	sudo npm install -g protoc-gen-grpc-web protoc-gen-js grpc-tools @protobuf-ts/plugin
	sudo rm -rf node_modules

install_openapi_generator:
	npm install -g @openapitools/openapi-generator-cli

http_go_generate:
	openapi-generator-cli generate \
                -i docs/swagger.yaml \
                -g go \
                -o ./gen/http/go \
                --additional-properties=packageName=not-games-sdk-public
http_js_generate:
	openapi-generator-cli generate \
                -i docs/swagger.yaml \
                -g typescript-axios \
                -o ./gen/http/nodejs \
                --additional-properties=supportsES6=true,npmName=not-games-sdk-public