package worker

import (
	"github.com/lognitor/entrypoint/internal/worker/models"
	"sync"
	"time"
)

type Batch struct {
	debounce   time.Duration
	lastUpdate time.Time
	maxLength  int
	token      string
	batch      []models.Log
	mu         sync.Mutex
}

func CreateBatch(token string, maxDebounce time.Duration, maxLen int) (*Batch, error) {
	// TODO: Check exists token
	b := Batch{
		debounce:   maxDebounce,
		token:      token,
		maxLength:  maxLen,
		lastUpdate: time.Now(),
		batch:      make([]models.Log, 0, maxLen),
	}

	go b.work()

	return &b, nil
}

func (b *Batch) Add(log models.Log) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if len(b.batch) >= b.maxLength {
		b.flush()
	}

	b.batch = append(b.batch, log)
	b.lastUpdate = time.Now()
}

func (b *Batch) GetCount() int {
	return len(b.batch)
}

func (b *Batch) GetDuration() time.Duration {
	return b.debounce - time.Now().Sub(b.lastUpdate)
}

func (b *Batch) work() {
	// TODO: Add close worker from timeout
	//if b.GetDuration() > time.Minute*-15 {
	//
	//}

	for {
		if b.GetDuration() <= 0 {
			b.flush()
		}

		if b.GetCount() >= b.maxLength {
			b.flush()
		}

		time.Sleep(time.Microsecond * 100)
	}
}

func (b *Batch) flush() {
	if len(b.batch) == 0 {
		return
	}
	// TODO: Send models

	b.batch = make([]models.Log, 0, b.maxLength)
	b.lastUpdate = time.Now()
	return
}
