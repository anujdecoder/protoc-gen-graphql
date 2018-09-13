package templates

import (
	"github.com/gobuffalo/packr"
	"io/ioutil"
)

func init() {
	fs := packr.NewBox("./static")

	if err := fs.Walk(func(s string, file packr.File) (err error) {
		dat, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		defer file.Close()

		T, err = T.Parse(string(dat))
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		panic(err)
	}
}