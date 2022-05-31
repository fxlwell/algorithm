# go groutine 

```
// 重要的全局变量
var (
	allgs      []*g   // 存储所有g的数组
	allm       *m     // 存储所有m的链表，通过 m.alllink链向下一个m
	allp       []*p   // len(allp) == gomaxprocs
	gomaxprocs int32  // 最大可创建p的个数，默认为cpu个数，可通过环境变量设置
	ncpu       int32  // cpu个数
	m0   m           // 代表进程的主线程
    g0   g           // m0的g0，也就是m0.g0 = &g0
	sched      schedt
)
type schedt struct {
   // 空闲状态的m
   midle muintptr // idle m's waiting for work
   // 空闲状态的m个数
   nmidle int32 // number of idle m's waiting for work
   // m允许的最大个数
   maxmcount int32 // maximum number of m's allowed (or die)
   nmsys     int32 // number of system m's not counted for deadlock
   nmfreed   int64 // cumulative number of freed m's
   // 系统中goroutine的数目，会自动更新
   ngsys uint32 // number of system goroutines; updated atomically
   // 空闲的p列表
   pidle puintptr // idle p's
   // 有多少个状态为空闲的p
   npidle uint32
   // 有多少个m自旋
   nmspinning uint32 // m处于自旋状态时，是当前m没有g可运行，正在寻找可运行的g
   // Global runnable queue.
   // 全局的可运行的g队列
   runqhead guintptr
   runqtail guintptr
   // 全局队列的大小
   runqsize int32
   // freem is the list of m's waiting to be freed when their
   // m.exited is set. Linked through m.freelink.
   freem *m

}

//G结构体
type g struct {
   // 简单数据结构，lo 和 hi 成员描述了栈的下界和上界内存地址
   stack       stack   // offset known to runtime/cgo
   stackguard0 uintptr // offset known to liblink
   stackguard1 uintptr // offset known to liblink
   // 当前的m
   m *m // current m; offset known to arm liblink
   // goroutine切换时，用于保存g的上下文，包括栈顶、栈低、pc等寄存器
   sched     gobuf 
   // 唯一的goroutine的ID
   goid int64
   // 标记是否可抢占
   preempt        bool  
}

//M 结构体
type m struct {
   // 用来执行调度指令的 goroutine
   g0      *g     // goroutine with scheduling stack
   morebuf gobuf  // gobuf arg to morestack
   // thread-local storage
   tls      [6]uintptr // thread-local storage (for x86 extern register)
   mstartfn func()
   // 当前运行的goroutine
   curg      *g       // current running goroutine
   // 关联p和执行的go代码
   p     puintptr // attached p for executing go code (nil if not executing go code)
   nextp puintptr
   id    int64
   // 用于标识睡眠还是唤醒
   park        note
   // 是否自旋，自旋就表示M正在找G来运行
   spinning bool // m is out of work and is actively looking for work
   // m是否被阻塞
   blocked bool // m is blocked on a note
   // 用于链接allm
   alllink   *m // on allm
   schedlink muintptr 
   //...
}

// P 结构体
type p struct {
   // id也是allp的数组下标
   id     int32
   status uint32 // one of pidle/prunning/...
   // 单向链表，指向下一个P的地址
   link puintptr
   // 每调度一次加1
   schedtick uint32 // incremented on every scheduler call
   // 回链到关联的m
   m       muintptr // back-link to associated m (nil if idle)

   // Queue of runnable goroutines. Accessed without lock.
   // 可运行的goroutine的队列
   runqhead uint32
   runqtail uint32
   runq     [256]guintptr
   // 下一个运行的g，优先级最高
   runnext guintptr  
}


// 程序启动时的初始化代码
......
for i := 0; i < N; i++ { // 创建N个操作系统线程执行schedule函数
     create_os_thread(schedule) // 创建一个操作系统线程执行schedule函数
}


// 定义一个线程私有全局变量，注意它是一个指向m结构体对象的指针
// ThreadLocal用来定义线程私有全局变量
ThreadLocal self *m
//schedule函数实现调度逻辑
func schedule() {
    // 创建和初始化m结构体对象，并赋值给私有全局变量self
    self = initm()  
    for { //调度循环
          if (self.p.runqueue is empty) {
                 // 根据某种算法从全局运行队列中找出一个需要运行的goroutine
                 g := find_a_runnable_goroutine_from_global_runqueue()
           } else {
                 // 根据某种算法从私有的局部运行队列中找出一个需要运行的goroutine
                 g := find_a_runnable_goroutine_from_local_runqueue()
           }
          run_g(g) // CPU运行该goroutine，直到需要调度其它goroutine才返回
          save_status_of_g(g) // 保存goroutine的状态，主要是寄存器的值
     }
}



```

### 基本关系
	1. P 的数量 GOMAXPROCS（最大2	56），实际的并发度，可以理解为调度器，负责对G的各种调度逻辑。
	2. M 的数量 SetMaxThreads(默认值 10000)
	2. G 全局队列和本地队列获取关系 ： 先从全局队列获取，再从其他队列获取（work stealing）
### 调度过程
	1. 新建一个G，会调用newproc，检查是否有空闲的p，如果有空闲p，就为p找一个空闲M，如果没有空闲M，则会创建新的M
	2. 新的一个P，再去执行的时候，首先小概率1/61从全局队列获取G（避免饥饿）,没有的话，再从本地队列，如果本地队列没有，则从全局队列获取，如果还没有，则会检测io，最后才会work stealing，
	3. M0 : 是启动程序后的编号为0的主线程，这个M对应的实例会在全局变量runtime.m0中，不需要在heap上分配，M0负责执行初始化操作和启动第一个G， 在之后M0就和其他的M一样了。
	4. 是每次启动一个M都会第一个创建的gourtine，G0仅用于负责调度的G，G0不指向任何可执行的函数, 每个M都会有一个自己的G0。在调度或系统调用时会使用G0的栈空间, 全局变量的G0是M0的G0。
	5. sysmons是系统后台监控线程，而且这个函数不符合GPM模型，该函数直接占用一个M，且不需要P，没有任何上下文切换，用不着P

### 问题
	1. PM什么时候绑定，什么时候解绑
	