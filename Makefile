.PHONY: gen_proto
gen_proto:
	buf generate

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