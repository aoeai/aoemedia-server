package testpath

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ProjectRoot 返回项目的根目录路径。
// 该函数通过获取当前工作目录，然后向上遍历目录树直到找到名为"aoemedia-server"的目录来确定项目根目录。
//
// 返回值:
//   - string: 项目根目录的绝对路径
//   - error: 如果获取当前工作目录失败，则返回相应的错误
func ProjectRoot() (string, error) {
	// 获取当前工作目录
	projectRoot, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("获取当前工作目录失败: %w", err)
	}
	// 移动到项目根目录
	maxIterations := 50
	iterationCount := 0
	for !strings.HasSuffix(projectRoot, "aoemedia-server") {
		if iterationCount >= maxIterations {
			return "", fmt.Errorf("无法找到项目根目录：超过最大循环次数限制（%d次）", maxIterations)
		}
		projectRoot = filepath.Dir(projectRoot)
		iterationCount++
	}

	return projectRoot, nil
}
