module github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation

go 1.12

replace github.com/ob-vss-ss19/blatt-4-sudo_blatt4/user => ../user

replace github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation => ../presentation

replace github.com/ob-vss-ss19/blatt-4-sudo_blatt4/cinema => ../cinema

require (
	github.com/armon/circbuf v0.0.0-20190214190532-5111143e8da2 // indirect
	github.com/golang/protobuf v1.3.1
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/micro/go-micro v1.2.0
	github.com/ob-vss-ss19/blatt-4-sudo_blatt4/cinema v0.0.0-20190618105109-2f19ae34c4b4
	github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation v0.0.0-20190616115933-e047b0ef1f44
	github.com/ob-vss-ss19/blatt-4-sudo_blatt4/user v0.0.0-20190616115933-e047b0ef1f44
	github.com/onsi/gomega v1.5.0 // indirect
)
