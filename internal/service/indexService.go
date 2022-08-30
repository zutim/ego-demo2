package service

import "sync"

type Index struct {

}

var index *Index
var once sync.Once
func GetInstanceIndex() *Index{
	once.Do(func() {
		index = &Index{}
	})
	return index
}

func (i *Index) GetAll() {

}
