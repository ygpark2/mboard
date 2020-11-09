
# linter and tool for proto files
# (if you use brew to install buf, skip next line)
GO111MODULE=on go get github.com/bufbuild/buf/cmd/buf
# prototool make it eazy to use protoc plugins
GO111MODULE=on go get github.com/uber/prototool/cmd/prototool@dev
# kind - kubernetes in docker (optional)
GO111MODULE=on go get sigs.k8s.io/kind
# go lang  build/publish/deploy tool (optional)
GO111MODULE=off go get github.com/google/ko/cmd/ko
# other way to get latest kustomize
GO111MODULE=on go get sigs.k8s.io/kustomize/kustomize/v3@v3.3.0
# pkger cli
go get github.com/markbates/pkger/cmd/pkger

# fetch protoc plugins into $GOPATH
GO111MODULE=off go get github.com/golang/protobuf/protoc
GO111MODULE=off go get github.com/golang/protobuf/protoc-gen-go

GO111MODULE=on go get github.com/micro/micro/v3/cmd/protoc-gen-micro@master

GO111MODULE=on go get github.com/gogo/protobuf/protoc-gen-gofast
GO111MODULE=on go get github.com/gogo/protobuf/protoc-gen-gogofast
GO111MODULE=on go get github.com/gogo/protobuf/protoc-gen-gogofaster
GO111MODULE=on go get github.com/gogo/protobuf/protoc-gen-gogoslick

GO111MODULE=off go get -u github.com/envoyproxy/protoc-gen-validate
GO111MODULE=off go get -u github.com/infobloxopen/protoc-gen-gorm

# (or) get latest 
go get github.com/golang/protobuf/protoc-gen-go

# goup checks if there are any updates for imports in your module.
# the main purpose is using it as a linter in continuous integration or in development process.
# Usage: goup -v -m ./...
GO111MODULE=on go get github.com/rvflash/goup
