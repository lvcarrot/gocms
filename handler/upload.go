package handler

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type upFile struct {
	*os.File
	hash.Hash
}

type Size interface {
	Size() int64
}

type Stat interface {
	Stat() (os.FileInfo, error)
}

var (
	endPoint   = ""
	accessID   = ""
	accessKey  = ""
	bucketName = ""
	bucket     *oss.Bucket
)

func init() {
	client, err := oss.New(endPoint, accessID, accessKey)
	if err != nil {
		panic(err)
	}
	bucket, err = client.Bucket(bucketName)
	if err != nil {
		panic(err)
	}
}

func (f upFile) Write(p []byte) (int, error) {
	if _, err := f.File.Write(p); err != nil {
		return -1, err
	}
	return f.Hash.Write(p)
}

// Upload 上传文件
func Upload(w http.ResponseWriter, r *http.Request) {
	f, h, err := r.FormFile("file_data")
	if err != nil {
		jFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	defer f.Close()

	var (
		fileSize int64
	)
	if sizeInterface, ok := f.(Size); ok {
		fileSize = sizeInterface.Size()
	} else if statInterface, ok := f.(Stat); ok {
		fileInfo, _ := statInterface.Stat()
		fileSize = fileInfo.Size()
	} else {
		jFailed(w, http.StatusBadRequest, "file size error")
		return
	}

	u := &upFile{Hash: md5.New()}
	if u.File, err = ioutil.TempFile("", "tempfile"); err != nil {
		jFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	defer os.Remove(u.File.Name())
	io.Copy(u, f)

	options := []oss.Option{
		oss.ContentType(h.Header.Get("Content-Type")),
	}
	bucket.PutObjectFromFile(path.Base(h.Filename), u.File.Name(), options...)

	jSuccess(w, map[string]string{
		"url":      "",
		"md5":      hex.EncodeToString(u.Hash.Sum(nil)),
		"pkg_size": fmt.Sprint("%d", fileSize),
	},
	)
}
