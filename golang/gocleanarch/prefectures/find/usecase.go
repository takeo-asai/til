package find

// UseCase :
type UseCase interface {
	Handle(req *Request) (*Response, error)
}

// Request :
type Request struct{}

// Response :
type Response struct {
	Prefectures []ResponsePrefecture
}

// ResponsePrefecture :
type ResponsePrefecture struct {
	Name string
	ID   int8
}
