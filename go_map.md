# go map

```
type hmap struct {
	count     int // 当前保存的元素个数
	...
	B         uint8  // 指示bucket数组的大小
	...
	buckets    unsafe.Pointer // bucket数组指针，数组的大小为2^B
	...
}

type bmap struct {
	tophash [8]uint8 //存储哈希值的高8位
	data    byte[1]  //key value数据:key/key/key/.../value/value/value...
	overflow *bmap   //溢出bucket的地址
}

```

1. 赋值
	高八位用于定位 bucket， 低八位用于定位 key，快速试错后再进行完整对比
	 
2. 扩容
	1. 触发 load factor 的最大值，负载因子已达到当前界限
		a. 负载因子是评估空间复杂度和时间复杂度的一种计算方法,默认值是6.5
		b. 当所有的bucket都装满时，装载因子=8	
		c. 对于这种情况，元素太多而bucket太少，扩容时将B加1，直接将bucket扩容两倍，老的buckets挂在oldbucket上
		d. 哈希因子过小，说明空间利用率低
		e. 哈希因子过大，说明冲突严重，存取效率低
		f. 负载因子 > 6.5时，也即平均每个bucket存储的键值对达到6.5个。
	2. 溢出桶 overflow buckets 过多
		overflow数量 > 2^15时，也即overflow数量超过32768时。


3. 迁移
	1. oldbuckets设计，增量扩容，扩容完毕才清空。
	2. 它只是分配好了新的buckets并把老的buckets挂到了oldbuckets字段上。只有在插入，修改，删除key的时候才会进行真正的搬迁。rehash