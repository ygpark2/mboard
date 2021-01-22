module github.com/ygpark2/mboard

go 1.15

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.3.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.3 // indirect
	github.com/google/uuid v1.1.2
	github.com/google/wire v0.4.0
	github.com/gosimple/slug v1.9.0
	github.com/infobloxopen/atlas-app-toolkit v0.22.1
	github.com/infobloxopen/protoc-gen-gorm v0.20.0
	github.com/jinzhu/gorm v1.9.16
	github.com/jinzhu/now v1.1.1 // indirect
	github.com/lib/pq v1.8.0
	github.com/markbates/pkger v0.17.1
	github.com/micro/go-micro v1.18.0
	github.com/micro/micro/v3 v3.0.3-0.20201203165704-226c32e1ea25
	github.com/micro/services v0.14.0
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.20.0
	github.com/sarulabs/di/v2 v2.4.0
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.6.1
	github.com/thoas/go-funk v0.7.0
	github.com/xmlking/configor v0.1.0
	github.com/xmlking/micro-starter-kit v0.3.7
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	golang.org/x/net v0.0.0-20201010224723-4f7140c49acb // indirect
	golang.org/x/sys v0.0.0-20201119102817-f84b799fce68 // indirect
	google.golang.org/genproto v0.0.0-20201119123407-9b1e624d6bc4
	google.golang.org/grpc v1.32.0
	google.golang.org/grpc/examples v0.0.0-20201106192519-9c2f82d9a79c // indirect
	google.golang.org/protobuf v1.25.1-0.20201020201750-d3470999428b
)

// exclude github.com/ygpark2/mboard v0.0.0-20201103090146-2c7cb3e5d3fa

// replace github.com/ygpark2/mboard/shared => ./shared/
