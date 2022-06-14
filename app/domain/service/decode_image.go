package service

import (
	"io"
)

// DecodeByImageURL 通过ImageUrl下载图片，并解析图片
func DecodeByImageURL(uploadUrl string) *UploadImage {
	return nil
}

// DecodeByUploadFile 通过上传的图片文件解析
func DecodeByUploadFile(uploadFile io.Writer) *UploadImage {
	return nil
}

// DecodeByRawData 通过图片二进制内容解析
func DecodeByRawData(rawData []byte) *UploadImage {
	return nil
}
