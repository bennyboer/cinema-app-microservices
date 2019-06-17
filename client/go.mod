module github.com/ob-vss-ss19/blatt-4-sudo_blatt4/client

go 1.12

replace github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation => ../reservation

replace github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation => ../presentation

replace github.com/ob-vss-ss19/blatt-4-sudo_blatt4/cinema => ../cinema

replace github.com/ob-vss-ss19/blatt-4-sudo_blatt4/movie => ../movie

replace github.com/ob-vss-ss19/blatt-4-sudo_blatt4/user => ../user

require (
	github.com/micro/go-micro v1.2.0
	github.com/ob-vss-ss19/blatt-4-sudo_blatt4/cinema v0.0.0-20190616115933-e047b0ef1f44
	github.com/ob-vss-ss19/blatt-4-sudo_blatt4/movie v0.0.0-20190616115933-e047b0ef1f44
	github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation v0.0.0-20190616115933-e047b0ef1f44
	github.com/ob-vss-ss19/blatt-4-sudo_blatt4/reservation v0.0.0-20190616115933-e047b0ef1f44
	github.com/ob-vss-ss19/blatt-4-sudo_blatt4/user v0.0.0-20190616115933-e047b0ef1f44
)
