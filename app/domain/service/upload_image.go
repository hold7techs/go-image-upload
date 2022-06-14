package service

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go-image-upload/app/domain/entity"
)

type IUploader interface {
	// TransformImageType 图片类型转换
	TransformImageType(dstImgType entity.ImgType) error

	// AddWatermark 在原图(posX,posY)位置增加word或者img图片
	AddWatermark(watermark *entity.Watermark) error

	// ResizeImageWithSameRatio 等比重置图片大小
	ResizeImageWithSameRatio(widths []int) error
}

// UploadImage 上传的图
type UploadImage struct {
	FilePath  string               // 存储图片路径
	Cfg       *entity.UploadConfig // 上传图片配置
	ImgFile   image.Image          // 上传图片
	ImgType   entity.ImgType       // 图片类型
	Watermark *entity.Watermark    // 水印图片
}

// TransformImageType 转成指定的图片类型
func (u *UploadImage) TransformImageType(dstImgType entity.ImgType) error {
	if u.ImgType == dstImgType {
		return nil
	}
	if u.ImgType == entity.ImageTypeGIF {
		return errors.New("transform fail, image type cannot be gif")
	}

	fpath := u.getImageFilePath()
	if err := os.MkdirAll(fpath, 0755); err != nil {
		return errors.Wrapf(err, "transform fail, mkdir [%s] got err", fpath)
	}
	f, err := os.Create(fpath)
	if err != nil {
		return errors.Wrapf(err, "transform fail, create [%s] got err", fpath)
	}

	switch dstImgType {
	case entity.ImageTypePNG:
		err = png.Encode(f, u.ImgFile)
	case entity.ImageTypeJPEG:
		quality := u.Cfg.JpegQuality
		if quality == 0 {
			quality = 75
		}
		err = jpeg.Encode(f, u.ImgFile, &jpeg.Options{Quality: quality})
	case entity.ImageTypeWEBP:
		// todo implement webp
		return errors.New("transform fail, webp image type conversion will be implemented in the future")
	}

	u.ImgType = dstImgType
	return nil
}

// AddWatermark 为
func (u *UploadImage) AddWatermark(watermark *entity.Watermark) error {
	// TODO implement me
	panic("implement me")
}

// getImageFilePath 获取图片上存储的路径
func (u *UploadImage) getImageFilePath() string {
	fileMonth := time.Now().Format("06/01")
	fileName := fmt.Sprintf("%s%s", uuid.New().String(), entity.ImgFileExtMap[u.ImgType])
	filePath := fmt.Sprintf("%s/%s/%s", u.Cfg.ImageRoot, fileMonth, fileName)
	return filePath
}

// getImageFilePathBySize 获取图片上存储的路径
func (u *UploadImage) getImageFilePathBySize(size int) string {
	fileMonth := time.Now().Format("06/01")
	fileName := fmt.Sprintf("%s@%d%s", uuid.New().String(), size, entity.ImgFileExtMap[u.ImgType])
	filePath := fmt.Sprintf("%s/%s/%s", u.Cfg.ImageRoot, fileMonth, fileName)
	return filePath
}

func (u *UploadImage) ResizeImageWithSameRatio(widths []int) error {
	// TODO implement me
	panic("implement me")
}
