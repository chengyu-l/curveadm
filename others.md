目录
===

* [部署 tgt](#部署-tgt)
* [升级 CurveAdm](#升级-curveadm)
* [错误码表](#错误码表)
* [问题与反馈](#问题与反馈)
* [参与 CurveAdm 的开发](#参与-curveadm-的开发)

部署 tgt
---

如果你要在除 Linux 外的其它操作系统中使用 CurveBS 卷，可以部署[网易高性能版本 tgt][curve-tgt]：

* 首先你需要部署一个 CurveBS 集群，详见[部署 CurveBS 集群][curvebs-cluster-deployment]
* 其次，你需要部署 iSCSI 服务端，详见[部署 tgt][curve-tgt-deployment]
* 最后，根据当前操作系统选择相应的方式连接 iSCSI target：
  * [Linux][linux-initiator]
  * [Windows][windows-initiator]
  * [MacOS][macos-initiator]

升级 CurveAdm
---

CurveAdm 支持一键升级到最新版本：

```shell
$ curveadm -u
```

同时，CurveAdm 也支持升级到指定版本：

```shell
$ CURVEADM_VERSION=v0.1.0 curveadm -u
```

| 稳定版本    | 发布时间           |
| :---        | :---               |
| **v0.1.0** | 2022 年 8 月 16 号 |
| **v0.0.13** | 2022 年 1 月 20 号 |

问题与反馈
---

如果您有任何建议或遇到问题需要帮助，可以通过以下 2 种途径反馈或找到我们：

* [提交 issue][issue]

* 通过扫描以下二维码，添加微信群

<img src="https://raw.githubusercontent.com/opencurve/curve/master/docs/images/curve-wechat.jpeg" width="300px" height="300px">

参与 CurveAdm 的开发
---

众人拾柴火焰高，我们希望更多的人能参与到 CurveAdm 的开发中，帮助我们改进 CurveAdm，使 CurveAdm 成为更好用的部署和运维工具。

CurveAdm 的整体实现逻辑较简单清晰，并具有较好的模块抽象，我们相信一般用户能在较短的时间内加入到 CurveAdm 的开发中。
为此，我们还特地写了以下文章来帮助大家快速上手开发，希望对你有帮助：

* [快速上手 CurveAdm 开发](develop)

[issue]: https://github.com/opencurve/curveadm/issues
[curve-tgt]: https://github.com/opencurve/curve-tgt
[curvebs-cluster-deployment]: https://github.com/opencurve/curveadm/wiki/curvebs-cluster-deployment
[curve-tgt-deployment]: https://github.com/opencurve/curveadm/wiki/curve-tgt-deployment
[linux-initiator]: https://www.unixmen.com/attach-iscsi-target-disks-linux-servers/
[windows-initiator]: https://jingyan.baidu.com/article/e4511cf37feade2b845eaff8.html
[macos-initiator]: https://apple.stackexchange.com/questions/324745/iscsi-mounts-in-macos
