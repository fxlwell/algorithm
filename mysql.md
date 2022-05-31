# mysql 
```
ACID：
原子性：
	undo log（记录修改之前的数据） + 锁
一致性
	
隔离性：
	锁机制
		表锁：
		行锁：
			行锁（Record Lock） 锁定单条数据
			间隙锁（Gap Lock）锁定部分数据
			临键锁（Next-Key Lock） （】防止幻读
	MVCC
持久性：一旦提交，数据库的改变就是持久的 
	redo log （更新之后的数据）
	redo log 是innodb引擎特有的，binlog是mysql底层的。

insert	（排它锁：自动加锁)
update	（排它锁：自动加锁)
delete	（排它锁：自动加锁)
select	（不加锁）
select ... for update （排它锁：手动加锁)
1. select ... in share mod （共享锁：)



```