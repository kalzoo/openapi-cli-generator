package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _templates_commands_tmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x59\x4b\x73\xe3\x36\x12\x3e\x93\xbf\x02\x61\x79\x52\xe4\x58\xa6\x26\xd9\xd4\x1e\xb4\xd1\x56\xcd\x78\x1e\x9e\xaa\xb1\xc7\xeb\x47\x72\xf0\xba\x32\x10\xd9\x92\x50\x06\x01\x1a\x84\xfc\x88\xc2\xff\xbe\xd5\x00\xf8\x12\x29\xd9\x9e\xdd\xcb\x5e\x6c\x1a\x0d\xa0\xbb\xbf\x7e\xc3\xe3\x31\x39\x94\x29\x90\x05\x08\x50\x54\x43\x4a\x66\x8f\x44\xe6\x20\x68\xce\x0e\x12\xce\x0e\x1c\x41\xaa\x98\xbc\xff\x4a\x4e\xbe\x5e\x90\x0f\xef\x3f\x5f\xc4\xfe\x78\x4c\xce\x01\xc8\x52\xeb\xbc\x98\x8c\xc7\x0b\xa6\x97\xab\x59\x9c\xc8\x6c\x7c\x43\xf9\x9f\x52\x8e\x07\x2f\xf1\xfd\x9c\x26\x37\x74\x01\x24\xa3\x4c\xf8\x3e\xcb\x72\xa9\x34\x09\x7d\x6f\xbd\x26\x6c\x4e\xe2\xcf\x66\xa1\x88\x3f\x66\x9a\x94\x65\x30\xcf\x74\xb0\x5e\x13\x10\x29\x29\xcb\xde\xa6\x73\xad\x98\x58\x14\xb8\xb1\xb0\x9f\x3b\x36\x5f\xb0\x0c\x70\xa7\x66\x19\xb4\xb6\xf9\x5e\xf0\x4c\xd9\xc7\x09\x67\x41\x77\x7b\x7e\xb3\x18\x83\x52\x52\x15\x1b\x04\x55\x8c\xff\x04\x25\xb9\x5c\x8c\xb9\x5c\x6c\x10\x8b\x7c\xfe\xd3\xdf\xc6\x89\x9c\x29\x3a\x48\xb9\x63\x39\x28\x43\x91\xf9\xcd\x22\x66\x62\xbc\xfc\x59\x48\x31\x5e\x80\xd0\x1c\x32\x2a\xe2\xbb\x9f\x03\x3f\xf2\xfd\xf5\x9a\xa4\x30\x67\x02\x48\x90\x53\x45\xb3\x22\x70\x8a\x1f\x10\x45\xc5\x02\x48\xfc\x35\xd7\x4c\x0a\xca\x4f\x0d\xd9\x50\x0d\x99\xcd\x09\xdc\x92\xf8\xe2\x31\x07\x12\xcc\xa4\xe4\x40\x85\x3d\xec\x79\x49\x96\xc6\x1f\x39\x5d\x14\x61\x14\xbf\x93\x92\x87\x88\x56\x7c\xf8\xe5\xf3\x09\xb5\x08\x8e\xc8\x9c\xf2\x02\x46\xc4\x10\xde\x43\x91\x28\x66\xf8\x20\x31\x72\x1c\x80\x17\xd0\x65\xc3\x84\xfe\xfb\x2f\x43\x4c\x3e\x23\x61\x80\xcb\x9b\x97\x72\x98\x73\x49\xb7\xf0\xf8\x68\x49\x43\x5c\xe2\xe7\xf0\xe9\xdf\x68\x9d\x6f\xe0\xc2\x20\x78\xe2\xbe\xda\x3f\x0f\x1a\x27\x6c\xd9\xec\x77\xca\x34\x28\x67\xac\xbe\x31\xee\x29\xd3\x07\x78\xbd\xdd\xb7\xdd\x30\x8e\x7e\xbe\xc4\x08\xb3\xfc\x3b\x2c\x13\xce\xe2\x73\xd0\x87\xab\x42\xcb\xcc\xf2\x48\xb2\x34\xf2\x7d\x8f\xcd\x49\x9b\xef\x11\x2d\xdc\x27\x59\xfb\x9e\x67\x5d\x2d\x7e\xc7\x44\x7a\x5a\x1f\xab\x36\x47\xbe\x87\x77\xcf\xa5\x22\x7f\x8c\xc8\x82\xcb\x19\xe5\x48\x22\x93\xa9\xd3\xaf\x59\x2b\xf0\x3a\x42\x48\x87\xd9\xdb\x34\xc5\xcf\xb0\xd9\x66\x48\xf6\xde\x56\xd8\xae\xd7\x64\x4f\xa0\xde\x93\x29\x89\x1d\x00\x66\x91\xe6\xcc\xac\x7d\x92\x1b\xab\xa7\xab\x19\x67\x89\xa1\xd9\xcf\x66\x87\x7f\x47\x15\xa9\x0e\x97\xe5\xf9\x6a\x96\xc8\x2c\xa3\x22\x25\x18\x1b\xbe\x3f\x5f\x89\xa4\x4d\x07\x75\x07\x0a\xf1\xb8\xba\xce\x68\x7e\x65\x93\xcf\xb5\xfd\x85\x4a\x29\xd0\x2b\x25\x86\xa8\x6b\xe3\x04\x0e\x8a\xbd\xc2\x5c\x64\x44\x72\x77\x3a\x47\x1b\x3c\xe7\x79\x41\xda\xb8\x54\x30\x31\x66\x76\x77\x6c\x3a\xdb\xc8\xee\x5f\x29\xbe\xb1\xef\xf2\xec\x4b\x4d\x2f\x47\x56\x9a\xca\x23\x4b\xdf\x02\xeb\xa4\x93\x39\x26\x3e\xbc\x10\x05\xfc\x5a\xfd\x65\x65\x1c\x8f\x49\x17\xd7\xb2\x44\xa7\xab\x31\x45\x6a\xed\x7b\xbe\xd7\x46\x70\xf8\x40\x58\x33\x8e\xcf\xe0\x76\xc5\x14\xa4\x75\xd6\xea\xde\x6c\x11\x19\x91\x5a\x72\xeb\x91\xe4\xb5\x49\x9c\xf1\x6f\xf8\xd3\x25\xff\x43\x2a\x8e\xe8\x1d\xbc\x93\xe9\x23\x29\xcb\x11\x99\xe1\x87\x43\xb4\x3a\x1d\x91\xf0\x75\x93\x5a\xcf\xa0\xc8\xa5\xc0\x20\x42\xa6\x67\xc6\x90\x26\xbb\xe0\x71\x93\xeb\x6d\x14\x2c\xa9\x48\x39\xa8\x53\xaa\x97\x08\x8f\x89\xb8\x23\xbb\x56\x85\xa3\xef\x61\x24\x0d\xba\x95\x31\x67\xfb\x0a\x7b\x83\xf5\xe8\xb2\x24\x01\xd9\x27\x2d\xb2\xef\x99\xa8\xf2\x1a\x7f\xc1\xe0\x3d\x5b\x89\x43\x29\xe6\x6c\x11\x7f\x02\x7d\xaa\xe4\x9c\x71\xc0\x08\xca\xd9\xe5\xd9\x17\xdc\xbe\x52\x1c\xf7\xda\x53\xfb\x46\x44\xc3\x0c\x65\x73\xc9\xc8\x99\x9a\x8d\xc8\x9e\x41\xd1\x98\xba\x87\x3f\x4a\xdb\x94\x0d\xbb\x33\xfe\x2c\xb0\xec\xe8\x65\x95\x6e\x0d\xbb\xa9\x43\xb7\x88\xcf\x20\xe7\x34\x81\x70\xa5\xb8\x49\x48\xdf\xd6\xdf\x8c\x21\xdd\x69\x07\xd2\x7a\xfd\xad\xfc\x66\x52\x57\x43\xaa\x2d\x3d\x22\x3f\x45\x15\xeb\xca\x47\xbb\x39\xcc\x53\x70\x5b\xc1\x71\xc8\x19\x08\x1d\xa3\x96\xc7\xa0\x97\x12\xb7\x84\x11\x3a\x3c\xca\x10\xf9\x9d\xc8\x7b\x96\xc2\x7d\x7d\x6f\x57\xa0\x1e\x6b\x85\x91\xf9\x94\x28\xb8\xc5\xac\xf5\x2f\x24\xd9\x72\xd0\xe4\xe3\x01\xa5\x9c\x46\xed\xf2\xd5\x62\xb0\x04\x9a\x82\x1a\xe6\x70\x64\x68\x2f\x61\xd1\x60\xd6\x82\xec\x89\xf6\xc0\xeb\x06\xdb\x64\x4a\x5c\xc2\xff\x04\x1a\x49\x26\x18\xfe\x22\x9a\x69\x6e\xc2\x76\xb3\x00\x1a\xe6\xd6\xef\x5b\xd7\xfc\x30\x25\xd5\xe1\x13\xc6\x4d\x7e\x30\x0a\xb6\x9a\x91\x3e\xbe\xcf\x00\x78\x9e\xe9\xf8\x3c\x57\x4c\xe8\x79\x18\xbc\xba\xb3\x78\xb4\x90\x88\x6a\x2e\xed\x76\x61\x00\xe9\xe7\x40\xfd\x02\x66\x15\xf2\x5e\xdf\x65\x07\x53\x93\x03\xcd\xe4\xa7\x1f\xa6\x24\x08\x1c\x3e\x83\x52\x1d\x4a\xa1\x41\xe8\x03\x44\xb3\x6a\x36\x8e\x21\x65\xd4\x25\xaa\x00\x7b\x85\xf4\xd1\x75\x28\x78\x67\xd4\x88\xd2\x92\x04\xa3\xc6\xe6\xac\x77\x30\x97\x0a\xc2\x56\xca\x19\x39\xb3\x8f\x90\x79\x64\x43\xad\xc8\x4d\x0a\x34\xb5\x1c\x6e\xe3\xf7\x32\x8c\x6c\x8e\xc3\xc5\x1f\xa6\x44\x30\x6e\xc5\x76\xf5\x4f\x30\x3e\xb2\x3f\x6c\x93\x1c\xff\xae\x68\x1e\x82\x52\x23\x12\x60\xc8\x41\xa1\xc9\x9c\x32\x0e\xa9\xf1\x1a\x23\x13\x56\xe2\x14\x12\x99\x42\xda\xcf\xc0\xbe\x65\x87\x92\xc4\xe7\x9a\xea\x55\x61\x86\x97\x5f\xc9\x2f\x6f\xde\x58\xce\x4e\x18\x97\x12\x2e\x45\x46\x55\xb1\xa4\xbc\xca\xea\xa1\x55\xe2\x47\xc7\x21\xfa\x47\x4f\xf4\xe7\xc8\x5e\x5f\xcb\xb1\xde\x2b\x77\x77\x5b\x15\x83\x75\x69\x7d\x6e\x27\x22\x1f\xf0\xd7\x3c\x0c\x8e\x2e\x2e\x4e\xc9\xab\x74\x42\x5e\x15\xc1\x68\x53\xc1\x7a\xc1\xd8\x33\xaa\xb1\xa2\x73\xdd\x54\x03\x6b\xc8\xb7\xb8\xb4\xcd\x8e\xa8\x7a\xa5\xb9\x45\xd2\xde\xd0\xd6\xbf\xc2\x7e\x6a\x69\xd6\x59\x05\x74\x0c\x81\xad\x3c\xa8\x39\x4d\x60\x5d\x62\x00\xc5\x61\xcf\x52\x51\x3b\xfd\xb8\x4c\x6d\x10\xe8\x48\x61\xb0\xc0\xbe\xa3\xdb\xd8\xb9\x1c\x7d\x6f\x1a\x58\x93\xa0\xdb\x3d\xf1\x8b\x5a\x89\xba\x6b\xf9\x5f\x34\x15\x91\xb5\x9a\x01\x8a\x6a\x0d\x59\xae\x51\xba\x37\xbe\x67\x1a\xde\x6a\xe9\x57\x23\x9d\x95\x3e\x7e\x6b\x17\x8b\x3a\xe5\xb9\x5d\xfb\xfb\xbe\xf5\x8b\x0e\x1c\xce\x77\x87\xb4\x6b\x34\xf9\x6f\xf4\xec\x29\x18\xb5\xa2\x66\x28\x0e\xfa\xde\x7f\x28\x57\x3c\x25\x42\x6a\x92\x50\xce\x89\xb3\x52\xdd\x2c\x56\xfe\x8f\x3f\x31\x98\x69\xa2\x57\x94\x93\x96\xcb\x54\x94\x8c\xea\x64\x69\x3b\x6c\xaf\x5d\x9b\xcd\xba\x33\xfc\xb1\xfd\xae\x6a\x93\x67\x6f\xb3\x40\x59\xbf\xff\x04\xda\x6c\xfa\x8d\xf2\x95\x8d\xef\xd8\xe4\xc7\x07\xed\x32\xe3\x39\x70\x48\xb4\xcd\xe0\xae\x94\xbd\xe5\xfc\x1c\xb4\xc6\x3e\x25\x8c\x3a\x31\x31\x8c\xc5\x73\xc0\x58\x80\x26\x95\xe4\x77\x28\x8b\x05\xc2\x21\xe1\x19\x52\x5b\x6e\x23\xb4\xad\x30\x17\x98\x08\x8d\x7c\x57\xd7\xb3\x47\x0d\x26\x9c\x3e\x3c\xe4\x90\x68\x48\xc9\x5f\xc4\x96\x1c\x12\xbc\xba\xc5\x68\x8b\x46\x0e\xd3\xef\x91\xf7\x77\x27\xa1\xc5\x1e\x33\xd6\x4a\xd5\x92\xd6\x35\xd2\x52\xdd\x5d\x75\x1f\x64\x32\x12\x8e\xd9\xee\x54\x5d\x3a\x37\xd8\x55\x39\xcd\x06\x2d\x49\xa8\x40\x7c\x14\xd0\x64\x49\x52\x28\xd0\x39\x49\x61\xae\x9a\x41\x42\x57\x05\x90\x57\x05\x61\x85\x4d\x7d\x3d\x93\xed\xc6\xa2\x16\xb1\x35\xa5\x7b\x9e\x37\x53\x40\x6f\x1a\x5a\x5d\x8d\xbd\xb2\xdb\x1a\xe1\x5f\x9a\x65\x10\x9f\x73\x80\x3c\xb4\x53\x3b\xa7\x58\x91\x5f\xdb\x75\x48\xa4\x48\xeb\x8c\x8b\x29\xd3\x45\xf9\x3f\xa7\x3b\xc3\xbc\x0b\xc9\x09\xdc\x87\xc1\x31\x7d\x60\xd9\x2a\xab\x6e\x28\x08\x3c\x24\x00\x69\xbb\xfa\x35\x65\x62\x23\x2b\x6e\x4c\xa0\x67\xb0\x60\x05\x66\xfa\xa2\x3b\xaa\x8e\x3a\xe3\xf5\xd5\xb5\x09\x90\x7a\xc5\x4c\x2e\x4a\x4a\x5d\x4f\x10\x52\x6a\x3b\xef\x17\xdd\xd9\xc4\x6c\x9a\x92\x1f\xcd\x23\x55\x7c\x68\x29\x46\xaf\xcb\x02\x26\x9d\x59\xc5\x8e\x92\x66\xd2\xb3\x84\xf8\xc2\xb5\x86\x96\xf2\x45\x8a\xc5\xc4\x79\xbc\xba\x49\xe5\xbd\x08\x07\x5f\x47\x46\x7e\xdd\xa1\xf4\xe7\xa5\x29\xd1\x6a\x05\x7e\xbb\xa4\x56\xf2\xbb\x21\x73\xba\xc1\xbb\xbd\x03\x45\xa8\xa3\x6e\x97\x0c\xf6\xf1\xc2\x35\x68\x9d\x77\x18\x34\x33\xa2\xb6\x15\x11\xdc\xd0\x85\x02\xcf\x13\x53\x17\x48\x02\x4a\x53\x26\x08\xdc\x81\xd0\x44\xaa\xda\xfd\xb1\xeb\x22\xd6\xe8\x4c\x2c\xda\x80\x05\xef\xb8\x4c\x6e\xd0\x47\x20\x59\x19\x01\x11\x87\x55\x01\x05\xc9\xa5\x6d\x3c\xb4\x24\x39\x28\x26\x53\x86\x89\xf8\x91\x24\x4b\x48\x6e\xbe\x83\x63\xe9\x0c\x8e\x2d\xa6\x53\x2c\x44\x75\x36\x46\xa6\x2d\xe5\xd8\xb3\x05\xd9\x3d\x0e\x55\xcf\x43\xb8\xcd\x56\x4e\x74\x7d\x1b\xa6\x49\x96\x6e\x81\xb0\xe5\x56\xf1\x65\xd1\xf8\x4e\x3d\x1f\xc4\x6f\x39\xa3\xa8\x7b\x1d\xe1\x6e\x61\x42\xae\x3a\x2f\x24\x5e\x67\xbe\xe9\x9d\xf2\x3c\xc3\xa3\xc5\x60\xb3\x5f\xaf\x9e\x45\x06\x08\x6d\x1f\xaf\x5f\xd5\xdc\xde\x6d\x5e\x6e\x5c\xaf\x72\x6f\x14\x5b\x2d\x8a\x09\xb1\x08\x1c\x33\x81\xf9\xe0\x04\xd7\x30\xf5\x70\x10\x3b\x0b\x79\x75\xc7\xd9\x4a\x4c\x08\x82\x1e\x22\xa2\xaf\x3b\x70\x8e\x08\x55\x26\xf2\x2d\x28\x95\x51\xda\x8d\xf1\x33\x5b\xa7\xbd\x87\xce\x88\xbc\x43\x2e\xe4\x78\x85\xb7\x3e\x90\xb2\xbc\xee\x77\x18\x03\x4d\xb6\xe7\x79\x5c\x2e\xe2\x8f\x54\x53\x1e\x46\x58\x31\xb0\x3e\x45\xf1\x71\xb1\x08\x03\x53\x3f\x4c\x5f\x81\x1e\x1a\x55\x56\xf1\xdb\xc6\xb1\x7f\xe1\x9e\xb6\xd7\xba\xf7\x4b\x9b\xe2\x31\xc9\x72\x53\xb0\xaa\xc7\xf1\x46\x89\x6a\x38\x0b\xa3\xee\x0b\x58\xbb\x32\x3c\xf3\x21\xac\xeb\xfe\xc3\xde\x5f\x75\x3c\xf0\x40\xb3\x9c\x43\xe1\xba\x4d\xbf\xdb\xf7\xc0\x83\xb9\xff\x43\xb5\xc9\xf9\x5d\x7d\x68\x7f\x4a\x02\x62\x9e\x86\xea\xcc\xe6\x14\xc7\x46\x3f\x8c\xc8\x3e\x09\x8c\x75\x6b\x79\x5d\x30\x99\x45\x40\xeb\xfc\x5b\x04\xfd\x02\xb8\x23\x2e\xb7\x84\xe5\xb6\xa8\xdc\x1a\x94\x3b\x63\xb2\x17\x92\x9b\x81\x57\x8e\x06\x06\xeb\x5d\xe1\xf8\xcc\x68\xac\xd4\x38\x62\x69\x0a\xa2\x66\x67\xff\x9c\x98\xce\xa3\x26\x0d\x8a\xe0\x4c\x35\xa9\x0d\x6b\x77\x3d\x19\xe4\xdb\x42\xfb\x7b\x22\xbb\x52\xa2\xff\x9e\xe0\x79\x38\xfa\x8f\xda\x03\xf1\x27\xd0\xb8\x21\xec\x3f\x19\xd8\xeb\xaf\x86\x25\x2c\xcb\xc9\xb5\x8b\xc3\xc1\x76\x73\x47\x30\x5f\x0a\x3a\xe3\x80\xb5\x0a\x1b\x64\x14\xa8\x8a\xe8\xb2\x97\x65\xdb\xfd\x66\x7c\x02\x90\x16\xd5\xdc\x4e\xca\x12\x3b\xfb\xa6\xcf\xfb\xa3\x76\xe1\xe7\x4d\x4f\x4f\x27\xb8\x97\xa6\xb5\x1d\x2f\xcc\xcd\xd3\xf2\xf7\x61\x66\x13\x60\xe2\xde\x17\x36\x26\xab\x7a\xa2\xd8\x78\xeb\xf8\x28\x55\x86\x9d\xa5\x72\x5f\xe1\x8e\x37\x8e\x5d\xcc\xdd\x3d\xc8\xb9\xfd\xa0\xd1\xb0\x1d\x6a\x8f\xec\x10\x70\x7b\x5a\x67\xbf\xa1\x41\xcb\xaf\xc3\x7b\xf0\x15\xb6\x95\xb6\x06\xdf\x63\x3b\x3c\xae\x82\xde\x53\x72\x70\x4d\xa6\xb5\x13\xef\x61\x03\x79\xdd\x30\xec\x78\xd8\x8e\xb9\x74\xd0\x1f\x87\xff\x35\xe7\x4c\xd0\xbc\x8e\x3e\xf9\xff\xb9\xa6\x1a\x7b\xf7\xdb\x0b\x45\x97\xa7\xf1\xd2\xc2\xcc\x43\xd2\x36\x5f\x98\x49\xda\x1d\xcd\xcb\x47\xe2\xe6\x3e\x13\xf8\x35\xa8\x1b\x93\xf0\x76\xd7\xdd\xe9\x40\x4f\x8f\xc4\x6d\x6f\x42\x05\xd4\xa2\x0a\x33\x54\xb0\xff\xca\xda\x1e\x77\x87\x2c\xd4\xfd\x57\x46\xba\x09\x58\xcf\x8b\xfe\xff\x21\x73\xee\x13\x9f\x83\x36\xb9\x7c\x8f\xa5\x2e\x89\x3f\x09\xd5\xee\x46\xd0\x39\xed\xb6\x74\x39\xe4\x88\x2d\xf3\xb5\x92\xe4\xfd\xce\xe6\x6f\x07\x16\xee\xd5\xc0\x0c\xce\x8d\xf2\x65\xb7\x0b\xdc\xa8\xc3\xb6\x3a\x9b\xef\xcd\x41\xa6\x6e\x09\x87\x3b\x42\xf7\x40\x19\x46\xed\xde\xaf\xf4\xff\x13\x00\x00\xff\xff\x98\xe4\xb3\x34\x17\x23\x00\x00")

func templates_commands_tmpl() ([]byte, error) {
	return bindata_read(
		_templates_commands_tmpl,
		"templates/commands.tmpl",
	)
}

var _templates_main_tmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x53\x41\x4f\xf3\x38\x10\x3d\x27\xbf\x62\xb0\xd4\x55\xac\x2d\x0e\x7b\xad\x94\x43\x05\x29\x20\xc1\xa6\x2a\x85\x2b\x32\xc9\x24\xb5\x70\xec\xc8\x76\x0a\x5a\xd4\xff\xbe\x72\x62\x1a\x28\x9c\xbe\x5e\x32\xe3\xf7\x3c\xf3\xe6\x79\xda\xf1\xf2\x95\x37\x08\x2d\x17\x2a\x8e\x45\xdb\x69\xe3\x20\x89\x01\x00\x48\xdd\x3a\x32\x46\x52\x37\x21\x32\xbd\x72\xa2\xc5\x90\x69\x1b\x02\xeb\x8c\x50\x8d\x25\x71\x1c\x91\x46\xb8\x5d\xff\xc2\x4a\xdd\xa6\xaf\x5c\xfe\xa7\x75\xaa\x3b\x54\xbc\x13\xe7\xa5\x14\xe7\x0d\x2a\x34\xdc\x69\x93\x96\x52\x90\x98\xc6\x71\xdd\xab\x72\xe8\x9f\x50\xf8\x18\xca\x95\x5a\xd5\xa2\x81\x45\x06\x7f\x95\x52\xb0\xcb\x21\x1d\x21\xff\x5b\x76\xdd\xbf\xbc\xc5\x85\x6f\xfc\xf1\x01\xcc\x27\x70\x38\x90\xf9\x91\x91\xab\xfd\xda\x60\x2d\xde\x17\x13\x23\x57\xfb\xef\xa4\x27\x34\x56\x68\x35\x94\xf9\x87\x5d\xb0\x8b\x80\x1d\x46\x0d\x52\xb0\x5b\x25\x5c\x32\x8a\xa1\xf1\x70\xda\x48\xfd\xc2\xe5\x4a\xf2\xc6\xce\x01\x8d\xf1\x1a\x3d\xf3\x9e\xbf\xe2\x52\x55\x6b\x6e\x2c\x5e\x4f\x9c\x84\x0e\xb7\x44\x3d\x70\xcf\x32\x50\x42\xc2\x34\x88\xd4\x0d\x5b\x71\xc7\x65\x82\xc6\xd0\xd0\x7c\xf8\xa4\x29\x6c\x8b\xab\x62\x01\xcb\xaa\x02\x83\x8d\xb0\x0e\x0d\x94\xba\x6d\xb9\xaa\x2c\xec\xd0\x20\x1b\x99\xbe\x70\x76\x54\x3b\x5a\xd5\x1b\xee\x84\x56\xc9\x8f\xd9\xa1\x41\xb7\x2d\xee\xef\x56\x42\xe2\x9a\xbb\x5d\x42\x2c\x3a\x37\xbc\x1c\xfd\x15\x2c\x0d\xba\x11\x9b\x86\xfa\x93\x99\x22\x2f\x70\x59\x55\xa3\xbe\xcb\x30\x47\xe2\x4f\x37\x5a\x3b\x7a\x24\x2c\x7b\xb7\xfb\x05\x1e\x71\x1f\xb3\xfc\x1d\xcb\xde\x61\x42\xe3\x43\xd8\x9d\x53\xdd\xb5\x90\xa8\xfc\x4e\x8c\x5b\x49\xc3\x37\x88\x44\xb5\x7f\xe2\xb2\x47\xff\x74\xda\xb2\x6b\x74\xa8\xf6\x49\x58\x60\xb6\xd5\x8f\x5d\x87\x26\xa9\x5b\xc7\x1e\x3a\x23\x94\xab\x13\x32\xb3\xcf\x33\xfb\xbc\x5e\x6e\x6f\xc8\xfc\xe7\x42\xc1\x67\x3f\x4a\x27\x6b\x3e\x9b\x9c\x65\x40\xc8\x17\x7b\x0c\xba\xde\xa8\x23\xfe\x65\xe1\x02\x72\xd2\x38\x65\x33\x9b\xce\x2c\x73\xba\x95\x64\x0e\xbd\x45\x73\xa3\x5b\xbc\x12\x26\xa1\x73\x98\x54\xdf\xe9\x37\x34\x3f\x1e\x9c\x7e\x11\x77\xb4\xeb\x5b\x8d\xc9\x9b\x48\xd4\x10\xfe\xdd\xec\xba\x28\x1e\x20\xcb\x80\xbc\x09\x55\xe9\x37\xeb\x27\x88\xa2\x9d\x6e\x4f\x5c\x23\x37\xc5\x7d\x7e\xb5\xb9\x7d\xca\x09\x85\xbf\x4f\x81\xc1\x30\x1a\x47\xbe\xf2\x70\x37\x0b\x66\x44\xa1\xd6\xb7\x52\x8f\x0f\xf9\x66\xbd\x29\x56\xb7\x77\xf9\x70\xe9\x10\x47\x51\xf0\xc4\x93\x63\x7f\x10\xf2\x93\x3e\xc4\xcf\xf6\x7f\x00\x00\x00\xff\xff\x65\x95\x7c\x9b\xc9\x04\x00\x00")

func templates_main_tmpl() ([]byte, error) {
	return bindata_read(
		_templates_main_tmpl,
		"templates/main.tmpl",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"templates/commands.tmpl": templates_commands_tmpl,
	"templates/main.tmpl": templates_main_tmpl,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"templates": &_bintree_t{nil, map[string]*_bintree_t{
		"commands.tmpl": &_bintree_t{templates_commands_tmpl, map[string]*_bintree_t{
		}},
		"main.tmpl": &_bintree_t{templates_main_tmpl, map[string]*_bintree_t{
		}},
	}},
}}
