package bind

import (
	"fmt"
	"testing"
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"

	"github.com/lunny/tango"
)

type BindExample struct {
	Id   int64
	Name string
}

func (a *BindExample) Get() string {
	return fmt.Sprintf("%d-%s", a.Id, a.Name)
}

func TestBind(t *testing.T) {
	buff := bytes.NewBufferString("")
	recorder := httptest.NewRecorder()
	recorder.Body = buff

	o := tango.Classic()
	o.Use(new(Binds))
	o.Get("/", new(BindExample))

	req, err := http.NewRequest("GET", "http://localhost:3000/?id=1&name=lllll", nil)
	if err != nil {
		t.Error(err)
	}

	o.ServeHTTP(recorder, req)
	expect(t, recorder.Code, http.StatusOK)
	refute(t, len(buff.String()), 0)
	expect(t, buff.String(), "1-lllll")
}

/* Test Helpers */
func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func refute(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		t.Errorf("Did not expect %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}