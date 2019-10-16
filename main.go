package main

import (
	"golang.org/x/sync/errgroup"
	"log"
	"sync"
	"time"
)

func main() {
	wg()
	eg()
}

func wg() {
	log.Printf("starting wg#main")
	var wg sync.WaitGroup
	go func() {
		defer wg.Done()
		log.Printf("starting wg#1")
		time.Sleep(1 * time.Second)
		go func() {
			defer wg.Done()
			log.Printf("starting wg#2")
			time.Sleep(3 * time.Second)
			log.Printf("finished wg#2")
		}()
		wg.Add(1)
		time.Sleep(1 * time.Second)
		log.Printf("finished wg#1")
	}()
	wg.Add(1)
	wg.Wait()
	log.Printf("finished wg#main")
}

func eg() {
	log.Printf("starting eg#main")
	var eg errgroup.Group
	eg.Go(func() error {
		log.Printf("starting eg#1")
		time.Sleep(1 * time.Second)
		eg.Go(func() error {
			log.Printf("starting eg#2")
			time.Sleep(3 * time.Second)
			log.Printf("finished eg#2")
			return nil
		})
		time.Sleep(1 * time.Second)
		log.Printf("finished eg#1")
		return nil
	})
	if err := eg.Wait(); err != nil {
		log.Fatalf("error: %s", err)
	}
	log.Printf("finished eg#main")
}
