package singleton

type Singleton interface {
	AddOne() int //this will help us to do something with the resource inside the instance of the structure
}
type singleton struct {
	count int //resource
}

var instance *singleton

func GetInstance() Singleton { // this will return the instance of the structure
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
