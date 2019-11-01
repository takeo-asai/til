package datastore

// PrefectureRepository :
type PrefectureRepository interface {
	FindAll() ([]Prefecture, error)
}

// Prefecture :
type Prefecture struct {
	Name string
	ID   int8
}
