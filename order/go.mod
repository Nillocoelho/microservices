module github.com/nillocoelho/microservices/order

go 1.23

require (
	github.com/nillocoelho/microservices-proto/golang/order v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.64.0
	gorm.io/driver/mysql v1.5.6
	gorm.io/gorm v1.25.7
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

replace github.com/nillocoelho/microservices-proto/golang/order => ../../microservices-proto/golang/order
