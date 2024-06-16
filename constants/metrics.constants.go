package constants

type Domain string

const (
	HealthCheck Domain = "health_check"
	Metrics     Domain = "metrics"
	File        Domain = "file"
)

func (d Domain) String() string {
	return string(d)
}

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	DELETE Method = "DELETE"
)

func (m Method) String() string {
	return string(m)
}

// type StatusCode uint16

// const (
// 	OK   StatusCode = 200
// 	_400 StatusCode = 400
// 	_401 StatusCode = 401
// 	_404 StatusCode = 404
// 	_500 StatusCode = 500
// )

// func (s StatusCode) String() string {
// 	return s.String()
// }
