package localuploader

import (
	"context"
	"encoding/base64"
	"github.com/ItsWewin/superfactory/aerror"
	"os"
	"path"
	"strings"
)

type LocalUploader struct {
	dir string
}

func NewUploader(dir string) *LocalUploader {
	return &LocalUploader{dir: dir}
}

func (u *LocalUploader) UploadBase64(ctx context.Context, base64str string, fileName string) (string, aerror.Error) {
	arr := strings.Split(base64str, ";base64,")
	if len(arr) != 2 {
		return "", aerror.NewErrorf(nil, aerror.Code.BUnexpectedData, "base64 string is invalid")
	}

	dec, err := base64.StdEncoding.DecodeString(arr[1])
	if err != nil {
		return "", aerror.NewErrorf(err, aerror.Code.BUnexpectedData, "DecodeString failed")
	}

	dirPath := path.Join(u.dir, "pictures")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0777)
		if err != nil {
			return "", aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "create dir failed")
		}
	}

	fileName = path.Join(dirPath, fileName)
	f, err := os.Create(fileName)
	if err != nil {
		return "", aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "create file failed")
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		return "", aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "file write failed")
	}
	if err := f.Sync(); err != nil {
		return "", aerror.NewErrorf(err, aerror.Code.SSqlExecuteErr, "file load to disk failed")
	}

	return fileName, nil
}
