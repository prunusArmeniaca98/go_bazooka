package msg

// FixedSizePool 解析流程中的对象池
type FixedSizePool struct {
	maxSize       int                      //最大数量
	pool          chan *BazookaStructure   // 用 channel 限制大小
	createHandler func() *BazookaStructure // 创建对象的函数
}

// Get 从池中获取对象（如果池为空，则新建）
func (p *FixedSizePool) Get() *BazookaStructure {
	select {
	case obj := <-p.pool: // 从 channel 取
		return obj
	default: // 如果池为空，新建对象
		return p.createHandler()
	}
}

// Put 放回对象（如果池已满，则丢弃）
func (p *FixedSizePool) Put(obj *BazookaStructure) {
	select {
	case p.pool <- obj: // 放回 channel
	default: // 如果池已满，丢弃对象
	}
}

// NewFixedSizePool 创建解析流程对象池
func NewFixedSizePool(maxSize int) *FixedSizePool {
	return &FixedSizePool{
		maxSize: maxSize,
		pool:    make(chan *BazookaStructure, maxSize),
		createHandler: func() *BazookaStructure {
			return &BazookaStructure{}
		},
	}
}
