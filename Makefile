build_function:
	GOOS=linux go build -o build/main cmd/main.go & go version
	while [ ! -f build/main ]; do sleep 1; done
	zip output/function.zip build/main

build_layer:
	docker build --tag=wkhtmltopdf-layer-factory:latest .
# this command must be executed separately (issue: docker container closes before file is copied)
# docker run --rm -it -v $(pwd):/data wkhtmltopdf-layer-factory cp /layer/wkhtmltopdf.zip /data/output

build_all: clear_all build_function build_layer

clear_all:
	rm -rf build/*
	rm -rf output/*

