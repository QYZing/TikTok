package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DBSource       string
	Cache          cache.CacheConf
	Snowflake      SnowflakeConf
	KqConsumerConf kq.KqConf
	//测试
	KqPusherConf struct {
		Brokers []string
		Topic   string
	}
}
type SnowflakeConf struct {
	StartTime int64
	Node      int64
}
