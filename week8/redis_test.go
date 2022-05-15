package main

//写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估,
//结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，
//平均每个 key 的占用内存空间。

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/shirou/gopsutil/mem"
	"strconv"
)

// redis

// 定义一个全局变量
var redisdDb *redis.Client

func initRedis() (err error) {
	redisdDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // 指定
		Password: "",
		DB:       0, // redis一共16个库，指定其中一个库即可
	})
	_, err = redisdDb.Ping().Result()
	return
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("connect redis failed! err : %v\n", err)
		return
	}
	fmt.Println("redis连接成功！")

	//写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

	startMemory, _ := mem.VirtualMemory()

	str := "test"
	n := 1000
	for i := 1; i <= n; i++ {
		str += strconv.Itoa(i)
	}
	n = 10000
	for i := 0; i < n; i++ {
		redisdDb.Set("key_"+strconv.Itoa(i+1), str, -1)
	}
	endMemory, _ := mem.VirtualMemory()

	consume := int64(endMemory.Used) - int64(startMemory.Used)

	avgSize := consume / int64(n)
	fmt.Printf("开始内存：%d  结束内存： %d   内存消耗：%d  写入数量：%d  "+
		"value大小：%d 平均大小：%d \n\n", startMemory.Used, endMemory.Used, consume, n, len(str), avgSize)
}
