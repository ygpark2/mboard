module github.com/ygpark2/mboard

go 1.16

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.3.0
	github.com/evanphx/json-patch/v5 v5.2.0 // indirect
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/google/uuid v1.1.2
	github.com/google/wire v0.5.0
	github.com/gosimple/slug v1.9.0
	github.com/infobloxopen/atlas-app-toolkit v0.22.1
	github.com/infobloxopen/protoc-gen-gorm v0.20.1
	github.com/jinzhu/gorm v1.9.16
	github.com/lib/pq v1.8.0
	github.com/markbates/pkger v0.17.1
	github.com/micro/go-micro v1.18.0
	github.com/micro/micro/v3 v3.3.1-0.20210713153811-bb922eccdbd3
	github.com/micro/services v0.14.0
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.20.0
	github.com/sarulabs/di/v2 v2.4.0
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.7.0
	github.com/thoas/go-funk v0.7.0
	github.com/xmlking/configor v0.1.0
	github.com/xmlking/micro-starter-kit v0.3.7
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	google.golang.org/genproto v0.0.0-20210708141623-e76da96a951f
	google.golang.org/grpc v1.40.0-dev.0.20210623211556-d9eb12feed7a
	google.golang.org/grpc/examples v0.0.0-20210715165331-ce7bdf50abb1 // indirect
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	gorm.io/driver/mysql v1.1.1
	gorm.io/driver/postgres v1.1.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.12
)

// exclude github.com/ygpark2/mboard v0.0.0-20201103090146-2c7cb3e5d3fa

// replace github.com/ygpark2/mboard/shared => ./shared/
