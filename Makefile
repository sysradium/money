protobuf:
	@echo "Making proto/src/go (GoLang package) ..."
	@mkdir -p proto/src/go
	@protoc -I . -I=${GOPATH}/src \
	--micro_out=. \
	--gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types:. \
	proto/server.proto
