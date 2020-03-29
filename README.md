# redis_sort
使用Redis实现数据库里面的排序及分页

开发语言：golang

Redis: list, set/hset

举例条件：以大量视频为例，根据视频的点赞数(upNum)，评论数(cmtNum)，发布时间(addTime)做组合排序

实现思路：

存一个list, 里面存入的是 id列表

如：

127.0.0.1:6379> lpush vids 1 2 3 4

(integer) 4

另一个 set 存所需要的排序，

如：

127.0.0.1:6379> set vid-1 11

OK

127.0.0.1:6379> set vid-2 22

OK

127.0.0.1:6379> set vid-3 13

OK

127.0.0.1:6379> set vid-4 14

OK

注：vid-id（id与list中存入的id一致），set只能存单一的某个（或多个字段其中一个组合的）条件排序

hset





