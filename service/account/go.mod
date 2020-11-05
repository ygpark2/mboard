module github.com/ygpark2/mboard/service/account

go 1.15

require (
	github.com/envoyproxy/protoc-gen-validate v0.4.1
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/infobloxopen/atlas-app-toolkit v0.22.1
	github.com/infobloxopen/protoc-gen-gorm v0.20.0
	github.com/jinzhu/gorm v1.9.16
	github.com/markbates/pkger v0.17.1
	github.com/micro/micro/v3 v3.0.0-beta.7
	github.com/miekg/dns v1.1.35 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.20.0
	github.com/sarulabs/di/v2 v2.4.0
	github.com/satori/go.uuid v1.2.0
	github.com/thoas/go-funk v0.7.0
	github.com/urfave/cli/v2 v2.3.0 // indirect
	github.com/xmlking/micro-starter-kit v0.3.7
	github.com/ygpark2/mboard/service/greeter v0.0.0-20201104020547-99881663cfa3
	github.com/ygpark2/mboard/shared v0.0.0-20201104020547-99881663cfa3 // indirect
	// github.com/ygpark2/mboard v0.0.0-20201030010300-0a58f2994974
	// github.com/ygpark2/mboard/service/greeter v0.0.0-20201103090146-2c7cb3e5d3fa
	// github.com/ygpark2/mboard/service/greeter v0.0.0-20201030051620-33d7831905d0
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	golang.org/x/net v0.0.0-20201031054903-ff519b6c9102 // indirect
	golang.org/x/sys v0.0.0-20201101102859-da207088b7d1 // indirect
	golang.org/x/text v0.3.4 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20201102152239-715cce707fb0
	google.golang.org/protobuf v1.25.0
)

// replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

// replace github.com/ygpark2/mboard => ../account
