package learn

type RedisClient struct {
	Name string
}

type DatabaseClient struct {
	Name string
}

type MyComponent struct {
	Db    *DatabaseClient `inject:""`
	Redis *RedisClient    `inject:""`
}

type MyApplication struct {
	Container        *MyComponent    `inject:"private"`
	Anon             *DatabaseClient `inject:"anon"`
	NonInjectedField string
}
