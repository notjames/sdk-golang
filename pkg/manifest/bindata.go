// Code generated by go-bindata.
// sources:
// pkg/manifest/data/package-manifest.schema.json
// DO NOT EDIT!

package manifest

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

var _pkgManifestDataPackageManifestSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5c\x6d\x6f\xdb\x38\xf2\x7f\xef\x4f\x41\x78\x03\xac\xf3\x6f\x6c\xa7\x8f\xff\x3b\x17\x8b\x20\xe8\x76\x71\x3d\x6c\x1f\x70\x69\x7b\xc0\xc6\x6e\x41\x4b\x63\x9b\x1b\x49\xd4\x91\x94\xd3\x74\xaf\xdf\xfd\x40\x52\x8f\x14\x25\x5b\x32\x7b\xc9\x75\xdb\x37\xbb\x11\x39\xd4\xcc\x8f\x3f\x0e\x67\x86\x94\xff\x18\x20\x34\x3c\xe2\xde\x06\x42\x3c\x9c\xa1\xe1\x46\x88\x78\x36\x9d\xfe\xce\x69\x34\xd6\x4f\x27\x94\xad\xa7\x3e\xc3\x2b\x31\x3e\x7d\x34\xd5\xcf\x7e\x18\x9e\x48\x39\xe2\x67\x22\x7c\x36\x9d\xd2\x98\xc7\xe0\x4d\x08\x9d\x9e\x4e\xee\x4f\x1e\x4f\xe3\xab\xf5\x38\xc4\x11\x59\x01\x17\x93\x74\x2c\x39\xae\x96\x15\x44\x04\x20\xc5\x63\xec\x5d\xe1\x35\xbc\x4c\x7b\xea\x56\x1f\xb8\xc7\x48\x2c\x08\x8d\x64\x9f\xac\x11\xad\x28\x43\x18\xa5\x22\xba\x6b\xcc\x68\x0c\x4c\x10\xe0\xc3\x19\x92\xe6\x20\x34\x8c\x70\x08\xf9\x5f\xf5\xe1\x5e\xe1\x10\x10\x5d\x21\xb1\x01\x44\x63\x35\x8c\xea\x26\x6e\x62\xa5\x12\x17\x8c\x44\xeb\xa1\x7a\xfc\x45\xb7\x1a\x43\x34\x8d\xfc\x73\xf1\x67\xd7\x17\x90\x28\x4e\x04\x2f\x8f\x7d\xc4\x60\x25\x7b\xff\x30\xf5\x61\x45\x22\x22\x47\xe5\xd3\x18\x33\x1c\x56\x45\x69\x22\x7a\xcb\xb2\x24\xda\x2d\xc7\x3d\x43\xd9\x2d\x30\xde\x8e\xc4\x7b\xdd\xc3\x86\x42\xc3\x3b\x20\x7c\x0f\x2c\x7d\xcd\x20\x7d\xd5\x90\xc1\xbf\x12\xc2\x40\x12\xed\xb2\x34\xb7\x03\x84\x16\xaa\x1d\xfb\xbe\x92\xc7\xc1\x9b\x32\x0f\x56\x38\xe0\x90\x32\x29\x7f\x45\xc1\x0f\x9f\xb0\x37\x0a\x8a\x92\xfe\x39\x21\xf3\xc6\x93\xa6\x49\x26\x0c\x3c\x41\xd9\x0d\x52\x78\x82\x00\x26\xcd\xc4\x91\x75\xae\xe9\xf2\x77\xf0\x44\xf1\xdc\xc2\xd7\x4c\xa7\xca\x83\xe6\xae\x2d\x7c\xcc\x9b\x6d\x4c\xcb\xfe\x7d\x39\x31\x87\x5a\xe1\x24\x10\xb6\x61\x6a\xf4\x56\x3d\xd1\x16\x07\x09\x3c\x95\x16\xe3\x25\xa7\x41\x22\x00\xc5\x58\x6c\xd0\x8a\xd1\x10\x31\x4a\x85\xc4\x23\xbe\x5a\x23\xca\x10\x83\x00\x0b\xb2\x4d\x7b\x90\x48\x00\x8b\x19\x08\xf0\x75\x6f\x49\x0e\x9f\x30\x44\x22\x74\xbd\x21\xde\x26\x65\x0b\x22\x1c\x49\x6a\x9e\x1c\x62\x18\xe1\x17\xe0\x31\xd8\xc7\xb2\x17\xab\x4c\x93\x74\x6a\x09\x47\x5c\x0b\x37\xaa\xb0\xa4\x34\x00\x1c\x19\x3a\x0c\x1a\xf4\x69\xa5\xea\xc0\x14\xcf\x45\x77\x32\x5c\x75\xaa\x2d\x13\x94\x51\x2a\xfd\x6b\x51\x59\xbe\x2b\x12\x40\xf3\x02\x28\x5a\x9b\x56\xc0\x2f\x24\x00\xa7\xe4\x97\xaf\xfc\xce\xfe\xdb\x66\xbf\x9c\x85\x6f\x82\xf8\x8a\x4e\x56\xe6\xc7\x49\x10\x3c\x63\xe0\xf3\xe1\xec\x8f\x5d\x64\x35\x40\x92\x62\x10\x09\x82\x03\x8e\x12\x0e\x3e\xf2\x13\x39\x09\x08\x27\x62\x23\x9f\x7b\x58\x6d\xfb\xd7\x44\xe8\x69\xe4\x34\x61\x1e\xa4\x8b\x83\x84\x78\x0d\x92\x10\xe5\xd8\x05\xb5\x2d\x89\x84\x03\x33\xe2\x18\xd4\xbc\x7b\x6a\x66\xd1\x00\x0b\xf0\x2f\x0c\x72\x94\x66\x63\x18\x63\xce\xaf\x29\xf3\x1d\x8c\x5a\x9b\xb0\x62\x2e\x2e\x2d\x66\xd8\x94\xc8\xa6\x68\xaf\x39\xaf\xcc\x63\x94\x84\x4b\x60\xcf\x68\xc4\x05\xc3\x24\xaa\x06\x3f\xb9\x27\xab\xf7\xea\xed\xa1\x70\x10\xbc\x5e\x99\xa8\x19\x0c\xf9\xfb\xc5\xeb\x57\xe8\x42\x45\xba\xe8\x52\x09\xa0\x2b\xb8\x91\x96\x2e\x46\x59\x94\x2c\x28\x0d\xf8\x84\x80\x58\xa9\xc0\x7a\x23\xc2\x20\x8d\xae\x3f\xdf\xc4\xe3\x52\xdc\x3d\x3e\x7d\xf4\x03\x07\x4f\x8e\x3d\x7e\x3c\x79\xf0\xe0\xb8\xb2\x1e\x73\xfd\x31\x63\xf8\xa6\xda\x44\x04\x84\x16\x9f\x69\x9f\xe2\x3a\x44\xe5\x75\x6c\xe5\x10\x8e\x6e\x3a\x42\x21\x05\xdc\x41\xf1\xf0\xee\x40\x01\x51\x12\x76\x41\x42\xf6\x77\x07\xc4\xe9\xc1\x40\x64\x12\xda\xf4\xdd\xf6\xae\x28\x0b\xb1\xb9\xa3\x0c\x69\x04\x8a\x10\x97\x95\xc1\xeb\xbb\x72\xb6\x2c\xa5\x5b\x59\x03\xab\x6f\x30\x06\x74\xff\xd0\x0e\x85\x2b\x77\xaa\x55\x44\x4b\x50\xee\xb4\x69\x04\x63\x83\xac\xb5\xa7\x13\x76\x69\x3c\x47\x85\x52\x46\xcb\xa2\x71\x93\x5b\x58\x11\x0a\xf1\x27\x12\x76\x23\x45\x2a\xe2\x8e\x17\x0d\xb4\x30\x27\xb9\xa2\x36\x89\x3a\xab\xad\x45\x9c\xa9\xfd\xa8\x8f\xda\x49\x20\x48\x1c\x40\x37\x7f\x54\x48\x39\x53\xfe\x7e\x0f\xe5\x23\x5a\x5b\x49\x6d\x5a\x47\x54\xb8\xa3\xc8\x63\x43\xdf\xce\x1e\xb1\x6c\x48\xb6\xfe\xf7\x36\x45\x09\xb8\x33\xa6\x89\x39\xff\xad\x0d\xa1\x53\xe8\x6a\x09\x63\x9a\x53\xb1\x72\x7b\x53\x5c\xfa\x4a\x3b\x46\x97\xe9\x58\xca\xda\x3b\x97\x90\x35\x6f\x57\x2e\xd2\x9f\x74\x87\xe9\x9d\x00\x19\x0a\x78\xd6\xe0\x34\x6f\xee\x41\xb9\xbb\x95\x64\x55\x27\xa1\x9a\x66\x71\xea\x5d\x81\x68\xe6\x75\xb9\x7d\x27\x4b\x8d\xd9\xba\x50\xb2\xad\x7c\x6f\xe2\xb5\x7e\xed\x2d\xf1\xba\x3b\x21\xb5\xba\xdf\x44\x46\x9e\x22\x6f\x27\x8b\x02\x6f\x57\x2e\x57\xef\xf5\x3d\x97\x33\x4b\xe8\x35\x88\xfe\xb4\xb9\x5c\x0f\x28\xbe\x95\x5c\xae\xee\x8b\xba\xe4\x72\xa6\xb3\x4d\x62\x60\x1c\x54\x49\xb3\x62\xbd\x96\x76\x64\xff\xff\x1b\xc6\x77\x4d\x28\x7d\x2c\x60\x2c\x48\xa5\xca\xb4\xc7\x0c\xe6\x62\x48\x5b\x73\xa8\x15\x93\x87\x66\x1e\x60\x9b\x98\x0e\x39\x69\x61\x57\x7b\x56\x7a\xb2\x2f\x4e\xd2\x0b\xb3\xb1\x2a\x46\x8e\xe5\xba\xd9\x05\xd7\x39\xd2\x22\x69\xfd\x92\xc1\x0a\x18\x44\x1e\x20\xcc\x91\x5a\x6e\xe0\xa3\xe5\x0d\xba\x5c\x13\xb1\x49\x96\x13\x8f\x86\x53\x2d\x30\xf5\x89\x34\x77\x99\xc8\x91\xa6\xb9\x5c\x81\xf0\x0e\x09\xc1\x00\xb2\x86\xfb\x93\xfb\x0f\x8b\x21\xdc\x02\x6c\x02\xe2\x06\x67\x08\x31\x09\xba\x71\x51\x89\xb8\xe3\xe1\x03\xa7\x30\x69\x7b\xdc\x60\xb3\xa1\x5c\x18\x05\xe1\x3d\xe0\xc9\xa4\xdc\x21\xf4\xd0\x29\x42\xb9\x55\x6e\x40\x22\xf1\xf6\x51\x37\x80\xa4\x84\x3b\x70\x1e\x39\x05\x47\x59\xe3\x0c\x98\x27\x9d\x81\x79\xe2\x0e\x98\xc7\xae\x81\x79\xe2\x08\x98\x84\x91\x6e\xb8\x24\x8c\xb8\x83\xe5\x89\x53\x58\xa4\x2d\x6e\x50\xe1\x10\x6e\xf7\x28\x34\x9f\x23\x0e\x21\x8e\x04\xf1\x50\x7a\xb3\xc5\xdc\xe0\xf4\x40\x12\x15\x8d\xd6\x6c\x3a\x2d\x1e\x4d\x9d\x5a\x9f\xea\xdc\x0e\xc0\xc0\xd6\x62\xd4\xa1\x7f\x85\x68\x2d\x36\x1d\x4b\xba\x5a\xc8\x59\x5c\xfb\xa4\x21\xac\xb5\x15\xf0\x4b\x45\xe8\xfb\x76\x9b\x32\xf5\x6e\xd7\x26\x33\x5a\xdd\xd3\xa6\xd3\x93\xaa\xca\x59\x95\xeb\xd4\x6a\xeb\xdd\x2f\x0e\xb7\xa4\x58\x7f\x86\xe2\x70\x8f\x0c\x33\xc6\x42\x00\x33\x0b\x49\xad\x60\xa4\x22\xce\xe0\xf8\x4b\x03\x1a\x16\x37\x55\x24\x88\x43\x06\x6b\xf8\xd4\x76\xf4\xbf\x77\xc1\x5b\xbf\xa7\xa5\x30\x58\x6a\xef\x5c\x18\x54\xb2\xfd\x0a\x83\xda\xfc\x3b\x5b\xf0\xfe\x4a\xf5\x45\x0d\xd8\xad\x16\xbc\x5b\x97\xd1\x1d\xab\x61\x56\x26\xc1\xb8\x57\x64\x32\xda\x40\xfc\x4d\x9f\xd3\x19\xbd\xf6\xdf\x34\x70\xf6\x72\xfc\x71\x82\xc7\x9f\xcf\xc7\xbf\x9d\x8e\xff\xba\xb8\xd7\xf3\x50\xde\x3e\x2b\xf9\x1d\xd8\x4e\x61\x97\x7d\xac\xe2\x3a\xa1\x83\xc1\xca\x07\x62\x0e\x86\x2b\x9f\x43\xb8\x18\xae\xe4\xbd\xf6\x89\xd7\x0e\xf0\xa3\xde\xda\xee\x3f\xbd\x75\x8b\x7f\xc4\x32\xc2\xf5\x70\x10\xa0\x35\xc3\xf1\x26\xe7\xe2\x53\xc4\x01\x50\xb6\xb5\x40\x34\xb9\x26\x57\x24\x06\x9f\xe8\xdb\xff\xf2\xaf\xe9\x33\x1c\x04\x1f\x95\x58\xf1\x82\x3a\xc7\xf6\xb8\x54\xe6\xd1\x48\x60\x12\x01\x93\x23\x5a\x63\x86\x3d\x06\xa1\xf1\x21\xd2\x72\xb5\x06\x01\x04\x87\x8c\xc1\x81\x11\x6c\x8e\x60\xf5\x0d\x55\x83\x6d\xb3\x56\xed\xd1\xfb\x5c\x23\x1f\xa6\xcb\x46\xe6\x85\xe6\x65\x40\x1b\x77\x9e\xd1\x30\xc4\x91\x8f\x58\x12\xc9\x7c\x08\xa3\xfc\x5d\x4f\x11\xdd\x02\x63\xc4\x07\x8e\x70\x74\x83\x38\x08\x84\x85\xda\x5f\x74\xd1\x30\x80\x2d\x58\x8a\x61\xcd\x71\x18\x6a\x8e\xc5\x5a\xe6\xa3\xed\xae\x62\x75\x76\x6a\xb3\xad\x0d\x26\xcc\xba\x6d\x35\x4c\x42\x13\x50\xd9\xc7\x01\x04\x38\x22\x91\x82\xa1\x98\x96\x9a\x70\xbb\x8b\xcf\xbb\x7d\x18\x5d\x6a\x4f\xbf\x98\x1d\x9f\x49\xbf\x3f\x9f\x4f\x4b\xae\xff\xc8\x2a\xd5\xb8\x07\x64\xff\x6c\x22\x36\x93\x46\xd7\x24\x08\xd0\x12\xd0\x92\x26\x91\x8f\x04\x45\x1c\x87\xf9\x8d\xe6\xec\x46\x6b\x3d\x07\xae\x41\x18\x25\x41\xad\x8a\xa8\xff\x7d\xb1\xcb\xee\xab\xa1\x4f\x18\x62\xb0\xd2\xf7\x6b\x2b\x5a\xed\x56\xca\x16\x5a\xe5\x6a\x59\x9e\x2e\x6a\xcf\xcc\x5e\x35\x63\xf6\x8b\x55\x2c\xa2\x43\x88\xb6\xef\xb1\x13\x5e\x3e\x8f\xb6\x84\xd1\x28\x84\x48\xa0\x2d\x66\x04\x2f\x03\xa7\x0c\xbd\xfc\xf0\xd3\x2d\x10\x91\x44\x88\x7b\x34\x56\x27\x13\xe8\x7a\xaa\x89\x19\xe1\xf0\x56\xd9\xa8\x19\x95\x86\x2a\x19\x2f\xd3\x58\x5b\x7d\x3b\xd0\xac\x5b\x5f\xd7\x96\xea\x7d\xdb\x74\x95\xa1\x9e\x0b\xb2\xfe\x42\xdc\x92\xf3\xbb\xfb\x6c\xd5\x50\x7d\xf7\xf1\x4d\xfa\x4f\x15\x81\xf4\x21\x64\x4b\xc4\x84\x74\x86\x68\xd6\xb4\xf2\xa6\xda\x25\xe6\xec\xcc\x54\xd0\xfc\x4b\x10\x2b\xac\x07\xac\x7e\x0b\x01\x6c\x9f\xba\x58\xde\x66\x46\xc6\xb9\x50\x8f\x69\xb2\xa6\xcd\x15\xc8\x8c\xa7\x0b\x77\x33\x6d\xf9\x54\x06\x59\x66\x43\x7d\xfa\x5b\xf1\x2a\xc8\xc3\x91\x5c\x9d\xf9\x01\xb3\x2a\xf8\xab\xaf\xb3\xa8\xd8\xc8\xf6\xac\x27\x3f\xec\x7b\xac\x98\x32\x7b\x61\xc4\x2c\x15\xc8\x7e\xa9\xb3\xc8\x3f\x12\x2b\xd4\x15\x54\x3d\xd8\x50\xde\x52\xaf\x69\xa4\xf5\x7e\x0e\xf3\x52\xf9\xc5\xd1\x58\xff\xf7\xf8\x6c\x24\xbc\xf8\xdf\x89\x1f\x1f\x9f\xed\x49\xfa\xbf\x51\x2e\x90\x34\x78\xc4\x8f\xa5\xc6\x4b\xa2\x3c\x9f\x9d\xf6\x3b\x8e\x4a\x50\xb5\x74\x5a\x53\xae\x0f\x53\x7b\xd3\x4c\x17\x0c\x7a\xed\x71\xfb\x62\x3f\x6b\x2e\xea\xe4\x9d\x6a\xd9\x59\xc6\x8e\xf4\xca\x20\xf6\x7d\x06\x9c\xa3\x10\xc7\x31\xa8\x3d\x07\x67\x4d\xb6\x6b\x1f\x68\x1f\x87\xfe\x35\x51\x15\xfe\x73\x66\x66\xad\x2e\x41\xfd\x30\x69\xde\xe5\x9b\xb1\x14\x3e\x30\x86\x62\x06\x2b\xf2\xa9\x0a\xa5\x0e\xe6\xee\x28\x94\xaf\x93\xd6\x5a\xf2\x2d\x41\x49\x13\xf1\x3f\x06\xe5\x35\x65\x57\x3f\xd7\xbe\xca\xb7\x19\xfa\x4f\xca\xae\xa4\x15\x7e\xe9\x97\x01\xc4\x06\x8d\xaa\x95\x91\xd2\x79\xb2\x0a\x02\x76\x9f\x1a\x0f\x9a\x2c\xad\x96\xc2\x1b\xf7\xdd\x34\xfe\x29\x3d\x5b\xb8\xa8\xa1\xdb\xcb\xe3\x45\x9c\x3e\x30\xde\xb5\x7f\x59\x33\x2d\xeb\xd9\x6a\x64\x69\x53\xef\xe2\x18\x8d\xcd\xaa\x58\xcb\x82\x68\xad\x98\xc5\x57\xeb\x7a\x48\x95\x8e\xb6\x57\x40\xe9\x30\x9e\x44\x52\x99\xde\xbb\x6a\x71\xc2\x97\x30\x32\xce\x23\xa0\xef\xf1\xa5\x25\x93\x30\x7f\x9e\x25\x6f\x31\xcf\xd6\xa4\xa9\x38\x20\x9f\x81\xa3\x17\xaf\xde\xbc\x7b\xfb\xf1\xd5\xf9\xcb\xe7\x3a\x92\x7b\x7f\xfe\xeb\xbb\xe7\x32\xbb\x4a\x6f\xcf\xfe\x58\x74\x98\xe9\xc6\x1f\x27\xe8\xc5\x2a\xeb\xc7\x91\x4c\xf9\x4e\x10\x11\xe8\xe5\xbb\x8b\xb7\xea\xdb\x47\xce\x93\x10\xfc\xb4\xc7\x4f\xe8\x68\x54\x0c\xd1\xe2\x4f\x0e\x8d\x49\x5a\x8f\x9a\xf2\x6e\x3d\xf3\x65\xf7\xf9\xed\x81\x65\x94\x8e\x6f\x6b\xfb\x2a\x29\x1f\xf3\xb6\x33\xe1\xfa\x6f\x04\xe5\x4d\x2d\x04\x2e\xa8\xfb\xfa\xdd\xdb\x9c\xcb\x25\x02\x6b\xea\x96\x1a\x35\x81\x2b\xbd\x5b\x68\xac\x3a\x48\x16\x97\x04\xbe\xd3\xd8\x1c\xf1\x6e\x95\x58\x1a\xc3\x90\x14\xb7\x1d\x07\xce\x65\x27\xae\xb6\xd2\x6e\x57\xfb\x2a\xe2\x72\x99\xb7\x5c\x8c\xfb\x7a\xb1\x0e\x8d\x0f\x08\x72\x2a\xa7\x8f\xb6\x50\xa7\xd2\xa1\x77\xc0\x93\x8d\xd2\x14\xf6\x1c\x7c\x01\xca\x6b\xfa\xc8\x64\x4f\x10\x73\x05\xfb\x43\x59\x3a\x84\xb5\x9e\x86\x17\xcd\xbd\x61\xd4\x63\xdc\x59\x10\x53\xf5\xfa\x43\x68\xd9\x15\x6d\x50\x5a\xba\xed\xb8\x87\xc5\x93\x38\xa6\x4c\xc8\xff\x3d\x1a\x29\xc7\x8e\x78\xb2\xe4\x82\x08\xf5\x89\x07\xe2\x37\x91\xc0\x9f\xd0\xf5\x06\x18\xe4\x3d\xb2\x5a\x39\x83\x38\xc0\x69\x2d\x4e\x6c\x40\x9f\x9b\x20\xba\xd2\x3b\x06\x16\x88\x25\x51\xe5\x83\x9f\xd6\xdf\xd6\x4b\x7f\x5b\xce\x4e\x11\xd5\xd4\x64\xcb\xf6\xc1\xe4\x74\x72\x5a\xbf\x1e\x3c\xca\x2e\x4c\x54\x2f\x02\xf3\x18\xbc\xa9\x96\x99\x6c\x44\x18\x1c\x37\xea\x67\xde\xf2\x91\x4d\x1f\x46\x69\x4d\x6b\x3e\x9f\x58\xfe\x77\x74\x36\x1b\xcd\xe7\xaa\xee\x75\x3e\xfe\x0d\x8f\x3f\x8f\x17\xf7\x46\x67\xb3\xf9\x7c\x52\x79\x74\xfc\x7f\xc7\xc7\x67\xea\xf9\xbd\xd2\xf3\xf9\x7c\x3c\x9f\x4f\x16\xf7\x8e\xcf\x8e\x4a\xbf\xb1\x37\xf8\x32\x18\xfc\x27\x00\x00\xff\xff\x19\xb4\xb9\x25\x03\x52\x00\x00")

func pkgManifestDataPackageManifestSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_pkgManifestDataPackageManifestSchemaJson,
		"pkg/manifest/data/package-manifest.schema.json",
	)
}

func pkgManifestDataPackageManifestSchemaJson() (*asset, error) {
	bytes, err := pkgManifestDataPackageManifestSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/manifest/data/package-manifest.schema.json", size: 20995, mode: os.FileMode(420), modTime: time.Unix(1499953039, 0)}
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
	"pkg/manifest/data/package-manifest.schema.json": pkgManifestDataPackageManifestSchemaJson,
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
	"pkg": &bintree{nil, map[string]*bintree{
		"manifest": &bintree{nil, map[string]*bintree{
			"data": &bintree{nil, map[string]*bintree{
				"package-manifest.schema.json": &bintree{pkgManifestDataPackageManifestSchemaJson, map[string]*bintree{}},
			}},
		}},
	}},
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
