package confgorm2

import (
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type gormDBKey int

var key = gormDBKey(100)

func WithContext(ctx context.Context, db *gorm.DB) context.Context {

	ginctx, ok := ctx.(*gin.Context)
	if ok {
		ctx = ginctx.Request.Context()
	}

	return context.WithValue(ctx, key, db)
}

func FromContext(ctx context.Context) *gorm.DB {

	ginctx, ok := ctx.(*gin.Context)
	if ok {
		ctx = ginctx.Request.Context()
	}

	val := ctx.Value(key)
	db, ok := val.(*gorm.DB)
	if ok {
		return db
	}

	return nil
}

func InjectGinContext(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := WithContext(c, db)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
