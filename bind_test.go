package bind

import (
	"fmt"
	"testing"
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"net/http"

	"github.com/lunny/tango"
)

type BindExample struct {
	Id   int64
	Name string
}

func (a *BindExample) Do() string {
	return fmt.Sprintf("%d-%s", a.Id, a.Name)
}

func TestBind(t *testing.T) {
	go func() {
		t := tango.Classic()
		t.Use(new(Binds))
		t.Get("/", new(BindExample))
		t.Run("0.0.0.0:9997")
	}()

	res, err := get("http://localhost:9997/?id=1&name=lllll")
	if err != nil {
		t.Error(err)
		return
	}

	if res != "1-lllll" {
		t.Error("not equal "+res+" != 1-lllll")
		return
	}
}

func gzipDecode(src []byte) ([]byte, error) {
	rd := bytes.NewReader(src)
	b, err := gzip.NewReader(rd)
	if err != nil {
		return nil, err
	}

	defer b.Close()

	data, err := ioutil.ReadAll(b)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.Header.Get("Content-Encoding") == "gzip" {
		data, err := gzipDecode(bs)
		if err != nil {
			return "", err
		}
		return string(data), nil
	}
	return string(bs), nil
}
