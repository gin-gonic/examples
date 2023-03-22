package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/examples/upload-file/chunk/model"
	"github.com/google/uuid"
)

func main() {
	filePath := "your slice chunk upload file path"
	fileName := filepath.Base(filePath)

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Fatalf("file stat fail: %v\n", err)
		return
	}

	const chunkSize = 1 << (10 * 2) * 30

	num := math.Ceil(float64(fileInfo.Size()) / float64(chunkSize))

	fi, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file fail: %v\n", err)
		return
	}

	fileId := uuid.NewString()

	fileKeys := make([]string, 0)
	for i := 1; i <= int(num); i++ {
		file := make([]byte, chunkSize)
		fi.Seek((int64(i)-1)*chunkSize, 0)
		if len(file) > int(fileInfo.Size()-(int64(i)-1)*chunkSize) {
			file = make([]byte, fileInfo.Size()-(int64(i)-1)*chunkSize)
		}
		fi.Read(file)

		key := fmt.Sprintf("%x", md5.Sum(file))

		fileKeys = append(fileKeys, key)

		req := model.ChunkFileRequest{
			FileId:    fileId,
			FileName:  fileName,
			FileIndex: i,
			FileCount: int(num),
			FileKey:   key,
			FileKeys:  fileKeys,
			File:      file,
		}
		body, _ := json.Marshal(req)

		res, err := http.Post("http://127.0.0.1:8080/chunkUploadFile", "application/json", bytes.NewBuffer(body))

		if err != nil {
			log.Fatalf("http post fail: %v", err)
			return
		}
		defer res.Body.Close()
		msg, _ := io.ReadAll(res.Body)
		fmt.Println(string(msg))
	}
}
