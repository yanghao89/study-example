# Elasticsearch 核心概念

#### 数据库的组成结构
```

链接器,分析器,优化器,执行器,cache 

存储引擎
```


#### es 文件


```
B-tree 多路平衡查找树

磁盘的存储都是16kb 


lucene


segment

posting list (词频）

term index
 
词频  

```


```
数字存储为
高16位  和 低16

array Container
65536 * 2bit / 1024b = 128kb

bitmap Container

65536 /8 bit /1024b = 8kb


4096


run contaier


FOR 算法

frame of reference

RBM 算法
Roaring Bitmap




```