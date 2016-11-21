## gglog

基于 **[glog](https://github.com/golang/glog)** 进行改进的轻量级日志系统。

### 改进内容

* 添加 DEGUG 级别，五种日志等级 DEBUG < INFO < WARING < ERROR < FATAL
* 增加 SetOL方法，通过 SetOL 设置 stderrThreshold 值，级别大于等于 stderrThreshold 则会在控制台输出信息，默认级别 ERROR
* 优化日志输出名字，把“级别”标签放在靠前的位置
* **其他操作与 [glog](https://github.com/golang/glog) 完全一样**

### demo
```
package main

import (
    "github.com/haozibi/gglog"
    "flag"
)

func main() {
    flag.Parse()
    defer gglog.Flush()

    // 设置控制台输出级别，比此级别大的都会在控制台输出
    // DEBUG < INFO < WARING < ERROR < FATAL , 默认ERROR级别
    gglog.SetOL("DEBUG")
    // 设置log输出路径
    gglog.SetLogDir("log")

    gglog.Info("Hello gglog")
    gglog.Debug("This is a Debug log")
    gglog.Warning("This is a Warning log")
    gglog.Error("This is a Error log")

    gglog.Infof("info %d", 1)
    gglog.Warningf("warning %d", 2)
    gglog.Errorf("error %d", 3)

    gglog.V(1).Infoln("level 1")
    gglog.V(2).Infoln("level 2")
}
```

### 截图

![](http://ww2.sinaimg.cn/large/005LNEMyjw1f9yvu4hb30j30m30c8gsj.jpg)
