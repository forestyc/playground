package workpool

import (
	"context"
	"runtime"
	"sync/atomic"
	"time"
)

type Job func(ctx context.Context)
type Option func(wp *WorkPool)

type WorkPool struct {
	concurrency, size int
	queue             chan Job
	concurrencyCtrl   chan struct{}
	cancel            context.CancelFunc
	ctx               context.Context
	closed            atomic.Bool
	ref               atomic.Int32
}

// NewWorkPool new a work pool.
func NewWorkPool(option ...Option) *WorkPool {
	var wp WorkPool
	// default values
	wp.concurrency = runtime.GOMAXPROCS(-1)
	wp.size = 4 * wp.concurrency

	for _, o := range option {
		o(&wp)
	}

	wp.queue = make(chan Job, wp.size)
	wp.concurrencyCtrl = make(chan struct{}, wp.concurrency)
	wp.ref.Store(0)
	wp.closed.Store(false)
	wp.ctx, wp.cancel = context.WithCancel(context.Background())

	return &wp
}

// AddJob add a job
func (wp *WorkPool) AddJob(job Job) {
	if !wp.closed.Load() {
		wp.queue <- job
	}
}

// Start work pool
func (wp *WorkPool) Start() {
	go func() {
		for {
			select {
			case job := <-wp.queue:
				go wp.worker(job)
			case <-wp.ctx.Done():
				return
			}
		}
	}()
}

// Stop work pool
func (wp *WorkPool) Stop() {
	wp.closed.Store(true)

	// block until all jobs done
	wp.waitWorker()

	// exit
	if wp.cancel != nil {
		wp.cancel()
	}
	close(wp.queue)
	close(wp.concurrencyCtrl)
}

// WithConcurrency set concurrency
func WithConcurrency(concurrency int) Option {
	return func(wp *WorkPool) {
		wp.concurrency = concurrency
	}
}

// WithSize set size
func WithSize(size int) Option {
	return func(wp *WorkPool) {
		wp.size = size
	}
}

// worker goroutine
func (wp *WorkPool) worker(job Job) {
	wp.ref.Add(1)
	defer wp.ref.Add(-1)

	wp.concurrencyCtrl <- struct{}{}
	job(wp.ctx)
	<-wp.concurrencyCtrl
}

// block until all workers done
func (wp *WorkPool) waitWorker() {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			if wp.ref.Load() == 0 {
				return
			}
		}
	}
}
