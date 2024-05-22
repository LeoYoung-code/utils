package common

import (
	"context"
	"fmt"
	"hash/fnv"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

func Num2Version(innerVersion string) (outVersion string) {
	innr, err := strconv.Atoi(innerVersion)
	if err != nil {
		return innerVersion
	}
	if len(innerVersion) >= 7 {
		innr = innr / 100
	}
	big := innr / 10000
	innr = innr % 10000
	mid := innr / 100
	innr = innr % 100
	min := ""
	if innr > 0 {
		min = "." + strconv.Itoa(innr)
	}
	outVersion = strconv.Itoa(big) + "." + strconv.Itoa(mid) + min
	return
}

func GetSum32(data string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(data))
	return h.Sum32()
}

func UntilSuccess(
	ctx context.Context,
	runnerName string,
	runner func(context.Context) (bool, error),
	minRetryPeriod,
	maxRetryPeriod time.Duration) {

	success, err := runSafelyWithSuccess(ctx, runnerName, runner)
	if success && err == nil {
		// Brief success!
		return
	}

	incrementalTimer := newIncrementalTimer(minRetryPeriod, maxRetryPeriod, 2)

	for !success || err != nil {

		if err != nil {

		} else {

		}

		select {
		case <-ctx.Done():
			return
		case <-incrementalTimer.next():
			// To guarantee *no* work is being done after ctx was cancelled.
			if IsCancelled(ctx) {
				return
			}

			success, err = runSafelyWithSuccess(ctx, runnerName, runner)
		}
	}
}

// RunSafely handles panic using defer.
func runSafelyWithSuccess(ctx context.Context, runnerName string, runner func(context.Context) (bool, error)) (success bool, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			success = false
			err = errors.Wrap(errors.WithStack(rec.(error)), fmt.Sprintf("Runner '%s' panicked", runnerName))
		}
	}()

	return runner(ctx)
}

type incrementalTimer struct {
	initialPeriod time.Duration
	maxPeriod     time.Duration
	multiplier    time.Duration

	currentPeriod time.Duration
	iteration     int
}

func newIncrementalTimer(initialPeriod, maxPeriod time.Duration, multiplier int) *incrementalTimer {
	return &incrementalTimer{
		initialPeriod: initialPeriod,
		maxPeriod:     maxPeriod,
		multiplier:    time.Duration(multiplier),

		currentPeriod: initialPeriod,
		iteration:     0,
	}
}

func (t *incrementalTimer) next() <-chan time.Time {
	result := time.After(t.currentPeriod)

	t.currentPeriod = t.currentPeriod * t.multiplier

	if t.currentPeriod > t.maxPeriod {
		t.currentPeriod = t.maxPeriod
	}

	t.iteration += 1

	return result
}

func IsCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
