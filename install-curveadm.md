软硬件环境需求
===

Linux 操作系统版本需求
---

| 发行版  | 版本需求    |
| :---   | :---         |
| Debian | 9 及以上版本 |     
| CentOS | 8 及以上版本 |
| Ubuntu | **未测试**   |

> 📢 **注意：**
> 
> CurveBS 客户端依赖内核 nbd 模块，请确保内核版本大于或等于 `4.18`。

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