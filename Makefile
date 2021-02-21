
render: $(wildcard protos/*.proto) $(wildcard *.go) $(wildcard template/*)
	protoc -I. --plugin=protoc-gen-doc=./run.sh --doc_out=html/ protos/*.proto

clean:
	find html/ -not -name html -not -regex html/static.* | xargs rm -rf

test:
	go test . -v

build:
	go build -o build/protoc-gen-doc .
