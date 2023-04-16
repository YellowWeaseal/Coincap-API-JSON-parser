
package Coincap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client,error) {
	if timeout == 0{
		return nil,errors.New("timepot cant be zero")
	}
	return &Client{
		client: &http.Client{
			Timeout: timeout,
		},
	},nil
}
func (c Client) GetAssets()( []Asset, error){

	resp, err:=c.client.Get("https://api.coincap.io/v2/assets")
	if err!=nil{
		return nil, err
	}
	defer resp.Body.Close()

	body, err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		return nil, err
	}
	var r assetsResponse
	err = json.Unmarshal(body, &r)
	if err!=nil{
		return nil, err
	}
	return r.Assets,nil
}
func (c Client) GetAsset(name string)( Asset, error){
	url:=fmt.Sprintf("https://api.coincap.io/v2/assets/%s", name)
	resp, err:=c.client.Get(url)
	if err!=nil{
		return Asset{}, err
	}
	defer resp.Body.Close()

	body, err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		return Asset{}, err
	}
	var r assetResponse
	err = json.Unmarshal(body, &r)
	if err!=nil{
		return Asset{}, err
	}
	return r.Asset,nil
}