package i

type Server interface {
	Route()
	Run(addr string) error
}
