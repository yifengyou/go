<!-- TOC -->

- [gorm](#gorm)
  - [db.Exec 执行原生SQL命令](#dbexec-执行原生sql命令)

<!-- /TOC -->

# gorm

## model定义

```
type Animal struct {
  ID   int64
  Name string `gorm:"default:'galeone'"`
  Age  int64
}
```

## 字段标签

```
`gorm:"<-:create;default:'galeone';column='test'"`
```

1. 多个字段分号隔开
2. 所有字段引号包括

声明 model 时，tag 是可选的，GORM 支持以下 tag： tag 名大小写不敏感，但建议使用 camelCase 风格

| 标签名                 | 说明                                                                                                                                                                                                                                                                                                                                   |
| ---------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| column                 | 指定 db 列名                                                                                                                                                                                                                                                                                                                           |
| type                   | 列数据类型，推荐使用兼容性好的通用类型，例如：所有数据库都支持 bool、int、uint、float、string、time、bytes 并且可以和其他标签一起使用，例如：not null、size, autoIncrement… 像 varbinary(8) 这样指定数据库数据类型也是支持的。在使用指定数据库数据类型时，它需要是完整的数据库数据类型，如：MEDIUMINT UNSIGNED not NULL AUTO_INCREMENT |
| size                   | 指定列大小，例如：size:256                                                                                                                                                                                                                                                                                                             |
| primaryKey             | 指定列为主键                                                                                                                                                                                                                                                                                                                           |
| unique                 | 指定列为唯一                                                                                                                                                                                                                                                                                                                           |
| default                | 指定列的默认值                                                                                                                                                                                                                                                                                                                         |
| precision              | 指定列的精度                                                                                                                                                                                                                                                                                                                           |
| scale                  | 指定列大小                                                                                                                                                                                                                                                                                                                             |
| not null               | 指定列为 NOT NULL                                                                                                                                                                                                                                                                                                                      |
| autoIncrement          | 指定列为自动增长                                                                                                                                                                                                                                                                                                                       |
| autoIncrementIncrement | 自动步长，控制连续记录之间的间隔                                                                                                                                                                                                                                                                                                       |
| embedded               | 嵌套字段                                                                                                                                                                                                                                                                                                                               |
| embeddedPrefix         | 嵌入字段的列名前缀                                                                                                                                                                                                                                                                                                                     |
| autoCreateTime         | 创建时追踪当前时间，对于 int 字段，它会追踪秒级时间戳，您可以使用 nano/milli 来追踪纳秒、毫秒时间戳，例如：autoCreateTime:nano                                                                                                                                                                                                         |
| autoUpdateTime         | 创建/更新时追踪当前时间，对于 int 字段，它会追踪秒级时间戳，您可以使用 nano/milli 来追踪纳秒、毫秒时间戳，例如：autoUpdateTime:milli                                                                                                                                                                                                   |
| index                  | 根据参数创建索引，多个字段使用相同的名称则创建复合索引，查看 索引 获取详情                                                                                                                                                                                                                                                             |
| uniqueIndex            | 与 index 相同，但创建的是唯一索引                                                                                                                                                                                                                                                                                                      |
| check                  | 创建检查约束，例如 check:age > 13，查看 约束 获取详情                                                                                                                                                                                                                                                                                  |
| <-                     | 设置字段写入的权限， <-:create 只创建、<-:update 只更新、<-:false 无写入权限、<- 创建和更新权限                                                                                                                                                                                                                                        |
| ->                     | 设置字段读的权限，->:false 无读权限                                                                                                                                                                                                                                                                                                    |
| -                      | 忽略该字段，- 无读写权限                                                                                                                                                                                                                                                                                                               |
| comment                | 迁移时为字段添加注释         

### 嵌套接口体

```
type User struct {
  gorm.Model
  Name string
}
// 等效于
type User struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
  Name string
}
```

对于正常的结构体字段，你也可以通过标签 embedded 将其嵌入，例如：


```
type Author struct {
    Name  string
    Email string
}

type Blog struct {
  ID      int
  Author  Author `gorm:"embedded"`
  Upvotes int32
}
// 等效于
type Blog struct {
  ID    int64
  Name  string
  Email string
  Upvotes  int32
}
```

### gorm.Model


GORM 定义一个 gorm.Model 结构体，其包括字段 ID、CreatedAt、UpdatedAt、DeletedAt

```
// gorm.Model 的定义
type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}

```


## 表创建/迁移


## 表删除




## db.Exec 执行原生SQL命令

* <https://v1.gorm.io/docs/sql_builder.html>

```
db.Exec("DROP TABLE users;")
db.Exec("UPDATE orders SET shipped_at=? WHERE id IN (?)", time.Now(), []int64{11,22,33})

// Scan
type Result struct {
  Name string
  Age  int
}

var result Result
db.Raw("SELECT name, age FROM users WHERE name = ?", 3).Scan(&result)
```






---
