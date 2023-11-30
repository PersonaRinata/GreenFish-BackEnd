package pkg

import (
	"GreenFish/server/service/api/config"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/minio/minio-go/v7"
	"io"
)

func MinioVideoUpgrade(suffix string, tmpFilePath string, fileName string) error {
	_, err := config.GlobalMinioClient.FPutObject(context.Background(), config.GlobalServerConfig.MinioInfo.Bucket, fileName, tmpFilePath, minio.PutObjectOptions{
		ContentType: "video/" + suffix,
	})
	if err != nil {
		hlog.Error("minio upgrade video failed,", err)
		return err
	}
	return nil
}

func MinioCoverUpgrade(tmpFilePath string, fileName string) error {
	_, err := config.GlobalMinioClient.FPutObject(context.Background(), config.GlobalServerConfig.MinioInfo.Bucket, fileName, tmpFilePath, minio.PutObjectOptions{
		ContentType: "image/png",
	})
	if err != nil {
		hlog.Error("minio upgrade cover failed,", err)
		return err
	}
	return nil
}

func MinioAvatarUpgrade(file io.Reader, fileName string, size int64) error {
	err := config.GlobalMinioClient.RemoveObject(context.Background(), config.GlobalServerConfig.MinioInfo.Bucket, fileName, minio.RemoveObjectOptions{})
	if err != nil {
		hlog.Error("minio delete avatar failed,", err)
		return err
	}
	_, err = config.GlobalMinioClient.PutObject(context.Background(), config.GlobalServerConfig.MinioInfo.Bucket, fileName, file, size, minio.PutObjectOptions{})
	if err != nil {
		hlog.Error("minio upgrade avatar failed,", err)
		return err
	}
	return nil
}
