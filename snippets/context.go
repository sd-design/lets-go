package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/context"
)

func main() {
	duration := 690 * time.Millisecond
	ctx := context.Background()
	// d := time.Now().Add(duration)
	// ctx, cancel := context.WithCancel(ctx)
	ctx = context.WithValue(ctx, "test", "hello")
	ctx, cancel := context.WithTimeout(ctx, duration)

	defer cancel()

	// go func() {
	// 	err := cancelRequest(ctx)
	// 	if err != nil {
	// 		cancel()
	// 	}
	// }()

	doRequest(ctx, "https://ya.ru")
}

func cancelRequest(ctx context.Context) error {
	time.Sleep(1000 * time.Millisecond)
	return fmt.Errorf("fail request")

}

func doRequest(ctx context.Context, RequestStr string) {
	req, _ := http.NewRequest(http.MethodGet, RequestStr, nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Printf("Response completed, status code=%d\n", res.StatusCode)
		fmt.Println(ctx.Value("test"))
	case <-ctx.Done():
		fmt.Println("Request not good")
	}
}
