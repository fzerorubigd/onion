package onion

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const validJSON = `
{
	"string": "str",
	"number" : 100, 
	"nested" : {
		"bool" : "true"
	}
}
`

const invalidJSON = `{INVALID}`

func createFile(content string) string {
	tmp, err := ioutil.TempFile(os.TempDir(), "json_test")
	if err != nil {
		panic(err)
	}
	defer tmp.Close()
	fmt.Fprint(tmp, content)
	return tmp.Name()
}

func TestNewJSONLayer(t *testing.T) {
	Convey("JSON Loader test", t, func() {
		v := createFile(validJSON)
		defer os.Remove(v)
		l, err := NewJSONFileLayer(v)
		So(err, ShouldBeNil)
		o := New(l)
		So(o.GetString("string"), ShouldEqual, "str")
		So(o.GetInt("number"), ShouldEqual, 100)
		So(o.GetBool("nested.bool"), ShouldEqual, true)
	})

	Convey("JSON Loader fail", t, func() {
		v := createFile(invalidJSON)
		defer os.Remove(v)
		_, err := NewJSONFileLayer(v)
		So(err, ShouldBeError)

		_, err = NewJSONFileLayer("this_file_is_not_available_is_it")
		So(err, ShouldBeError)
	})
}
