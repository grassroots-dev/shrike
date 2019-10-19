module github.com/grassroots-dev/shrike

go 1.13

require (
	github.com/go-pg/pg/v9 v9.0.0-beta.9
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/golang/protobuf v1.3.2
	github.com/grassroots-dev/shrike/api v0.0.0-20190915185703-018b8fc494d7
	github.com/grassroots-dev/shrike/service v0.0.0-20190915174156-3353fac14c5b
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/satori/go.uuid v1.2.0 // indirect
	google.golang.org/grpc v1.24.0
)

replace github.com/grassroots-dev/shrike/service => ./service

replace github.com/grassroots-dev/shrike/api => ./api
