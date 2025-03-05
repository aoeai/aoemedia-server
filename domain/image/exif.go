package image

import (
	"fmt"
	"github.com/dsoprea/go-exif/v3"
	"time"
)

/**
## 通用

颜色模式: RGB
深度: 8
DPI高度: 72
DPI宽度: 72
方向: 1（正常）
像素高度: 3,648
像素宽度: 2,736
描述文件名称: sRGB IEC61966-2.1

## Exif

光圈值: 1.69
亮度值: 0
色彩空间: sRGB
色彩组合方案: 1，2，3，0
每像素的压缩位: 0.95
对比度: 正常
自定义渲染: 自定义过程
数字化日期时间: 2025年1月19日 19:25:39
原始日期时间: 2025年1月19日 19:25:39
数字缩放比例: 1
Exif版本: 2.1
曝光偏移值: 0
曝光模式: 自动曝光
曝光程序: 正常程序
曝光时间: 1/50
文件源: DSC
闪光灯: 无闪光
FlashPix版本: 1.0
光圈系数: 1.8
焦距: 5.58
按35毫米胶卷计的焦距: 26
增益控制: 无
ISO感光度等级: 250
光源: 日光
最大的光圈值: 1.69
测光模式: 图案
横向像素数: 2,736
纵向像素数: 3,648
饱和度: 正常
场景捕捉类型: 标准
场景类型: 直接拍摄的图像
感知方法: 单片色彩区域感应器
清晰度: 正常
快门速度值: 0/1
主体距离范围: 未知
次秒级时间: 339247
数字化次秒级时间: 339247
原始次秒级时间: 339247
白平衡: 自动白平衡

## GPS

海拔高度: 1,470.81 米 (4,825.49 英尺)
海拔高度参考: 高于海平面
日期戳: 2025年1月19日
GPS版本: 2.2.0.0
纬度: 北纬36° 3’ 5.655”
经度: 东经103° 50’ 16.38”
时间戳: 11:25:39世界标准时间

## JFIF

密度单位: 1
JFIF版本: 1.0.1
X密度: 96
Y密度: 96

## TIFF

日期时间: 2025年1月19日 19:25:39
品牌: HUAWEI
型号: TEL-AN00a
方向: 1（正常）
分辨率单位: 英寸
软件: TEL-AN00a 3.0.0.165(C00E165R5P2)
X分辨率: 72
Y分辨率: 72

通用、Exif、GPS、JFIF、TIFF

## 角色
Go专家、TDD专家、图片处理专家

## 目标

- 使用Go读取图片文件信息
- 使用方法名代替注释
- 使用函数式编程避免并发错误

## 任务

### 通用信息

- 颜色模式
- 深度
- DPI高度
- DPI宽度
- 方向
- 像素高度
- 像素宽度
- 描述文件名称
*/

// extractExifCreateTime 从图片数据中获取EXIF创建时间
func extractExifCreateTime(imageData []byte) (time.Time, error) {
	// 创建EXIF读取器
	rawExif, err := exif.SearchAndExtractExif(imageData)
	if err != nil {
		return time.Time{}, fmt.Errorf("读取EXIF数据失败: %w", err)
	}

	// 解析EXIF数据
	exifTags, _, err := exif.GetFlatExifData(rawExif, nil)
	if err != nil {
		return time.Time{}, fmt.Errorf("解析EXIF数据失败: %w", err)
	}

	// 尝试获取创建时间
	for _, ifd := range exifTags {
		if ifd.TagName == "DateTime" || ifd.TagName == "DateTimeOriginal" || ifd.TagName == "DateTimeDigitized" {
			strValue, ok := ifd.Value.(string)
			if !ok {
				continue
			}

			createTime, err := time.ParseInLocation("2006:01:02 15:04:05", strValue, time.Local)
			if err != nil {
				continue
			}

			return createTime, nil
		}
	}

	return time.Time{}, fmt.Errorf("未找到EXIF创建时间")
}
