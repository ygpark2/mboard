module github.com/ygpark2/mboard/service/account

go 1.15

require (
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/envoyproxy/protoc-gen-validate v0.4.1
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/infobloxopen/atlas-app-toolkit v0.22.1
	github.com/infobloxopen/protoc-gen-gorm v0.20.0
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/micro/v3 v3.0.0-beta.7
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.20.0
	github.com/sarulabs/di/v2 v2.4.0
	github.com/satori/go.uuid v1.2.0
	github.com/thoas/go-funk v0.7.0
	github.com/xmlking/micro-starter-kit v0.3.7
	github.com/ygpark2/mboard v0.0.0-20201030010300-0a58f2994974
	github.com/ygpark2/mboard/service/greeter v0.0.0-20201030010300-0a58f2994974
	// github.com/ygpark2/mboard/service/account v0.0.0-20201027090308-71883ec3a716 // indirect
	// github.com/ygpark2/mboard/service/greeter
	google.golang.org/genproto v0.0.0-20201028140639-c77dae4b0522
	google.golang.org/grpc v1.33.1 // indirect
	google.golang.org/protobuf v1.25.0
	gorm.io/gorm v1.20.5 // indirect
)

// replace github.com/ygpark2/mboard => ../account
