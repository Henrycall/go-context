package main

import (
	"context"
	"fmt"
	"time"
)

func main(){

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func ()  {
		time.Sleep(time.Second * 3)
		cancel()	
	}()
	bookHotel(ctx)
	
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done(): 
		fmt.Println("tempo execido para bookar o quarto")

	case <-time.After(time.Second * 5):
				fmt.Println("Quarto Reservado com sucesso")

	}
}