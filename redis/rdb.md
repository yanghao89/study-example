## Redis

#### 持久化 RDB
```
RDB快照是一次全量备份,存储的是内存数据的二进制序列化,
存储上非常紧凑。当进行快照持久化时,会开启一个子进程专门负责快照
持久化,子进程会拥有父进程的内存数据,父进程修改内存子进程不会反应出来,
所以在快照持久化期间修改的数据不会被保存,可能丢失数据
```

#### 持久化AOF

```

```

#### Redis AFK

```
强一致性：就必须是同步

解决可用性

通过异步方式

```


#### Redis CAP 

```

```