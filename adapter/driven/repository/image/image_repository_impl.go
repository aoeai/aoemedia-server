package image

import (
	"path/filepath"
	"sync"

	"github.com/sirupsen/logrus"

	localFileStorage "github.com/aoemedia-server/adapter/driven/persistence/local_storage/file"
	"github.com/aoemedia-server/adapter/driven/persistence/mysql/db"
	mysqlimage "github.com/aoemedia-server/adapter/driven/persistence/mysql/image"
	repofile "github.com/aoemedia-server/adapter/driven/repository/file"
	"github.com/aoemedia-server/common/eventbus"
	"github.com/aoemedia-server/config"
	"github.com/aoemedia-server/domain/file"
	"github.com/aoemedia-server/domain/image"
)

type Repository struct {
	fileLocalStorage file.Storage
	fileRepository   file.Repository
}

var (
	instance *Repository
	once     sync.Once
)

func Inst() *Repository {
	once.Do(func() {
		instance = &Repository{localFileStorage.NewLocalFileStorage(), repofile.Inst()}
	})
	return instance
}

func (r *Repository) Upload(domainImage *image.DomainImage, userId int64) (result *image.UploadResult, err error) {
	fullStoragePath, err := r.storeLocally(domainImage)
	if err != nil {
		logrus.Errorf("上传图片失败，存储本地失败 filename:%v userId:%v %v", domainImage.FileName, userId, err)
		return nil, err
	}

	// 开启事务
	tx := db.Inst().Begin()
	if tx.Error != nil {
		logrus.Errorf("开启事务失败: %v", tx.Error)
		return nil, tx.Error
	}

	defer func() {
		if err != nil {
			logrus.Errorf("上传图片失败，回滚 filename:%v userId:%v %v", domainImage.FileName, userId, err)
			tx.Rollback()
			return
		}
		if commitErr := tx.Commit().Error; commitErr != nil {
			logrus.Errorf("上传图片提交事务失败，filename:%v userId:%v %v", domainImage.FileName, userId, commitErr)
			err = commitErr
		}
	}()

	// 保存文件
	fileId, err := r.fileRepository.Save(domainImage.DomainFile, tx)
	if err != nil {
		logrus.Errorf("保存文件失败 filename:%v userId:%v %v", domainImage.FileName, userId, err)
		return nil, err
	}

	// 创建图片上传记录
	imageUploadRecordId, err := mysqlimage.Create(userId, fileId, tx)
	if err != nil {
		logrus.Errorf("创建图片上传记录失败 filename:%v userId:%v %v", domainImage.FileName, userId, err)
		return nil, err
	}

	logrus.Infof("上传图片成功 userId:%v fileId:%v imageUploadRecordId:%v path:%v",
		userId, fileId, imageUploadRecordId, filepath.Join(fullStoragePath, domainImage.FileName))

	return &image.UploadResult{
		FileId:              fileId,
		ImageUploadRecordId: imageUploadRecordId,
		FullStoragePath:     fullStoragePath,
	}, nil
}

// storeLocally 保存图片到本地
func (r *Repository) storeLocally(image *image.DomainImage) (fullStoragePath string, error error) {
	fullDirPath := filepath.Join(config.Inst().Storage.ImageRootDir, createTimeOf(image))
	image.StorageDir = fullDirPath

	fullStoragePath, err := r.fileLocalStorage.Save(image.DomainFile)
	if err != nil {
		return "", err
	}
	return fullStoragePath, nil
}

func createTimeOf(image *image.DomainImage) string {
	return YearMonthOf(image.ModifiedTime)
}

func (r *Repository) PublishImageUploadedEvent(params *image.ImageUploadedEventPublishParams) error {
	event, err := newImageUploadedEvent(*params)
	if err != nil {
		logrus.Errorf("创建图片上传事件失败: %v %v", params, err)
		return err
	}

	eventbus.Inst().Publish(image.ImageUploadedEventType, event)
	logrus.Infof("发布图片上传事件成功: %v", event)

	return nil
}
