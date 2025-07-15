# 短链接项目

## 搭建项目骨架

1. 建库建表

发号器表
```sql
// 对stub字段添加唯一索引，目的是为了保证插入的值不重复
CREATE TABLE `sequence`
(
    `id`        bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `stub`      varchar(1)          NOT NULL,
    `timestamp` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_uniq_stub` (`stub`)
) ENGINE = MyISAM DEFAULT CHARSET = utf8 COMMENT = '序号表';
```

长链接短链接映射表

```sql
// 由于lurl可能过长，所以引入md5，为其添加索引
CREATE TABLE `short_url_map`
(
    `id`        BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `create_at` DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_by` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '创建者',
    `is_del`    tinyint UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否删除：0正常1删除',
    `lurl`      varchar(2048)        DEFAULT NULL COMMENT '⻓链接',
    `md5`       char(32)             DEFAULT NULL COMMENT '⻓链接MD5',
    `surl`      varchar(11)          DEFAULT NULL COMMENT '短链接',
    PRIMARY KEY (`id`),
    INDEX(`is_del`),
    UNIQUE (`md5`),
    UNIQUE (`surl`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COMMENT = '⻓短链映射表';
```

2. 搭建go-zero框框的骨架

2.1 编写`api`文件，使用goctl命令生成代码
```api
syntax = "v1"
// 短链接项目

type ConvertRequest{
    LongUrl string `json:"longUrl" `
}

type ConvertResponse{
    ShortUrl string `json:"shortUrl" `
}

type ShowRequest{
    ShortUrl string `json:"shortUrl"`
}

type ShowResponse{
   LongUrl string `json:"longUrl"`
}

service shortener-api {
    @handler ConvertHandler
    post /convert (ConvertRequest) returns (ConvertResponse)

    @handler ShowHandler
        get /:showUrl  (ShowRequest) returns (ShowResponse)
}
```
2.2 根据api文件生成代码
```bash
goctl api go -api shortener.api -dir .
```

3. 根据数据表生成model层代码

```bash
goctl model mysql datasource -url="root:root@tcp(127.0.0.1:3306)/shortener" -table="short_url_map" -dir="./model"
goctl model mysql datasource -url="root:root@tcp(127.0.0.1:3306)/shortener" -table="sequence" -dir="./model" 
```

4. 下载项目依赖
```bash
go mod tidy
```

5. 运行项目
```bash
go run shortener.go
```

6. 修改配置结构和配置文件
注意：两边一定一定要对齐！！！

配置数据库`config.go`文件时，注意，方法一与方法二效果相同：
```
// 方式一：匿名结构体
type Config struct {
rest.RestConf

ShortUrlDB struct { // 匿名结构体
    DSN string
    }
}

// 方式二
type Config struct {
rest.RestConf

ShortUrlDB ShortUrlDB
}
type ShortUrlDB struct {
    DSN string
}
```

## 参数校验
使用validator库

https://pkg.go.dev/github.com/go-playground/validator/v10

下载依赖
```bash
go get -u github.com/go-playground/validator/v10
```
在`shortener.api`中为结构体添加校验规则tag

## 查看短链

### 缓存版
有两种方式：
1. 使用自己的实现的缓存， surl -> lurl ，缓存的数据量少，节省缓存空间
2. 使用go-zero自带的缓存， surl -> 数据行 ，不需要自己实现                 √(use)

这里使用第二种方案：
1. 添加缓存配置
 - 配置文件yaml
 - 配置config结构体
2. 删除旧的model层代码
 - 删除 shorturlmapmodel.go文件
3. 重新生成model代码

```bash
goctl model mysql datasource -url="root:root@tcp(127.0.0.1:3306)/shortener" -table="short_url_map" -dir="./model" -c
```
4.修改svccontext层代码

