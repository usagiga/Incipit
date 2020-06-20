package messages

type Response interface {
	GetHTTPStatusCode() int
}
