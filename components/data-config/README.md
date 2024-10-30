# data-config组件
- 自定义数据源
- 读取数据
- 热更新数据

## Install

### Prerequisites
- GO >= 1.17

### Using go get
```
go get gameserver/cherry/components/data-config@latest
```


## Quick Start
```
import cherryDataConfig "gameserver/cherry/components/data-config"
```

```
package demo
import (
	"cherry"
	cherryDataConfig "gameserver/cherry/components/data-config"
)

// RegisterComponent 注册struct到data-config
func RegisterComponent() {
	dataConfig := cherryDataConfig.NewComponent()
	dataConfig.Register(
		&DropList,
		&DropOne,
	)

	//data-config组件注册到cherry引擎
	cherry.RegisterComponent(dataConfig)
}

```

## example
- [示例代码跳转](https://cherry-game/examples/tree/master/test_data_config)