
render: $(wildcard protos/*.proto) $(wildcard *.go) $(wildcard template/*)
	protoc -I. --plugin=protoc-gen-doc=./run.sh --doc_out=html/ protos/*.proto

test:
	go test . -v

build:
	go build -o build/protoc-gen-doc .

clean:
	find html/ -not -name html -not -name .gitkeep | xargs rm -rf
