package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
*
  - What we will be improving on from the previous version:
  - What if genGreeting wants to abandon the call to local after one second?
  - If printGreeting is unsuccessful, then printFarewell should also fail (can't say bye before hi!)
*/
func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printGreetingV2(ctx); err != nil {
			fmt.Printf("cannot print greeting: %v\n", err)
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewellV2(ctx); err != nil {
			fmt.Printf("cannot print farewell: %v\n", err)
			return
		}
	}()

	wg.Wait()
}

func printGreetingV2(ctx context.Context) error {
	greeting, err := genGreetingV2(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}

func printFarewellV2(ctx context.Context) error {
	farewell, err := genFarewellV2(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", farewell)
	return nil
}

func genGreetingV2(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	switch locale, err := localeV2(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func genFarewellV2(ctx context.Context) (string, error) {
	switch locale, err := localeV2(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "goodbye", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func localeV2(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(1 * time.Minute):
	}
	return "EN/US", nil
}
