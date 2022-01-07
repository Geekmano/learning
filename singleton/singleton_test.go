package singleton

import "sync"

var singleInstance *Singleton
var once sync.Once

type Singleton struct {

}

func GetSingleton()*Singleton  {
	once.Do(func() {
		singleInstance=new(Singleton)
	})
	return singleInstance
}
