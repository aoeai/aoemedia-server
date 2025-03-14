package image

import (
	"testing"

	"github.com/aoemedia-server/adapter/driven/persistence/mysql/db"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	// 测试用例1：正常插入记录
	t.Run("成功插入记录并返回记录ID", func(t *testing.T) {
		// 准备测试数据
		userId := int64(1)
		fileId := int64(1)

		// 执行测试
		recordId, err := Create(userId, fileId, db.Inst())

		// 验证结果
		assert.NoError(t, err)
		assert.Greater(t, recordId, int64(0))

		// 验证插入的数据
		var savedRecord ImageUploadRecord
		db.Inst().First(&savedRecord, recordId)

		// 验证字段是否正确保存
		assert.Equal(t, userId, savedRecord.UserId, "用户ID不匹配")
		assert.Equal(t, fileId, savedRecord.FileId, "文件ID不匹配")
		assert.False(t, savedRecord.CreatedAt.IsZero(), "创建时间不应为空")

		// 清理测试数据
		db.Inst().Delete(&ImageUploadRecord{}, recordId)
	})

	// 测试用例2：重复插入记录
	t.Run("重复的UserId和FileId返回错误", func(t *testing.T) {
		// 准备测试数据
		userId := int64(2)
		fileId := int64(2)

		// 第一次插入
		recordId, err := Create(userId, fileId, db.Inst())
		assert.NoError(t, err)
		assert.Greater(t, recordId, int64(0))

		// 第二次插入相同的数据
		_, err = Create(userId, fileId, db.Inst())
		assert.Error(t, err)

		// 清理测试数据
		db.Inst().Delete(&ImageUploadRecord{}, recordId)
	})
}
