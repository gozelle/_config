# _config 配置文件读取

使用示例：
```go
config := Config{}
err := _config.Load("/path/to/config.toml", &config)
if err != nil {
	panic(err)
}

_config.PrintJson(config)
```