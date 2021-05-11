module github.com/ygpark2/mboard

go 1.16

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0 // indirect
	github.com/armon/consul-api v0.0.0-20180202201655-eb2c6b5be1b6 // indirect
	github.com/bufbuild/buf v0.41.0 // indirect
	github.com/coreos/go-etcd v2.0.0+incompatible // indirect
	github.com/cpuguy83/go-md2man v1.0.10 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.3.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.1.2
	github.com/google/wire v0.4.0
	github.com/gosimple/slug v1.9.0
	github.com/infobloxopen/atlas-app-toolkit v0.22.1
	github.com/infobloxopen/protoc-gen-gorm v0.20.1
	github.com/jinzhu/gorm v1.9.16
	github.com/jinzhu/now v1.1.1 // indirect
	github.com/lib/pq v1.8.0
	github.com/markbates/pkger v0.17.1
	github.com/micro/go-micro v1.18.0
	github.com/micro/micro/v3 v3.2.2-0.20210510150101-c2dad95e51af
	github.com/micro/services v0.14.0
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.20.0
	github.com/rvflash/goup v0.4.1 // indirect
	github.com/sarulabs/di/v2 v2.4.0
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.7.0
	github.com/thoas/go-funk v0.7.0
	github.com/uber/prototool v1.10.1-0.20200519182255-a6d064684c01 // indirect
	github.com/ugorji/go/codec v0.0.0-20181204163529-d75b2dcb6bc8 // indirect
	github.com/xmlking/configor v0.1.0
	github.com/xmlking/micro-starter-kit v0.3.7
	github.com/xordataexchange/crypt v0.0.3-0.20170626215501-b2862e3d0a77 // indirect
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	google.golang.org/genproto v0.0.0-20210426193834-eac7f76ac494
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
	sigs.k8s.io/kind v0.10.0 // indirect
	sigs.k8s.io/kustomize/kustomize/v3 v3.3.0 // indirect
	sigs.k8s.io/kustomize/v3 v3.1.1-0.20190821175718-4b67a6de1296 // indirect
)

// exclude github.com/ygpark2/mboard v0.0.0-20201103090146-2c7cb3e5d3fa

// replace github.com/ygpark2/mboard/shared => ./shared/
