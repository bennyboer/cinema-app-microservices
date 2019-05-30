module github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation

go 1.12

replace github.com/ob-vss-ss19/blatt-4-sudo_blatt4/user => ../user

replace github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation => ../reservation

replace github.com/ob-vss-ss19/blatt-4-sudo_blatt4/movie => ../movie

require (
	github.com/golang/protobuf v1.3.1
	github.com/micro/go-micro v1.2.0
	github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation v0.0.0-00010101000000-000000000000
)
