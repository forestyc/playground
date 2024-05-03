package db

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type SessionOption func(opt *gorm.Session)

// Session new a session with prepared statement and manually controlled transactions.
// To switch auto-commit call mdb.Session(db.WithSessionSkipDefaultTransaction(true))
func (mdb *Mysql) Session(opts ...SessionOption) *gorm.DB {
	// default
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(mdb.config.OperationTimeout)*time.Second)
	s := gorm.Session{
		PrepareStmt:            true,
		NewDB:                  true,
		Context:                ctx,
		SkipDefaultTransaction: true,
	}

	for _, opt := range opts {
		if opt != nil {
			opt(&s)
		}
	}
	return mdb.db.Session(&s)
}

func WithSessionDryRun(DryRun bool) SessionOption {
	return func(opt *gorm.Session) {
		opt.DryRun = DryRun
	}
}

func WithSessionPrepareStmt(PrepareStmt bool) SessionOption {
	return func(opt *gorm.Session) {
		opt.PrepareStmt = PrepareStmt
	}
}

func WithSessionNewDB(NewDB bool) SessionOption {
	return func(opt *gorm.Session) {
		opt.NewDB = NewDB
	}
}

func WithSessionInitialized(Initialized bool) SessionOption {
	return func(opt *gorm.Session) {
		opt.Initialized = Initialized
	}
}

func WithSessionSkipHooks(SkipHooks bool) SessionOption {
	return func(opt *gorm.Session) {
		opt.SkipHooks = SkipHooks
	}
}

func WithSessionSkipDefaultTransaction(SkipDefaultTransaction bool) SessionOption {
	return func(opt *gorm.Session) {
		opt.SkipDefaultTransaction = SkipDefaultTransaction
	}
}

func WithSessionDisableNestedTransaction(DisableNestedTransaction bool) SessionOption {
	return func(opt *gorm.Session) {
		opt.DisableNestedTransaction = DisableNestedTransaction
	}
}

func WithSessionAllowGlobalUpdate(AllowGlobalUpdate bool) SessionOption {
	return func(opt *gorm.Session) {
		opt.AllowGlobalUpdate = AllowGlobalUpdate
	}
}

func WithSessionFullSaveAssociations(FullSaveAssociations bool) SessionOption {
	return func(opt *gorm.Session) {
		opt.FullSaveAssociations = FullSaveAssociations
	}
}

func WithSessionQueryFields(QueryFields bool) SessionOption {
	return func(opt *gorm.Session) {
		opt.QueryFields = QueryFields
	}
}

func WithSessionContext(Context context.Context) SessionOption {
	return func(opt *gorm.Session) {
		opt.Context = Context
	}
}

func WithSessionLogger(Logger logger.Interface) SessionOption {
	return func(opt *gorm.Session) {
		opt.Logger = Logger
	}
}

func WithSessionNowFunc(NowFunc func() time.Time) SessionOption {
	return func(opt *gorm.Session) {
		opt.NowFunc = NowFunc
	}
}

func WithSessionCreateBatchSize(CreateBatchSize int) SessionOption {
	return func(opt *gorm.Session) {
		opt.CreateBatchSize = CreateBatchSize
	}
}
