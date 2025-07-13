# 短链接项目

## 搭建项目骨架

1. 建库建表

发号器表
```sql
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
1. 使用validator库

https://pkg.go.dev/github.com/go-playground/validator/v10

下载依赖
```bash
go get -u github.com/go-playground/validator/v10
```
在`shortener.api`中为结构体添加校验规则tag
