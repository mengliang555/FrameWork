package impl

import (
	"FrameWork/util/file_util"
	"FrameWork/util/method_util"
	"archive/zip"
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type fileUtil struct {
	maxSize int64
}

func (f *fileUtil) GenerateFileName(_ context.Context, filePath, targetPath string) string {
	ansPath := []string{targetPath}
	if !filepath.IsAbs(filePath) {
		ansPath = append(ansPath, filePath)
	} else {
		filePathList := strings.Split(filePath, string(os.PathSeparator))
		if filePathList == nil || len(filePathList) <= 0 {
			return ""
		}
		if index := strings.LastIndex(filePath, string(os.PathSeparator)); index == len(filePath) {
			ansPath = append(ansPath, filePathList[len(filePathList)-1]+"/")
		} else {
			ansPath = append(ansPath, filePathList[len(filePathList)-1])
		}
		ansPath = append(ansPath, filePath[strings.Index(filePath, string(os.PathSeparator)):])
	}
	return strings.Join(ansPath, string(os.PathSeparator))
}

func (f *fileUtil) CompressionFile(_ context.Context, targetPath string, zipFileName string, recursion bool) {
	method_util.ErrToPanic(os.RemoveAll(zipFileName))

	// 创建：zip文件
	zipFile, err := os.Create(zipFileName)
	method_util.ErrToPanic(err)
	defer func() {
		method_util.ErrToPanic(zipFile.Close())
	}()

	// 打开：zip文件
	archive := zip.NewWriter(zipFile)
	defer func() {
		method_util.ErrToPanic(archive.Close())
	}()

	// 遍历路径信息
	method_util.ErrToPanic(filepath.Walk(targetPath, func(path string, info os.FileInfo, _ error) error {

		// 如果是源路径，提前进行下一个遍历
		if (path == targetPath) || (info.IsDir() && !recursion) {
			return nil
		}

		header, err := zip.FileInfoHeader(info)
		method_util.ErrToPanic(err)

		header.Name = strings.TrimPrefix(path, targetPath+string(os.PathSeparator))

		// 判断：文件是不是文件夹
		if info.IsDir() {
			header.Name += string(os.PathSeparator)
		} else {
			// 设置：zip的文件压缩算法
			header.Method = zip.Deflate
		}

		// 创建：压缩包头部信息
		writer, _ := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer func() {
				method_util.ErrToPanic(file.Close())
			}()

			_, err = io.Copy(writer, file)
			method_util.ErrToPanic(err)
		}
		return nil
	}))
}

func (f *fileUtil) UnCompressionFile(ctx context.Context, targetFile string, targetPath string) {
	fr, err := zip.OpenReader(targetFile)
	if err != nil {
		panic(err)
	}
	defer func() { method_util.ErrToPanic(fr.Close()) }()
	//r.reader.file 是一个集合，里面包括了压缩包里面的所有文件
	for _, file := range fr.Reader.File {
		//判断文件该目录文件是否为文件夹
		if file.FileInfo().IsDir() {
			err := os.MkdirAll(f.GenerateFileName(ctx, file.Name, targetPath), 0644)
			method_util.ErrToPanic(err)
			continue
		}

		r, err := file.Open()
		method_util.ErrToPanic(err)

		NewFile, err := os.Create(f.GenerateFileName(ctx, file.Name, targetPath))
		method_util.ErrToPanic(err)

		_, err = io.Copy(NewFile, r)
		method_util.ErrToPanic(err)
		//关闭文件
		method_util.ErrToPanic(NewFile.Close())
		method_util.ErrToPanic(r.Close())
	}
}

func (f *fileUtil) CheckFileType(_ context.Context, filePath string, targetType string) bool {
	return filepath.Ext(filePath) == targetType
}

func (f *fileUtil) FileIsExist(_ context.Context, filePath string) (exist bool) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	method_util.ErrToPanic(err)
	return
}

func (f *fileUtil) SetFileMaxSize(_ context.Context, fileSize int64) {
	f.maxSize = fileSize
}

func (f *fileUtil) GetCurrentMaxSize(_ context.Context) int64 {
	return f.maxSize
}

func (f *fileUtil) GetFileInfoList(_ context.Context, dir string) []os.FileInfo {
	file, err := os.Open(dir)
	method_util.ErrToPanic(err)
	list, err := file.Readdir(-1)
	method_util.ErrToPanic(err)
	defer func() {
		method_util.ErrToPanic(file.Close())
	}()
	return list
}

func init() {
	file_util.InjectFileManagerFactory(&fileUtil{})
}
