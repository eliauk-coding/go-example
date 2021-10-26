package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.111.175:6379",
		Password: "",
		DB:       0,
	})
	val, err := rdb.Get(ctx, "test_yj").Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("key", val)

	val1 := make(map[string]interface{})
	val1["code_name"] = "qygsjnhhdqzq20200426"
	val1["sensor_type"] = "AnemometerContData"
	val1["sensor_ids"] = []string{"2", "3"}
	val1["start"] = 1629598292.835
	val1["end"] = 1629684692.835
	val1["zoom"] = "d"
	val1["taskid"] = 666
	val1["ts"] = 1629685717
	var column []interface{}
	column = append(column, map[string]string{
		"key":   "wind_speed",
		"value": "风速",
	})
	column = append(column, map[string]string{
		"key":   "wind_direction",
		"value": "风向",
	})
	val1["column_arr"] = column
	val1["filename"] = "AnemometerContData_['2', '3']"

	input, _ := json.Marshal(val1)
	cmd := rdb.RPush(context.Background(), "REDIS_SENSORCMD_TASKER_DOWNLOAD_TEST", input)
	fmt.Println(cmd)
}
