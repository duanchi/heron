# Wand-Go

## 项目初始化

wand-go 目前支持本地`go mod`初始化

```
require (
	heurd.com/wand-go/wand v0.0.0
)

replace heurd.com/wand-go/wand => [path to gopath]/src/heurd.com/wand-go/wand
```



## 配置文件

### 自动初始化配置

项目中所有的配置文件可以通过定义一个`struct`类型的变量实现, 该变量可以引用`heurd.com/wand-go/wand/types`下的`Config`结构体进行组合。

```go
package config

import "heurd.com/wand-go/wand/types"

var Config = struct {
	types.Config
	Env string `env:"ENV,development"`
	ServerPort string `env:"SERVER_PORT,9080"`
	Db struct{
		Enabled bool `value:"true"`
		Dsn string `value:"host=172.31.16.1 port=3308 user=tb_cloud password=123456 dbname=thingsboard sslmode=disable"`
	}
	// 其他项目自定义配置
}{}
```

在配置变量初始化时, 可以通过常规变量初始化方式进行初始化, 也可以通过`wand-go`提供的简化方式(通过字段标签)通过初始化匿名结构体进行变量初始化,。

目前支持`env`和`value`两组标签

- `value` 设置静态配置项的属性值，标签的值即为属性值的描述值[*]

- `env` 将配置项的属性值设置为标签值所定义的环境变量, 配合docker可以方便的完成不同环境配置

> `value`和`env`实际赋值可根据字段基本类型自动进行类型转换, 支持的基本类型有`int/int64` `float64` `string` `bool`



### 获取配置值

在项目初始化后，可以直接引入该变量读取初始化后的配置。也可以通过配置获取方法`config.Get`获取配置

#### 直接获取配置

```go
import xxx/config
Dsn := Config.Db.Dsn
// host=172.31.16.1 port=3308 user=tb_cloud password=123456 dbname=thingsboard sslmode=disable
```

> 推荐使用配置变量直接获取配置, 既可以准确定位配置元素, 又可以省去类型强制转换。



#### 通过Get方法获取配置

```go
package config

func Get(key string) interface{} {}
```

其中`key`是以`.`分割的配置定义层级，获取配置后，需要进行强制类型转换。

```go
import "heurd.com/config"

func getConfig () {
    fmt.Println(config.Get("Db.Enabled").(bool))
    fmt.Println(config.Get("Db.Dsn").(string))
}

// true
// host=172.31.16.1 port=3308 user=tb_cloud password=123456 dbname=thingsboard sslmode=disable
```

> 读取配置后, 需手动进行类型转换

## 项目文件结构

建议采用三层分离的文件结构，即`控制器`、`业务逻辑服务`、`实体关系映射`三层，分别对应`controller`、`sevice`、`mapper`三个包。

root

- main.go
- config.go
- bean.go
- controller
  - xxx.go `控制器文件`
- service
  - xxx.go `业务逻辑文件`
- mapper
  - xxx.go `数据库mapper文件`

## Bean容器

借鉴于`Java Spring`的`Bean`概念，可以通过定义`Bean`进行项目中实体实例的管理，在初始化后可在任何位置调用。

### Bean定义

使用类似配置文件定义时使用的自定义结构体进行`Bean`的定义，字段名为`Bean`的`name`, 字段类型为`Bean`实例的类型， 可通过标签配置扩展`Bean`配置信息

```go
package config

var Beans = struct {
	DataDevices controller.DevicesController `route:"data/devices"`
}{}
```

项目初始化后可直接通过获取`Bean`变量字段调用`Bean`

```go
import config

func Test() {
    config.Beans.Fecth()
}

```

## Http服务

项目初始化后，可以使用`wand.HttpServer`进行Http服务器的相关操作

服务器端口配置可以使用配置`ServerPort` 指定

> Http服务器使用[Gin](https://github.com/gin-gonic/gin)

### 路由配置

Restful路由目前可根据资源进行路由配置，是通过扩展`Bean`定义实现的，因此只需在`Bean`变量中添加名为`route`的标签, 并设置其属性值为资源路径即可，`Bean`字段类型为相应的处理结构体。

配置完成后，可自动配置 `[resource]/:id`和`[resource]/`的路由解析

### Restful控制器

创建Restful控制器只需新建一个结构体，组合`wand.abstract.RestController`，并实现当前控制器关心的处理方法即可。

> 若不组合`wand.abstract.RestController`，则需实现所有`wand.interface.RestControllerInterface`方法

可实现的方法有

#### Fetch(获取) - GET

```go
func (controller RestController)
Fetch (
    id string,
    resource string,
    parameters *gin.Params,
    ctx *gin.Context
) (result interface{}, err types.Error) {}
```



#### Create(创建) - POST

```go
func (controller RestController)
Create (
    id string,
    resource string,
    parameters *gin.Params,
    ctx *gin.Context
) (result interface{}, err types.Error) {}
```



#### Update(更新) - PUT

```go
func (controller RestController)
Update (
    id string,
    resource string,
    parameters *gin.Params,
    ctx *gin.Context
) (result interface{}, err types.Error) {}
```



#### Remove(删除) - DELETE

```go
func (controller RestController)
Remove (
    id string,
    resource string,
    parameters *gin.Params,
    ctx *gin.Context
) (result interface{}, err types.Error) {}
```

> 方法返回值可以是任意可转换为`json`的数据格式

> 其余方法未完成

## 数据库

若设置配置 `Db.Enabled=true`，则可开始数据库支持。在项目中可以使用`wand.Db`进行和数据库有关的操作。

> 数据库使用[xorm](https://github.com/go-xorm/xorm)

## 异常

人为抛出异常时，建议使用`wand.types.Error`类型，并通过`Message`和`Code`字段定义异常

Code需使用HTTP协议支持的错误码。

## 运行

main.go 必须包含以下结构

```go
wand.Bootstrap(&Config, &Beans)
```

其中, `Config`为实际定义配置的变量，`Beans`为实际定义`Bean`的变量。项目初始化完成后，可通过引用调用配置和`Bean`的实际值。

项目开发环境可以使用 [fresh](https://github.com/gravityblast/fresh) 等hot-reload热加载方案。