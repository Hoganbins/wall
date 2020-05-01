<p align="center">
<a title="logo" target="_blank" href="https://github.com/q191201771/lal">
<img alt="Wide" src="https://pengrl.com/images/other/lallogo.png">
</a>
<br>
Go语言编写的直播流媒体网络传输服务器
<br><br>
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

已支持RTMP，HTTP-FLV，H264/AVC，H265/HEVC，AAC，GOP缓存。



### README 目录

* 运行
* 配置文件说明
* 仓库目录框架
* Roadmap
* 文档
* 联系我
* 性能测试
* 测试过的第三方客户端

### 运行

#### 方式1，直接下载对应平台二进制可执行文件，体验功能

TODO

#### 方式2，自己编译

```shell
# 不使用 Go module
$go get -u github.com/q191201771/lal
$cd $GOPATH/src/github.com/q191201771/lal
$./build.sh

# 使用 Go module
$export GOPROXY=https://goproxy.cn
$git clone https://github.com/q191201771/lal.git
$cd lal
$./build.sh

# 运行
$./bin/lals -c conf/lals.conf.json
```

### 配置文件说明

```
{
  "rtmp": {
    "addr": ":19350", // RTMP服务监听的端口，客户端向lals推拉流都是这个地址
    "gop_num": 2      // rtmp拉流的GOP缓存数量，加速秒开
  },
  "httpflv": {
    "sub_listen_addr": ":8080", // HTTP-FLV拉流地址
    "gop_num": 2
  },
  "log": {
    "level": 1,                    // 日志级别，1 debug, 2 info, 3 warn, 4 error, 5 fatal
    "filename": "./logs/lals.log", // 日志输出文件
    "is_to_stdout": true,          // 是否打印至标志控制台输出
    "is_rotate_daily": true,       // 日志按天翻滚
    "short_file_flag": true,       // 日志末尾是否携带源码文件名以及行号的信息
    "assert_behavior": 1           // 日志断言的行为，1 只打印错误日志 2 打印并退出程序 3 打印并panic
  },
  "pprof": {
    "addr": ":10001" // Go pprof web地址
  }
}
```

其它放在代码中的配置：

- [rtmp/var.go](https://github.com/q191201771/lal/blob/master/pkg/rtmp/var.go)
- [httpflv/var.go](https://github.com/q191201771/lal/blob/master/pkg/httpflv/var.go)

### 仓库目录框架

简单来说，源码在`app/`和`pkg/`两个目录下，后续我再画些源码架构图。

```
pkg/                  ......源码包
|-- rtmp/             ......RTMP协议
|-- httpflv/          ......HTTP-FLV协议
|-- logic/            ......lals服务器的上层业务
|-- aac/              ......音频AAC编解码格式相关
|-- avc/              ......视频H264/AVC编解码格式相关

app/                  ......各种main包的源码文件，一个子目录对应一个main包，也即对应可生成一个可执行文件
|-- lals/             ......[最重要的]流媒体服务器
|-- flvfile2rtmppush  ......// RTMP推流客户端，读取本地FLV文件，使用RTMP协议推送出去
                            //
                            // 支持循环推送：文件推送完毕后，可循环推送（RTMP push流并不断开）
                            // 支持推送多路流：相当于一个RTMP推流压测工具
                            
|-- rtmppull          ......// RTMP拉流客户端，从远端服务器拉取RTMP流，存储为本地FLV文件
                            //
                            // 另外，作为一个RTMP拉流压测工具，已经支持：
                            // 1. 对一路流拉取n份
                            // 2. 拉取n路流
                            
|-- httpflvpull       ......HTTP-FLV拉流客户端
|-- modflvfile        ......修改本地FLV文件
|-- flvfile2es        ......将本地FLV文件分离成H264/AVC和AAC的ES流文件
bin/                  ......可执行文件编译输出目录
conf/                 ......配置文件目录
```

目前唯一的第三方依赖（我自己写的Go基础库）： [github.com/q191201771/naza](https://github.com/q191201771/naza)


### Roadmap

#### 项目原则：

* 代码可读可维护（是的，所有写这的，都是认真的）
* 框架清晰，模块化。业务与协议隔离。协议、网络传输等基础功能都是功能纯粹，可独立使用的库。
* 高性能
* 提供各种client代码，即使你使用其他流媒体服务器，这些client也是非常好用的
* 依托Go语言，提供所有平台下最简单的编译、调试、发布方式
* 不依赖第三方代码
* 后续可快速集成各种网络传输协议，流媒体封装协议

- 接收RTMP推流 [DONE]
- 转发给RTMP拉流 [DONE]
- 转发给HTTP-FLV拉流 [DONE]
- AAC [DONE]
- H264/AVC [DONE]
- H265/HEVC [DONE]
- GOP缓存 [DONE]
- RTMP转推
- RTMP回源
- HTTP-FLV回源
- 静态转推、回源
- 动态转推、回源
- hls
- rtsp
- rtp/rtcp
- webrtc
- udp quic
- udp srt
- udp kcp
- 分布式。提供与外部调度系统交互的接口。应对多级分发场景，或平级源站类型场景
- 调整框架代码
- 各种推流、拉流客户端兼容性测试
- 和其它主流服务器的性能对比测试
- 整理日志
- 稳定性测试
- mp4

### 文档

* [流媒体音视频相关的点我](https://pengrl.com/categories/%E6%B5%81%E5%AA%92%E4%BD%93%E9%9F%B3%E8%A7%86%E9%A2%91/)
* [Go语言相关的点我](https://pengrl.com/categories/Go/)
* [我写的其他文章](https://pengrl.com/all/)

### 联系我

扫码加我微信，进行技术交流或扯淡。微信群已开放。加我后拉你进群。

<img src="https://pengrl.com/images/yoko_vx.jpeg" width="180" height="180" />

### 性能测试

测试场景一：持续推送n路RTMP流至lals（没有拉流）

| 推流数量 | CPU占用 | 内存占用（RES） |
| - | - | - |
| 1000 | （占单个核的）16% | 104MB |

测试场景二：持续推送1路RTMP流至lals，使用RTMP协议从lals拉取n路流

| 拉流数量 | CPU占用 | 内存占用（RES） |
| - | - | - |
| 1000 | （占单个核的）30% | 120MB |

测试场景三： 持续推送n路RTMP流至lals，使用RTMP协议从lals拉取n路流（推拉流为1对1的关系）

| 推流数量 | 拉流数量 | CPU占用 | 内存占用（RES） |
| - | - | - | - |
| 1000 | 1000 | 125% | 464MB |

* 测试机：32核16G（lals服务器和压测工具同时跑在这一个机器上）
* 压测工具：lal中的 `/app/flvfile2rtmppush` 以及 `/app/rtmppull`
* 推流码率：使用`srs-bench`中的FLV文件，大概200kbps
* lals版本：基于 git commit: xxx

*由于测试机是台共用的机器，上面还跑了许多其他服务，这里列的只是个粗略的数据，还待做更多的性能分析以及优化。如果你对性能感兴趣，欢迎进行测试并将结果反馈给我。*

性能和可读，有时是矛盾的，存在取舍。我会保持思考，尽量平衡好两者。

### 测试过的第三方客户端

```
RTMP推流端：
- OBS 21.0.3(macos)
- OBS 24.0.3(win10 64 bit)
- ffmpeg 3.4.2(macos)
- srs-bench (macos srs项目配套的一个压测工具)
- flvfile2rtmppush (macos lal app中的RTMP推流客户端)

RTMP / HTTP-FLV 拉流端：
- VLC 3.0.8(macos 10.15.1)
- VLC 3.0.8(win10)
- MPV 0.29.1(macos)
- ffmpeg 3.4.2(macos)
- srs-bench (macos srs项目配套的一个压测工具)
```
