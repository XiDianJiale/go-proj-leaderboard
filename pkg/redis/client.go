package redis

import (
	"github.com/redis/go-redis/v9"
)

func NewClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "123456",
		DB:       1,
	})
}

/** JAVA
public RedisClient newClient() {
	RedisOptions options = new RedisOptions();
	options.setAddr("localhost:6379");
	options.setDb(0);

	RedisClient client = new RedisClient(options);
	return client;
}
🎯 为什么 Go 函数常返回指针 *T，而 Java 不需要？
Java 的对象本质就是指针（引用类型）Java 的所有对象都是引用（指针），不是值。
Java 对象 = 永远是引用。
Go struct = 默认是值类型，需要指针才能像 Java。
所以go如果直接值传递而不是指针引用，这里每次调用都复制一整个 Redis 客户端
*/
