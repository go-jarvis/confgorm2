// migrate 提供一个临时存储自动创建或更新表的仓库。
package migrate

import "gorm.io/gorm"

type bucket struct {
	tables []any
}

func New() *bucket {
	return &bucket{
		tables: make([]any, 0),
	}
}

func (b *bucket) AddTable(tables ...any) {
	b.tables = append(b.tables, tables...)
}

func (b *bucket) Migrate(m gorm.Migrator) error {
	return m.AutoMigrate(b.tables...)
}
