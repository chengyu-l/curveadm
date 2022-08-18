软硬件环境需求
===

Linux 操作系统版本需求
---

| 发行版  | 版本需求      |
| :---   | :---        |
| Debian | 9 及以上版本  |
| CentOS | 7 及以上版本  |
| Ubuntu | 20 及以上版本 |

> 📢 **注意：**
>
> CurveBS 客户端在以下几种使用场景下依赖内核 nbd 模块，若当前内核不提供 nbd 模块，用户需自行编译并导入：
>   * 使用 [curveadm map][map] 命令映射 nbd 设备
>   * 使用 K8S CSI
>
> 而以下几种场景则可以在没有 nbd 模块的情况下正常工作：
>   * 使用 iSCSI 客户端
>   * 自行编译 QUME 对接 CurveBS

网络需求
---

目前通过 CurveAdm 部署集群有以下 2 方面网络需求：

* 安装 CurveAdm 的中控机需要通过 SSH 连接到部署服务的服务器
* 每台服务器需要从 Docker 仓库拉取镜像

安装依赖
===

CurveBS/CurveFS 所有服务都运行在 Docker 容器内，用户需要在每台服务器上[安装 Docker][install-docker]，并确保 Docker Daemon 已经运行。

你可以在服务器上运行以下命令来检测：

```shell
$ sudo docker run --rm hello-world
```

这个命令会下载一个测试映像，并在容器中运行它。当容器运行时，它打印一条消息并退出。


安装 CurveAdm
===

CurveAdm 支持一键安装：

```shell
$ bash -c "$(curl -fsSL https://curveadm.nos-eastchina1.126.net/script/install.sh)"
```

默认安装路径为当前用户主目录下的 `.curevadm` 目录，即 `~/.curveadm`。

> :bulb: **提醒：**
>
> CurveAdm 内置了命令补全命令，执行以下命令并根据相应提示即可加载补全：
> ```shell
> $ curveadm completion -h
> ```

[install-docker]: https://yeasy.gitbook.io/docker_practice/install

CurveAdm 配置文件
===

CurveAdm 配置文件为安装目录下的 `curveadm.cfg` 文件，即 `~/.curveadm/curveadm.cfg`。CurveAdm 配置文件中的配置项作用于 CurveAdm 所有命令的执行。

### 示例

```ini
[defaults]
log_level = info
sudo_alias = "sudo"
timeout = 300
auto_upgrade = true

[ssh_connections]
retries = 3
timeout = 10
```

### 配置项说明

| 区块            | 配置项       | 默认值 | 说明                                                                                                                                                                                   |
| :---            | :---         | :---   | :---                                                                                                                                                                                   |
| defaults        | log_level    | info   | 日志等级。支持 `debug`、`info`、`warn`、`error` 4 个等级                                                                                                                               |
| defaults        | sudo_alias   | sudo   | sudo 别名。CurveAdm 在执行某些命令时需要 root 权限，默认会以 sudo 方式去执行，如果用户想控制这个行为，可以修改该配置项，用户在需要 sudo 执行命令的场景下都会替换成 `sudo_alias` 去执行 |
| defaults        | timeout      | 300    | 执行命令超时时间（单位：秒）                                                                                                                                                           |
| defaults        | auto_upgrade | true   | 自动升级到 CurveAdm 最新版本                                                                                                                                                           |
| ssh_connections | retries      | 3      | SSH 连接失败重试次数                                                                                                                                                                   |
| ssh_connections | timeout      | 10     | SSH 连接超时时间（单位：秒）                                                                                                                                                           |

[map]: https://github.com/opencurve/curveadm/wiki/curvebs-client-deployment#%E7%AC%AC-4-%E6%AD%A5%E6%98%A0%E5%B0%84-curvebs-%E5%8D%B7
