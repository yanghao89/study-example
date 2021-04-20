## KafKa 基础知识

```
1、kafka分组中consumer 大于分区数后, 多于的consumer是挂起的。
 
```


```
分区:桶
消息

如果没有顺序上的约束的话:水平扩展


拉取粒度   
按批次的拉取

kafka consumer 是以什么粒度更新&持久化 offset？


```



```shell
//创建 topic
kafka-topics.sh --create --bootstrap-server kafka1:9092,kafka2:9092,kafka3:9092 --topic mykafka3 --partitions 3 --replication-factor 3

//删除 topic
kafka-topics.sh --delete --bootstrap-server kafka1:9092,kafka2:9092,kafka3:9092 --topic mykafka

// 查看topic
kafka-topics.sh --zookeeper zoo1:2181 --topic mykafka3 --describe

```