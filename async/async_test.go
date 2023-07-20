package async

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func TestRecoverGO(t *testing.T) {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:        "https://50e55437f5774080a6eac8cc4a8edfc8@sentry.qimao.com/3",
		SampleRate: 1.0,
	}); err != nil {
		log.Infof("Sentry initialization failed: %v\n", err)
	}
	type args struct {
		f func()
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{f: func() {
				panic("1111")
			}},
		},
		{
			name: "2",
			args: args{f: func() {
				var x [10]int
				b := 10
				x[b] = 10
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var g sync.WaitGroup
			g.Add(1)
			RecoverGO(func() {
				defer g.Done()
				tt.args.f()
			})
			g.Wait()
			t.Log(tt.name, "结束")
		})
	}
	<-time.After(3 * time.Second)
}

func TestRecoverWrap(t *testing.T) {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:        "https://50e55437f5774080a6eac8cc4a8edfc8@sentry.qimao.com/3",
		SampleRate: 1.0,
	}); err != nil {
		log.Infof("Sentry initialization failed: %v\n", err)
	}
	type args struct {
		f1 func() error
		f2 func() error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				f1: func() error {
					panic("panic 1111")
				},
				f2: func() error {
					panic("panic 2222")
				},
			},
		},
		{
			name: "2",
			args: args{
				f1: func() error {
					return errors.New("正常错误1")
				},
				f2: func() error {
					return errors.New("正常错误2")
				},
			},
		},
		{
			name: "3",
			args: args{
				f1: func() error {
					var x [10]int
					b := 10
					x[b] = 10
					return nil
				},
				f2: func() error {
					return nil
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, _ := errgroup.WithContext(context.Background())
			g.Go(RecoverFunc(tt.args.f1))
			g.Go(RecoverFunc(tt.args.f2))
			err := g.Wait()
			t.Log("接收到", err)
		})
	}
	<-time.After(3 * time.Second)
}
