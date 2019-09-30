package avc

import (
	"bytes"
	"github.com/q191201771/naza/pkg/assert"
	"testing"
)

func TestCaptureAVC(t *testing.T) {
	b := &bytes.Buffer{}
	CaptureAVC(b, []byte{0x17, 0x0, 0x0, 0x0, 0x0, 0x1, 0x64, 0x0, 0x1f, 0xff, 0xe1, 0x0, 0xa, 0x27, 0x64, 0x0, 0x1f, 0xac, 0x56, 0x80, 0xb4, 0xa, 0x19, 0x1, 0x0, 0x4, 0x28, 0xee, 0x3c, 0xb0})
	CaptureAVC(b, []byte{0x27, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x8, 0x57, 0x21, 0xe1, 0x4, 0x43, 0x5f, 0xf4, 0x34, 0xb, 0x0, 0x4, 0x7d, 0xe1, 0x97, 0x46, 0xdf, 0xb1, 0xd1, 0xb5, 0xad, 0xd2, 0xf8, 0xf4, 0x5f, 0x39, 0x7d, 0x26, 0x8e, 0xec, 0xa3, 0xaf, 0xc0, 0xe7, 0xad, 0x35, 0x86, 0xfc, 0x53, 0xdb, 0xd6, 0x27, 0x72, 0x4e, 0x34, 0xd4, 0xbb, 0x83, 0x68, 0x29, 0xbd, 0xa8, 0xa0, 0xb9, 0x5b, 0x3c, 0xe3, 0xae, 0x2b, 0xd6, 0x2e, 0xb0, 0x6, 0x10, 0xa4, 0x3b, 0xb2, 0x7b, 0x15, 0x99, 0x8f, 0x83, 0x7d, 0xa7, 0x28, 0x82, 0x5f, 0x37, 0x16, 0x38, 0x7a, 0x9f, 0x49, 0xf9, 0x7c, 0xd2, 0xad, 0xd4, 0xf, 0x5b, 0xac, 0x90, 0x33, 0x83, 0x46, 0x29, 0xf, 0xae, 0x36, 0x3e, 0x32, 0x95, 0x85, 0x25, 0x66, 0xa3, 0x29, 0xc6, 0xa5, 0x1f, 0xb, 0xcc, 0xa1, 0x54, 0x63, 0x8a, 0xe, 0xa2, 0xd, 0x9e, 0x5d, 0x70, 0xd3, 0x2e, 0xeb, 0xbe, 0x71, 0x39, 0xf7, 0x7, 0xc0, 0x7e, 0xba, 0xf2, 0xdb, 0x69, 0xd, 0xba, 0xfd, 0x27, 0xba, 0x64, 0x6e, 0x89, 0xe4, 0x17, 0x9b, 0x5c, 0xc3, 0xb9, 0x23, 0x16, 0xe0, 0xf6, 0xae, 0x8, 0x28, 0x21, 0x46, 0x62, 0x7f, 0x5f, 0x69, 0x9, 0x4c, 0xf3, 0x6a, 0x78, 0xbc, 0xf3, 0x1b, 0x0, 0x48, 0xda, 0xbb, 0x1e, 0x1d, 0xdb, 0x6e, 0xdf, 0x84, 0x70, 0xa1, 0x2a, 0xac, 0x11, 0x3, 0x32, 0x95, 0x16, 0x36, 0xcf, 0x0, 0xbb, 0xd0, 0x43, 0x7, 0xc1, 0x19, 0x24, 0x5a, 0x6b, 0x7d, 0x70, 0x4a, 0x62, 0x8e, 0x1a, 0xa8, 0x82, 0xca, 0x58, 0x10, 0x4e, 0x21, 0x24, 0xff, 0xae, 0x7, 0x49, 0x12, 0xab, 0x19, 0x96, 0x3e, 0x6, 0x5, 0x3d, 0x82, 0x81, 0xf4, 0x42, 0xc1, 0x31, 0x1a, 0x70, 0x18, 0xe3, 0x4c, 0x86, 0x22, 0x4d, 0xd9, 0x65, 0xbd, 0x4f, 0xc3, 0x14, 0xff, 0x76, 0x95, 0xb0, 0xe2, 0x8f, 0x52, 0xde, 0xd7, 0x4e, 0x55, 0xb4, 0xf1, 0x7e, 0x9, 0x9e, 0xf, 0x55, 0xd1, 0xff, 0x8, 0x8c, 0xb8, 0x51, 0x82, 0x73, 0xde, 0xc6, 0xf1, 0x23, 0x10, 0x40, 0xed, 0xbf, 0x25, 0x60, 0x14, 0x4e, 0x18, 0x9f, 0x58, 0x82, 0xf, 0x80, 0xa7, 0x55, 0x2d, 0x87, 0xf7, 0x5, 0x53, 0x60, 0xf0, 0xdf, 0xe8, 0x42, 0xab, 0x4f, 0xcc, 0xc5, 0xdc, 0x6e, 0x25, 0x35, 0x81, 0xea, 0x1a, 0x5, 0x81, 0x7a, 0x36, 0x71, 0xe8, 0xef, 0x9f, 0x33, 0xb0, 0xf0, 0x82, 0x14, 0xe5, 0xec, 0xb9, 0x66, 0x6e, 0xbf, 0x12, 0x1d, 0x6, 0x3f, 0x5a, 0x17, 0xfd, 0x6, 0x90, 0x3f, 0x22, 0x7c, 0xb1, 0x7d, 0x31, 0x4f, 0x7c, 0x5a, 0x5a, 0xce, 0x3c, 0x70, 0x97, 0xb2, 0x3e, 0x90, 0x89, 0x17, 0x48, 0xb9, 0x8c, 0x26, 0x58, 0xa, 0xff, 0x26, 0x43, 0xd1, 0xcd, 0x52, 0xe9, 0x64, 0xd2, 0x79, 0xc5, 0x8c, 0xa3, 0x10, 0x3f, 0x34, 0xa3, 0xd8, 0x14, 0xfa, 0x75, 0xc9, 0x7c, 0x58, 0xc4, 0x6, 0x7f, 0xac, 0xe7, 0x81, 0x5b, 0x21, 0xab, 0x2b, 0x25, 0x8f, 0x1e, 0xa9, 0xcd, 0xef, 0xd4, 0xd0, 0x87, 0x75, 0xe3, 0x4b, 0xe6, 0x8c, 0x6e, 0x81, 0x4, 0xb3, 0x8e, 0xf, 0xd1, 0x24, 0x16, 0x48, 0xa4, 0x5d, 0x17, 0x6d, 0xf1, 0x9a, 0xfd, 0x92, 0x8e, 0xf9, 0x69, 0x8, 0x7b, 0x5a, 0x85, 0x8, 0xa1, 0x40, 0x3, 0x85, 0x88, 0x8e, 0x17, 0xa8, 0x86, 0x71, 0x37, 0xc9, 0x7e, 0x45, 0x75, 0xe, 0xeb, 0xdb, 0xd3, 0x6, 0x3c, 0x49, 0x5e, 0x2b, 0xfa, 0xc7, 0x25, 0x5c, 0x34, 0x26, 0x7b, 0xf4, 0xfb, 0x44, 0x5f, 0xd9, 0xcb, 0x22, 0xab, 0x53, 0xec, 0xd7, 0x1e, 0xb5, 0x9f, 0x15, 0x69, 0x95, 0x7, 0xcd, 0x2b, 0x81, 0x4a, 0xa7, 0x37, 0x8c, 0xe1, 0xc5, 0x6c, 0xce, 0xb3, 0xf8, 0xa8, 0x9b, 0x4d, 0xb6, 0x53, 0x85, 0x1, 0x25, 0xd, 0x32, 0x60, 0x7e, 0xdc, 0xf6, 0x6c, 0xe4, 0xf8, 0x3, 0x9f, 0x8a, 0x6b, 0xdf, 0x71, 0x96, 0x2f, 0x24, 0x98, 0x19, 0xcd, 0xda, 0x42, 0xd8, 0x75, 0x91, 0x1a, 0x32, 0xcd, 0xd8, 0x8c, 0xfe, 0x19, 0xc3, 0x58, 0xec, 0x30, 0xc9, 0xa2, 0xc5, 0x5e, 0x88, 0xde, 0xb5, 0xcb, 0x18, 0x5a, 0x1c, 0x30, 0x7b, 0xa3, 0xb0, 0x5c, 0xd3, 0x3d, 0x54, 0xe7, 0xbb, 0x85, 0x97, 0x75, 0x9c, 0x91, 0x6e, 0x66, 0x2a, 0x2d, 0xf2, 0x7b, 0x4c, 0xfc, 0xe, 0xf9, 0x3d, 0xa0, 0x73, 0xd2, 0x80, 0xb0, 0xd, 0xb0, 0xc8, 0x3b, 0x18, 0x31, 0x31, 0xbe, 0x64, 0xbd, 0x9c, 0xd1, 0x55, 0x2d, 0x86, 0xf8, 0xb9, 0x7c, 0x6, 0xcd, 0x3e, 0xac, 0xcc, 0xba, 0x1b, 0x35, 0xca, 0xda, 0xa6, 0xbc, 0x9e, 0xab, 0xc7, 0xed, 0xa3, 0xcf, 0x34, 0x7, 0x37, 0x33, 0x89, 0x5f, 0x73, 0xc8, 0xa0, 0x19, 0x3e, 0x45, 0x37, 0x3b, 0x3f, 0xc4, 0xf3, 0x38, 0x16, 0x72, 0xea, 0x4a, 0x21, 0x74, 0x0, 0xc3, 0xee, 0x41, 0x7c, 0x81, 0x5a, 0xcc, 0xab, 0x46, 0xd3, 0x1b, 0xe3, 0x67, 0xdf, 0xd5, 0x43, 0xc3, 0x48, 0xf5, 0xbc, 0xe8, 0xfb, 0x70, 0x65, 0x72, 0x91, 0x41, 0x2d, 0x92, 0x61, 0x38, 0x21, 0x6c, 0x27, 0xb6, 0x32, 0xf1, 0x9a, 0xd7, 0x16, 0xf9, 0x97, 0xb, 0xd6, 0x7b, 0x93, 0xc4, 0xa1, 0x63, 0x2f, 0x4f, 0xd9, 0xe3, 0xd6, 0x97, 0xa0, 0x6d, 0x1e, 0xee, 0xe8, 0xb8, 0xbc, 0x46, 0x85, 0x99, 0x75, 0x3, 0xc1, 0xf, 0x3f, 0xb3, 0x68, 0xe9, 0x6b, 0x41, 0x35, 0x9b, 0x46, 0x1e, 0xe6, 0x4c, 0x27, 0x67, 0x7f, 0x2a, 0xb, 0x83, 0x55, 0x22, 0x4f, 0xfc, 0x69, 0x52, 0x28, 0xf5, 0x91, 0x49, 0x24, 0xbf, 0xd, 0xd9, 0x24, 0x74, 0x31, 0xe1, 0xa2, 0x9c, 0xf4, 0x2e, 0x83, 0x55, 0xaf, 0x12, 0xd1, 0x77, 0xbb, 0xcd, 0x1f, 0xa7, 0xb1, 0x44, 0xd0, 0x29, 0x14, 0x5b, 0xc6, 0x34, 0xd0, 0x5b, 0x2d, 0x6d, 0xf3, 0xd5, 0x2c, 0x97, 0xb5, 0xc1, 0x1, 0x8b, 0x29, 0x32, 0x76, 0x54, 0xde, 0x66, 0xb2, 0x26, 0xbf, 0x74, 0xa2, 0x8e, 0x98, 0xd6, 0xa6, 0x9a, 0x72, 0x63, 0x14, 0xee, 0x66, 0x30, 0x1c, 0x20, 0x5b, 0x35, 0x3c, 0x1b, 0x50, 0xde, 0xe2, 0x2f, 0x5, 0x36, 0x35, 0x61, 0x68, 0xd9, 0x23, 0xa6, 0x63, 0x6d, 0x78, 0x88, 0x6a, 0x8d, 0x41, 0x42, 0xc6, 0x49, 0x42, 0xc4, 0xaf, 0x7d, 0x8e, 0xb, 0x0, 0x2f, 0x19, 0xe0, 0x90, 0xfd, 0x95, 0x3b, 0xa9, 0x22, 0xb9, 0x78, 0x97, 0x3b, 0x20, 0xf3, 0x10, 0xd, 0xb9, 0x96, 0x2, 0xb7, 0xd8, 0x0, 0x5f, 0x6c, 0x52, 0xb7, 0xd1, 0x86, 0x7a, 0xb1, 0x40, 0x5, 0x23, 0xf8, 0x90, 0xaf, 0xa5, 0x83, 0x29, 0x29, 0x31, 0x25, 0xbe, 0x6d, 0xee, 0x5, 0x3a, 0xb1, 0xb0, 0xc7, 0xe6, 0xe5, 0x80, 0xd6, 0x72, 0x1c, 0x73, 0x87, 0xed, 0x81, 0xf1, 0x46, 0x10, 0x6, 0xd8, 0x90, 0x27, 0xcb, 0x44, 0x4f, 0x40, 0x49, 0x7e, 0x2d, 0xd9, 0x7f, 0xd, 0x2, 0xc8, 0x28, 0xc0, 0x73, 0xf5, 0x93, 0x38, 0xf5, 0xce, 0x19, 0xeb, 0xed, 0x65, 0xe, 0x54, 0x5e, 0x33, 0xd7, 0xdc, 0xb0, 0xb5, 0xa0, 0x62, 0xb0, 0xde, 0xfb, 0x0, 0x8c, 0xf8, 0xad, 0x5d, 0x9b, 0xef, 0xe8, 0xd8, 0x6b, 0x74, 0x85, 0x2b, 0x1a, 0xcb, 0xd4, 0x62, 0x19, 0x17, 0x1a, 0x90, 0x7c, 0xae, 0xdc, 0xcd, 0x7f, 0x71, 0xd2, 0xba, 0x22, 0x6, 0x1f, 0x80, 0xaf, 0xca, 0x1, 0xa, 0x15, 0x32, 0x92, 0x93, 0x14, 0x65, 0x20, 0xdb, 0xee, 0x17, 0x67, 0xa5, 0x41, 0x59, 0xbc, 0xee, 0xe4, 0x3f, 0xad, 0x1c, 0x32, 0xc6, 0xad, 0xb1, 0x51, 0x20, 0xc3, 0xeb, 0xa5, 0xc4, 0xa4, 0xc7, 0x6e, 0xff, 0xd4, 0x83, 0x6c, 0x8a, 0xeb, 0xee, 0x22, 0x12, 0x5d, 0xc5, 0xc2, 0x88, 0x3a, 0xcb, 0xa3, 0xf5, 0x58, 0x10, 0x12, 0x98, 0xb0, 0xa6, 0xbb, 0xc8, 0xf9, 0xb1, 0xd7, 0xf5, 0x8b, 0x92, 0x48, 0x61, 0x6d, 0xc5, 0x42, 0x8c, 0x3, 0x17, 0x55, 0xd7, 0x2a, 0x95, 0x5, 0xd9, 0x2a, 0x51, 0x52, 0x12, 0x18, 0x55, 0x6e, 0x2e, 0xc8, 0xfa, 0x41, 0x73, 0xb4, 0x53, 0x42, 0x85, 0xb7, 0xac, 0xca, 0x44, 0x2b, 0x3f, 0xac, 0xc7, 0x12, 0x4d, 0xb1, 0x1d, 0x80, 0xd2, 0xe5, 0x85, 0x2a, 0xd, 0x19, 0x84, 0x38, 0xcb, 0x4f, 0x52, 0xc7, 0x1a, 0x39, 0x88, 0xf0, 0xed, 0xde, 0xf4, 0x14, 0x26, 0x24, 0x8c, 0xcd, 0x11, 0xcd, 0x58, 0x70, 0x3, 0x87, 0x8f, 0xb1, 0xe4, 0x98, 0x49, 0xdf, 0xef, 0xe6, 0x5b, 0xfc, 0xff, 0x0, 0xba, 0xa2, 0x32, 0xe6, 0xf2, 0x6, 0x3e, 0xb0, 0x6f, 0xc3, 0xab, 0xcb, 0x3a, 0xc3, 0x5c, 0x5f, 0x9d, 0x68, 0x70, 0x16, 0x68, 0xf, 0x17, 0x7f, 0xd9, 0x4f, 0xfd, 0x3, 0x2d, 0x5b, 0x97, 0xc9, 0x88, 0xe, 0x5c, 0x8a, 0x4d, 0xfb, 0x4b, 0x86, 0xed, 0x9a, 0x9, 0xa7, 0x7, 0xec, 0xa7, 0x87, 0x88, 0xe, 0x27, 0x80, 0xe2, 0x55, 0x47, 0x15, 0xe2, 0x2e, 0xa8, 0x4d, 0x29, 0xb, 0x20, 0x1a, 0x79, 0xae, 0x37, 0x68, 0x9f, 0x64, 0xad, 0x61, 0xcf, 0x2, 0x15, 0x7f, 0x90, 0xf5, 0x94, 0x35, 0xc0, 0xc8, 0x13, 0x22, 0x30, 0xfe, 0xbb, 0x7d, 0x92, 0x80, 0x30, 0x70, 0x27, 0xfc, 0x8e, 0xb6, 0x2a, 0xe8, 0xfa, 0x42, 0xe3, 0x84, 0xdb, 0xe4, 0x66, 0x5a, 0x23, 0xf9, 0x55, 0x62, 0x1a, 0xe5, 0xa7, 0xf, 0x64, 0x5e, 0x66, 0x11, 0x81, 0x2d, 0x9c, 0xb3, 0x41, 0x4f, 0x3e, 0x8b, 0x66, 0xa2, 0x75, 0x72, 0x7, 0x80, 0xde, 0xd3, 0xd9, 0xbd, 0x4, 0xb8, 0x9c, 0x8b, 0x67, 0xdf, 0x48, 0x9, 0xb1, 0x88, 0xf0, 0x74, 0x5c, 0xa, 0xa6, 0x82, 0xba, 0x38, 0x72, 0x29, 0x1d, 0xa7, 0x46, 0xb6, 0xae, 0x72, 0x4e, 0x3c, 0xde, 0x2, 0x9b, 0x47, 0xef, 0xc9, 0x4e, 0x11, 0x78, 0xdf, 0x79, 0xa0, 0x64, 0xfe, 0x5e, 0xdc, 0x7b, 0xb5, 0xad, 0x39, 0x1e, 0xe0, 0x8d, 0x2c, 0x5e, 0xa3, 0x98, 0x3e, 0xd2, 0x4a, 0x8b, 0x17, 0x19, 0x4b, 0xbe, 0x75, 0xf4, 0xa, 0x12, 0xf0, 0x31, 0x99, 0xe7, 0x82, 0x29, 0xe9, 0xef, 0x11, 0x43, 0xdf, 0x96, 0x6, 0x3e, 0x32, 0xe5, 0x52, 0x12, 0x98, 0xc6, 0x61, 0x7c, 0xee, 0x7c, 0xda, 0x99, 0x8b, 0x19, 0xcd, 0x83, 0x10, 0xec, 0xd9, 0xcb, 0xf7, 0x1d, 0xfc, 0x23, 0x95, 0xf9, 0xa, 0x61, 0x47, 0x69, 0xd5, 0x55, 0x5, 0x63, 0x99, 0x71, 0xe8, 0x13, 0x14, 0x9c, 0x27, 0x9, 0x8, 0x78, 0x42, 0xe6, 0xbd, 0x59, 0x26, 0xa8, 0x6, 0x17, 0xff, 0xf8, 0x99, 0xb3, 0x74, 0xdd, 0x70, 0x5e, 0x23, 0xec, 0x36, 0x65, 0x83, 0x67, 0xec, 0x81, 0xa4, 0x71, 0xf6, 0x3e, 0x19, 0x63, 0x95, 0xfd, 0x1, 0x44, 0x3d, 0x54, 0x1c, 0xf4, 0x15, 0xe1, 0xc, 0x97, 0x4c, 0x40, 0xee, 0x83, 0x72, 0xa, 0x56, 0x82, 0x52, 0x12, 0xe7, 0x56, 0xe2, 0xf, 0x51, 0xa, 0xac, 0xa0, 0x81, 0xaa, 0xda, 0x91, 0x2a, 0x61, 0xa3, 0x80, 0x2a, 0x9f, 0x94, 0x98, 0xe4, 0xf3, 0xa, 0x1, 0x58, 0xd9, 0x97, 0x8e, 0x74, 0x5d, 0xdd, 0x70, 0x5c, 0x1a, 0x41, 0x16, 0x27, 0xdb, 0x3e, 0xdd, 0x1e, 0xd2, 0xce, 0xd2, 0x70, 0xbb, 0x8b, 0x5d, 0x4c, 0xa1, 0x66, 0x9, 0xa9, 0x2a, 0x60, 0xa7, 0xb8, 0xda, 0xb2, 0xbe, 0xf4, 0xcc, 0x37, 0x4d, 0xad, 0x97, 0x9e, 0x62, 0xc8, 0xcf, 0x1b, 0xf6, 0x72, 0x10, 0x3f, 0xbf, 0x9e, 0x21, 0x96, 0x32, 0xe2, 0xf1, 0x8d, 0x82, 0x6d, 0xcf, 0xea, 0x41, 0xdd, 0xc, 0xb6, 0xa0, 0x39, 0x64, 0x9f, 0x37, 0xee, 0xa6, 0x6, 0x62, 0x91, 0xb4, 0x13, 0x79, 0x4d, 0x15, 0x53, 0xd4, 0xa3, 0x2e, 0xd3, 0x54, 0x15, 0x3e, 0x9a, 0x4, 0xfc, 0x85, 0xfc, 0x30, 0x97, 0x9f, 0x19, 0xac, 0x2, 0xbb, 0xe2, 0x8c, 0xae, 0x6f, 0xa0, 0x16, 0x41, 0xe5, 0xd8, 0xf8, 0xb5, 0x5a, 0x4a, 0x7f, 0xc7, 0x73, 0xf2, 0xd4, 0x7b, 0x18, 0xf0, 0x9a, 0x61, 0x17, 0xb3, 0xfd, 0xc5, 0x8, 0x23, 0x96, 0x10, 0x2b, 0x3e, 0xc, 0x57, 0x66, 0x35, 0xc, 0x93, 0x83, 0xb1, 0x13, 0xf7, 0xcf, 0xa2, 0xf7, 0x4d, 0x29, 0xfb, 0x2f, 0x96, 0xef, 0x69, 0x48, 0x86, 0x4d, 0xc0, 0x7f, 0xa3, 0x4c, 0x5d, 0xc7, 0x96, 0x36, 0x80, 0xeb, 0xa9, 0x69, 0x7a, 0x5d, 0x83, 0x52, 0x28, 0xae, 0x51, 0xa6, 0x3e, 0xf9, 0x7a, 0x9d, 0x77, 0xcf, 0x3b, 0x87, 0x2e, 0x3, 0xf1, 0x13, 0x5d, 0x5f, 0x2d, 0x93, 0xa5, 0xec, 0x2c, 0x7a, 0xa2, 0x40, 0xf, 0x8f, 0xf7, 0x72, 0xb5, 0xe2, 0xc2, 0x95, 0x2e, 0x15, 0xb8, 0x75, 0xa, 0x16, 0x63, 0x21, 0xe7, 0x2c, 0x20, 0xd1, 0xde, 0x4f, 0x49, 0xaa, 0xf3, 0x1f, 0x1f, 0x50, 0xf9, 0xa4, 0x19, 0xd7, 0xdf, 0x86, 0x1a, 0x65, 0x35, 0x42, 0x1e, 0x69, 0xcb, 0x9, 0x92, 0x6f, 0xa5, 0x10, 0xe3, 0x8a, 0xd, 0xad, 0xfa, 0x19, 0x98, 0xb1, 0xfb, 0x7, 0xa8, 0x7f, 0x86, 0x83, 0x40, 0xb1, 0x51, 0x3d, 0x86, 0x8c, 0x2e, 0x9d, 0xac, 0x72, 0x3b, 0x93, 0xf0, 0xf8, 0x28, 0x55, 0x33, 0x8e, 0xfc, 0x74, 0x29, 0xfc, 0xa7, 0xd2, 0x66, 0x28, 0xc2, 0xcf, 0xf3, 0x42, 0xb5, 0xe3, 0x7e, 0x32, 0x5, 0x66, 0xfa, 0xea, 0xe0, 0xf7, 0x7d, 0xa8, 0xf5, 0xcc, 0x3a, 0xb9, 0x9c, 0xb1, 0x33, 0x6c, 0x75, 0xee, 0xb2, 0xb4, 0x1, 0x8d, 0x3a, 0xaf, 0xa3, 0xd0, 0xa3, 0x75, 0x5e, 0x2c, 0x40, 0x5e, 0x42, 0xae, 0x9c, 0xab, 0x8f, 0x8d, 0x2a, 0xb3, 0xf5, 0xfe, 0x0, 0x92, 0xd6, 0x24, 0x63, 0x3d, 0x3c, 0xd9, 0xc, 0xe6, 0x7d, 0x98, 0xc4, 0xd4, 0xdf, 0xb9, 0x2c, 0xe1, 0x3a, 0xb9, 0x96, 0x92, 0x7c, 0xd1, 0xa8, 0x7a, 0xb2, 0x3b, 0x24, 0xc7, 0x6, 0xb4, 0x51, 0x2b, 0xd3, 0x55, 0xf3, 0x6a, 0xda, 0x82, 0x55, 0x38, 0xb3, 0x1f, 0x96, 0x66, 0x35, 0xb0, 0xf8, 0x3a, 0xba, 0x44, 0xcf, 0x42, 0x7c, 0xf7, 0xb8, 0x91, 0xda, 0x11, 0x84, 0x2e, 0x75, 0xd, 0xc9, 0xaf, 0xce, 0x2d, 0xbf, 0x86, 0xb4, 0x6f, 0x9a, 0x19, 0x60, 0x33, 0x16, 0x9b, 0x34, 0x2d, 0x4, 0x4c, 0x53, 0x12, 0x79, 0x85, 0x7f, 0x2, 0x7d, 0xf, 0xbf, 0x1, 0xc1, 0x94, 0x17, 0xd8, 0xf2, 0x93, 0x0, 0xf3, 0x33, 0x18, 0x9d, 0xea, 0x20, 0x49, 0x88, 0x0, 0xc7, 0x3c, 0xc8, 0x99, 0x8e, 0x16, 0xd5, 0x20, 0xa4, 0x89, 0xb5, 0x5c, 0x70, 0xfc, 0xbf, 0x1e, 0x11, 0x32, 0x27, 0x5e, 0xb1, 0x6a, 0xf8, 0x99, 0x90, 0xe4, 0x12, 0x18, 0xde, 0x55, 0x34, 0xaf, 0x42, 0x8c, 0x2b, 0x2e, 0xad, 0xed, 0xc2, 0xc3, 0xb1, 0x3, 0xca, 0xd0, 0xbc, 0x13, 0x6d, 0xda, 0xcc, 0xca, 0x4e, 0x3a, 0x58, 0x6d, 0x0, 0xd3, 0xca, 0xf1, 0x2e, 0x63, 0x96, 0xfe, 0x54, 0xfb, 0x71, 0x77, 0xfe, 0x3f, 0x5a, 0x53, 0xa, 0xc0, 0x88, 0x3e, 0x97, 0xf2, 0x3f, 0xb8, 0x6, 0x9c, 0xfb, 0x71, 0x41, 0x6c, 0x82, 0x46, 0x74, 0x9a, 0x16, 0x92, 0xaa, 0x68, 0x61, 0x3f, 0xae, 0x1a, 0x15, 0xea, 0xee, 0x16, 0xa3, 0x46, 0xa6, 0xc9, 0x53, 0x22, 0x36, 0xce, 0x5b, 0xec, 0x7d, 0x38, 0xf, 0xe, 0xf0, 0x8a, 0xcc, 0xd7, 0xd3, 0x94, 0xf1, 0x29, 0x86, 0x64, 0x49, 0xc0, 0x2a, 0x39, 0x65, 0x24, 0x46, 0x4b, 0xe6, 0x21, 0xba, 0xe4, 0x51, 0x39, 0x88, 0xac, 0x25, 0x4a, 0xa3, 0xc0, 0x8d, 0xc7, 0xa9, 0x29, 0xe3, 0xe0, 0x6e, 0xe0, 0xe0, 0x9b, 0x54, 0xf1, 0xed, 0xa4, 0xcc, 0x8b, 0x8d, 0xc9, 0xd7, 0xe6, 0x27, 0xa9, 0x2d, 0x7d, 0x36, 0x8f, 0x41, 0xdc, 0x63, 0x31, 0x51, 0x7d, 0x83, 0x29, 0x2, 0x1e, 0x23, 0xbc, 0x96, 0x33, 0x6b, 0x39, 0xca, 0xb1, 0xfd, 0xc3, 0xb1, 0xec, 0xe5, 0xbd, 0x60, 0xa5, 0xc6, 0x3b, 0x1b, 0xdf, 0x34, 0x71, 0x87, 0x14, 0xd, 0xfe, 0x4d, 0xce, 0x3c, 0xeb, 0x65, 0x24, 0x5b, 0xee, 0x51, 0xb7, 0x4e})
	expected := []byte{0x0, 0x0, 0x0, 0x1, 0x27, 0x64, 0x0, 0x1f, 0xac, 0x56, 0x80, 0xb4, 0xa, 0x19, 0x0, 0x0, 0x0, 0x1, 0x28, 0xee, 0x3c, 0xb0, 0x0, 0x0, 0x0, 0x1, 0x21, 0xe1, 0x4, 0x43, 0x5f, 0xf4, 0x34, 0xb, 0x0, 0x4, 0x7d, 0xe1, 0x97, 0x46, 0xdf, 0xb1, 0xd1, 0xb5, 0xad, 0xd2, 0xf8, 0xf4, 0x5f, 0x39, 0x7d, 0x26, 0x8e, 0xec, 0xa3, 0xaf, 0xc0, 0xe7, 0xad, 0x35, 0x86, 0xfc, 0x53, 0xdb, 0xd6, 0x27, 0x72, 0x4e, 0x34, 0xd4, 0xbb, 0x83, 0x68, 0x29, 0xbd, 0xa8, 0xa0, 0xb9, 0x5b, 0x3c, 0xe3, 0xae, 0x2b, 0xd6, 0x2e, 0xb0, 0x6, 0x10, 0xa4, 0x3b, 0xb2, 0x7b, 0x15, 0x99, 0x8f, 0x83, 0x7d, 0xa7, 0x28, 0x82, 0x5f, 0x37, 0x16, 0x38, 0x7a, 0x9f, 0x49, 0xf9, 0x7c, 0xd2, 0xad, 0xd4, 0xf, 0x5b, 0xac, 0x90, 0x33, 0x83, 0x46, 0x29, 0xf, 0xae, 0x36, 0x3e, 0x32, 0x95, 0x85, 0x25, 0x66, 0xa3, 0x29, 0xc6, 0xa5, 0x1f, 0xb, 0xcc, 0xa1, 0x54, 0x63, 0x8a, 0xe, 0xa2, 0xd, 0x9e, 0x5d, 0x70, 0xd3, 0x2e, 0xeb, 0xbe, 0x71, 0x39, 0xf7, 0x7, 0xc0, 0x7e, 0xba, 0xf2, 0xdb, 0x69, 0xd, 0xba, 0xfd, 0x27, 0xba, 0x64, 0x6e, 0x89, 0xe4, 0x17, 0x9b, 0x5c, 0xc3, 0xb9, 0x23, 0x16, 0xe0, 0xf6, 0xae, 0x8, 0x28, 0x21, 0x46, 0x62, 0x7f, 0x5f, 0x69, 0x9, 0x4c, 0xf3, 0x6a, 0x78, 0xbc, 0xf3, 0x1b, 0x0, 0x48, 0xda, 0xbb, 0x1e, 0x1d, 0xdb, 0x6e, 0xdf, 0x84, 0x70, 0xa1, 0x2a, 0xac, 0x11, 0x3, 0x32, 0x95, 0x16, 0x36, 0xcf, 0x0, 0xbb, 0xd0, 0x43, 0x7, 0xc1, 0x19, 0x24, 0x5a, 0x6b, 0x7d, 0x70, 0x4a, 0x62, 0x8e, 0x1a, 0xa8, 0x82, 0xca, 0x58, 0x10, 0x4e, 0x21, 0x24, 0xff, 0xae, 0x7, 0x49, 0x12, 0xab, 0x19, 0x96, 0x3e, 0x6, 0x5, 0x3d, 0x82, 0x81, 0xf4, 0x42, 0xc1, 0x31, 0x1a, 0x70, 0x18, 0xe3, 0x4c, 0x86, 0x22, 0x4d, 0xd9, 0x65, 0xbd, 0x4f, 0xc3, 0x14, 0xff, 0x76, 0x95, 0xb0, 0xe2, 0x8f, 0x52, 0xde, 0xd7, 0x4e, 0x55, 0xb4, 0xf1, 0x7e, 0x9, 0x9e, 0xf, 0x55, 0xd1, 0xff, 0x8, 0x8c, 0xb8, 0x51, 0x82, 0x73, 0xde, 0xc6, 0xf1, 0x23, 0x10, 0x40, 0xed, 0xbf, 0x25, 0x60, 0x14, 0x4e, 0x18, 0x9f, 0x58, 0x82, 0xf, 0x80, 0xa7, 0x55, 0x2d, 0x87, 0xf7, 0x5, 0x53, 0x60, 0xf0, 0xdf, 0xe8, 0x42, 0xab, 0x4f, 0xcc, 0xc5, 0xdc, 0x6e, 0x25, 0x35, 0x81, 0xea, 0x1a, 0x5, 0x81, 0x7a, 0x36, 0x71, 0xe8, 0xef, 0x9f, 0x33, 0xb0, 0xf0, 0x82, 0x14, 0xe5, 0xec, 0xb9, 0x66, 0x6e, 0xbf, 0x12, 0x1d, 0x6, 0x3f, 0x5a, 0x17, 0xfd, 0x6, 0x90, 0x3f, 0x22, 0x7c, 0xb1, 0x7d, 0x31, 0x4f, 0x7c, 0x5a, 0x5a, 0xce, 0x3c, 0x70, 0x97, 0xb2, 0x3e, 0x90, 0x89, 0x17, 0x48, 0xb9, 0x8c, 0x26, 0x58, 0xa, 0xff, 0x26, 0x43, 0xd1, 0xcd, 0x52, 0xe9, 0x64, 0xd2, 0x79, 0xc5, 0x8c, 0xa3, 0x10, 0x3f, 0x34, 0xa3, 0xd8, 0x14, 0xfa, 0x75, 0xc9, 0x7c, 0x58, 0xc4, 0x6, 0x7f, 0xac, 0xe7, 0x81, 0x5b, 0x21, 0xab, 0x2b, 0x25, 0x8f, 0x1e, 0xa9, 0xcd, 0xef, 0xd4, 0xd0, 0x87, 0x75, 0xe3, 0x4b, 0xe6, 0x8c, 0x6e, 0x81, 0x4, 0xb3, 0x8e, 0xf, 0xd1, 0x24, 0x16, 0x48, 0xa4, 0x5d, 0x17, 0x6d, 0xf1, 0x9a, 0xfd, 0x92, 0x8e, 0xf9, 0x69, 0x8, 0x7b, 0x5a, 0x85, 0x8, 0xa1, 0x40, 0x3, 0x85, 0x88, 0x8e, 0x17, 0xa8, 0x86, 0x71, 0x37, 0xc9, 0x7e, 0x45, 0x75, 0xe, 0xeb, 0xdb, 0xd3, 0x6, 0x3c, 0x49, 0x5e, 0x2b, 0xfa, 0xc7, 0x25, 0x5c, 0x34, 0x26, 0x7b, 0xf4, 0xfb, 0x44, 0x5f, 0xd9, 0xcb, 0x22, 0xab, 0x53, 0xec, 0xd7, 0x1e, 0xb5, 0x9f, 0x15, 0x69, 0x95, 0x7, 0xcd, 0x2b, 0x81, 0x4a, 0xa7, 0x37, 0x8c, 0xe1, 0xc5, 0x6c, 0xce, 0xb3, 0xf8, 0xa8, 0x9b, 0x4d, 0xb6, 0x53, 0x85, 0x1, 0x25, 0xd, 0x32, 0x60, 0x7e, 0xdc, 0xf6, 0x6c, 0xe4, 0xf8, 0x3, 0x9f, 0x8a, 0x6b, 0xdf, 0x71, 0x96, 0x2f, 0x24, 0x98, 0x19, 0xcd, 0xda, 0x42, 0xd8, 0x75, 0x91, 0x1a, 0x32, 0xcd, 0xd8, 0x8c, 0xfe, 0x19, 0xc3, 0x58, 0xec, 0x30, 0xc9, 0xa2, 0xc5, 0x5e, 0x88, 0xde, 0xb5, 0xcb, 0x18, 0x5a, 0x1c, 0x30, 0x7b, 0xa3, 0xb0, 0x5c, 0xd3, 0x3d, 0x54, 0xe7, 0xbb, 0x85, 0x97, 0x75, 0x9c, 0x91, 0x6e, 0x66, 0x2a, 0x2d, 0xf2, 0x7b, 0x4c, 0xfc, 0xe, 0xf9, 0x3d, 0xa0, 0x73, 0xd2, 0x80, 0xb0, 0xd, 0xb0, 0xc8, 0x3b, 0x18, 0x31, 0x31, 0xbe, 0x64, 0xbd, 0x9c, 0xd1, 0x55, 0x2d, 0x86, 0xf8, 0xb9, 0x7c, 0x6, 0xcd, 0x3e, 0xac, 0xcc, 0xba, 0x1b, 0x35, 0xca, 0xda, 0xa6, 0xbc, 0x9e, 0xab, 0xc7, 0xed, 0xa3, 0xcf, 0x34, 0x7, 0x37, 0x33, 0x89, 0x5f, 0x73, 0xc8, 0xa0, 0x19, 0x3e, 0x45, 0x37, 0x3b, 0x3f, 0xc4, 0xf3, 0x38, 0x16, 0x72, 0xea, 0x4a, 0x21, 0x74, 0x0, 0xc3, 0xee, 0x41, 0x7c, 0x81, 0x5a, 0xcc, 0xab, 0x46, 0xd3, 0x1b, 0xe3, 0x67, 0xdf, 0xd5, 0x43, 0xc3, 0x48, 0xf5, 0xbc, 0xe8, 0xfb, 0x70, 0x65, 0x72, 0x91, 0x41, 0x2d, 0x92, 0x61, 0x38, 0x21, 0x6c, 0x27, 0xb6, 0x32, 0xf1, 0x9a, 0xd7, 0x16, 0xf9, 0x97, 0xb, 0xd6, 0x7b, 0x93, 0xc4, 0xa1, 0x63, 0x2f, 0x4f, 0xd9, 0xe3, 0xd6, 0x97, 0xa0, 0x6d, 0x1e, 0xee, 0xe8, 0xb8, 0xbc, 0x46, 0x85, 0x99, 0x75, 0x3, 0xc1, 0xf, 0x3f, 0xb3, 0x68, 0xe9, 0x6b, 0x41, 0x35, 0x9b, 0x46, 0x1e, 0xe6, 0x4c, 0x27, 0x67, 0x7f, 0x2a, 0xb, 0x83, 0x55, 0x22, 0x4f, 0xfc, 0x69, 0x52, 0x28, 0xf5, 0x91, 0x49, 0x24, 0xbf, 0xd, 0xd9, 0x24, 0x74, 0x31, 0xe1, 0xa2, 0x9c, 0xf4, 0x2e, 0x83, 0x55, 0xaf, 0x12, 0xd1, 0x77, 0xbb, 0xcd, 0x1f, 0xa7, 0xb1, 0x44, 0xd0, 0x29, 0x14, 0x5b, 0xc6, 0x34, 0xd0, 0x5b, 0x2d, 0x6d, 0xf3, 0xd5, 0x2c, 0x97, 0xb5, 0xc1, 0x1, 0x8b, 0x29, 0x32, 0x76, 0x54, 0xde, 0x66, 0xb2, 0x26, 0xbf, 0x74, 0xa2, 0x8e, 0x98, 0xd6, 0xa6, 0x9a, 0x72, 0x63, 0x14, 0xee, 0x66, 0x30, 0x1c, 0x20, 0x5b, 0x35, 0x3c, 0x1b, 0x50, 0xde, 0xe2, 0x2f, 0x5, 0x36, 0x35, 0x61, 0x68, 0xd9, 0x23, 0xa6, 0x63, 0x6d, 0x78, 0x88, 0x6a, 0x8d, 0x41, 0x42, 0xc6, 0x49, 0x42, 0xc4, 0xaf, 0x7d, 0x8e, 0xb, 0x0, 0x2f, 0x19, 0xe0, 0x90, 0xfd, 0x95, 0x3b, 0xa9, 0x22, 0xb9, 0x78, 0x97, 0x3b, 0x20, 0xf3, 0x10, 0xd, 0xb9, 0x96, 0x2, 0xb7, 0xd8, 0x0, 0x5f, 0x6c, 0x52, 0xb7, 0xd1, 0x86, 0x7a, 0xb1, 0x40, 0x5, 0x23, 0xf8, 0x90, 0xaf, 0xa5, 0x83, 0x29, 0x29, 0x31, 0x25, 0xbe, 0x6d, 0xee, 0x5, 0x3a, 0xb1, 0xb0, 0xc7, 0xe6, 0xe5, 0x80, 0xd6, 0x72, 0x1c, 0x73, 0x87, 0xed, 0x81, 0xf1, 0x46, 0x10, 0x6, 0xd8, 0x90, 0x27, 0xcb, 0x44, 0x4f, 0x40, 0x49, 0x7e, 0x2d, 0xd9, 0x7f, 0xd, 0x2, 0xc8, 0x28, 0xc0, 0x73, 0xf5, 0x93, 0x38, 0xf5, 0xce, 0x19, 0xeb, 0xed, 0x65, 0xe, 0x54, 0x5e, 0x33, 0xd7, 0xdc, 0xb0, 0xb5, 0xa0, 0x62, 0xb0, 0xde, 0xfb, 0x0, 0x8c, 0xf8, 0xad, 0x5d, 0x9b, 0xef, 0xe8, 0xd8, 0x6b, 0x74, 0x85, 0x2b, 0x1a, 0xcb, 0xd4, 0x62, 0x19, 0x17, 0x1a, 0x90, 0x7c, 0xae, 0xdc, 0xcd, 0x7f, 0x71, 0xd2, 0xba, 0x22, 0x6, 0x1f, 0x80, 0xaf, 0xca, 0x1, 0xa, 0x15, 0x32, 0x92, 0x93, 0x14, 0x65, 0x20, 0xdb, 0xee, 0x17, 0x67, 0xa5, 0x41, 0x59, 0xbc, 0xee, 0xe4, 0x3f, 0xad, 0x1c, 0x32, 0xc6, 0xad, 0xb1, 0x51, 0x20, 0xc3, 0xeb, 0xa5, 0xc4, 0xa4, 0xc7, 0x6e, 0xff, 0xd4, 0x83, 0x6c, 0x8a, 0xeb, 0xee, 0x22, 0x12, 0x5d, 0xc5, 0xc2, 0x88, 0x3a, 0xcb, 0xa3, 0xf5, 0x58, 0x10, 0x12, 0x98, 0xb0, 0xa6, 0xbb, 0xc8, 0xf9, 0xb1, 0xd7, 0xf5, 0x8b, 0x92, 0x48, 0x61, 0x6d, 0xc5, 0x42, 0x8c, 0x3, 0x17, 0x55, 0xd7, 0x2a, 0x95, 0x5, 0xd9, 0x2a, 0x51, 0x52, 0x12, 0x18, 0x55, 0x6e, 0x2e, 0xc8, 0xfa, 0x41, 0x73, 0xb4, 0x53, 0x42, 0x85, 0xb7, 0xac, 0xca, 0x44, 0x2b, 0x3f, 0xac, 0xc7, 0x12, 0x4d, 0xb1, 0x1d, 0x80, 0xd2, 0xe5, 0x85, 0x2a, 0xd, 0x19, 0x84, 0x38, 0xcb, 0x4f, 0x52, 0xc7, 0x1a, 0x39, 0x88, 0xf0, 0xed, 0xde, 0xf4, 0x14, 0x26, 0x24, 0x8c, 0xcd, 0x11, 0xcd, 0x58, 0x70, 0x3, 0x87, 0x8f, 0xb1, 0xe4, 0x98, 0x49, 0xdf, 0xef, 0xe6, 0x5b, 0xfc, 0xff, 0x0, 0xba, 0xa2, 0x32, 0xe6, 0xf2, 0x6, 0x3e, 0xb0, 0x6f, 0xc3, 0xab, 0xcb, 0x3a, 0xc3, 0x5c, 0x5f, 0x9d, 0x68, 0x70, 0x16, 0x68, 0xf, 0x17, 0x7f, 0xd9, 0x4f, 0xfd, 0x3, 0x2d, 0x5b, 0x97, 0xc9, 0x88, 0xe, 0x5c, 0x8a, 0x4d, 0xfb, 0x4b, 0x86, 0xed, 0x9a, 0x9, 0xa7, 0x7, 0xec, 0xa7, 0x87, 0x88, 0xe, 0x27, 0x80, 0xe2, 0x55, 0x47, 0x15, 0xe2, 0x2e, 0xa8, 0x4d, 0x29, 0xb, 0x20, 0x1a, 0x79, 0xae, 0x37, 0x68, 0x9f, 0x64, 0xad, 0x61, 0xcf, 0x2, 0x15, 0x7f, 0x90, 0xf5, 0x94, 0x35, 0xc0, 0xc8, 0x13, 0x22, 0x30, 0xfe, 0xbb, 0x7d, 0x92, 0x80, 0x30, 0x70, 0x27, 0xfc, 0x8e, 0xb6, 0x2a, 0xe8, 0xfa, 0x42, 0xe3, 0x84, 0xdb, 0xe4, 0x66, 0x5a, 0x23, 0xf9, 0x55, 0x62, 0x1a, 0xe5, 0xa7, 0xf, 0x64, 0x5e, 0x66, 0x11, 0x81, 0x2d, 0x9c, 0xb3, 0x41, 0x4f, 0x3e, 0x8b, 0x66, 0xa2, 0x75, 0x72, 0x7, 0x80, 0xde, 0xd3, 0xd9, 0xbd, 0x4, 0xb8, 0x9c, 0x8b, 0x67, 0xdf, 0x48, 0x9, 0xb1, 0x88, 0xf0, 0x74, 0x5c, 0xa, 0xa6, 0x82, 0xba, 0x38, 0x72, 0x29, 0x1d, 0xa7, 0x46, 0xb6, 0xae, 0x72, 0x4e, 0x3c, 0xde, 0x2, 0x9b, 0x47, 0xef, 0xc9, 0x4e, 0x11, 0x78, 0xdf, 0x79, 0xa0, 0x64, 0xfe, 0x5e, 0xdc, 0x7b, 0xb5, 0xad, 0x39, 0x1e, 0xe0, 0x8d, 0x2c, 0x5e, 0xa3, 0x98, 0x3e, 0xd2, 0x4a, 0x8b, 0x17, 0x19, 0x4b, 0xbe, 0x75, 0xf4, 0xa, 0x12, 0xf0, 0x31, 0x99, 0xe7, 0x82, 0x29, 0xe9, 0xef, 0x11, 0x43, 0xdf, 0x96, 0x6, 0x3e, 0x32, 0xe5, 0x52, 0x12, 0x98, 0xc6, 0x61, 0x7c, 0xee, 0x7c, 0xda, 0x99, 0x8b, 0x19, 0xcd, 0x83, 0x10, 0xec, 0xd9, 0xcb, 0xf7, 0x1d, 0xfc, 0x23, 0x95, 0xf9, 0xa, 0x61, 0x47, 0x69, 0xd5, 0x55, 0x5, 0x63, 0x99, 0x71, 0xe8, 0x13, 0x14, 0x9c, 0x27, 0x9, 0x8, 0x78, 0x42, 0xe6, 0xbd, 0x59, 0x26, 0xa8, 0x6, 0x17, 0xff, 0xf8, 0x99, 0xb3, 0x74, 0xdd, 0x70, 0x5e, 0x23, 0xec, 0x36, 0x65, 0x83, 0x67, 0xec, 0x81, 0xa4, 0x71, 0xf6, 0x3e, 0x19, 0x63, 0x95, 0xfd, 0x1, 0x44, 0x3d, 0x54, 0x1c, 0xf4, 0x15, 0xe1, 0xc, 0x97, 0x4c, 0x40, 0xee, 0x83, 0x72, 0xa, 0x56, 0x82, 0x52, 0x12, 0xe7, 0x56, 0xe2, 0xf, 0x51, 0xa, 0xac, 0xa0, 0x81, 0xaa, 0xda, 0x91, 0x2a, 0x61, 0xa3, 0x80, 0x2a, 0x9f, 0x94, 0x98, 0xe4, 0xf3, 0xa, 0x1, 0x58, 0xd9, 0x97, 0x8e, 0x74, 0x5d, 0xdd, 0x70, 0x5c, 0x1a, 0x41, 0x16, 0x27, 0xdb, 0x3e, 0xdd, 0x1e, 0xd2, 0xce, 0xd2, 0x70, 0xbb, 0x8b, 0x5d, 0x4c, 0xa1, 0x66, 0x9, 0xa9, 0x2a, 0x60, 0xa7, 0xb8, 0xda, 0xb2, 0xbe, 0xf4, 0xcc, 0x37, 0x4d, 0xad, 0x97, 0x9e, 0x62, 0xc8, 0xcf, 0x1b, 0xf6, 0x72, 0x10, 0x3f, 0xbf, 0x9e, 0x21, 0x96, 0x32, 0xe2, 0xf1, 0x8d, 0x82, 0x6d, 0xcf, 0xea, 0x41, 0xdd, 0xc, 0xb6, 0xa0, 0x39, 0x64, 0x9f, 0x37, 0xee, 0xa6, 0x6, 0x62, 0x91, 0xb4, 0x13, 0x79, 0x4d, 0x15, 0x53, 0xd4, 0xa3, 0x2e, 0xd3, 0x54, 0x15, 0x3e, 0x9a, 0x4, 0xfc, 0x85, 0xfc, 0x30, 0x97, 0x9f, 0x19, 0xac, 0x2, 0xbb, 0xe2, 0x8c, 0xae, 0x6f, 0xa0, 0x16, 0x41, 0xe5, 0xd8, 0xf8, 0xb5, 0x5a, 0x4a, 0x7f, 0xc7, 0x73, 0xf2, 0xd4, 0x7b, 0x18, 0xf0, 0x9a, 0x61, 0x17, 0xb3, 0xfd, 0xc5, 0x8, 0x23, 0x96, 0x10, 0x2b, 0x3e, 0xc, 0x57, 0x66, 0x35, 0xc, 0x93, 0x83, 0xb1, 0x13, 0xf7, 0xcf, 0xa2, 0xf7, 0x4d, 0x29, 0xfb, 0x2f, 0x96, 0xef, 0x69, 0x48, 0x86, 0x4d, 0xc0, 0x7f, 0xa3, 0x4c, 0x5d, 0xc7, 0x96, 0x36, 0x80, 0xeb, 0xa9, 0x69, 0x7a, 0x5d, 0x83, 0x52, 0x28, 0xae, 0x51, 0xa6, 0x3e, 0xf9, 0x7a, 0x9d, 0x77, 0xcf, 0x3b, 0x87, 0x2e, 0x3, 0xf1, 0x13, 0x5d, 0x5f, 0x2d, 0x93, 0xa5, 0xec, 0x2c, 0x7a, 0xa2, 0x40, 0xf, 0x8f, 0xf7, 0x72, 0xb5, 0xe2, 0xc2, 0x95, 0x2e, 0x15, 0xb8, 0x75, 0xa, 0x16, 0x63, 0x21, 0xe7, 0x2c, 0x20, 0xd1, 0xde, 0x4f, 0x49, 0xaa, 0xf3, 0x1f, 0x1f, 0x50, 0xf9, 0xa4, 0x19, 0xd7, 0xdf, 0x86, 0x1a, 0x65, 0x35, 0x42, 0x1e, 0x69, 0xcb, 0x9, 0x92, 0x6f, 0xa5, 0x10, 0xe3, 0x8a, 0xd, 0xad, 0xfa, 0x19, 0x98, 0xb1, 0xfb, 0x7, 0xa8, 0x7f, 0x86, 0x83, 0x40, 0xb1, 0x51, 0x3d, 0x86, 0x8c, 0x2e, 0x9d, 0xac, 0x72, 0x3b, 0x93, 0xf0, 0xf8, 0x28, 0x55, 0x33, 0x8e, 0xfc, 0x74, 0x29, 0xfc, 0xa7, 0xd2, 0x66, 0x28, 0xc2, 0xcf, 0xf3, 0x42, 0xb5, 0xe3, 0x7e, 0x32, 0x5, 0x66, 0xfa, 0xea, 0xe0, 0xf7, 0x7d, 0xa8, 0xf5, 0xcc, 0x3a, 0xb9, 0x9c, 0xb1, 0x33, 0x6c, 0x75, 0xee, 0xb2, 0xb4, 0x1, 0x8d, 0x3a, 0xaf, 0xa3, 0xd0, 0xa3, 0x75, 0x5e, 0x2c, 0x40, 0x5e, 0x42, 0xae, 0x9c, 0xab, 0x8f, 0x8d, 0x2a, 0xb3, 0xf5, 0xfe, 0x0, 0x92, 0xd6, 0x24, 0x63, 0x3d, 0x3c, 0xd9, 0xc, 0xe6, 0x7d, 0x98, 0xc4, 0xd4, 0xdf, 0xb9, 0x2c, 0xe1, 0x3a, 0xb9, 0x96, 0x92, 0x7c, 0xd1, 0xa8, 0x7a, 0xb2, 0x3b, 0x24, 0xc7, 0x6, 0xb4, 0x51, 0x2b, 0xd3, 0x55, 0xf3, 0x6a, 0xda, 0x82, 0x55, 0x38, 0xb3, 0x1f, 0x96, 0x66, 0x35, 0xb0, 0xf8, 0x3a, 0xba, 0x44, 0xcf, 0x42, 0x7c, 0xf7, 0xb8, 0x91, 0xda, 0x11, 0x84, 0x2e, 0x75, 0xd, 0xc9, 0xaf, 0xce, 0x2d, 0xbf, 0x86, 0xb4, 0x6f, 0x9a, 0x19, 0x60, 0x33, 0x16, 0x9b, 0x34, 0x2d, 0x4, 0x4c, 0x53, 0x12, 0x79, 0x85, 0x7f, 0x2, 0x7d, 0xf, 0xbf, 0x1, 0xc1, 0x94, 0x17, 0xd8, 0xf2, 0x93, 0x0, 0xf3, 0x33, 0x18, 0x9d, 0xea, 0x20, 0x49, 0x88, 0x0, 0xc7, 0x3c, 0xc8, 0x99, 0x8e, 0x16, 0xd5, 0x20, 0xa4, 0x89, 0xb5, 0x5c, 0x70, 0xfc, 0xbf, 0x1e, 0x11, 0x32, 0x27, 0x5e, 0xb1, 0x6a, 0xf8, 0x99, 0x90, 0xe4, 0x12, 0x18, 0xde, 0x55, 0x34, 0xaf, 0x42, 0x8c, 0x2b, 0x2e, 0xad, 0xed, 0xc2, 0xc3, 0xb1, 0x3, 0xca, 0xd0, 0xbc, 0x13, 0x6d, 0xda, 0xcc, 0xca, 0x4e, 0x3a, 0x58, 0x6d, 0x0, 0xd3, 0xca, 0xf1, 0x2e, 0x63, 0x96, 0xfe, 0x54, 0xfb, 0x71, 0x77, 0xfe, 0x3f, 0x5a, 0x53, 0xa, 0xc0, 0x88, 0x3e, 0x97, 0xf2, 0x3f, 0xb8, 0x6, 0x9c, 0xfb, 0x71, 0x41, 0x6c, 0x82, 0x46, 0x74, 0x9a, 0x16, 0x92, 0xaa, 0x68, 0x61, 0x3f, 0xae, 0x1a, 0x15, 0xea, 0xee, 0x16, 0xa3, 0x46, 0xa6, 0xc9, 0x53, 0x22, 0x36, 0xce, 0x5b, 0xec, 0x7d, 0x38, 0xf, 0xe, 0xf0, 0x8a, 0xcc, 0xd7, 0xd3, 0x94, 0xf1, 0x29, 0x86, 0x64, 0x49, 0xc0, 0x2a, 0x39, 0x65, 0x24, 0x46, 0x4b, 0xe6, 0x21, 0xba, 0xe4, 0x51, 0x39, 0x88, 0xac, 0x25, 0x4a, 0xa3, 0xc0, 0x8d, 0xc7, 0xa9, 0x29, 0xe3, 0xe0, 0x6e, 0xe0, 0xe0, 0x9b, 0x54, 0xf1, 0xed, 0xa4, 0xcc, 0x8b, 0x8d, 0xc9, 0xd7, 0xe6, 0x27, 0xa9, 0x2d, 0x7d, 0x36, 0x8f, 0x41, 0xdc, 0x63, 0x31, 0x51, 0x7d, 0x83, 0x29, 0x2, 0x1e, 0x23, 0xbc, 0x96, 0x33, 0x6b, 0x39, 0xca, 0xb1, 0xfd, 0xc3, 0xb1, 0xec, 0xe5, 0xbd, 0x60, 0xa5, 0xc6, 0x3b, 0x1b, 0xdf, 0x34, 0x71, 0x87, 0x14, 0xd, 0xfe, 0x4d, 0xce, 0x3c, 0xeb, 0x65, 0x24, 0x5b, 0xee, 0x51, 0xb7, 0x4e}
	assert.Equal(t, expected, b.Bytes())
}

func TestCorner(t *testing.T) {
	sps, pps, err := ParseAVCSeqHeader([]byte{0})
	assert.Equal(t, nil, sps)
	assert.Equal(t, nil, pps)
	assert.Equal(t, err, avcErr)

	b := &bytes.Buffer{}
	err = CaptureAVC(b, []byte{0x17, 0x0, 0x1})
	assert.Equal(t, nil, b.Bytes())
	assert.Equal(t, err, avcErr)
}
