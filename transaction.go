package utils

import (
	"context"
	"runtime"

	"github.com/pkg/errors"
	"xorm.io/xorm"
)

// 出于对事务处理，将多个表的事务同时处理

// TransactionUtil 主要是为了service可以方便的调用事务而做出来的一个结构体
type TransactionUtil interface {
	// Transaction
	//  开启一个非自动提交的事务, 可以嵌套调用，每次都会从ctx中找是否有合适的事务，如果没有
	//  合适的则会新建一个事务，并且注入到ctx，开启一个事务，任意一个fc抛出错误都会进行回滚
	//  ctx 上下文，会尝试从ctx中获取事务
	Transaction(ctx context.Context, fc func(ctx context.Context) error) error

	// WithContextForTransaction
	//  repo 自己内部使用的方法，如果你希望你的repo方法支持上面那种事务,那么需要用这个方法来获取
	// 	session。  先尝试从这个ctx获取session，如果获取到了就返回这个session，如果没有获取
	//	到session，会自己创建一个新自动提交的session
	WithContextForTransaction(ctx context.Context) *xorm.Session
}

type (
	mysqlClient xorm.EngineInterface
)

type transactionImpl struct {
	mysqlClient mysqlClient // 这里的mysqlClient是一个接口，里面有一个NewSession()方法，返回一个*xorm.Session
}

func (t *transactionImpl) Transaction(ctx context.Context, fc func(ctx context.Context) error) (err error) {

	if _, ok := t.getTx(ctx); ok {
		// 已经有事务了，继续执行就好了
		return fc(ctx)
	}
	tx := t.mysqlClient.NewSession()
	if err = tx.Begin(); err != nil {
		err = errors.WithMessage(err, "[TransactionImpl.Transaction] begin transaction failed")
		return
	}
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 4096)
			buf = buf[:runtime.Stack(buf, false)]
			err = errors.Errorf("db transaction panic: %v, stack: \n%s", e, buf)
		}
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	ctx = context.WithValue(ctx, t.mysqlClient, tx)
	err = fc(ctx)
	if err == nil {
		err = tx.Commit()
	}
	return
}

func (t *transactionImpl) getTx(ctx context.Context) (*xorm.Session, bool) {
	txAny := ctx.Value(t.mysqlClient)
	tx, ok := txAny.(*xorm.Session)
	if !ok {
		return nil, false
	}
	return tx, true
}

func (t *transactionImpl) WithContextForTransaction(ctx context.Context) *xorm.Session {
	tx, ok := t.getTx(ctx)
	if ok {
		// 已经有tx了，用这个tx即可
		return tx
	}
	// 没有tx, 返回一个自动提交的transaction
	return t.mysqlClient.Context(ctx)

}
