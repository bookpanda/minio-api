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
