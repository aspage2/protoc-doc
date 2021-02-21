
render: $(wildcard protos/*.proto) $(wildcard *.go) $(wildcard template/*)
	protoc -I. --plugin=protoc-gen-doc=./run.sh --doc_out=html/ protos/*.proto

clean:
	find html/ -name *.html | xargs rm
