// Code generated by go-bindata.
// sources:
// sample-ilxd.conf
// DO NOT EDIT!

package repo

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _sampleIlxdConf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x55\x51\x6f\xdb\x36\x10\x7e\xd7\xaf\xb8\x87\x14\xd8\x00\xd7\x8a\xd3\x6c\x45\xa2\x79\x80\x9b\x64\x45\x3a\xaf\x16\x62\x27\xcd\xf2\x46\x93\x27\x89\x33\xcd\xd3\x48\xca\xb6\x36\x2c\xbf\x7d\x38\x4a\x4e\x9c\xb5\x01\x06\xbf\xd8\xbc\xef\xbe\xbb\xfb\x3e\xf2\x9c\xc1\xa2\x42\x50\xda\xa1\x0c\xe4\x5a\x08\x04\x3e\x90\x43\x50\x22\x08\xf0\x8d\xac\x40\x78\x08\x15\x02\x2d\x77\xf1\x70\x29\x3c\x0e\x93\x3e\x0f\x0b\xd1\x98\x00\xda\xc3\x63\x3a\x64\x04\x59\xc8\x67\xf3\xeb\x7b\x98\xcd\xd1\x0f\xe0\x68\x3a\xbb\x98\x4c\x27\x79\x7e\x39\x59\x4c\xd2\x1e\xf0\x45\x5b\x45\x5b\x3f\x48\x32\x78\x4c\xa7\x7a\xe9\x84\x6b\xd3\x49\x5d\x1b\x2d\x45\xd0\x64\x61\xde\xd4\x35\xb9\xb0\xc7\xff\x26\x24\xcc\xe6\x03\x10\x56\xc1\x51\x45\x6b\xec\x03\x49\x06\xb9\x11\xf6\x6c\x08\x70\x65\x37\xda\x91\x5d\xa3\x0d\xb0\x11\x4e\x8b\xa5\x41\x0f\xc2\x21\xe0\xae\x16\x56\xa1\x02\x4f\x3c\x46\x0b\x6b\xd1\xc2\x12\xa1\xf1\xa8\x86\x00\x9f\x67\x8b\xab\xf3\x7d\x47\x49\x06\xf8\x2a\x51\x68\x6b\x2d\x85\x31\x2d\xbc\xb9\x9b\xdc\x5c\x4f\x3e\x4c\xaf\xde\x0c\x60\xd9\x84\x9e\xb6\xf1\x81\x79\x85\x94\xe8\x3d\x2a\xd8\xea\x50\x25\x19\x1c\xed\xc1\x50\xa1\xc3\x21\xc0\xc4\x78\x1a\xc0\x23\x6b\xf6\xd4\x5b\xa0\x97\x4a\x1d\xa8\xc4\x52\xb3\xec\x4a\xbb\xf1\x63\x3a\xd4\x66\xa7\x92\x24\x83\x5b\x8f\x10\xd0\x07\x8b\x81\x11\xfd\xd7\xf1\x28\xc6\xac\xde\xa0\xf3\xc2\x40\x6e\x9a\x32\xca\x96\x1b\xd1\xc2\x77\xb7\xb9\xcd\xbf\x07\xd1\x04\x5a\x8b\xd0\x0f\x43\x35\xda\xce\x60\xa3\x7d\x40\x0b\xac\x3c\xd0\x32\x08\x6d\x59\x10\x8e\xe0\x2e\xa0\xb3\xc2\xc0\x75\x0e\x42\x29\x87\xde\x43\xe1\x68\x0d\xbe\x33\x0a\x15\x28\xdc\x68\x89\x7e\x08\x8b\x4a\x7b\xa0\x3a\xfa\xa8\xb4\xef\xf4\xd3\xb1\x49\x4b\x4d\x6d\xeb\xae\xc7\x39\xa2\xda\x73\xf5\x02\xb3\x23\xac\xc4\x1f\xa4\x6d\x2c\x6b\x31\x6c\xc9\xad\x86\x30\xb3\xe0\x83\x70\xa1\x3b\x25\x85\xb0\xd5\xc6\xc0\x5a\xac\x30\xc9\x80\x9a\x50\x92\xb6\x25\x48\xb2\x16\x25\x57\xf6\xcc\xc3\xe0\x25\x51\xf0\xc1\x89\x1a\x6a\x44\xe7\xa3\x16\x0d\x4b\x57\xe1\x9a\x31\x4a\x7b\x49\x1b\x74\x40\xa1\x42\x97\x64\x3d\xec\x3f\x0d\x24\x19\x78\x44\xc5\xed\x8e\x53\x5d\x9f\xa6\xbb\x61\xfc\xa4\x41\xd6\xe9\xe9\xf1\xf1\x28\xad\x4f\xea\x74\x74\x72\xf9\xee\x57\xa2\x2f\xf9\xc3\xbb\xdd\x87\xcf\x37\x1f\x77\xa7\x45\x75\xb3\x2c\x7e\x9f\xc8\xfb\xdb\x4a\x3e\x54\x8b\x87\x93\xe9\xc5\xea\xd3\xfb\xd3\xd5\xa7\xfb\x8f\xc5\x5f\x67\x8b\xbb\xe9\x82\xa5\x98\x76\xba\xbf\x14\x83\xcb\x1f\x9c\x58\x05\xb5\xa3\x40\x92\x8c\x7f\x12\xaa\x37\xac\x20\x07\xda\x4a\x5a\x6b\x5b\xb2\xca\x5d\xd3\x87\x6a\xf0\x00\x1d\xf8\x79\x84\xe3\x61\xfc\xc4\x11\xce\x8e\x8f\x47\x5f\x41\x7e\x4c\xcf\xcf\x5f\x8f\x3e\x13\x34\xaa\x83\xa4\x7f\x36\x5a\x7e\x9b\xe5\x25\x24\xc9\xe0\x12\x97\x4d\x09\x86\xca\x92\x7d\x33\xb8\x41\xc3\x3d\xde\x09\xa3\x55\xf7\xb3\x53\xe1\x6f\xc5\xc0\x01\x68\x5b\xd0\x00\x2c\x05\x2d\x71\x00\x5b\xe1\xac\xb6\xe5\x00\xd0\x39\x72\x03\x90\x4e\xc7\xdb\xfc\x0f\x57\xa7\x32\xe6\x8f\x39\x25\x49\x5e\x5d\x6f\x86\x4a\x28\xb4\x41\xdf\xe5\x1c\xbc\xae\xd4\x50\xe9\xe3\x15\xad\x51\xea\xa2\x8d\x56\x94\x37\xf9\x05\x68\x1b\xd0\x15\x42\x62\xe7\x07\xbf\x93\x67\x17\xc8\x82\x2e\xa0\xa5\x06\xb6\xc2\xc6\x40\x7f\xcf\xba\xdc\x49\x7e\xcd\x13\x96\xae\x96\x5d\xc2\xb8\xd7\xef\xfc\x07\x96\x37\xc9\x60\x62\xf9\x65\x56\x68\xc3\x7e\x0d\x06\x5a\xf5\xf6\x1e\xd2\x30\xf5\x01\x10\x41\x1a\x8d\x36\xf8\x3d\x3d\xc7\x62\xe6\xf8\x27\xe2\xef\x27\x6f\xe3\xaf\x9f\xb9\xc6\x2f\xda\x20\x5f\x0c\x7e\xdc\xac\x3c\xf3\x4a\x74\x41\x17\x1d\x15\x2b\x92\x64\xe0\x6a\xc9\xa7\xe3\xb8\xd3\xa5\x13\x5b\x83\x2e\x75\xb5\x1c\xf2\xe9\xff\xe1\x59\x61\xdb\xd1\xac\xb0\xfd\x9a\x85\xa3\xd1\x99\xe7\x4d\xe1\x2b\x6a\x8c\xda\xef\xe5\x68\xd3\x81\xfa\xdf\xda\x3f\xba\x80\xc6\xef\x6b\xf3\x4a\x7b\x5b\xa2\x45\x27\x78\x19\xcd\xe7\xd3\xc3\x76\x58\x99\xeb\x02\xc2\x41\x3d\xed\xf9\x36\x75\xc5\xb6\x15\xda\xa7\x09\x38\xf2\x4c\xa4\x43\xb7\x69\x8c\x5e\xa1\x89\x7f\x1b\xc1\x61\x8c\x08\x5e\x12\x1b\xbe\xaf\xcc\xbe\x6f\x50\xd7\x7e\x3c\x3a\x79\x1f\x9d\x1d\xfd\x1b\x00\x00\xff\xff\x35\xf3\xc7\x11\x5b\x07\x00\x00")

func sampleIlxdConfBytes() ([]byte, error) {
	return bindataRead(
		_sampleIlxdConf,
		"sample-ilxd.conf",
	)
}

func sampleIlxdConf() (*asset, error) {
	bytes, err := sampleIlxdConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sample-ilxd.conf", size: 1883, mode: os.FileMode(436), modTime: time.Unix(1645988622, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"sample-ilxd.conf": sampleIlxdConf,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"sample-ilxd.conf": &bintree{sampleIlxdConf, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}