package main

import (
	"io"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"../pool"
)

const (
	maxGoroutines   = 100
	pooledResources = 10
)

type dbConnection struct {
	ID int32
}

func (dbCon *dbConnection) Close() error {

	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {

	id := atomic.AddInt32(&idCounter, 1)
	return &dbConnection{
		id,
	}, nil
}

func main() {
	runtime.GOMAXPROCS(10)
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Fatal(err)
	}
	for query := 0; query < maxGoroutines; query++ {
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Microsecond)

		go func(query int) {
			performQueries(query, p)
			wg.Done()
		}(query)
	}
	wg.Wait()
	p.Close()

}

func performQueries(query int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer p.Release(conn)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
	log.Printf("QID:[%d]-----CID[%d]", query, conn.(*dbConnection).ID)

}
