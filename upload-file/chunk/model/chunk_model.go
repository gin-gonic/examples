package model

import (
	"crypto/md5"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type ChunkFileRequest struct {
	FileId    string   `json:"fileId"`    // client create uuid
	FileName  string   `json:"fileName"`  // file name
	FileIndex int      `json:"fileIndex"` // file index
	FileCount int      `json:"fileCount"` // file slice size
	FileKeys  []string `json:"fileKeys"`  // file slice all key md5 (accumulation file key if fileIndex == fileCount then merge after verification)
	FileKey   string   `json:"fileKey"`   // file now key to md5 - if server read the slice to md5 eq key not eq then fail
	File      []byte   `json:"file"`      // now file
}

func (cf *ChunkFileRequest) BindingForm(ctx *gin.Context) error {
	if err := ctx.ShouldBind(cf); err != nil {
		return err
	}

	return cf.md5()
}

func (cf *ChunkFileRequest) md5() error {
	hash := fmt.Sprintf("%x", md5.Sum(cf.File))
	if hash != cf.FileKey {
		return errors.New("current file slice key error")
	}
	return nil
}

func (cf *ChunkFileRequest) SaveUploadedFile(tempPath, path string) (string, error) {
	tempFolder := filepath.Join(tempPath, cf.FileId)

	_, err := os.Stat(tempFolder)
	if os.IsNotExist(err) {
		err := os.MkdirAll(tempFolder, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	out, err := os.Create(filepath.Join(tempFolder, cf.FileKey))
	if err != nil {
		return "", err
	}
	defer out.Close()
	if _, err := out.Write(cf.File); err != nil {
		return "", err
	}

	fmt.Println(cf.FileIndex, cf.FileCount)
	if cf.FileIndex != cf.FileCount {
		return "", nil
	}
	for _, fileKey := range cf.FileKeys {
		tempFile := filepath.Join(tempFolder, fileKey)
		if _, err := os.Stat(tempFile); err != nil {
			return "", errors.New("file " + fileKey + " is emtpy")
		}
	}

	base := filepath.Dir(path)
	if _, err := os.Stat(base); err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(base, os.ModePerm)
			if err != nil {
				return "", err
			}
		}
	}

	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		return "", err
	}

	defer file.Close()

	for _, fileKey := range cf.FileKeys {
		tempFile := filepath.Join(tempFolder, fileKey)
		bt, err := os.ReadFile(tempFile)
		if err != nil {
			return "", err
		}
		file.Write(bt)
	}

	return tempFolder, nil
}

// request method is json
// param: fileId
// param: fileName
// param: fileIndex the file slice index
// param: fileCount the file slice size
// param: fileKeys the file slice all file key md5 (accumulation file key if fileIndex == fileCount then merge after verification)
// param: fileKey  now file slice key md5
// param: file     now slice file
func ChunkUploadFile(ctx *gin.Context) {
	var cf ChunkFileRequest

	if err := cf.BindingForm(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "400", "msg": "bad file param", "err": err.Error()})
		return
	}

	tempFolder, err := cf.SaveUploadedFile("./temp", "./uploads/"+cf.FileName)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"code": "503", "msg": "bad save upload file", "err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": "200", "msg": "success"})
	if tempFolder != "" {
		defer func(tempFolder string) {
			os.RemoveAll(tempFolder)
		}(tempFolder)
	}
}
