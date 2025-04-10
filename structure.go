package msg

const (
	START_CHAR byte = 0x68
	END_CHAR   byte = 0x16
)

var TITLE_CHARS = [4]byte{'B', 'Z', 'K', 'A'}

// BazookaStructure 通讯报文结构
type BazookaStructure struct {
	StartChar    byte   //开始字符
	TitleChars   []byte //头字符
	FrameIndex   byte   //报文id
	Control      byte   //控制域 bit0:0-服务端，1-客户端; bit6~bit1 发送方的逻辑地址长度；bit7为0
	SenderAddr   []byte //发送方逻辑地址
	FramingFlag  byte   //分帧标志，0-完整帧，1-分帧，2-结束帧
	FramingIndex uint64 //分帧序号，当为完整帧或结束帧时，就不存在该标志
	SegmentChar  byte   //分割字符
	Hcs          uint16 //头校验位
	FrameType    byte   //报文类型
	DataSize     uint16 //报文链路数据长度
	Data         []byte //报文链路数据
	Fcs          uint16 //帧校验位
	EndChar      byte   //结束字符
}
