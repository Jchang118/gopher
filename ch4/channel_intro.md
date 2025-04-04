channel的创建
    msgChannel := make(cha int)
    msgChannel := make(chan int, 100) // 100代表缓冲区

channel的写入
    msgChannel <- i

channel的读取
    j := <- msgChannel

channel的数据结构--hchan
    type hchan struct {
    qcount          uint                // channel里的元素数量
    dataqsiz        uint                // 缓冲区大小
    buf             unsafe.Pointer      // 缓冲区指针
    elemsize        uint16              // 元素的大小
    closed          uint32              // 关闭状态
    elemtype        *_type              // 元素的数据类型
    sendx           uint                // 缓冲区中，已被读取的数据的位置
    recvx           uint                // 缓冲区中，当前已经写入数据的位置
    recvq           waitq               // 等待从channel中读取数据，正在被阻塞的协程队列
    sendq           waitq               // 等待从channel中写入数据，正在被阻塞的协程队列
    lock            mutex
    }
