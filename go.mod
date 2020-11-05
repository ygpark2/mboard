module github.com/ygpark2/mboard

go 1.15

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.4.1
	github.com/golang/protobuf v1.4.3
	github.com/gosimple/slug v1.9.0
	github.com/grpc-ecosystem/grpc-gateway v1.9.5 // indirect
	github.com/infobloxopen/atlas-app-toolkit v0.22.1
	github.com/infobloxopen/protoc-gen-gorm v0.20.0
	github.com/jinzhu/gorm v1.9.16
	github.com/jinzhu/now v1.1.1 // indirect
	github.com/lib/pq v1.8.0 // indirect
	github.com/markbates/pkger v0.17.1
	github.com/mattn/go-sqlite3 v2.0.1+incompatible // indirect
	github.com/micro/micro/v3 v3.0.0-beta.7
	github.com/micro/services v0.14.0
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.20.0
	github.com/sarulabs/di/v2 v2.4.0
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.6.1
	github.com/thoas/go-funk v0.7.0
	github.com/xmlking/micro-starter-kit v0.3.7
	github.com/ygpark2/mboard/shared v0.0.0-20201104020547-99881663cfa3
	// github.com/ygpark2/mboard/shared v0.0.0-20201104020547-99881663cfa3
	// github.com/ygpark2/mboard/shared v0.0.1
	// github.com/ygpark2/mboard/shared v0.0.0-20201105013738-aa04de279772
	golang.org/x/net v0.0.0-20201010224723-4f7140c49acb // indirect
	google.golang.org/genproto v0.0.0-20201001141541-efaab9d3c4f7
	google.golang.org/grpc v1.33.1
	google.golang.org/grpc/examples v0.0.0-20201103233412-9a3c6618eeee // indirect
	google.golang.org/protobuf v1.25.0
)

// exclude github.com/ygpark2/mboard v0.0.0-20201103090146-2c7cb3e5d3fa

// replace github.com/ygpark2/mboard/shared => ./shared/
