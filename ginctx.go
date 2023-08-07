package confgorm2

import (
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InjectGinContext(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx = WithContext(ctx, db)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func ExtractGinContext(ctx context.Context) *gorm.DB {

	ginctx, ok := ctx.(*gin.Context)
	if ok {
		ctx = ginctx.Request.Context()
	}

	return FromContext(ctx)
}
