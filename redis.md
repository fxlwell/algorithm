# redis

```
单线程是针对内存的读写，采用单线程
网络io不是单线程


瓶颈在于内存和IO，而不是CPU
没有锁的开销
没有上下文切换的开销
没有竞争
io多路复用

```