package net

import (
	"github.com/trustyboy/goutils/ioutil"
	"io"
	"net/http"
	"os"
)

func DownloadFile(url string, savePath string, fileName string) error {
	bExist, _ := ioutil.PathExists(savePath + "/" + fileName)
	if bExist {
		return nil
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	_ = os.MkdirAll(savePath, os.ModePerm)
	fs, err := os.Create(savePath + "/" + fileName)
	if err != nil {
		return err
	}
	_, err = io.Copy(fs, res.Body)
	if err != nil {
		return err
	} else {
		return nil
	}
}
