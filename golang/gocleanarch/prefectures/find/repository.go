package find

// Repository :
type Repository interface {
	FindAll() ([]Prefecture, error)
}

// Prefecture :
type Prefecture struct {
	Name string
	ID   int8
}
