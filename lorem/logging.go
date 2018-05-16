package lorem

import (
	"github.com/rs/zerolog"
	"time"
	"context"
)

//struct passing the logger
type LoggingMiddleware struct {
	Logger zerolog.Logger
	Next   Service
}

//each method will have its own logger for app logs
func (mw LoggingMiddleware) Lorem(ctx context.Context, requestType string, min int, max int) (output string, err error) {
	defer func(begin time.Time) {
		mw.Logger.Info().Str(
			"method", "lorem").Str("requesttype", requestType).Int("min", min).Int("max",
			max).Str("output", output).Err(err).Dur("took",
			time.Since(begin)).Msg("")

	}(time.Now())
	output, err = mw.Next.Lorem(ctx, requestType, min, max)
	return
}
