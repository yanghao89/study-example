#### mysql 锁

#### 表锁
```
加锁快,不会出现死锁；锁定粒度大。并发度低
```
#### 行锁
```
开销大,加锁慢,会出现死锁,锁定粒最小。发生锁冲突的概率最低
并发度最高
```

#### innodb 锁
```
支持行级锁,也支持表级锁。但是在默认情况下是采用行级锁
表里有索引,会是行锁。 innodb 是针对索引加锁的
如果表里没有索引,innodb会退化为表级锁
```

#### myisam 表锁
```
支持表级锁
表共享读锁  表共享写锁

lock table mylock write  加表独占写锁

lock table mylock read 加表共享读锁
```

#### 排它锁

```sql
select * from tables lock in share mode
```

#### 共享锁

```sql
select * from tables for update
```


#### mysql ACID 特性

```sql
A:原子性
C:一致性
I:隔离性
D:持久性

mysql 如何保证原子性
undolog 为回滚日志 是实现原子性的关键,当事务回滚时能够撤销
所以已经成功执行的SQL语句,它需要记录你要回滚的相应日志信息
```



