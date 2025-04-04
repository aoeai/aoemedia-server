package upload

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// postFile 上传文件并返回响应结果
// 返回值:
//   - int: HTTP响应状态码
//   - map[string]interface{}: 响应内容，包含message（上传结果消息）、filename（文件名）、size（文件大小）和hash（文件哈希值）
func postFile(t *testing.T, testFilePath, url string, token string) (int, map[string]interface{}) {
	file := openTestFile(t, testFilePath)
	defer closeTestFile(t, file)

	// 创建multipart表单
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(testFilePath))
	if err != nil {
		t.Fatalf("创建表单文件失败: %v", err)
	}

	// 复制文件内容到表单
	if _, err := io.Copy(part, file); err != nil {
		t.Fatalf("复制文件内容失败: %v", err)
	}
	errClose := writer.Close()
	if errClose != nil {
		t.Fatalf("关闭multipart表单失败: %v", errClose)
	}

	// 创建测试请求
	req := httptest.NewRequest(http.MethodPost, url, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", token)
	w := httptest.NewRecorder()

	// 执行请求
	newTestRouter().ServeHTTP(w, req)

	// 解析响应内容
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("解析响应失败: %v", err)
	}

	return w.Code, response
}

func newTestRouter() *gin.Engine {
	// 设置测试环境
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST(File, NewFileController().Upload)
	r.POST(Image, NewImageController().Upload)

	return r
}

func openTestFile(t *testing.T, testFilePath string) *os.File {
	file, err := os.Open(testFilePath)
	if err != nil {
		t.Fatalf("无法打开测试文件: %v", err)
	}
	return file
}

func closeTestFile(t *testing.T, file *os.File) {
	if err := file.Close(); err != nil {
		t.Fatalf("关闭测试文件失败: %v", err)
	}
}

func assertSuccess(t *testing.T, code int, response map[string]interface{}, expectedFilename, expectedHash string, expectedSize float64) {
	// 验证响应状态码
	assert.Equal(t, http.StatusOK, code)

	// 验证响应字段
	assert.Equal(t, "文件上传成功", response["message"])
	assert.Equal(t, expectedFilename, response["filename"])
	assert.Equal(t, expectedSize, response["size"])
	assert.Equal(t, expectedHash, response["hash"])
}

func assertBadRequest(t *testing.T, testFilePath, expectedErrorMsg, url string) {
	code, response := postFile(t, testFilePath, url, "b")

	// 验证响应状态码
	assert.Equal(t, http.StatusBadRequest, code)

	// 验证响应字段
	assert.Equal(t, expectedErrorMsg, response["error"])
}
