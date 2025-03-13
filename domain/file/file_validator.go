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
