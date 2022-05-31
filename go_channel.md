# golang channel 原理
```
type hchan struct {
    qcount   uint           // total data in the queue
    dataqsiz uint           // size of the circular queue
    buf      unsafe.Pointer // points to an array of dataqsiz elements
    elemsize uint16
    closed   uint32
    elemtype *_type // element type
    sendx    uint   // send index
    recvx    uint   // receive index
    recvq    waitq  // list of recv waiters
    sendq    waitq  // list of send waiters

    // lock protects all fields in hchan, as well as several
    // fields in sudogs blocked on this channel.
    //
    // Do not change another G's status while holding this lock
    // (in particular, do not ready a G), as this can deadlock
    // with stack shrinking.
    lock mutex
}
```

#### 基本原理
	1. buf 存放数据，是一个循环链表
	2. sendx，recvx存放的是读写位置下表。
#### 其他
	1. go是用户态协程调度，使用go的scheduler完成调度,<- q 或者 q <- 会自动触发调度
	2. 如果G2先去recv数据，然后G1在写入数据，此时channle并不会被锁住。而是直接从G1的数据栈copy到G2的数据栈。
		a.G2在唤醒的时候，不需要再去获取等待锁。
		b.G2无需从环形队列里面拿数据，减少数据的copy次数。
	3. 唤醒顺序是FIFO原则，先阻塞先唤醒。chansend
		```
		if sg := c.recvq.dequeue(); sg != nil {
			// Found a waiting receiver. We pass the value we want to send
			// directly to the receiver, bypassing the channel buffer (if any).
			send(c, sg, ep, func() { unlock(&c.lock) }, 3)
			return true
		}
		```
	4. 当channel关闭的时候，会唤醒所有阻塞协程。