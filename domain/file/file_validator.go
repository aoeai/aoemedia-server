package file

import "errors"

func (f *DomainFile) validate() error {
	result := []error{f.Content.validate(), f.Metadata.validate()}

	for _, err := range result {
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Content) validate() error {
	if c == nil {
		return newError("文件内容不能为空")
	}
	if c.SizeInBytes <= 0 {
		return newError("文件内容大小必须大于0")
	}
	if c.HashValue == "" {
		return newError("文件内容哈希值不能为空")
	}
	if len(c.HashValue) != 64 {
		return newError("文件内容哈希值长度必须是64")
	}

	return nil
}

func (m *Metadata) validate() error {
	if m == nil {
		return newError("文件元数据不能为空")
	}
	if m.FileName == "" {
		return newError("文件名不能为空")
	}
	if m.StorageDir == "" {
		return newError("存储路径不能为空")
	}

	if m.Source == 0 {
		return newError("文件来源不能为空")
	}
	// 来源 1:相机 2:微信
	sourceList := []uint8{1, 2}
	if !contains(sourceList, m.Source) {
		return newError("文件来源无效")
	}

	if m.ModifiedTime.IsZero() {
		return newError("文件修改时间不能为空")
	}

	return nil

}

func newError(text string) error {
	return errors.New(text)
}

func contains(list []uint8, target uint8) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}
	return false
}
