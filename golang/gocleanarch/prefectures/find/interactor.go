package find

import (
	"sampleapp/datastore"
)

// Interactor implements FindUseCase
type Interactor struct {
	repository datastore.PrefectureRepository
}

// NewInteractor : constructor
func NewInteractor(repo datastore.PrefectureRepository) *Interactor {
	return &Interactor{repo}
}

// Handle :
func (it *Interactor) Handle(req *Request) (*Response, error) {
	prefs, err := it.repository.FindAll()
	if err != nil {
		return nil, err
	}

	prefectures := []ResponsePrefecture{}
	for _, p := range prefs {
		prefectures = append(prefectures, ResponsePrefecture{Name: p.Name, ID: p.ID})
	}
	return &Response{Prefectures: prefectures}, nil
}
