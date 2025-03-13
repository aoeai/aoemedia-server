package file

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

	if err := ValidateSource(m.Source); err != nil {
		return err
	}

	if m.ModifiedTime.IsZero() {
		return newError("文件修改时间不能为空")
	}

	return nil
}

func ValidateSource(source uint8) error {
	if source == 0 {
		return newError("来源不能为空")
	}

	// 来源 1:相机 2:微信
	sourceList := []uint8{1, 2}
	if !contains(sourceList, source) {
		return newError("来源无效")
	}

	return nil
}
