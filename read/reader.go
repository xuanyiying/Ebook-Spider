package read

type Reader interface {
	Read(id int) (any, error)

	ReadAll() ([]any, error)
}
