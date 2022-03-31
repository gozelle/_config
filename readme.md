# _config 配置文件读取

使用示例：
```go
configer,err := _config.Load("/path/to/config.toml")
if err != nil {
	panic(err)
}

err = configer.Bind(&config)
if err != nil {
	panic(err)
}

//_config.PrintJson(config)
```