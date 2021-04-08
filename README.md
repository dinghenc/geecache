# GeeCache

> 参考于 [geektutu](https://geektutu.com/post/geecache.html) 系列文章

基于Go实现的分布式缓存  

支持的特性有:

* LRU(Least Recently Used) 淘汰机制
* 支持并发缓存
* 基于 `Http` 访问
* 一致性哈希负载均衡
* 分布式节点
* 防止缓存击穿
* 基于 `Protobuf` 通信数据
