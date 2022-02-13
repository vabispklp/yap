package shortener

import (
	"context"
	"fmt"
	"sync"
)

const workersCount = 10

func (s *Shortener) DeleteRedirectLinks(_ context.Context, ids []string, userID string) error {
	inputCh := make(chan string)

	go func() {
		for _, id := range ids {
			inputCh <- id
		}

		close(inputCh)
	}()

	chs := s.deleteFanOut(inputCh, userID)
	outCh := s.deleteFanIn(chs)

	s.scheduledDelete(outCh, userID)

	return nil
}
func (s *Shortener) scheduledDelete(ch chan string, userID string) {
	ctx := context.Background()
	deleteIDs := make([]string, 0)
	for id := range ch {
		deleteIDs = append(deleteIDs, id)
	}
	if err := s.storage.Delete(ctx, deleteIDs, userID); err != nil {
		fmt.Println(err)
	}
}

func (s *Shortener) deleteFanOut(inputCh <-chan string, userID string) []chan string {
	chs := make([]chan string, 0, workersCount)
	for i := 0; i < workersCount; i++ {
		ch := make(chan string)
		chs = append(chs, ch)
	}

	go func() {
		defer func(chs []chan string) {
			for _, ch := range chs {
				close(ch)
			}
		}(chs)

		for i := 0; ; i++ {
			if i == len(chs) {
				i = 0
			}

			id, ok := <-inputCh
			if !ok {
				return
			}

			shortenURL, err := s.storage.Get(context.Background(), id)
			if err != nil {
				continue
			}
			if shortenURL == nil || shortenURL.Deleted {
				continue
			}

			chs[i] <- id
		}
	}()

	return chs

}

func (s *Shortener) deleteFanIn(inputChs []chan string) chan string {
	outCh := make(chan string)

	go func() {
		wg := &sync.WaitGroup{}

		for _, inputCh := range inputChs {
			wg.Add(1)

			go func(inputCh chan string) {
				defer wg.Done()
				for id := range inputCh {
					outCh <- id
				}
			}(inputCh)
		}

		wg.Wait()
		close(outCh)
	}()

	return outCh
}
