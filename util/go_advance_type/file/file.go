package ftil

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/optim-corp/cios-golang-sdk/util/go_advance_type/convert"
)

type (
	FileService struct {
		Path string
	}
	DirByt struct {
		Value   []byte
		AbsPath string
	}
)

func Path(path string) FileService {
	return FileService{Path: path}
}

func (file FileService) ReadFile() ([]byte, error) {
	return ioutil.ReadFile(file.Path)
}

func (file FileService) ReadString() (string, error) {
	b, e := file.ReadFile()
	return string(b), e
}

func (file FileService) ReadDirInfo() ([]os.FileInfo, error) {
	return ioutil.ReadDir(file.Path)
}

func (file FileService) ReadDirRec() ([]DirByt, error) {
	var (
		result = []DirByt{}
		fn     func(path string, files []os.FileInfo)
		err    error
	)
	absPath, err := filepath.Abs(file.Path)
	files, err := file.ReadDirInfo()
	if err != nil {
		return nil, err
	}
	fn = func(path string, files []os.FileInfo) {
		for _, f := range files {
			_path := filepath.Join(path, f.Name())
			if f.IsDir() {
				fs, err := ioutil.ReadDir(_path)
				if err == nil {
					fn(_path, fs)
				}
			} else {
				byts, err := ioutil.ReadFile(_path)
				if err == nil {
					result = append(result, DirByt{
						Value:   byts,
						AbsPath: _path,
					})
				}
			}
		}
	}
	fn(absPath, files)
	return result, err
}

func (file FileService) ReadDir() ([]DirByt, error) {
	var (
		result = []DirByt{}
		err    error
	)
	absPath, err := filepath.Abs(file.Path)
	files, err := file.ReadDirInfo()
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if !f.IsDir() {
			_path := filepath.Join(absPath, f.Name())
			byts, err := ioutil.ReadFile(_path)
			if err == nil {
				result = append(result, DirByt{
					Value:   byts,
					AbsPath: _path,
				})
			}
		}
	}
	return result, err
}

func (file FileService) ReadJsonMap() (interface{}, error) {
	return readJsonMap(file.Path)
}
func (file FileService) LoadJsonStruct(st interface{}) error {
	byts, err := ioutil.ReadFile(file.Path)
	if err != nil {
		return err
	}
	return convert.UnMarshalJson(byts, &st)
}

func (file FileService) WriteFileAsString(value string) error {
	if err := ioutil.WriteFile(file.Path, []byte(value), 0777); err != nil {
		return err
	}
	return nil
}

func (file FileService) WriteJson(data interface{}) error {
	var buf bytes.Buffer
	str, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = json.Indent(&buf, str, "", "  ")
	if err != nil {
		return err
	}
	return file.WriteFileAsString(buf.String())
}

func (file FileService) CreateDir() error {
	if _, err := os.Stat(file.Path); os.IsNotExist(err) {
		if err := os.Mkdir(file.Path, 0777); err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}

func (file FileService) CreateFile() error {
	f, err := os.OpenFile(file.Path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func (file FileService) IsDir() bool {
	fileInfo, err := os.Stat(file.Path)
	if err != nil {
		panic(err)
	}
	return fileInfo.IsDir()
}

func (file FileService) IsFile() bool {
	return !file.IsDir()
}

func readJsonMap(path string) (result interface{}, err error) {
	byts, _err := ioutil.ReadFile(path)
	if _err != nil {
		err = _err
		return
	}
	err = convert.UnMarshalJson(byts, &result)
	return
}
