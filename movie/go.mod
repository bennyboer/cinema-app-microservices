module github.com/ob-vss-ss19/blatt-4-sudo_blatt4/movie

go 1.12

replace github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation => ../presentation

replace github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation => ../reservation

require (
	github.com/golang/protobuf v1.3.1
	github.com/micro/go-micro v1.2.0
	github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation v0.0.0-20190619104653-e9d594d5f0ee
)
