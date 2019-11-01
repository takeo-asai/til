package datastore

// PrefectureRepositoryImpl :
type PrefectureRepositoryImpl struct{}

// NewPrefectureRepositoryImpl :
func NewPrefectureRepositoryImpl() *PrefectureRepositoryImpl {
	return &PrefectureRepositoryImpl{}
}

// FindAll :
func (repo *PrefectureRepositoryImpl) FindAll() ([]Prefecture, error) {
	prefs := []Prefecture{}
	prefs = append(prefs, Prefecture{Name: "北海道", ID: 1})
	prefs = append(prefs, Prefecture{Name: "青森県", ID: 2})
	prefs = append(prefs, Prefecture{Name: "岩手県", ID: 3})
	prefs = append(prefs, Prefecture{Name: "宮城県", ID: 4})
	prefs = append(prefs, Prefecture{Name: "秋田県", ID: 5})
	prefs = append(prefs, Prefecture{Name: "山形県", ID: 6})
	prefs = append(prefs, Prefecture{Name: "福島県", ID: 7})
	prefs = append(prefs, Prefecture{Name: "茨城県", ID: 8})
	prefs = append(prefs, Prefecture{Name: "栃木県", ID: 9})
	prefs = append(prefs, Prefecture{Name: "群馬県", ID: 10})
	prefs = append(prefs, Prefecture{Name: "埼玉県", ID: 11})
	prefs = append(prefs, Prefecture{Name: "千葉県", ID: 12})
	prefs = append(prefs, Prefecture{Name: "Tokyo", ID: 13})
	return prefs, nil
}