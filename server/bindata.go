// Code generated by go-bindata.
// sources:
// cards.chain
// DO NOT EDIT!

package main

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

var _cardsChain = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\xcb\x8e\xa3\x46\x14\x86\x29\x8f\xd5\xa3\xa0\x28\x8b\x3c\x41\x67\x3b\x9d\xb4\x31\xd8\x60\x76\xa9\xe2\x62\x30\x17\x83\x69\x7b\x30\x51\x16\x65\x28\x03\x85\x8d\x6d\xf0\xa5\x6d\x29\x52\x32\x99\xc9\x23\xe4\x35\xf2\x20\x79\x9b\x48\x91\xb2\x88\x9a\x88\x5e\x64\x35\x67\x77\xbe\xef\x9c\x7f\xf3\x7f\xdb\xfc\xf2\x06\x80\xaf\x14\x5c\x25\x68\xb3\x8b\x0b\x25\xc3\x79\x09\x9a\x0f\x0c\x00\x80\x6d\xe9\xaa\xa5\x35\x68\x7e\x63\x18\xe6\x9b\xe6\x53\x07\x80\xaf\x7f\xf8\xf1\x5d\x9d\xe1\x8a\x3c\xfe\xff\xd5\x6a\xd0\x7c\x64\x98\xb8\xf9\xf5\x0d\xe8\x34\x1f\x19\xf0\x16\xbc\x5d\x90\xaa\xce\x77\x25\xf8\x92\x01\x5d\x03\x57\x09\xe8\x32\xe0\xce\x3b\xad\x2c\x72\x6d\xd9\x17\x4f\xf9\x96\xd4\x47\xbc\xdd\xb7\x82\x9d\xe1\x32\x71\x4f\xdb\x15\xa9\x5e\x57\xaf\x22\xe7\x36\xdf\x54\xdb\xdb\x3b\x83\xe4\x69\x76\x04\x5d\x86\x61\x5e\x3a\xbd\xe6\x03\x00\xe0\xee\xcc\x3d\xf2\x8f\x5c\xe7\x05\xfc\xf9\x5d\x3b\x48\x1b\x9b\xee\xfd\x2c\x80\xf7\xde\x1c\xd9\xa6\x72\x6f\x69\xcb\x57\xc1\x3a\xa6\x89\x4c\x0a\x5d\x94\x16\x87\xac\xc8\xc7\xf2\x85\x43\xd0\xd7\x74\x08\xa7\x0a\xf4\x47\xb0\xf5\x4a\x6a\x29\xd0\xd7\xe0\xa5\x2f\x47\xb7\x90\xf3\xf5\xa8\x08\xd5\xec\x44\xbd\x59\xb9\x63\xa5\xc8\xf4\x8b\x3c\xa2\x83\x43\x08\xe9\x3e\xa7\x48\x47\x5a\x2f\x39\xcd\x9d\x64\x16\xf4\xa7\xb7\x50\x4e\x63\x53\xf7\x4b\x3a\xe3\x27\xb6\xda\xc3\x68\xe9\x8d\x38\x2c\x42\xab\xdc\x1a\x15\x82\x7a\xc1\xce\xd3\x78\xf3\xb0\x1e\x48\xcb\x5c\xf1\x6f\x39\xb7\x59\xad\x5c\xe8\xd8\x9b\xfd\x6c\xb7\xbd\x0a\x75\x99\x66\x72\x54\x51\x4a\x94\x69\x02\x47\x10\x53\x67\x31\xc1\xc6\x5e\x0e\x07\xb8\xd2\x91\x5a\x5a\x88\xed\x55\xd0\x3c\x18\x85\xfe\x7e\xaf\xd8\xe2\x93\x7f\xe9\x1d\xcc\xc2\x70\xd1\xb3\x67\xd7\xfc\xdc\xbf\x9c\x83\xda\x5c\xb8\x16\xa1\x76\xfd\x00\xd3\xf3\x60\x17\xdc\x64\x64\x19\xee\x82\x3c\x60\x8f\xde\xd0\x98\xad\xad\xf1\x2e\x3e\xd0\x43\x55\x47\xd7\xfe\xcc\x9e\xcb\x2a\xb1\x63\xbc\x7c\x0a\x0e\xc1\xd4\x19\x06\x19\x8d\xb7\x23\x4e\x1c\x8a\x83\x65\xb0\x77\x0b\x2c\x50\xba\xbc\x16\xce\xf5\xc8\x1d\x97\x3c\xaa\xfa\x1e\x3b\x35\xa2\xba\x0e\x75\xfe\x6a\x97\x46\x9a\x39\x53\x0d\x73\x1c\x3d\x07\x21\xdd\x9f\xd6\x9c\xa4\x0e\xe6\x56\x1a\x44\x81\x18\x4c\x42\x6b\xa2\x59\x0e\xf4\x32\x07\x5f\x4b\xcf\x16\x1d\x57\x7f\xe6\xb3\x01\xfb\xec\x9b\x2a\xf4\x21\x62\x5f\x4b\xd1\x5c\xf5\xb3\x5d\x81\x7f\xde\x75\x3f\xfd\xfd\x3e\xfa\xe9\x77\xf0\xef\xcf\x7f\xfd\xc1\x80\xef\x65\xbc\xee\x0f\x57\x82\x20\x12\x11\xcb\x62\x5f\x96\xf9\xd1\x50\x90\x92\xb5\xc0\xad\x78\x22\x62\x5e\x90\xc4\xa1\x28\xaf\xe3\xb5\x9c\x48\x44\x92\x04\x12\xc7\x24\x21\xe2\x50\xe4\xc4\x21\x2f\x63\x8e\x61\xfe\x0b\x00\x00\xff\xff\xff\xf8\x1d\x74\xe5\x02\x00\x00")

func cardsChainBytes() ([]byte, error) {
	return bindataRead(
		_cardsChain,
		"cards.chain",
	)
}

func cardsChain() (*asset, error) {
	bytes, err := cardsChainBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cards.chain", size: 741, mode: os.FileMode(384), modTime: time.Unix(1513927530, 0)}
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
	"cards.chain": cardsChain,
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
	"cards.chain": &bintree{cardsChain, map[string]*bintree{}},
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

