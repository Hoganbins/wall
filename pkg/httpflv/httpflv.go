// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package httpflv

import (
	"errors"
	"strings"
)

var ErrHTTPFLV = errors.New("lal.httpflv: fxxk")

const (
	TagHeaderSize int = 11

	flvHeaderSize        int = 13
	prevTagSizeFieldSize int = 4
)

const (
	TagTypeMetadata uint8 = 18
	TagTypeVideo    uint8 = 9
	TagTypeAudio    uint8 = 8
)

const (
	frameTypeKey   uint8 = 1
	frameTypeInter uint8 = 2
)

const (
	codecIDAVC  uint8 = 7
	codecIDHEVC uint8 = 12
)

const (
	AVCKeyFrame   = frameTypeKey<<4 | codecIDAVC
	AVCInterFrame = frameTypeInter<<4 | codecIDAVC

	HEVCKeyFrame   = frameTypeKey<<4 | codecIDHEVC
	HEVCInterFrame = frameTypeInter<<4 | codecIDHEVC
)

const (
	AVCPacketTypeSeqHeader uint8 = 0
	AVCPacketTypeNalu      uint8 = 1

	HEVCPacketTypeSeqHeader uint8 = 0
	HEVCPacketTypeNalu      uint8 = 1

	AACPacketTypeSeqHeader uint8 = 0
	AACPacketTypeRaw       uint8 = 1
)

const (
	SoundFormatAAC uint8 = 10
)

type LineReader interface {
	ReadLine() (line []byte, isPrefix bool, err error)
}

// @return firstLine: request 的 request line 或 response 的 status line
// @return headers: 头中的键值对
func parseHTTPHeader(r LineReader) (firstLine string, headers map[string]string, err error) {
	headers = make(map[string]string)

	var line []byte
	var isPrefix bool
	line, isPrefix, err = r.ReadLine()
	if err != nil {
		return
	}
	if len(line) == 0 || isPrefix {
		err = ErrHTTPFLV
		return
	}
	firstLine = string(line)

	for {
		line, isPrefix, err = r.ReadLine()
		if len(line) == 0 { // 读到一个空的 \r\n 表示http头全部读取完毕了
			break
		}
		if isPrefix {
			err = ErrHTTPFLV
			return
		}
		if err != nil {
			return
		}
		l := string(line)
		pos := strings.Index(l, ":")
		if pos == -1 {
			err = ErrHTTPFLV
			return
		}
		headers[strings.Trim(l[0:pos], " ")] = strings.Trim(l[pos+1:], " ")
	}
	return
}

func parseRequestLine(line string) (method string, uri string, version string, err error) {
	items := strings.Split(line, " ")
	if len(items) != 3 {
		err = ErrHTTPFLV
		return
	}
	return items[0], items[1], items[2], nil
}

func parseStatusLine(line string) (version string, statusCode string, reason string, err error) {
	items := strings.Split(line, " ")
	if len(items) != 3 {
		err = ErrHTTPFLV
		return
	}
	return items[0], items[1], items[2], nil
}
