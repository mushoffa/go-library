package log_test

import (
	"testing"
	"time"

	"github.com/mushoffa/go-library/log"
	"go.uber.org/zap"
)

func TestNewZapLogger(t *testing.T) {
	l := log.NewZapLogger()
	logger := l.GetInstance().(*zap.Logger)
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "adfadfadf"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
