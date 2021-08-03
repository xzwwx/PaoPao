package common

const (
	CmdHeaderSize = 2
)

// protobuf解码数据
func GetCmd(buf []byte) uint16 {
	if len(buf) < CmdHeaderSize {
		return 0
	}
	return uint16(buf[0]) | uint16(buf[1])<<8
}
