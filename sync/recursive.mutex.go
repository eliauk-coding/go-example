package main

import (
	"fmt"
	"github.com/petermattis/goid"
	"sync"
	"sync/atomic"
)

// Two ways: goroutine id and token

type RecursiveMutex struct {
	sync.Mutex
	owner     int64 // The id of the goroutine currently holding the lock
	recursion int32 // The number of times this goroutine has reentered
}

func (m *RecursiveMutex) Lock() {
	gid := goid.Get()
	if atomic.LoadInt64(&m.owner) == gid {
		m.recursion++
		return
	}
	m.Mutex.Lock()
	atomic.StoreInt64(&m.owner, gid)
	m.recursion = 1
}

func (m *RecursiveMutex) Unlock() {
	gid := goid.Get()
	if atomic.LoadInt64(&m.owner) == gid {
		panic(fmt.Sprintf("wrong the owner(%d):%d!", m.owner, gid))
	}
	m.recursion--
	if m.recursion != 0 {
		return
	}
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}

type TokenRecursiveMutex struct {
	sync.Mutex
	token     int64
	recursion int32
}

func (m *TokenRecursiveMutex) Lock(token int64) {
	if atomic.LoadInt64(&m.token) == token {
		m.recursion++
		return
	}
	m.Mutex.Lock()
	atomic.StoreInt64(&m.token, token)
	m.recursion = 1
}

func (m *TokenRecursiveMutex) Unlock(token int64) {
	if atomic.LoadInt64(&m.token) == token {
		panic(fmt.Sprintf("wrong the token(%d):%d!", m.token, token))
	}
	m.recursion--
	if m.recursion != 0 {
		return
	}
	m.Mutex.Unlock()
	atomic.StoreInt64(&m.token, token)
}
