package obj_pool

import (
	"errors"
	"time"
)

type ReusableObj struct {

}
type ObjPool struct {
	bufChan chan *ReusableObj
}

func NewObjPoll(num int) *ObjPool {
	bufChan:=make(chan *ReusableObj,num)
	for i:=0;i<num;i++{
		bufChan<-new(ReusableObj)
	}
	return &ObjPool{
		bufChan: bufChan,
	}
}
func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj,error) {
	select {
	case ret:=<-p.bufChan:
		return ret,nil
	case <-time.After(timeout):
		return nil,errors.New("time out")
	default:
		return nil,errors.New("can not find")
	}
}

func (p *ObjPool) Put(obj *ReusableObj) error {
	select {
	case p.bufChan<-obj:
		return nil
	default:
		return errors.New("overflow")

	}
}