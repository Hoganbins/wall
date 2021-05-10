// Copyright 2021, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtprtcp

import (
	"github.com/q191201771/lal/pkg/avc"
	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/hevc"
	"github.com/q191201771/naza/pkg/bele"
	"github.com/q191201771/naza/pkg/nazalog"
)

type RTPUnpackerAVCHEVC struct {
	payloadType base.AVPacketPT
	clockRate   int
	onAVPacket  OnAVPacket
}

func NewRTPUnpackerAVCHEVC(payloadType base.AVPacketPT, clockRate int, onAVPacket OnAVPacket) *RTPUnpackerAVCHEVC {
	return &RTPUnpackerAVCHEVC{
		payloadType: payloadType,
		clockRate:   clockRate,
		onAVPacket:  onAVPacket,
	}
}

func (unpacker *RTPUnpackerAVCHEVC) CalcPositionIfNeeded(pkt *RTPPacket) {
	switch unpacker.payloadType {
	case base.AVPacketPTAVC:
		calcPositionIfNeededAVC(pkt)
	case base.AVPacketPTHEVC:
		calcPositionIfNeededHEVC(pkt)
	}
}

func (unpacker *RTPUnpackerAVCHEVC) TryUnpackOne(list *RTPPacketList) (unpackedFlag bool, unpackedSeq uint16) {
	first := list.Head.Next
	if first == nil {
		return false, 0
	}

	switch first.Packet.positionType {
	case PositionTypeSingle:
		var pkt base.AVPacket
		pkt.PayloadType = unpacker.payloadType
		pkt.Timestamp = first.Packet.Header.Timestamp / uint32(unpacker.clockRate/1000)

		pkt.Payload = make([]byte, len(first.Packet.Raw)-int(first.Packet.Header.payloadOffset)+4)
		bele.BEPutUint32(pkt.Payload, uint32(len(first.Packet.Raw))-first.Packet.Header.payloadOffset)
		copy(pkt.Payload[4:], first.Packet.Raw[first.Packet.Header.payloadOffset:])

		list.Head.Next = first.Next
		list.Size--
		unpacker.onAVPacket(pkt)
		return true, first.Packet.Header.Seq

	case PositionTypeSTAPA:
		var pkt base.AVPacket
		pkt.PayloadType = unpacker.payloadType
		pkt.Timestamp = first.Packet.Header.Timestamp / uint32(unpacker.clockRate/1000)

		// 跳过首字节，并且将多nalu前的2字节长度，替换成4字节长度
		buf := first.Packet.Raw[first.Packet.Header.payloadOffset+1:]

		// 使用两次遍历，第一次遍历找出总大小，第二次逐个拷贝，目的是使得内存块一次就申请好，不用动态扩容造成额外性能开销
		totalSize := 0
		for i := 0; i != len(buf); {
			if len(buf)-i < 2 {
				nazalog.Errorf("invalid STAP-A packet.")
				return false, 0
			}
			naluSize := int(bele.BEUint16(buf[i:]))
			totalSize += 4 + naluSize
			i += 2 + naluSize
		}

		pkt.Payload = make([]byte, totalSize)
		j := 0
		for i := 0; i != len(buf); {
			naluSize := int(bele.BEUint16(buf[i:]))
			bele.BEPutUint32(pkt.Payload[j:], uint32(naluSize))
			copy(pkt.Payload[j+4:], buf[i+2:i+2+naluSize])
			j += 4 + naluSize
			i += 2 + naluSize
		}

		list.Head.Next = first.Next
		list.Size--
		unpacker.onAVPacket(pkt)

		return true, first.Packet.Header.Seq

	case PositionTypeFUAStart:
		prev := first
		p := first.Next
		for {
			if prev == nil || p == nil {
				return false, 0
			}
			if SubSeq(p.Packet.Header.Seq, prev.Packet.Header.Seq) != 1 {
				return false, 0
			}

			if p.Packet.positionType == PositionTypeFUAMiddle {
				prev = p
				p = p.Next
				continue
			} else if p.Packet.positionType == PositionTypeFUAEnd {
				var pkt base.AVPacket
				pkt.PayloadType = unpacker.payloadType
				pkt.Timestamp = p.Packet.Header.Timestamp / uint32(unpacker.clockRate/1000)

				var naluTypeLen int
				var naluType []byte
				if unpacker.payloadType == base.AVPacketPTAVC {
					naluTypeLen = 1
					naluType = make([]byte, naluTypeLen)

					fuIndicator := first.Packet.Raw[first.Packet.Header.payloadOffset]
					fuHeader := first.Packet.Raw[first.Packet.Header.payloadOffset+1]
					naluType[0] = (fuIndicator & 0xE0) | (fuHeader & 0x1F)
				} else {
					naluTypeLen = 2
					naluType = make([]byte, naluTypeLen)

					buf := first.Packet.Raw[first.Packet.Header.payloadOffset:]
					fuType := buf[2] & 0x3f
					naluType[0] = (buf[0] & 0x81) | (fuType << 1)
					naluType[1] = buf[1]
				}

				// 使用两次遍历，第一次遍历找出总大小，第二次逐个拷贝，目的是使得内存块一次就申请好，不用动态扩容造成额外性能开销
				totalSize := 0
				pp := first
				for {
					totalSize += len(pp.Packet.Raw) - int(pp.Packet.Header.payloadOffset) - (naluTypeLen + 1)
					if pp == p {
						break
					}
					pp = pp.Next
				}

				pkt.Payload = make([]byte, totalSize+4+naluTypeLen)
				bele.BEPutUint32(pkt.Payload, uint32(totalSize+naluTypeLen))
				var index int
				if unpacker.payloadType == base.AVPacketPTAVC {
					pkt.Payload[4] = naluType[0]
					index = 5
				} else {
					pkt.Payload[4] = naluType[0]
					pkt.Payload[5] = naluType[1]
					index = 6
				}
				packetCount := 0
				pp = first
				for {
					copy(pkt.Payload[index:], pp.Packet.Raw[int(pp.Packet.Header.payloadOffset)+(naluTypeLen+1):])
					index += len(pp.Packet.Raw) - int(pp.Packet.Header.payloadOffset) - (naluTypeLen + 1)
					packetCount++

					if pp == p {
						break
					}
					pp = pp.Next
				}

				list.Head.Next = p.Next
				list.Size -= packetCount
				unpacker.onAVPacket(pkt)

				return true, p.Packet.Header.Seq
			} else {
				// 不应该出现其他类型
				nazalog.Errorf("invalid position type. position=%d", p.Packet.positionType)
				return false, 0
			}
		}

	case PositionTypeFUAMiddle:
		// noop
	case PositionTypeFUAEnd:
		// noop
	default:
		nazalog.Errorf("invalid position. pos=%d", first.Packet.positionType)
	}

	return false, 0
}
func calcPositionIfNeededAVC(pkt *RTPPacket) {
	b := pkt.Raw[pkt.Header.payloadOffset:]

	// rfc3984 5.3.  NAL Unit Octet Usage
	//
	// +---------------+
	// |0|1|2|3|4|5|6|7|
	// +-+-+-+-+-+-+-+-+
	// |F|NRI|  Type   |
	// +---------------+

	outerNALUType := avc.ParseNALUType(b[0])
	if outerNALUType <= NALUTypeAVCSingleMax {
		pkt.positionType = PositionTypeSingle
		return
	} else if outerNALUType == NALUTypeAVCFUA {

		// rfc3984 5.8.  Fragmentation Units (FUs)
		//
		// 0                   1                   2                   3
		// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
		// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		// | FU indicator  |   FU header   |                               |
		// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+                               |
		// |                                                               |
		// |                         FU payload                            |
		// |                                                               |
		// |                               +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		// |                               :...OPTIONAL RTP padding        |
		// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		//
		// FU indicator:
		// +---------------+
		// |0|1|2|3|4|5|6|7|
		// +-+-+-+-+-+-+-+-+
		// |F|NRI|  Type   |
		// +---------------+
		//
		// Fu header:
		// +---------------+
		// |0|1|2|3|4|5|6|7|
		// +-+-+-+-+-+-+-+-+
		// |S|E|R|  Type   |
		// +---------------+

		fuIndicator := b[0]
		_ = fuIndicator
		fuHeader := b[1]

		startCode := (fuHeader & 0x80) != 0
		endCode := (fuHeader & 0x40) != 0

		if startCode {
			pkt.positionType = PositionTypeFUAStart
			return
		}

		if endCode {
			pkt.positionType = PositionTypeFUAEnd
			return
		}

		pkt.positionType = PositionTypeFUAMiddle
		return
	} else if outerNALUType == NALUTypeAVCSTAPA {
		pkt.positionType = PositionTypeSTAPA
	} else {
		nazalog.Errorf("unknown nalu type. outerNALUType=%d", outerNALUType)
	}

	return
}

func calcPositionIfNeededHEVC(pkt *RTPPacket) {
	b := pkt.Raw[pkt.Header.payloadOffset:]

	// +---------------+---------------+
	// |0|1|2|3|4|5|6|7|0|1|2|3|4|5|6|7|
	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	// |F|   Type    |  LayerId  | TID |
	// +-------------+-----------------+

	outerNALUType := hevc.ParseNALUType(b[0])

	switch outerNALUType {
	case hevc.NALUTypeVPS:
		fallthrough
	case hevc.NALUTypeSPS:
		fallthrough
	case hevc.NALUTypePPS:
		fallthrough
	case hevc.NALUTypeSEI:
		fallthrough
	case hevc.NALUTypeSliceTrailN:
		fallthrough
	case hevc.NALUTypeSliceTrailR:
		fallthrough
	case hevc.NALUTypeSliceIDRNLP:
		pkt.positionType = PositionTypeSingle
		return
	case NALUTypeHEVCFUA:
		// Figure 1: The Structure of the HEVC NAL Unit Header

		// 0                   1                   2                   3
		// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
		// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		// |    PayloadHdr (Type=49)       |   FU header   | DONL (cond)   |
		// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-|
		// | DONL (cond)   |                                               |
		// |-+-+-+-+-+-+-+-+                                               |
		// |                         FU payload                            |
		// |                                                               |
		// |                               +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		// |                               :...OPTIONAL RTP padding        |
		// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

		// Figure 9: The Structure of an FU

		// +---------------+
		// |0|1|2|3|4|5|6|7|
		// +-+-+-+-+-+-+-+-+
		// |S|E|  FuType   |
		// +---------------+

		// Figure 10: The Structure of FU Header

		startCode := (b[2] & 0x80) != 0
		endCode := (b[2] & 0x40) != 0

		if startCode {
			pkt.positionType = PositionTypeFUAStart
			return
		}

		if endCode {
			pkt.positionType = PositionTypeFUAEnd
			return
		}

		pkt.positionType = PositionTypeFUAMiddle
		return
	default:
		// TODO chef: 没有实现 AP 48
		nazalog.Errorf("unknown nalu type. outerNALUType=%d(%d), header=%+v, len=%d",
			b[0], outerNALUType, pkt.Header, len(pkt.Raw))
	}

}
