
render: $(wildcard protos/*.proto) $(wildcard *.go) $(wildcard template/*)
	@mkdir html || true
	protoc -I. --plugin=protoc-gen-doc=./run.sh --doc_out=html/ protos/*.proto

eh: $(wildcard protos/*.proto) $(wildcard *.go) $(wildcard template/*)
	protoc -I. --plugin=protoc-gen-doc=./run_test.sh --doc_out=. protos/*.proto

test:
	go test . -v

build:
	go build -o build/protoc-gen-doc .

clean:
	rm -rf build html
