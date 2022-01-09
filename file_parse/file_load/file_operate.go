package file_load

import "context"

type FileOperateManager interface {
	BatchDownLoad(ctx context.Context, files []string, storePath string, compression bool, zipFileName string)
	BatchUpload(ctx context.Context, files []string)
	DealFile(ctx context.Context, behave func(ctx context.Context, singleLine string))
	BatchDealFile(ctx context.Context, behave func(ctx context.Context, lines []string), count int8)
	ExportDataToFile(ctx context.Context, data interface{}, filePath string, model int8, format string)
}

var fileLoadManager FileOperateManager

func RefFileOperateManager() FileOperateManager {
	return fileLoadManager
}

func InjectFileOperateManager(impl FileOperateManager) {
	fileLoadManager = impl
}
