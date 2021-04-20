#mysql调优

## 性能监控

```
链接器 --->  控制用户的链接
分析器 ---> 词法分析
优化器 --->
执行器
存储引擎



性能监控
https://dev.mysql.com/doc/refman/5.7/en/show-profile.html
SELECT @@profiling;
SET profiling = 1;
https://dev.mysql.com/doc/refman/5.7/en/performance-schema.html
preformance shema 
https://dev.mysql.com/doc/refman/5.7/en/show-processlist.html
show processlist;



 
```


## schema与数据类型优化
```
1 数据类型的优化  
2 合理使用范式和反范式
3 主键的选择
4 字符集的选择
5 适当的数据
6 适当拆分

```