package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	m       sync.Mutex
	res     chan io.Closer
	factory func() (io.Closer, error)
	closed  bool
}

var ErrPoolClosed = errors.New("资源池已经被关闭")

//创建一个资源池
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size的值太小了。")
	}
	return &Pool{
		factory: fn,
		res:     make(chan io.Closer, size),
	}, nil
}

//从资源池里获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.res:
		log.Println("Acquire:共享资源")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:新生成资源")
		return p.factory()
	}
}
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}
	p.closed = true
	close(p.res)
	for r := range p.res {
		r.Close()
	}
}
func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		r.Close()
		return
	}
	select {
	case p.res <- r:
		log.Println("release")
	default:
		log.Print("release:Close")
		r.Close()
	}

}
