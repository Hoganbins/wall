<p align="center">
<a title="logo" target="_blank" href="https://github.com/q191201771/lal">
<img alt="Wide" src="https://pengrl.com/images/other/lallogo.png">
</a>
<br>
<a title="TravisCI" target="_blank" href="https://www.travis-ci.org/q191201771/lal"><img src="https://www.travis-ci.org/q191201771/lal.svg?branch=master"></a>
<a title="codecov" target="_blank" href="https://codecov.io/gh/q191201771/lal"><img src="https://codecov.io/gh/q191201771/lal/branch/master/graph/badge.svg?style=flat-square"></a>
<a title="goreportcard" target="_blank" href="https://goreportcard.com/report/github.com/q191201771/lal"><img src="https://goreportcard.com/badge/github.com/q191201771/lal?style=flat-square"></a>
<br>
<a title="codeline" target="_blank" href="https://github.com/q191201771/lal"><img src="https://sloc.xyz/github/q191201771/lal/?category=code"></a>
<a title="license" target="_blank" href="https://github.com/q191201771/lal/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square"></a>
<a title="lastcommit" target="_blank" href="https://github.com/q191201771/lal/commits/master"><img src="https://img.shields.io/github/commit-activity/m/q191201771/lal.svg?style=flat-square"></a>
<a title="commitactivity" target="_blank" href="https://github.com/q191201771/lal/graphs/commit-activity"><img src="https://img.shields.io/github/last-commit/q191201771/lal.svg?style=flat-square"></a>
<br>
<a title="pr" target="_blank" href="https://github.com/q191201771/lal/pulls"><img src="https://img.shields.io/github/issues-pr-closed/q191201771/lal.svg?style=flat-square&color=FF9966"></a>
<a title="hits" target="_blank" href="https://github.com/q191201771/lal"><img src="https://hits.b3log.org/q191201771/lal.svg?style=flat-square"></a>
<a title="language" target="_blank" href="https://github.com/q191201771/lal"><img src="https://img.shields.io/github/languages/count/q191201771/lal.svg?style=flat-square"></a>
<a title="toplanguage" target="_blank" href="https://github.com/q191201771/lal"><img src="https://img.shields.io/github/languages/top/q191201771/lal.svg?style=flat-square"></a>
<a title="godoc" target="_blank" href="https://godoc.org/github.com/q191201771/lal"><img src="http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square"></a>
<br><br>
<a title="watcher" target="_blank" href="https://github.com/q191201771/lal/watchers"><img src="https://img.shields.io/github/watchers/q191201771/lal.svg?label=Watchers&style=social"></a>&nbsp;&nbsp;
<a title="star" target="_blank" href="https://github.com/q191201771/lal/stargazers"><img src="https://img.shields.io/github/stars/q191201771/lal.svg?label=Stars&style=social"></a>&nbsp;&nbsp;
<a title="fork" target="_blank" href="https://github.com/q191201771/lal/network/members"><img src="https://img.shields.io/github/forks/q191201771/lal.svg?label=Forks&style=social"></a>&nbsp;&nbsp;
</p>

---

app/lalserver服务器支持的协议：

| - | sub rtmp | sub http-flv | sub http-ts | sub hls | sub rtsp | relay push rtmp |
| - | - | - | - | - | - | - |
| pub rtmp        | ✔ | ✔ | ✔ | ✔ | - | ✔ |
| pub rtsp        | ✔ | ✔ | ✔ | ✔ | - | ✔ |
| relay pull rtmp | ✔ | ✔ | ✔ | ✔ | - | . |

| 编码类型 | rtmp | rtsp | hls | http-flv | http-ts |
| - | - | - | - | - | - |
| aac       | ✔ | ✔ | ✔ | ✔ | ✔ |
| avc/h264  | ✔ | ✔ | ✔ | ✔ | ✔ |
| hevc/h265 | ✔ | - | - | ✔ | - |

表格含义见： [《流媒体传输连接类型之session client, server, pub, sub, push, pull》](https://pengrl.com/p/20080)

基于lal开发的客户端示例程序见： [《lal/app/demo/README.md》](https://github.com/q191201771/lal/blob/master/app/demo/README.md)

<img alt="Wide" src="https://pengrl.com/images/other/lalmodule.jpg?date=0829">

### 编译，运行，体验功能

#### 编译

方式1，从源码自行编译

```shell
$export GO111MODULE=on && export GOPROXY=https://goproxy.cn,https://goproxy.io,direct
$make
```

方式2，直接下载编译好的二进制可执行文件

[点我打开《github lal最新release版本页面》](https://github.com/q191201771/lal/releases/latest)，下载对应平台编译好的二进制可执行文件的zip压缩包。

#### 运行

```shell
$./bin/lalserver -c conf/lalserver.conf.json
```

#### 体验功能

快速体验lalserver服务器见： [《常见推拉流客户端软件的使用方式》](https://pengrl.com/p/20051/)

lalserver详细配置见： [《配置注释文档》](https://github.com/q191201771/lal/blob/master/conf/lalserver.conf.json.brief)

### 源码框架

<br>

简单来说，源码在`pkg/`，`app/lalserver/`，`app/demo/`三个目录下。

- `pkg/`：存放各package包，供本repo的程序以及其他业务方使用
- `app/lalserver`：基于lal编写的一个通用流媒体服务器程序入口
- `app/demo/`：存放各种基于`lal/pkg`开发的小程序（小工具），一个子目录是一个程序，详情见各源码文件中头部的说明

目前唯一的第三方依赖（我自己写的Go基础库）： [github.com/q191201771/naza](https://github.com/q191201771/naza)

### 联系我

扫码加我微信（微信号： q191201771），进行技术交流或扯淡。微信群已开放，加我好友后可拉进群。

也欢迎大家通过github issue交流，提PR贡献代码。提PR前请先阅读：[《yoko版本PR规范》](https://pengrl.com/p/20070/)

<img src="https://pengrl.com/images/yoko_vx.jpeg" width="180" height="180" />

### 性能测试，测试过的第三方客户端

见[TEST.md](https://github.com/q191201771/lal/blob/master/TEST.md)

### 项目star趋势图

觉得这个repo还不错，就点个star支持一下吧 :)

[![Stargazers over time](https://starchart.cc/q191201771/lal.svg)](https://starchart.cc/q191201771/lal)

