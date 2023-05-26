package upload

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"wood/pkg/file"
	"wood/pkg/logging"
	"wood/pkg/setting"
	"wood/pkg/util"
)

// GetFileFullUrl get the full access path
func GetFileFullUrl(name string) string {
	return setting.AppSetting.Domain + "/" + GetFilePath() + name
}

// GetFileName get image name
func GetFileName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName, "")

	return fileName + ext
}

// GetFilePath get save path
func GetFilePath() string {
	return setting.AppSetting.FileSavePath
}

// GetFileFullPath get full save path
func GetFileFullPath() string {
	dir, _ := os.Getwd()
	return dir + "/" + setting.AppSetting.RuntimeRootPath + GetFilePath()
}

// CheckFileExt check image file ext
func CheckFileExt(fileName string) bool {
	ext := file.GetExt(fileName)
	fmt.Println("文件后缀：", ext, setting.AppSetting.FileAllowExt)
	for _, allowExt := range setting.AppSetting.FileAllowExt {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

// CheckFileSize check image size
func CheckFileSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	fmt.Println("file size: ", size)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	return size <= setting.AppSetting.FileMaxSize
}

// CheckFile check if the file exists
func CheckFile(src string) error {
	dir, err := os.Getwd()
	fmt.Println("get wd: ", dir)
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
