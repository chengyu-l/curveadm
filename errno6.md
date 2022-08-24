## 600000

**类别：** [命令执行][execute-command]、[命令通用][command-common]

**描述：** 命令执行超时

**如何解决：** 用户可修改 [CurveAdm 配置文件][curveadm-configure]中的 `timeout` 配置项来增大命令执行的超时时间

## 600001

**类别：** [命令执行][execute-command]、[命令通用][command-common]

**描述：** 读取文件失败

**如何解决：** 命令在执行的时候，需要一些读取文件的操作，用户可根据错误码报告的错误线索以及查看对应的日志文件来确定问题，并进行相关的修改。

一般这类错误不常出现，我们已知的报告该错误码的场景有以下几个：
* 磁盘空间已满

## 600002

**类别：** [命令执行][execute-command]、[命令通用][command-common]

**描述：** 写入文件失败

**如何解决：** 命令在执行的时候，需要一些写入文件的操作，用户可根据错误码报告的错误线索以及查看对应的日志文件来确定问题，并进行相关的修改。

一般这类错误不常出现，我们已知的报告该错误码的场景有以下几个：
* 磁盘空间已满

## 600003

**类别：** [命令执行][execute-command]、[命令通用][command-common]

**描述：** 编译正则表达式失败

**如何解决：** 这个错误属于代码逻辑错误，正常情况下不应该报告给用户，若用户收到该错误码，代表该版本 CurveAdm 代码逻辑存在问题，用户可通过以下 2 种方法反馈该问题：

* [提交 issue][issue]

* 添加 `opencurve_bot` 微信号，进入微信群



## 600004

**类别：** [命令执行][execute-command]、[命令通用][command-common]

**描述：** 构建命令模版失败

**如何解决：** 这个错误属于代码逻辑错误，正常情况下不应该报告给用户，若用户收到该错误码，代表该版本 CurveAdm 代码逻辑存在问题，用户可通过以下 2 种方法反馈该问题：

* [提交 issue][issue]

* 添加 `opencurve_bot` 微信号，进入微信群







## 610000

**类别：** [命令执行][execute-command]、[SSH 连接][ssh-connection]

**描述：** 利用 SSH 连接下载远端主机文件到本地失败

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 610001

**类别：** [命令执行][execute-command]、[SSH 连接][ssh-connection]

**描述：** 利用 SSH 连接上传文件到远端主机失败

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 610002

**类别：** [命令执行][execute-command]、[SSH 连接][ssh-connection]

**描述：** 通过 SSH 连接交互式地登录远端主机失败

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620000

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 利用 sed 编辑文件失败 (*sed*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620001

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 获取当前目录文件列表失败 (*ls*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620002

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 创建目录失败 (*mkdir*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620003

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 删除文件或目录失败 (*rm*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620004

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 重命名文件或目录失败 (*mv*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620005

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 复制文件和目录失败 (*cp*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620006

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 改变文件权限失败 (*chmod*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620007

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 获取文件状态失败 (*stat*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620008

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 打印文件内容失败 (*cat*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620009

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 创建 Linux 文件系统失败 (*mkfs*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620010

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 挂载一个 Linux 文件系统失败 (*mount*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620011

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 卸载一个 Linux 文件系统失败 (*umount*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620012

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 查询哪一个进程正在使用文件失败 (*fuser*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620013

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 获取磁盘空间使用量失败 (*df*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620014

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 获取所有块设备失败 (*lsblk*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620015

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 获取连接信息失败 (*ss*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620016

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 发送 ICMP ECHO 请求到目标主机失败 (*ping*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620017

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 传送数据到服务器或从服务器获取数据失败 (*curl*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620018

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 获取系统时间失败 (*date*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620019

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 获取系统信息失败 (*uname*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620020

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 获取内核模块信息失败 (*modinfo*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620021

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 从 Linux 内核加载模块失败 (*modinfo*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620022

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 获取主机名失败 (*hostname*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620023

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 压缩或解压文件失败 (*tar*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620024

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 安装或者移除 debian 包失败 (*dpkg*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620025

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 安装或者移除 rpm 包失败 (*rpm*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620026

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 拷贝文件到远端主机失败 (*rpm*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620998

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 执行一个 bash 脚本失败 (*bash script.sh*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 620999

**类别：** [命令执行][execute-command]、[shell 命令][shell-command]

**描述：** 运行一个 bash 命令失败 (*bash -c*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 630000

**类别：** [命令执行][execute-command]、[shell 命令][docker-command]

**描述：** 获取 docker 信息失败 (*docker info*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 630001

**类别：** [命令执行][execute-command]、[docker 命令][docker-command]

**描述：** 拉取镜像失败 (*docker pull IMAGE*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 630002

**类别：** [命令执行][execute-command]、[docker 命令][docker-command]

**描述：** 创建容器失败 (*docker create IMAGE*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 630003

**类别：** [命令执行][execute-command]、[docker 命令][docker-command]

**描述：** 启动容器失败 (*docker start CONTAINER*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 630004

**类别：** [命令执行][execute-command]、[docker 命令][docker-command]

**描述：** 停止容器失败 (*docker stop CONTAINER*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 630005

**类别：** [命令执行][execute-command]、[docker 命令][docker-command]

**描述：** 重启容器失败 (*docker restart CONTAINER*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 630006

**类别：** [命令执行][execute-command]、[docker 命令][docker-command]

**描述：** 等待容器停止失败 (*docker wait CONTAINER*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 630007

**类别：** [命令执行][execute-command]、[docker 命令][docker-command]

**描述：** 移除容器失败 (*docker rm CONTAINER*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 630008

**类别：** [命令执行][execute-command]、[docker 命令][docker-command]

**描述：** 获取所有容器失败 (*docker ps*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 630009

**类别：** [命令执行][execute-command]、[docker 命令][docker-command]

**描述：** 在容器内运行命令失败 (*docker exec CONTAINER COMMAND*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 630010

**类别：** [命令执行][execute-command]、[docker 命令][docker-command]

**描述：** 从容器内拷贝文件失败 (*docker cp CONTAINER:SRC_PATH DEST_PATH*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 630011

**类别：** [命令执行][execute-command]、[docker 命令][docker-command]

**描述：** 拷贝文件进容器失败 (*docker cp SRC_PATH CONTAINER:DEST_PATH*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 630012

**类别：** [命令执行][execute-command]、[docker 命令][docker-command]

**描述：** 获取容器底层信息失败 (*docker inspect ID*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 630013

**类别：** [命令执行][execute-command]、[docker 命令][docker-command]

**描述：** 获取容器日志失败 (*docker logs ID*)

**如何解决：** 详见命令执行类错误码[通用解决方法][execute-command-common-solution]

## 690000

**类别：** [命令执行][execute-command]、[其他命令][others-command]

**描述：** 在容器内启动 crontab 定时任务失败

**如何解决：** 我们会在每个服务容器内启动一个 crontab 定时任务来定时上报集群信息，包括集群 UUID 以及集群使用量，来帮助 curve 团队更好地了解用户及改进服务，具体参见 [CurveBS/CurveFS 重要配置项][important-configure] 中的 `report_usage` 配置项。

出现这类错误一般是对应的服务容器不正常或异常退出，导致在容器内执行命令失败导致的，用户可查看对应的服务容器的日志，排除服务异常的原因后，重新运行 `deploy、start、restart...` 等命令。

[init]: https://github.com/opencurve/curveadm/wiki/errno#主类
[database]: https://github.com/opencurve/curveadm/wiki/errno#主类
[command-options]: https://github.com/opencurve/curveadm/wiki/errno#主类
[configure]: https://github.com/opencurve/curveadm/wiki/errno#主类
[common-logic]: https://github.com/opencurve/curveadm/wiki/errno#主类
[execute-command]:  https://github.com/opencurve/curveadm/wiki/errno#主类
[cluster-options]: https://github.com/opencurve/curveadm/wiki/errno#子类
[client-common-options]:  https://github.com/opencurve/curveadm/wiki/errno#子类
[curvebs-client-options]:  https://github.com/opencurve/curveadm/wiki/errno#子类
[curvefs-client-options]:  https://github.com/opencurve/curveadm/wiki/errno#子类
[playground-options]: https://github.com/opencurve/curveadm/wiki/errno#子类
[configure-common]: https://github.com/opencurve/curveadm/wiki/errno#子类
[curveadm-configure]: https://github.com/opencurve/curveadm/wiki/errno#子类
[hosts-configure]: https://github.com/opencurve/curveadm/wiki/errno#子类
[topology-configure]: https://github.com/opencurve/curveadm/wiki/errno#子类
[format-configure]: https://github.com/opencurve/curveadm/wiki/errno#子类
[client-configure]: https://github.com/opencurve/curveadm/wiki/errno#子类
[hosts-logic]:  https://github.com/opencurve/curveadm/wiki/errno#子类
[service-logic]: https://github.com/opencurve/curveadm/wiki/errno#子类
[curvefs-client-logic]: https://github.com/opencurve/curveadm/wiki/errno#子类
[curvebs-client-logic]: https://github.com/opencurve/curveadm/wiki/errno#子类
[polarfs-logic]: https://github.com/opencurve/curveadm/wiki/errno#子类
[playground-logic]: https://github.com/opencurve/curveadm/wiki/errno#子类
[topology-precheck]: https://github.com/opencurve/curveadm/wiki/errno#子类
[ssh-precheck]: https://github.com/opencurve/curveadm/wiki/errno#子类
[permission-precheck]: https://github.com/opencurve/curveadm/wiki/errno#子类
[kernel-precheck]: https://github.com/opencurve/curveadm/wiki/errno#子类
[network-precheck]: https://github.com/opencurve/curveadm/wiki/errno#子类
[date-precheck]: https://github.com/opencurve/curveadm/wiki/errno#子类
[service-precheck]: https://github.com/opencurve/curveadm/wiki/errno#子类
[command-common]: https://github.com/opencurve/curveadm/wiki/errno#子类
[others-command]: https://github.com/opencurve/curveadm/wiki/errno#子类
[docker-command]:  https://github.com/opencurve/curveadm/wiki/errno#子类
[ssh-connection]:  https://github.com/opencurve/curveadm/wiki/errno#子类
[shell-command]:  https://github.com/opencurve/curveadm/wiki/errno#子类
[database-common-solution]: https://github.com/opencurve/curveadm/wiki/errno#数据库类错误码
[insatall-directory]: https://github.com/opencurve/curveadm/wiki/overview#安装目录
[precheck]: https://github.com/opencurve/curveadm/wiki/precheck
[clean-cluster]: https://github.com/opencurve/curveadm/wiki/maintain-curve#清理集群
[map]: https://github.com/opencurve/curveadm/wiki/curvebs-client-deployment#第-4-步映射-curvebs-卷
[mount]: https://github.com/opencurve/curveadm/wiki/curvefs-client-deployment#第-4-步挂载-curvefs-文件系统
[variable]: https://github.com/opencurve/curveadm/wiki/topology#变量
[curveadm-configure-file]: https://github.com/opencurve/curveadm/wiki/install-curveadm#curveadm-配置文件
[hosts]: https://github.com/opencurve/curveadm/wiki/hosts
[topology]: https://github.com/opencurve/curveadm/wiki/topology
[replicas]: https://github.com/opencurve/curveadm/wiki/topology#replicas
[scale-curve]: https://github.com/opencurve/curveadm/wiki/scale-curve
[migrate-curve]: https://github.com/opencurve/curveadm/wiki/migrate-curve
[format-disk]: https://github.com/opencurve/curveadm/wiki/curvebs-cluster-deployment#第-4-步格式化磁盘
[curvebs-client-configure]: https://github.com/opencurve/curveadm/wiki/curvebs-client-deployment#第-3-步准备客户端配置文件
[curvefs-client-configure]: https://github.com/opencurve/curveadm/wiki/curvefs-client-deployment#第-3-步准备客户端配置文件
[add-and-checkout]: https://github.com/opencurve/curveadm/wiki/curvebs-cluster-deployment#第-6-步添加集群并切换集群
[umap]: https://github.com/opencurve/curveadm/wiki/curvebs-client-deployment#其他取消-curvebs-卷映射
[umount]: https://github.com/opencurve/curveadm/wiki/curvefs-client-deployment#其他卸载文件系统
[chunkfilepool]: https://github.com/opencurve/curve/blob/master/docs/cn/chunkserver_design.md#224-chunkfilepool
[config-template]: https://github.com/opencurve/curveadm/tree/master/configs
[execute-command-common-solution]: https://github.com/opencurve/curveadm/wiki/errno#命令执行类错误码
[start-docker-daemon]: https://yeasy.gitbook.io/docker_practice/install/ubuntu#qi-dong-docker
[issue]: https://github.com/opencurve/curveadm/issues
[topology]: https://github.com/opencurve/curveadm/wiki/topology
[install-requirements]: https://github.com/opencurve/curveadm/wiki/install-curveadm#安装依赖
[important-configure]: https://github.com/opencurve/curveadm/wiki/topology#curvebs-重要配置项
[curvebs-important-configure]: https://github.com/opencurve/curveadm/wiki/topology#curvebs-重要配置项
[curvefs-important-configure]: https://github.com/opencurve/curveadm/wiki/topology#curvefs-重要配置项
