module github.com/nillocoelho/microservices/order

go 1.24.0

toolchain go1.24.11

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/nillocoelho/microservices-proto/golang/order v0.0.0-00010101000000-000000000000
	github.com/nillocoelho/microservices-proto/golang/payment v0.0.0-20251216175506-f7927d35c2b1
	github.com/nillocoelho/microservices-proto/golang/shipping v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.77.0
	gorm.io/driver/postgres v1.5.9
	gorm.io/gorm v1.31.1
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/crypto v0.43.0 // indirect
	golang.org/x/net v0.46.1-0.20251013234738-63d1a5100f82 // indirect
	golang.org/x/sync v0.17.0 // indirect
	golang.org/x/sys v0.37.0 // indirect
	golang.org/x/text v0.30.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251022142026-3a174f9686a8 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

replace github.com/nillocoelho/microservices-proto/golang/order => ../../microservices-proto/golang/order

replace github.com/nillocoelho/microservices-proto/golang/payment => ../../microservices-proto/golang/payment

replace github.com/nillocoelho/microservices-proto/golang/shipping => ../../microservices-proto/golang/shipping
