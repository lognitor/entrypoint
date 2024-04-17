package worker

import (
	"github.com/lognitor/entrypoint/internal/worker/models"
	"testing"
	"time"
)

func TestBatch_Add(t *testing.T) {
	batch, err := CreateBatch("test", time.Second, 20)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		for {
			t.Log(batch.GetDuration(), batch.GetCount())

			time.Sleep(time.Millisecond * 100)
		}
	}()

	for i := 0; i < 19; i++ {
		time.Sleep(time.Millisecond * 500)
		batch.Add(models.Log{})
	}

	if batch.GetCount() != 19 {
		t.Fatalf("batch count should be 19, but %d", batch.GetCount())
	}

	time.Sleep(time.Second * 2)

	for i := 0; i < 21; i++ {
		time.Sleep(time.Millisecond * 500)
		batch.Add(models.Log{})
	}

	if batch.GetCount() != 1 {
		t.Fatalf("batch count should be 0, but %d", batch.GetCount())
	}

}
