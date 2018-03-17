package Pool
import (
	"io"
	"log"
	"sync"
	"errors"
)
type Pool struct {
	m sync.Mutex
	resource chan io.Closer
	factory func() (io.Closer,error)
	close bool
}	
var poolCloseError = errors.New("The Pool has already been closed!")
func New(f func()(io.Closer,error),size uint8) (*Pool,error) {
	if size < 0 {
		return nil,errors.New("The size of Pool is negative!")
	}
	
	
	return &Pool{
		factory : f,
		resource:make(chan io.Closer,size),

	},nil
}
func (p *Pool) Acquire() (io.Closer,error) {
	if p.close {
		return nil,poolCloseError
	}
	select {
	case r,ok := <- p.resource: {
		log.Println("Acquire:","shared resource")
		if !ok {
			return nil,poolCloseError
		} 
		return r,nil
	}
	default: {
		return p.factory()
	}
	}

}
func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()
	if p.close  {
		r.Close()
	}
	select {
	case p.resource <- r: {
		log.Println("Release","in queue.")
	}
	default: {
		log.Println("Release","closing")
		r.Close()
	}
	}
}
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	if p.close {
		return
	}
	p.close = true
	close(p.resource)
	for r := range p.resource {
		r.Close()
	}
}