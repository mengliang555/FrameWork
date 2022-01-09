package file_util

import (
	"context"
	"os"
)

type FileManagerFactory interface {
	CheckFileType(ctx context.Context, filePath string, targetType string) bool
	FileIsExist(ctx context.Context, filePath string) bool
	SetFileMaxSize(ctx context.Context, fileSize int64)
	GetCurrentMaxSize(ctx context.Context) int64
	GetFileInfoList(ctx context.Context, dir string) []os.FileInfo

	CompressionFile(ctx context.Context, targetPath string, zipFileName string, recursion bool)
	UnCompressionFile(ctx context.Context, targetFile, targetPath string)

	GenerateFileName(ctx context.Context, filePath, targetPath string)string
}

var fileManagerFactory FileManagerFactory

func RefFileManagerFactory() FileManagerFactory {
	return fileManagerFactory
}

func InjectFileManagerFactory(impl FileManagerFactory) {
	fileManagerFactory = impl
}
