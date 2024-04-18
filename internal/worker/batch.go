package worker

import (
	"github.com/lognitor/entrypoint/pkg/structs"
	"sync"
	"time"
)

type Batch struct {
	debounce   time.Duration
	lastUpdate time.Time
	maxLength  int
	token      string
	batch      []structs.Log
	mu         sync.Mutex
	out        chan<- []structs.Log
}

func CreateBatch(token string, debounce time.Duration, maxLen int) (*Batch, chan []structs.Log, error) {
	// TODO: Check exists token
	ch := make(chan []structs.Log)

	b := Batch{
		debounce:   debounce,
		token:      token,
		maxLength:  maxLen,
		lastUpdate: time.Now(),
		out:        ch,
		batch:      make([]structs.Log, 0, maxLen),
	}

	go b.work()

	return &b, ch, nil
}

func (b *Batch) Add(log structs.Log) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if len(b.batch) >= b.maxLength {
		b.flush() //TODO
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

// TODO: Dev
func (b *Batch) Close() {
	close(b.out)
}
func (b *Batch) work() {
	// TODO: Add close worker from timeout
	//if b.GetDuration() > time.Minute*-15 {
	//
	//}

	for {
		if b.GetDuration() <= 0 {
			b.mu.Lock()
			b.flush() //TODO:
			b.mu.Unlock()
		}

		if b.GetCount() >= b.maxLength {
			b.mu.Lock()
			b.flush() //TODO:
			b.mu.Unlock()
		}

		time.Sleep(time.Millisecond * 100)
	}
}

func (b *Batch) flush() error {
	if len(b.batch) == 0 {
		return nil
	}

	b.out <- b.batch

	b.batch = make([]structs.Log, 0, b.maxLength)
	b.lastUpdate = time.Now()
	return nil
}
