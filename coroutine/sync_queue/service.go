package sync_queue

import "sync"

type Limiter struct {
	SLock    sync.Mutex
	LimitMap map[int64]bool
	LimitNum chan bool
}

/** Open|Close Used to restrict the execution of programs with the same ID,
only one program with the same ID can be executed at a time **/

func (p *Limiter) Open(id int64) bool {
	p.SLock.Lock()
	defer p.SLock.Unlock()
	if wait, ok := p.LimitMap[id]; ok {
		if wait {
			return false
		}
	}
	p.LimitMap[id] = true
	return true
}

func (p *Limiter) Close(id int64) {
	p.SLock.Lock()
	defer p.SLock.Unlock()
	p.LimitMap[id] = false
}

/** Add|Done  quantitative restriction **/

func (p *Limiter) Add() {
	p.LimitNum <- true
}

func (p *Limiter) Done() {
	<-p.LimitNum
}

func NewLimiter(limit int64) *Limiter {
	obj := new(Limiter)
	obj.LimitNum = make(chan bool, limit)
	obj.LimitMap = make(map[int64]bool)
	return obj
}
