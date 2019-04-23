module github.com/synerex/synerex_alpha/api

require (
	github.com/golang/protobuf v1.2.0
	github.com/stretchr/testify v1.2.2
	github.com/synerex/synerex_alpha/api/adservice v0.0.0
	github.com/synerex/synerex_alpha/api/common v0.0.0-20181104051513-17262833074c
	github.com/synerex/synerex_alpha/api/fleet v0.0.0
	github.com/synerex/synerex_alpha/api/library v0.0.0
	github.com/synerex/synerex_alpha/api/ptransit v0.0.0
	github.com/synerex/synerex_alpha/api/rideshare v0.0.0
	github.com/synerex/synerex_alpha/api/routing v0.0.0
	//	github.com/synerex/synerex_alpha/api/agv v0.0.0
	//	C:/Users/sakakibara/go/project/synerex_alpha/api/agv v0.0.0
	golang.org/x/net v0.0.0-20181102091132-c10e9556a7bc
	google.golang.org/grpc v1.16.0
)

replace (
	github.com/synerex/synerex_alpha/api/adservice => ./adservice
	github.com/synerex/synerex_alpha/api/fleet => ./fleet
	github.com/synerex/synerex_alpha/api/library => ./library
	github.com/synerex/synerex_alpha/api/marketing => ./marketing
	github.com/synerex/synerex_alpha/api/ptransit => ./ptransit
	github.com/synerex/synerex_alpha/api/rideshare => ./rideshare
	github.com/synerex/synerex_alpha/api/routing => ./routing
	github.com/synerex/synerex_alpha/api/agv => ./agv
//	C:/Users/sakakibara/go/project/synerex_alpha/api/agv => ./agv
)
