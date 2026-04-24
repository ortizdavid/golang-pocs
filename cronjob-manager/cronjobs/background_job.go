package cronjobs
import "context"

type BackgroundJob interface {
	Execute(ctx context.Context) error  // execute the job
    Schedule() string //5m, 5s, 2h
	Name() string // JOb name - for logging pruposes
}