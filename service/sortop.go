package service

import (
	"fmt"

	"github.com/go-redis/redis"
)

const (
	listKey = "vlist"
	hsetKey = "vsort-%d"
)

//HsetField .
type HsetField struct {
	TimeDesc string `json:"timeDesc"`
	TimeAsc  string `json:"timeAsc"`
}

//DataField .
type DataField struct {
	VID     int64 `json:"vid"`
	UpNum   int64 `json:"upNum"`
	CmtNum  int64 `json:"cmtNum"`
	AddTime int64 `json:"addTime"`
}

//ListReq .
type ListReq struct {
	SortAct string `json:"sortAct"`
	Page    int64  `json:"page"`
	Size    int64  `json:"size"`
}

//RedisDao .
type RedisDao struct {
	redis *redis.Client
}

func getRedis() *RedisDao {
	d := &RedisDao{
		redis: redis.NewClient(&redis.Options{
			Network:      "tcp",
			Addr:         "127.0.0.1:6379",
			Password:     "",
			PoolSize:     100,
			DB:           0,
			DialTimeout:  1000000000,
			ReadTimeout:  1000000000,
			WriteTimeout: 1000000000,
			IdleTimeout:  300000000000,
		}),
	}
	return d
}

func getDKey(vid int64) string {
	return fmt.Sprintf(hsetKey, vid)
}

func calSort(df *DataField, s string) int64 {
	var i int64
	i += df.UpNum * 7258089600 * 100000000
	i += df.CmtNum * 7258089600 //2020年1月1日
	if s == "desc" {
		i += df.AddTime
	} else {
		i += (7258089600 - df.AddTime)
	}
	return i
}

//AddData 添加/修改数据
func AddData(df *DataField) (int64, error) {
	//先判断是否在list中
	rd := getRedis()
	k := getDKey(df.VID)
	isHave, err := rd.redis.Exists(k).Result()
	if err != nil {
		return 0, err
	}
	if isHave == 0 { //不存在
		//添加
		_, err := rd.redis.LPush(listKey, df.VID).Result()
		if err != nil {
			return 0, err
		}
	}
	//更新
	ms := make(map[string]interface{}, 0)
	ms["timeDesc"] = calSort(df, "desc")
	ms["timeAsc"] = calSort(df, "asc")
	_, err = rd.redis.HMSet(k, ms).Result()
	return 0, err
}

//DelData 清除数据
func DelData(df *DataField) (int64, error) {
	rd := getRedis()
	k := getDKey(df.VID)
	//从list中删除
	res, err := rd.redis.LRem(listKey, 1, df.VID).Result()
	if err != nil {
		return res, err
	}
	//
	res, err = rd.redis.Del(k).Result()
	return res, err
}

//ListData 查询数据
func ListData(req *ListReq) ([]string, error) {
	rd := getRedis()
	sortBy := &redis.Sort{
		By:     "vsort-*->" + req.SortAct,
		Offset: (req.Page - 1) * req.Size,
		Count:  req.Size,
	}
	// By            string
	// Offset, Count int64
	// Get           []string
	// Order         string
	// Alpha         bool
	li, err := rd.redis.Sort(listKey, sortBy).Result()
	if err != nil {
		return nil, err
	}
	return li, nil
}
