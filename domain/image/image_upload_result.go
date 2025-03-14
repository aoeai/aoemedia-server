package image

// UploadResult 图片上传结果
type UploadResult struct {
	// FileId 文件 ID
	FileId int64

	// ImageUploadRecordId 图片上传记录 ID
	ImageUploadRecordId int64

	// FullStoragePath 文件的完整存储路径
	FullStoragePath string
}
