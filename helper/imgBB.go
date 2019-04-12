package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ImgBB struct {
	APIKey string
}

func NewImgBBConn() *ImgBB {
	return &ImgBB{
		APIKey: "a31b9828d1404c7042f1ba2500c8c1f5",
	}
}

func (i *ImgBB) Upload(base64 string) (imgURL string, err error) {

	request_url := "https://api.imgbb.com/1/upload"

	form := url.Values{
		"key":   {i.APIKey},
		"image": {base64},
		"name":  {"Image"},
	}

	body := bytes.NewBufferString(form.Encode())
	rsp, err := http.Post(request_url, "application/x-www-form-urlencoded", body)
	if err != nil {
		return imgURL, err
	}

	if rsp.StatusCode != 200 {
		return "", errors.New("Upload image fail")
	}

	defer rsp.Body.Close()
	body_byte, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}

	mapRes := make(map[string]interface{})

	err = json.Unmarshal(body_byte, &mapRes)

	if err != nil {
		return imgURL, err
	}

	url := mapRes["data"].(map[string]interface{})["image"].(map[string]interface{})["url"]

	imgURL = url.(string)

	return imgURL, nil
}
