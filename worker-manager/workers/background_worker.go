package workers

import "context"

type BackgroundWorker interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Name() string
}