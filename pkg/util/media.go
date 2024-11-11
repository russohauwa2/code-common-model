package util

import "strings"

var coverCodecs = map[string]struct{}{
	"mjpeg": {},
	"png":   {},
	"jpeg":  {},
	"jpg":   {},
	"gif":   {},
	"bmp":   {},
	"tiff":  {},
	"webp":  {},
	"apic":  {}, // ID3v2 标签中的附加图片
}

// IsCoverCodec 判断给定的 codeName
func IsCoverCodec(codeName string) bool {
	_, exists := coverCodecs[strings.ToLower(codeName)]
	return exists
}

// GetAllCoverCodecs GetAllCoverCodecs
func GetAllCoverCodecs() []string {
	codecs := make([]string, 0, len(coverCodecs))
	for codec := range coverCodecs {
		codecs = append(codecs, codec)
	}
	return codecs
}
