package promise

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func TestStringOrError_Execute(t *testing.T) {
	promise := &Promise{}
	t.Run("Success result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		go timeout(t, &wg)

		promise.OnSuccess(func(s string) {
			t.Log(s)
			wg.Done()
		}).OnError(func(_ error) {
			t.Fail()
			wg.Done()
		}).Execute(func() (string, error) {
			return "Hello World!", nil
		})

		wg.Wait()
	})

	t.Run("Error result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		go timeout(t, &wg)

		promise.OnSuccess(func(_ string) {
			t.Fail()
			wg.Done()
		}).OnError(func(e error) {
			t.Log(e.Error())

			wg.Done()
		}).Execute(func() (string, error) {
			return "", errors.New("Error ocurred")
		})
		wg.Wait()
	})
}

// Add timeout to tests
func timeout(t *testing.T, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	t.Log("Timeout!")
	t.Fail()
	wg.Done()
}
