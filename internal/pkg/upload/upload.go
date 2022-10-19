package upload

import (
	"bufio"
	"github.com/golang-module/carbon/v2"
	"github.com/google/wire"
	"github.com/spf13/cast"
	"io/ioutil"
	"os"
)

type Upload struct {
	File *os.File
}

func (f *Upload) CreateName() error {
	nanoSecond:=carbon.Now().TimestampNano()
	name := cast.ToString(nanoSecond)
	file, err := ioutil.TempFile("", name)
	if err != nil {
		return err
	}
	f.File = file
	return nil
}

func (f *Upload) WriteContent(content []byte) (int, error) {
	err:=f.CreateName()
	if err!=nil{
		return 0,err
	}
	n, err1 := f.File.Write(content)
	if err1 != nil {
		return 0, err1
	}
	return n, nil
}

func (f *Upload) GetFileName() string  {
	return f.File.Name()
}

func (f *Upload) Close() {
	buf := bufio.NewWriter(f.File)
	_ = buf.Flush()
	_ = f.File.Close()
	name := f.File.Name()
	_ = os.Remove(name)
}

func NewUpload() *Upload {
	return &Upload{}
}

var Provider=wire.NewSet(NewUpload)



