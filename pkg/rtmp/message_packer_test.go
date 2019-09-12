package rtmp

import (
	"bytes"
	"github.com/q191201771/nezha/pkg/assert"
	"github.com/q191201771/nezha/pkg/mockwriter"
	"testing"
)

func TestWriteMessageHandler(t *testing.T) {
	//buf := &bytes.Buffer{}
	packer := NewMessagePacker()
	packer.writeMessageHeader(1, 2, 3, 4)
	assert.Equal(t, []byte{1, 0, 0, 0, 0, 0, 2, 3, 4, 0, 0, 0}, packer.b.Bytes())
}

func TestWrite(t *testing.T) {
	var (
		err    error
		result []byte
	)
	buf := &bytes.Buffer{}
	packer := NewMessagePacker()

	err = packer.writeProtocolControlMessage(buf, 1, 2)
	assert.Equal(t, nil, err)
	assert.Equal(t, []byte{2, 0, 0, 0, 0, 0, 4, 1, 0, 0, 0, 0, 0, 0, 0, 2}, buf.Bytes())
	buf.Reset()

	err = packer.writeChunkSize(buf, 1)
	assert.Equal(t, nil, err)
	assert.Equal(t, []byte{2, 0, 0, 0, 0, 0, 4, 1, 0, 0, 0, 0, 0, 0, 0, 1}, buf.Bytes())
	buf.Reset()

	err = packer.writeWinAckSize(buf, 1)
	assert.Equal(t, nil, err)
	assert.Equal(t, []byte{2, 0, 0, 0, 0, 0, 4, 5, 0, 0, 0, 0, 0, 0, 0, 1}, buf.Bytes())
	buf.Reset()

	err = packer.writePeerBandwidth(buf, 1, 2)
	assert.Equal(t, nil, err)
	assert.Equal(t, []byte{2, 0, 0, 0, 0, 0, 5, 6, 0, 0, 0, 0, 0, 0, 0, 1, 2}, buf.Bytes())
	buf.Reset()

	err = packer.writeConnect(buf, "live", "rtmp://127.0.0.1/live")
	assert.Equal(t, nil, err)
	result = []byte{
		0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x81, 0x14, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x07, 0x63,
		0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x00, 0x3f, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03,
		0x00, 0x03, 0x61, 0x70, 0x70, 0x02, 0x00, 0x04, 0x6c, 0x69, 0x76, 0x65, 0x00, 0x04, 0x74, 0x79,
		0x70, 0x65, 0x02, 0x00, 0x0a, 0x6e, 0x6f, 0x6e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x00,
		0x08, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x56, 0x65, 0x72, 0x02, 0x00, 0x1f, 0x46, 0x4d, 0x4c, 0x45,
		0x2f, 0x33, 0x2e, 0x30, 0x20, 0x28, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x6c, 0x65,
		0x3b, 0x20, 0x4c, 0x61, 0x6c, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x29, 0x00, 0x05, 0x74, 0x63, 0x55,
		0x72, 0x6c, 0x02, 0x00, 0x15, 0x72, 0x74, 0x6d, 0x70, 0x3a, 0x2f, 0x2f, 0x31, 0x32, 0x37, 0x2e,
		0x30, 0x2e, 0x30, 0x2e, 0x31, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x00, 0x00, 0x09,
	}
	assert.Equal(t, result, buf.Bytes())
	buf.Reset()

	err = packer.writeConnectResult(buf, 1)
	assert.Equal(t, nil, err)
	result = []byte{0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0xbe, 0x14, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x7, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x0, 0x3f, 0xf0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0x0, 0x6, 0x66, 0x6d, 0x73, 0x56, 0x65, 0x72, 0x2, 0x0, 0xd, 0x46, 0x4d, 0x53, 0x2f, 0x33, 0x2c, 0x30, 0x2c, 0x31, 0x2c, 0x31, 0x32, 0x33, 0x0, 0xc, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x0, 0x40, 0x3f, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x9, 0x3, 0x0, 0x5, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x2, 0x0, 0x6, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x0, 0x4, 0x63, 0x6f, 0x64, 0x65, 0x2, 0x0, 0x1d, 0x4e, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x0, 0xb, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2, 0x0, 0x15, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x2e, 0x0, 0xe, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x9}
	assert.Equal(t, result, buf.Bytes())
	buf.Reset()

	err = packer.writeCreateStream(buf)
	assert.Equal(t, nil, err)
	result = []byte{0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x19, 0x14, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0xc, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x0, 0x40, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5}
	assert.Equal(t, result, buf.Bytes())
	buf.Reset()

	err = packer.writeCreateStreamResult(buf, 1)
	assert.Equal(t, nil, err)
	result = []byte{0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1d, 0x14, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x7, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x0, 0x3f, 0xf0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5, 0x0, 0x3f, 0xf0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	assert.Equal(t, result, buf.Bytes())
	buf.Reset()

	err = packer.writePlay(buf, "test", 1)
	assert.Equal(t, nil, err)
	result = []byte{0x5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x18, 0x14, 0x1, 0x0, 0x0, 0x0, 0x2, 0x0, 0x4, 0x70, 0x6c, 0x61, 0x79, 0x0, 0x40, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5, 0x2, 0x0, 0x4, 0x74, 0x65, 0x73, 0x74}
	assert.Equal(t, result, buf.Bytes())
	buf.Reset()

	err = packer.writePublish(buf, "live", "test", 1)
	assert.Equal(t, nil, err)
	result = []byte{0x5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x22, 0x14, 0x1, 0x0, 0x0, 0x0, 0x2, 0x0, 0x7, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x0, 0x40, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5, 0x2, 0x0, 0x4, 0x74, 0x65, 0x73, 0x74, 0x2, 0x0, 0x4, 0x6c, 0x69, 0x76, 0x65}
	assert.Equal(t, result, buf.Bytes())
	buf.Reset()

	err = packer.writeOnStatusPublish(buf, 1)
	assert.Equal(t, nil, err)
	result = []byte{0x5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x69, 0x14, 0x1, 0x0, 0x0, 0x0, 0x2, 0x0, 0x8, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5, 0x3, 0x0, 0x5, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x2, 0x0, 0x6, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x0, 0x4, 0x63, 0x6f, 0x64, 0x65, 0x2, 0x0, 0x17, 0x4e, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x0, 0xb, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2, 0x0, 0x10, 0x53, 0x74, 0x61, 0x72, 0x74, 0x20, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x0, 0x0, 0x9}
	assert.Equal(t, result, buf.Bytes())
	buf.Reset()

	err = packer.writeOnStatusPlay(buf, 1)
	assert.Equal(t, nil, err)
	result = []byte{0x5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x60, 0x14, 0x1, 0x0, 0x0, 0x0, 0x2, 0x0, 0x8, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5, 0x3, 0x0, 0x5, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x2, 0x0, 0x6, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x0, 0x4, 0x63, 0x6f, 0x64, 0x65, 0x2, 0x0, 0x14, 0x4e, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x0, 0xb, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2, 0x0, 0xa, 0x53, 0x74, 0x61, 0x72, 0x74, 0x20, 0x6c, 0x69, 0x76, 0x65, 0x0, 0x0, 0x9}
	assert.Equal(t, result, buf.Bytes())
	buf.Reset()
	//
	//var str string
	//for i := 0; i < len(buf.Bytes()); i++ {
	//	a := fmt.Sprintf("0x%x, ", buf.Bytes()[i])
	//	str += a
	//}
	//log.Info(str)
}

func TestPackCorner(t *testing.T) {
	func() {
		defer func() {
			recover()
		}()
		packer := NewMessagePacker()
		// 测试csid超过63的情况
		packer.writeMessageHeader(128, 0, 0, 0)
	}()

	var err error
	mw := mockwriter.NewMockWriter(mockwriter.WriterTypeReturnError)
	packer := NewMessagePacker()

	err = packer.writeProtocolControlMessage(mw, 1, 2)
	assert.IsNotNil(t, err)
	err = packer.writeChunkSize(mw, 1)
	assert.IsNotNil(t, err)
	err = packer.writeWinAckSize(mw, 1)
	assert.IsNotNil(t, err)
	err = packer.writePeerBandwidth(mw, 1, 2)
	assert.IsNotNil(t, err)
	err = packer.writeConnect(mw, "live", "rtmp://127.0.0.1/live")
	assert.IsNotNil(t, err)
	err = packer.writeConnectResult(mw, 1)
	assert.IsNotNil(t, err)
	err = packer.writeCreateStream(mw)
	assert.IsNotNil(t, err)
	err = packer.writeCreateStreamResult(mw, 1)
	assert.IsNotNil(t, err)
	err = packer.writePlay(mw, "test", 1)
	assert.IsNotNil(t, err)
	err = packer.writePublish(mw, "live", "test", 1)
	assert.IsNotNil(t, err)
	err = packer.writeOnStatusPublish(mw, 1)
	assert.IsNotNil(t, err)
	err = packer.writeOnStatusPlay(mw, 1)
	assert.IsNotNil(t, err)
}

func BenchmarkMessagePacker(b *testing.B) {
	mw := mockwriter.NewMockWriter(mockwriter.WriterTypeDoNothing)
	packer := NewMessagePacker()
	for i := 0; i < b.N; i++ {
		_ = packer.writeConnect(mw, "live", "rtmp://127.0.0.1/live")
	}
}
