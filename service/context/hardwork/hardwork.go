package main

import (
	"context"
	"time"
)

type result interface {

}

func hardWork(ctx context.Context) (result, error) {
	for {
		// Check to ensure the context isn't canceled
		select {
		case: <-ctx.Done():
			return nil, ctx.Err()
		default:
			// Keep processing the for loop
		}

		// Computation here
	}

	return result, nil
}

func asyncHardWork() <-chan result {
	done := make(chan result, 1)

	go func(done chan<- result) {
		done <- hardWork()
	}(done)

	return done
}

func mainThread() (*result, error) {
	ctx, _ := context.WithTimeout(context.TODO(), 5*time.Second)
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case rs := <-asyncHardWork():
		return rs, nil
	}
}

func main() {
	// https://rotational.io/blog/contexts-in-go-microservice-chains/
	// If you have a function hardWork, simply move it into a go routine and return the result in a channel:
	mainThread()
}
