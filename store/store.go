package store

type Store interface {
	Store(data any) (bool, error)
}
