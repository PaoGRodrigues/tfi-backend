package host

type HostBlocker interface {
	Block(string) (*string, error)
}
