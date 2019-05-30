module github.com/ob-vss-ss19/blatt-4-sudo_blatt4/user

go 1.12

replace github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation => ../reservation

require (
	github.com/golang/protobuf v1.3.1
	github.com/micro/go-micro v1.2.0
	github.com/micro/protoc-gen-micro v0.8.0 // indirect
	github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation v0.0.0-20190530171822-2c90f3e94a08
)
