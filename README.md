# redis_sort
使用Redis实现数据库里面的排序及分页

开发语言：golang

Redis: list, set/hset

举例条件：以大量视频为例，根据视频的点赞数(upNum)，评论数(cmtNum)，发布时间(addTime)做组合排序

实现思路：

存一个list, 里面存入的是 id列表，本例在查询时只取ID，拿到ID集合后，根据ID去取视频详情

如：

127.0.0.1:6379> lpush vids 1 2 3 4

(integer) 4

(1) set 存所需要的排序

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

(2) hset 可以存多个字段（多种排序组合）

如：

 - 根据upNum 倒序，cmtNum倒序，addTime 正序

 - 根据upNum 倒序，cmtNum倒序，addTime 倒序


127.0.0.1:6379> hset vidsort-1 timedesc 12

(integer) 1

127.0.0.1:6379> hset vidsort-1 timeasc 12

(integer) 1

127.0.0.1:6379> hset vidsort-2 timedesc 22

(integer) 1

127.0.0.1:6379> hset vidsort-2 timeasc 32

(integer) 1
127.0.0.1:6379>

使用：

127.0.0.1:6379> sort vids by vidsort-*->timedesc

1) "3"
2) "4"
3) "1"
4) "2"

127.0.0.1:6379> sort vids by vidsort-*->timedasc

1) "1"
2) "2"
3) "3"
4) "4"

127.0.0.1:6379> sort vids by vidsort-*->timedasc limit 0 2

1) "1"
2) "2"

