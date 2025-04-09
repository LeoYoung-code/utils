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

/*
UntilSuccess 会不断调用 runner，直到 runner 返回 true 或一个非 nil 的错误。

在每次重试 runner 之前的暂停时间，最初等于 minRetryPeriod，并且每次重试后暂停时间变为之前的两倍，但不会超过 maxRetryPeriod。

通常不应该在 runner 内部记录错误，你应该返回错误——从 runner 函数返回的非 nil 错误会带有堆栈信息和 runnerName 被记录。

如果 runner 恐慌（panic），panic值将被转换为错误并记录堆栈信息，这种情况下重试逻辑与 runner 返回 false 和错误的情况相同。

只有在 runner 返回成功且没有错误，或者 ctx 被取消时，UntilSuccess 才会返回。
*/
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
			switch v := rec.(type) {
			case error:
				err = errors.Wrap(errors.WithStack(v), fmt.Sprintf("Runner '%s' panicked", runnerName))
			default:
				err = errors.New(fmt.Sprintf("Runner '%s' panicked: %v", runnerName, v))
			}
		}
	}()

	return runner(ctx)
}

type incrementalTimer struct {
	initialPeriod time.Duration
	maxPeriod     time.Duration
	multiplier    int

	currentPeriod time.Duration
	iteration     int
}

func newIncrementalTimer(initialPeriod, maxPeriod time.Duration, multiplier int) *incrementalTimer {
	return &incrementalTimer{
		initialPeriod: initialPeriod,
		maxPeriod:     maxPeriod,
		multiplier:    multiplier,

		currentPeriod: initialPeriod,
		iteration:     0,
	}
}

func (t *incrementalTimer) next() <-chan time.Time {
	result := time.After(t.currentPeriod)

	t.currentPeriod = t.currentPeriod * time.Duration(t.multiplier)

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
