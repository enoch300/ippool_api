package peers

import (
	"context"
	"fmt"
	"ippool_api/db/redis"
	"strings"
)

var ctx = context.Background()

func getByAppidProvinceIsp(network, appid, province, isp string) (results []string, err error) {
	key := fmt.Sprintf("%s_%s_%s_%s", network, appid, province, isp)
	n, err := redis.RDB.Exists(ctx, key).Result()

	if err != nil {
		return results, err
	}

	if n == 0 {
		return results, fmt.Errorf("hkey %s not exist", key)
	}
	results, err = redis.RDB.HVals(ctx, key).Result()

	return
}

func Get(network, appid, machineid string) (res []string, err error) {
	isExist, err := redis.RDB.HExists(ctx, "machine_ipip", machineid).Result()
	if err != nil {
		return res, err
	}

	if !isExist {
		return res, fmt.Errorf("hash %s key %s not exist", "machine_ipip", machineid)
	}

	provinceIsp, _ := redis.RDB.HGet(ctx, "machine_ipip", machineid).Result()
	strArray := strings.Split(provinceIsp, "_")
	if len(strArray) != 2 {
		return res, fmt.Errorf("%s provinceIsp len != 2", machineid)
	}

	province := strings.TrimSpace(strArray[0])
	isp := strings.TrimSpace(strArray[1])

	results, err := getByAppidProvinceIsp(network, appid, province, isp)
	if err != nil {
		return res, err
	}
	return results, nil
}
