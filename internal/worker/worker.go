package worker

import (
	"e-shop-fiber/config"
	"e-shop-fiber/internal/worker/jobs"
)

type Worker struct {
	cfg         *config.Config
	emailJob    *jobs.EmailJob
	inventoryJob *jobs.InventoryJob
}

func New(cfg *config.Config) *Worker {
	return &Worker{
		cfg:         cfg,
		emailJob:    jobs.NewEmailJob(),
		inventoryJob: jobs.NewInventoryJob(),
	}
}

func (w *Worker) Start() {
	// Set up your worker to process jobs, e.g., using a message queue or a scheduler

	// Example: Process email and inventory update jobs periodically
	// You can use a package like https://github.com/robfig/cron for scheduling tasks
	// c := cron.New()
	// c.AddFunc("0 */1 * * * *", w.emailJob.SendEmail)
	// c.AddFunc("0 0 */1 * * *", w.inventoryJob.UpdateInventory)
	// c.Start()
}
