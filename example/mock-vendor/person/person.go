// NOTE: we don't want to use DelName().
package person

//go:generate mockgen --build_flags=--mod=mod -destination=mocks/Client.go -package=mocks . Client

type Client interface {
	SetName(name string)
	GetName() string
}
