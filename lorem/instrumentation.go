package lorem

import (
	"github.com/go-kit/kit/metrics"
	"time"
	"fmt"
	"context"
)

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
	Next           Service
}


func (mw InstrumentingMiddleware) Lorem(ctx context.Context, requestType string, min int, max int) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "lorem", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Lorem(ctx, requestType,min, max)
	return
}

