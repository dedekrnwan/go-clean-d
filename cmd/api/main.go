package main

import (
	"context"
	"fmt"
	"go-clean/internal/contract"
	"go-clean/internal/deliveries/api"
	"os"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/super-saga/go-x/ctxval"
	"github.com/super-saga/go-x/graceful"
)

func main() {

	var (
		ctx, cancel = context.WithCancel(context.Background())
		starters    []graceful.ProcessStarter
		stoppers    []graceful.ProcessStopper
		uid         = uuid.New()
		TZ          = "Asia/Jakarta"
	)
	validator.New()

	ctx = ctxval.Sets(ctx, ctxval.SetCorrelationId(uid.String()))
	os.Setenv("TZ", TZ)

	dep := contract.NewDependency()

	//api
	api := api.New(dep)
	starterApi, stopperApi := api.Prepare()
	starters = append(starters, starterApi)
	stoppers = append(stoppers, stopperApi)

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		graceful.StartProcessAtBackground(starters...)
		graceful.StopProcessAtBackground(time.Second*30, stoppers...)
		cancel()
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(ctx, "services %s stopped!", "go-clean")
}
