<p align="center">
<img alt="Wide" src="https://pengrl.com/images/other/lallogo.png">
<br>
Go语言编写的流媒体 库 / 客户端 / 服务端
<br><br>
<a title="TravisCI" target="_blank" href="https://www.travis-ci.org/q191201771/lal"><img src="https://www.travis-ci.org/q191201771/lal.svg?branch=master"></a>
<a title="codecov" target="_blank" href="https://codecov.io/gh/q191201771/lal"><img src="https://codecov.io/gh/q191201771/lal/branch/master/graph/badge.svg?style=flat-square"></a>
<a title="goreportcard" target="_blank" href="https://goreportcard.com/report/github.com/q191201771/lal"><img src="https://goreportcard.com/badge/github.com/q191201771/lal?style=flat-square"></a>
<br>
<a title="codesize" target="_blank" href="https://github.com/q191201771/lal"><img src="https://img.shields.io/github/languages/code-size/q191201771/lal.svg?style=flat-square?style=flat-square"></a>
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

#### 工程目录说明

简单来说，主要源码在`app/`和`pkg/`两个目录下，后续我再画些源码架构图。

```
app/                  ......各种main包的源码文件，一个子目录对应一个main包，即对应可生成一个可执行文件
|-- lal/              ......[最重要的] 流媒体服务器
|-- flvfile2rtmppush  ......将本地flv文件使用rtmp推送出去
|-- rtmppull          ......rtmp拉流客户端
|-- httpflvpull       ......http-flv拉流客户端
|-- modflvfile        ......修改本地flv文件
|-- flvfile2es        ......将本地flv文件分离成h264/avc es流文件以及aac es流文件
pkg/                  ......源码包
|-- httpflv/          ......http-flv协议
|-- rtmp/             ......rtmp协议
|-- util/             ......帮助类包
    |-- bele/         ......大小端操作
    |-- bininfo/      ......可执行文件版本等信息
    |-- connstat/     ......连接超时信息
    |-- errors/       ......错误处理
    |-- log/          ......日志
    |-- unique/       ......对象唯一ID
bin/                  ......可执行文件输出目录
conf/                 ......配置文件目录
```

#### 编译和运行

```
$go get -u github.com/q191201771/lal
# cd into lal
$./build.sh

$./bin/lal -c conf/lal.conf.json
```

#### 配置文件说明

```
{
  "rtmp": {
    "addr": ":19350" // rtmp服务监听的端口
  }
}
```

#### roadmap

第一阶段：实现rtmp转发服务器

最终目标：

* 实现一个支持多种流媒体协议（比如rtmp, http-flv, hls, rtp/rtcp 等），多种底层传输协议（比如tcp, udp, srt, quic 等）的服务器
* 所有协议都以模块化的库形式提供给需要的用户使用
* 提供多种协议的推流客户端、拉流客户端，或者说演示demo

#### 依赖

目前不依赖任何第三方库

#### 文档

* [rtmp handshake | rtmp握手简单模式和复杂模式](https://pengrl.com/p/20027/)
* [rtmp协议中的chunk stream id, message stream id, transaction id, message type id](https://pengrl.com/p/25610/)
