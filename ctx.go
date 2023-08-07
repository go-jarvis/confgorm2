package confgorm2

import (
	"context"

	"gorm.io/gorm"
)

type gormDBKey int

var key = gormDBKey(100)

func WithContext(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, key, db)
}

func FromContext(ctx context.Context) *gorm.DB {

	val := ctx.Value(key)
	db, ok := val.(*gorm.DB)
	if ok {
		return db
	}

	return nil
}
