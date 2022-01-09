package test

import (
	"FrameWork/util/file_util"
	"FrameWork/util/file_util/file_common"
	_ "FrameWork/util/file_util/impl"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var fileMaxSize = int64(1024 * 1024)
var ctx = context.Background()

func TestFileSize(t *testing.T) {
	file_util.RefFileManagerFactory().SetFileMaxSize(ctx, fileMaxSize)
	assert.Equal(t, fileMaxSize, file_util.RefFileManagerFactory().GetCurrentMaxSize(ctx))
}

func TestFileExist(t *testing.T) {
	filePath := "/Users/mengliang.yang/selfcode/FrameWork"
	assert.NotPanics(t, func() {
		assert.Equal(t, true, file_util.RefFileManagerFactory().FileIsExist(ctx, filePath))
	})
}

func TestFileNotExist(t *testing.T) {
	filePath := "/Users/mengliang.yang/selde"
	assert.NotPanics(t, func() {
		assert.Equal(t, false, file_util.RefFileManagerFactory().FileIsExist(ctx, filePath))
	})
}

func TestGetFileInfoList(t *testing.T) {
	filePath := "/Users/mengliang.yang/test"
	fileInfoMap := map[string]int{
		"data.txt": 0,
		"test.csv": 1,
	}
	length := file_util.RefFileManagerFactory().GetFileInfoList(ctx, filePath)
	assert.Equal(t, 2, len(length))
	for _, v := range length {
		_, ok := fileInfoMap[v.Name()]
		assert.Equal(t, true, ok)
	}
}

func TestCheckFileType(t *testing.T) {
	filePath := "/Users/mengliang.yang/test/data.txt"
	assert.Equal(t, true, file_util.RefFileManagerFactory().CheckFileType(ctx, filePath, file_common.FILE_TYPE_TXT))
}
