package upload

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"

	imgtype "github.com/shamsher31/goimgtype"
	"woods/pkg/file"
	"woods/pkg/logging"
	"woods/pkg/setting"
	"woods/pkg/util"
)

// GetImageFullUrl get the full access path
func GetImageFullUrl(name string) string {
	return setting.AppSetting.Domain + "/" + GetImagePath() + name
}

// GetImageName get image name
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName, "")

	return fileName + ext
}

// GetImagePath get save path
func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

// GetImageFullPath get full save path
func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

// CheckImageExt check image file ext
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExt {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

// CheckImageSize check image size
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	fmt.Println("file size: ", size)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	return size <= setting.AppSetting.ImageMaxSize
}

// CheckImage check if the file exists
func CheckImage(src string) error {
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

func GetImageMime(fileName string) (string, error) {
	// 读取图片
	file, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("Open Image Failed: %s", err.Error())
	}
	defer file.Close()

	// 获取图片的类型
	datatype, err := imgtype.Get(fileName)
	if err != nil {
		return "", fmt.Errorf("Not Image File: %s", err.Error())
	}
	return datatype, nil
}
