// migrate 提供一个临时存储自动创建或更新表的仓库。
// 这段代码不应该被引用。 只是作为案例片段。
// 应该在自己的代码中创建不会冲突的仓库。
package migrate

import "gorm.io/gorm"

var bucket = make([]any, 0)

func AppendTables(tables ...any) {
	bucket = append(bucket, tables...)
}

func Migrate(migrator gorm.Migrator) error {
	return migrator.AutoMigrate(bucket...)
}
