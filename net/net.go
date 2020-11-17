package net

import (
	"github.com/trustyboy/goutils/ioutil"
	"io"
	"net/http"
	"os"
)

func DownloadFile(url string, fileName string) error {
	bExist, _ := ioutil.PathExists("images/" + fileName)
	if bExist {
		return nil
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	bExist, _ = ioutil.PathExists("./images")
	if !bExist {
		os.Mkdir("./images", os.ModePerm)
	}
	fs, err := os.Create("images/" + fileName)
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
