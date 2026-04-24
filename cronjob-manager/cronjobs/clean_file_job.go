package cronjobs

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

type TempFileJob struct {
	path 		string
	logger 		*slog.Logger
}

func NewTempFileJob(path string, logger *slog.Logger) *TempFileJob {
	return &TempFileJob{
		path: path,
		logger: logger,
	}
}

func (j *TempFileJob) Name() string {
	return "TempFileJob"
}

func (j *TempFileJob) Schedule() string {
	return "10s"
}

func (j *TempFileJob) Execute(ctx context.Context) error {
    j.logger.Info("Cleaning up directory", "path", j.path)

    entries, err := os.ReadDir(j.path)
    if err != nil {
        return err
    }

    now := time.Now()
    for _, entry := range entries {
        info, _ := entry.Info()
        
        if now.Sub(info.ModTime()) > 20 * time.Second {
            _ = os.RemoveAll(filepath.Join(j.path, entry.Name()))
        }
    }
    return nil
}
