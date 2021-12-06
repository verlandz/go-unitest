package generate

//go:generate mockgen --build_flags=--mod=mod -destination=mocks/Client.go -package=mocks . Client

type Client interface {
	GetSquareNumber(n int) (int, error)
}

type repository struct{}

func New() Client {
	return &repository{}
}
