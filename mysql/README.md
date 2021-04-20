# mysql 知识

#### MyISAM 和 InnoDB

```
锁引类型      MyISAM         InnoDB
支持事务        否              是
支持表锁        是              否
支持行锁        否              是
支持外键        否              是
支持全文锁引     是              是
```

```
适当拆分
1 水平拆分

2 垂直拆分
```


```
索引优化
存储引擎  memory  是用的hasindexces
 
右旋： 顺时针旋转，父节点被自己左孩子取代，而自己成为自己的右孩子
左旋： 逆时针旋转, 父节点被自己右孩子取代，而自己成为自己的左孩子


```


#### mysql 基础知识

```
1、大大减少了服务器需要扫描的数据量
2、帮助服务器避免排序和临时表
3、将随机IO变为顺序IO

B+树的 叶分裂和叶合并




覆盖索引：就是select id form table where name = '123', 找到id， 不需要回表查询数据，
就是覆盖索引
回表：就是 select * from table where key = '1234' 查询到id 后再找id主键b+树。 就是回表
最左匹配：

索引分类：
1、主键索引
2、唯一索引
3、普通索引
4、全文索引
5、组合索引
```


#### 覆盖索引

```
基本介绍:
1、如果一个索引包含所有需要查询字段的值,称之为覆盖索引
2、不是所有类型的索引都可以称为覆盖索引，覆盖索引必须要存在索引列的值
3、不同的存储引擎实现覆盖索引的方式不同,不是所有的引擎都支持覆盖索引,memory不支持覆盖索引


```


#### explain 中type 解释
```
system  const的特列,仅返回一条数据的时候。
const 常量连接
eq_ref 主键索引(primary key)或者非空唯一索引(unique not null)等值扫描
ref 非主键非唯一索引等值扫描
range 范围扫描
index 索引树扫描
all 全表扫描

system > const > eq_ref > ref > range > index > all

```

```
删除数据,需要记录binlog
redlog 
undolog
```

#### 强制用索引

```sql
CREATE TABLE `itdragon_order_list` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '主键id，默认自增长',
  `transaction_id` varchar(150) DEFAULT NULL COMMENT '交易号',
  `gross` double DEFAULT NULL COMMENT '毛收入(RMB)',
  `net` double DEFAULT NULL COMMENT '净收入(RMB)',
  `stock_id` int(11) DEFAULT NULL COMMENT '发货仓库',
  `order_status` int(11) DEFAULT NULL COMMENT '订单状态',
  `descript` varchar(255) DEFAULT NULL COMMENT '客服备注',
  `finance_descript` varchar(255) DEFAULT NULL COMMENT '财务备注',
  `create_type` varchar(100) DEFAULT NULL COMMENT '创建类型',
  `order_level` int(11) DEFAULT NULL COMMENT '订单级别',
  `input_user` varchar(20) DEFAULT NULL COMMENT '录入人',
  `input_date` varchar(20) DEFAULT NULL COMMENT '录入时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10003 DEFAULT CHARSET=utf8;

INSERT INTO itdragon_order_list VALUES ('10000', '81X97310V32236260E', '6.6', '6.13', '1', '10', 'ok', 'ok', 'auto', '1', 'itdragon', '2017-08-28 17:01:49');
INSERT INTO itdragon_order_list VALUES ('10001', '61525478BB371361Q', '18.88', '18.79', '1', '10', 'ok', 'ok', 'auto', '1', 'itdragon', '2017-08-18 17:01:50');
INSERT INTO itdragon_order_list VALUES ('10002', '5RT64180WE555861V', '20.18', '20.17', '1', '10', 'ok', 'ok', 'auto', '1', 'itdragon', '2017-09-08 17:01:49');

select 
/**
  全表扫描
 */
explain select * from itdragon_order_list order by order_level,input_date

/**
  创建索引
 */
create index idx_order_level on itdragon_order_list(order_level,input_date)
/*
优化过后  强制索引
*/
explain select * from  itdragon_order_list force index(idx_order_level) order by order_level,input_date;







```