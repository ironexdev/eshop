package main

import (
	"e-shop-fiber/config"
	"e-shop-fiber/internal/worker"
)

func main() {
	cfg := config.Load()
	wrk := worker.New(cfg)
	wrk.Start()
}
