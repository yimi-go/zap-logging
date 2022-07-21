zap-logging
===
Uber [zap](https://github.com/uber-go/zap) 的 [logging](https://github.com/yimi-go/logging) 适配
# 特性
* 可配置
  * 核心配置允许 json/yaml 格式文件 Unmarshal。
  * 支持 Option 模式编码配置。
* 允许动态重新加载配置。
  * 可动态配置 logger level。
  * 动态配置是否打印 caller、stacktrace。
  * ...
* logger 按 name 分别控制 minimum level. 相比全局、按 module、V 模式，配置更灵活，控制更精准。
