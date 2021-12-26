package storage

import (
	"context"
	"fmt"
	"github.com/jlaffaye/ftp"
	"github.com/romaxa83/mst-app/gin-app/pkg/logger"
	"time"
)

type FtpStorage struct {
	baseUrl  string
	host     string
	port     string
	username string
	password string
	timeOut  int
}

func NewFileStorage(baseUrl, host, port, username, password string, timeOut int) *FtpStorage {
	return &FtpStorage{
		baseUrl:  baseUrl,
		host:     host,
		port:     port,
		username: username,
		password: password,
		timeOut:  timeOut,
	}
}

func (fs *FtpStorage) Upload(ctx context.Context, input UploadInput) (string, error) {

	//logger.Warnf("INPUT - %+v", input)

	addr := fmt.Sprintf("%s:%s", fs.host, fs.port)
	c, err := ftp.Dial(addr, ftp.DialWithTimeout(30*time.Second))
	if err != nil {
		logger.Error(err)
	}

	err = c.Login(fs.username, fs.password)
	if err != nil {
		logger.Error(err)
	}

	err = c.MakeDir(fmt.Sprintf("%s", input.Type))
	if err != nil {
		logger.Warn(err)
	}

	err = c.Stor(input.Name, input.File)
	if err != nil {
		logger.Error(err)
	}

	return fs.generateFileURL(input.Name), nil
}

func (fs *FtpStorage) generateFileURL(filename string) string {
	return fmt.Sprintf("%s/%s", fs.baseUrl, filename)
}
