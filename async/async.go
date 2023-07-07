package async

import (
	"context"

	`github.com/getsentry/sentry-go`
)

func RecoverGO(f func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				hub := sentry.CurrentHub().Clone()
				hub.Recover(r)
			}
		}()
		f()
	}()
}

func RecoverFunc(f func() error) func() error {
	return func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				hub := sentry.CurrentHub().Clone()
				hub.Recover(r)
			}
		}()
		return f()
	}
}

func RunWithContext(ctx context.Context, f func() error) error {
	errChan := make(chan error)
	go func() {
		errChan <- f()
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}
