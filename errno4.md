## 400000

**类别：** [通用逻辑][common-logic]、[主机相关][hosts-logic]

**描述：** 未找到对应的主机

**如何解决：** 用户需确保在配置文件中或命令行中指定的主机都已被导入，若有主机未被导入，用户可在主机配置文件中添加相应的主机，并将其导入，详见[主机管理][hosts]


显示当前已导入主机列表

```shell
$ curveadm hosts ls
```

```shell
Host          Hostname  User   Port  Private Key File         Forward Agent  Become User  Labels
----          --------  ----   ----  ----------------         -------------  -----------  ------
server-host1  10.0.1.1  curve  22    /home/curve/.ssh/id_rsa  N              -            -
server-host2  10.0.1.2  curve  22    /home/curve/.ssh/id_rsa  N              -            -
server-host3  10.0.1.3  curve  22    /home/curve/.ssh/id_rsa  N              -            -
client-host   10.0.1.4  curve  22    /home/curve/.ssh/id_rsa  Y              admin        client
```

配置文件中使用主机示例：

```yaml
kind: curvefs
global:
  variable:
    machine1: server-host1  # 正确，server-host1 主机已被导入
    machine2: server-host2  # 正确，server-host2 主机已被导入
    machine3: server-host3  # 正确，server-host3 主机已被导入
    machine4: server-host4  # 错误，server-host4 主机未被导入
...
```

命令行中使用主机示例：

```shell
$ curveadm mount /fs1 /mnt/test1 --host client-host   # 正确，client-host 主机已被导入
$ curveadm mount /fs1 /mnt/test1 --host client-host2  # 错误，client-host2 主机未被导入
```

## 410001

**类别：** [通用逻辑][common-logic]、[服务相关][service-logic]

**描述：** 没有指定的集群

**如何解决：** 用户在进行集群相关的操作时，需切换到一个特定的集群，若当前没有集群，用户可新增一个集群后再切换过去，详见[添加集群并切换集群][add-and-checkout]

示例：

##### 1. 添加 'my-cluster' 集群，并指定集群拓扑文件
```shell
$ curveadm cluster add my-cluster -f topology.yaml
```

##### 2. 切换 'my-cluster' 集群为当前管理集群
```shell
$ curveadm cluster checkout my-cluster
```

## 410002

**类别：** [通用逻辑][common-logic]、[服务相关][service-logic]

**描述：** 没有主机或磁盘列表用于格式化

**如何解决：** CurveAdm 在解析完格式化配置文件后，并未发现任何用于格式化的主机或列表，出现此类情况是由于用户提供的格式化配置文件未按规定格式填写，用户可参见[磁盘格式化][format-disk]来准备磁盘格式化配置文件

示例：

```yaml
host:
  - server-host1
  - server-host2
  - server-host3
disk:
  - /dev/sda:/data/chunkserver0:90  # device:mount_path:format_percent
  - /dev/sdb:/data/chunkserver1:90
  - /dev/sdc:/data/chunkserver2:90
```

## 410003

**类别：** [通用逻辑][common-logic]、[服务相关][service-logic]

**描述：** 获取磁盘 UUID 失败

**如何解决：** 在进行磁盘格式化的时候，CurveAdm 会尝试去获取磁盘的 UUID，以用来将相应的磁盘挂载点记录添加到 `fstab` 文件中，用户可根据错误码提供的错误线索以及查看对应的日志文件来排除相应的问题

## 410004

**类别：** [通用逻辑][common-logic]、[服务相关][service-logic]

**描述：** 磁盘格式化配置文件中提供的设备并不是磁盘设备

**如何解决：** 磁盘列表中的每一项由设备、挂载路径、格式化百分比这 3 部分组成，并由 `:` 符号间隔开，其中的设备必须为磁盘，用户可根据错误线索来锁定哪一项不符合规范并进行相应修正

示例：

```yaml
host:
  - server-host1
  - server-host2
  - server-host3
disk:
  - /dev/sda:/data/chunkserver0:90   # 正确，/dev/sda 为磁盘设备
  - /dev/sdb1:/data/chunkserver0:90  # 正确，/dev/sdb1 为磁盘设备
  - /tmp:/data/chunkserver1:90       # 错误, /tmp 并不是磁盘设备
```

## 410005

**类别：** [通用逻辑][common-logic]、[服务相关][service-logic]

**描述：** 无效的磁盘 UUID

**如何解决：** *该错误码暂未被使用*

## 410006

**类别：** [通用逻辑][common-logic]、[服务相关][service-logic]

**描述：** 没有服务用于预检

**如何解决：** *该错误码暂未被使用*

## 410007

**类别：** [通用逻辑][common-logic]、[服务相关][service-logic]

**描述：** 没有服务用于部署

**如何解决：** *该错误码暂未被使用*

## 410008

**类别：** [通用逻辑][common-logic]、[服务相关][service-logic]

**描述：** 未找到服务对应的容器 ID

**如何解决：** 这个错误属于代码逻辑错误，正常情况下不应该报告给用户，若用户收到该错误码，代表该版本 CurveAdm 代码逻辑存在问题，用户可通过以下 2 种方法反馈该问题：

* [提交 issue][issue]

* 添加 `opencurve_bot` 微信号，进入微信群



## 410009

**类别：** [通用逻辑][common-logic]、[服务相关][service-logic]

**描述：** 集群已存在

**如何解决：** 用户在新增集群时，指定的集群名已经存在。用户可通过 `curveadm cluster ls` 查看已经存在的集群列表，并指定一个未存在的集群名进行添加

## 410010

**类别：** [通用逻辑][common-logic]、[服务相关][service-logic]

**描述：** 未找到指定集群

**如何解决：** 用户在切换集群、删除集群等操作时需确保指定的集群已存在，用户可通过 `curveadm cluster ls` 查看已存在的集群列表

## 410011

**类别：** [通用逻辑][common-logic]、[服务相关][service-logic]

**描述：** 删除集群前需清理该集群中所有的服务

**如何解决：** 为保证集群服务不残留在主机上，从而不影响其他集群的部署，用户在删除集群前需清理该集群下的所有服务。当然用户也可以添加 `-f` 选项强制删除集群，但这是我们不推介的

示例：

#### 1. 清理集群服务
```shell
$ curveadm stop
$ curveadm clean
```

#### 2. 删除集群

```shell
$ curveadm cluster rm my-cluster
```

## 410012

**类别：** [通用逻辑][common-logic]、[服务相关][service-logic]

**描述：** 不支持的配置类型

**如何解决：** 这个错误属于代码逻辑错误，正常情况下不应该报告给用户，若用户收到该错误码，代表该版本 CurveAdm 代码逻辑存在问题，用户可通过以下 2 种方法反馈该问题：

* [提交 issue][issue]

* 添加 `opencurve_bot` 微信号，进入微信群



## 410013

**类别：** [通用逻辑][common-logic]、[服务相关][service-logic]

**描述：** 位置的任务类型

**如何解决：** 这个错误属于代码逻辑错误，正常情况下不应该报告给用户，若用户收到该错误码，代表该版本 CurveAdm 代码逻辑存在问题，用户可通过以下 2 种方法反馈该问题：

* [提交 issue][issue]

* 添加 `opencurve_bot` 微信号，进入微信群




## 410014

**类别：** [通用逻辑][common-logic]、[服务相关][service-logic]

**描述：** 需停止的服务对应的容器已被删除

**如何解决：** CurveAdm 在停止服务时，发现其对应的容器已被删除，造成这一情况的原因一般是用户手动删除了指定的容器或其他原因导致容器被清理。如果接下来用户打算清理集群，那么可以忽略这个错误，否则用户可以需要重新使用部署命令来重新创建这个服务

示例：

```shell
$ curveadm deploy -k
```

部署命令为幂等操作，可重复使用

## 410015

**类别：** [通用逻辑][common-logic]、[服务相关][service-logic]

**描述：** 服务对应的容器处于不正常状态

**如何解决：** 用户可通过错误线索来锁定哪个服务容器处于不正常状态，并查看该服务对应的日志来排除相应的问题来使服务恢复正常。服务对应的日志目录可通过 `curveadm status -v` 命令来查看

示例：

```shell
$ curveadm status -v
```

```shell
Get Service Status: [OK]

cluster name      : my-cluster
cluster kind      : curvefs
cluster mds addr  : 10.0.1.1:6700,10.0.1.2:6700,10.0.1.3:6700
cluster mds leader: 10.0.1.1:6700 / 6d7a415caf90

Id            Role        Host          Replicas  Container Id  Status       Ports      Log Dir                Data Dir
--            ----        ----          --------  ------------  ------       -----      -------                --------
e5f42b2055c5  etcd        debian10-001  1/1       348c7725bb13  Up 29 hours  2379,2380  /tmp/logs/etcd0        /tmp/data/etcd0
a9600f9c65fd  etcd        debian10-002  1/1       dd2d26c5dd19  Up 29 hours  2379,2380  /tmp/logs/etcd0        /tmp/data/etcd0
811219cda513  etcd        debian10-003  1/1       8976d1a44c51  Up 29 hours  2379,2380  /tmp/logs/etcd0        /tmp/data/etcd0
6d7a415caf90  mds         debian10-001  1/1       68cdf36f0445  Up 29 hours  6700,7700  /tmp/logs/mds0         /tmp/data/mds0
f2977583365e  mds         debian10-002  1/1       30fa351b59c4  Up 29 hours  7700       /tmp/logs/mds0         /tmp/data/mds0
d229ce4fbe77  mds         debian10-003  1/1       f83363341980  Up 29 hours  7700       /tmp/logs/mds0         /tmp/data/mds0
c1ed2e71523d  metaserver  debian10-001  1/1       ba4149ab74d8  Up 29 hours  6800       /tmp/logs/metaserver0  /tmp/data/metaserver0
c81fa5a0c805  metaserver  debian10-002  1/1       9c6d16f787c0  Up 29 hours  6800       /tmp/logs/metaserver0  /tmp/data/metaserver0
95a6772ba8ad  metaserver  debian10-003  1/1       6a58f7d05457  Up 29 hours  6800       /tmp/logs/metaserver0  /tmp/data/metaserver0
```

## 410016

**类别：** [通用逻辑][common-logic]、[服务相关][service-logic]

**描述：** 编码集群逻辑池 JSON 字符串失败

**如何解决：** 这个错误属于代码逻辑错误，正常情况下不应该报告给用户，若用户收到该错误码，代表该版本 CurveAdm 代码逻辑存在问题，用户可通过以下 2 种方法反馈该问题：

* [提交 issue][issue]

* 添加 `opencurve_bot` 微信号，进入微信群



## 410017

**类别：** [通用逻辑][common-logic]、[服务相关][service-logic]

**描述：** 等待 mds 服务选主成功超时

**如何解决：** 在创建集群物理池、逻辑池时需要等待 mds 选主成功之后才能执行，目前我们最长等待 20 秒。用户可通过 `curveadm status` 确定成功选主后再重新执行一下部署操作即可。若 mds 长时间未选主，有可能是 etcd 或 mds 服务出现异常，用户需根据服务日志来进一步排查问题

示例：

```shell
$ curveadm status -v
```

```shell
Get Service Status: [OK]

cluster name      : my-cluster
cluster kind      : curvefs
cluster mds addr  : 10.0.1.1:6700,10.0.1.2:6700,10.0.1.3:6700
cluster mds leader: 10.0.1.1:6700 / 6d7a415caf90  # 此时代表 mds 服务已成功选主

Id            Role        Host          Replicas  Container Id  Status       Ports      Log Dir                Data Dir
--            ----        ----          --------  ------------  ------       -----      -------                --------
e5f42b2055c5  etcd        debian10-001  1/1       348c7725bb13  Up 29 hours  2379,2380  /tmp/logs/etcd0        /tmp/data/etcd0
a9600f9c65fd  etcd        debian10-002  1/1       dd2d26c5dd19  Up 29 hours  2379,2380  /tmp/logs/etcd0        /tmp/data/etcd0
811219cda513  etcd        debian10-003  1/1       8976d1a44c51  Up 29 hours  2379,2380  /tmp/logs/etcd0        /tmp/data/etcd0
6d7a415caf90  mds         debian10-001  1/1       68cdf36f0445  Up 29 hours  6700,7700  /tmp/logs/mds0         /tmp/data/mds0
f2977583365e  mds         debian10-002  1/1       30fa351b59c4  Up 29 hours  7700       /tmp/logs/mds0         /tmp/data/mds0
d229ce4fbe77  mds         debian10-003  1/1       f83363341980  Up 29 hours  7700       /tmp/logs/mds0         /tmp/data/mds0
c1ed2e71523d  metaserver  debian10-001  1/1       ba4149ab74d8  Up 29 hours  6800       /tmp/logs/metaserver0  /tmp/data/metaserver0
c81fa5a0c805  metaserver  debian10-002  1/1       9c6d16f787c0  Up 29 hours  6800       /tmp/logs/metaserver0  /tmp/data/metaserver0
95a6772ba8ad  metaserver  debian10-003  1/1       6a58f7d05457  Up 29 hours  6800       /tmp/logs/metaserver0  /tmp/data/metaserver0
```

## 410018

**类别：** [通用逻辑][common-logic]、[服务相关][service-logic]

**描述：** 等待全部 chunkserver 在线超时

**如何解决：** 在创建集群逻辑池时需要等待全部 chunkserver 在线后才能执行，目前我们最长等待 30 秒。用户在等待超时后，并确定 chunkserver 服务都正常的情况下，可尝试再次执行 `deploy` 部署命令，若多次尝试后仍无法解决，用户需查看 chunkserver 的服务日志来进一步排查 chunkserver 的问题，确保 chunkserver 服务正常后再进行部署

## 410019

**类别：** [通用逻辑][common-logic]、[服务相关][service-logic]

**描述：** 创建集群逻辑池失败

**如何解决：** 用户可通过错误线索和查看对应的错误日志来排查相应问题，不过此类问题一般难以排查，用户可通过以下方式提交 issue 或联系 Curve 团队小伙伴协助排查

* [提交 issue][issue]

* 添加 `opencurve_bot` 微信号，进入微信群



## 410020

**类别：** [通用逻辑][common-logic]、[服务相关][service-logic]

**描述：** 无效的磁盘使用量

**如何解决：** 在查看格式化进度时，CurveAdm 会通过 `df` 命令获取磁盘的使用量，而获取到的使用量是无效的，此类错误大部分是系统环境造成的，用户可根据错误线索和查看对应的日志来进行排查

## 410021

**类别：** [通用逻辑][common-logic]、[服务相关][service-logic]

**描述：** 加密文件失败

**如何解决：** *此错误码暂未使用*

## 410022

**类别：** [通用逻辑][common-logic]、[服务相关][service-logic]

**描述：** 未找到 ID 对应的客户端

**如何解决：** 用户在使用 `support` 命令上报服务、客户端相关日志时，通过 `-c` 命令指定客户端 ID，用户需确保该 ID 对应的客户端存在。用户可通过 `curveadm client status` 获取客户端 ID

```shell
$ curveadm client status
```

```shell
Get Client Status: [OK]

Id            Kind     Host          Container Id  Status       Aux Info
--            ----     ----          ------------  ------       --------
362d538778ad  curvebs  client-host1  cfa00fd01ae8  Up 36 hours  {"user":"curve","volume":"/test1"}
b0d56cfaad14  curvebs  client-host2  c0301eff2af0  Up 36 hours  {"user":"curve","volume":"/test2"}
c700e1f6acab  curvebs  client-host3  52554173a54f  Up 36 hours  {"user":"curve","volume":"/test3"}
```

## 420000

**类别：** [通用逻辑][common-logic]、[CurveBS 客户端相关][curvebs-client-logic]

**描述：** 卷已被映射

**如何解决：** 一个卷在同一台主机上只能被映射一次，用户需要取消映射后才可再次映射，详见[取消 CurveBS 卷映射][umap]

## 420001

**类别：** [通用逻辑][common-logic]、[CurveBS 客户端相关][curvebs-client-logic]

**描述：** 卷对应的容器丢失

**如何解决：** 出现这类情况是由于卷对应的容器被用户手动清除或因为其原因被清理，用户可[取消 CurveBS 卷映射][umap]后，再进行[映射 CurveBS 卷][map] 即可

## 420002

**类别：** [通用逻辑][common-logic]、[CurveBS 客户端相关][curvebs-client-logic]

**描述：** 卷对应的容器处于不正常状态

**如何解决：** 用户可通过查看该卷所对应的客户端日志文件来排查问题。客户端日志文件位于客户端配置文件 `client.yaml` 中 `log_dir` 指定的目录中，**特别需要注意的是**，若用户未配置该配置项，日志默认保存在 docker 容器内 `/curvebs/nebd/logs` 目录下，用户可通过 `docker cp` 命令将容器内的日志目录拷贝出来

## 420003

**类别：** [通用逻辑][common-logic]、[CurveBS 客户端相关][curvebs-client-logic]

**描述：** 创建卷失败

**如何解决：** 用户可通过错误线索和查看对应的日志文件来锁定具体的错误，并进行相应的排查

## 420004

**类别：** [通用逻辑][common-logic]、[CurveBS 客户端相关][curvebs-client-logic]

**描述：** 映射卷为 nbd 设备失败

**如何解决：** 用户可通过错误线索和查看对应的日志文件来锁定具体的错误，并进行相应的排查

## 420005

**类别：** [通用逻辑][common-logic]、[CurveBS 客户端相关][curvebs-client-logic]

**描述：** 解码卷的部署信息失败

**如何解决：** 这个错误属于代码逻辑错误，正常情况下不应该报告给用户，若用户收到该错误码，代表该版本 CurveAdm 代码逻辑存在问题，用户可通过以下 2 种方法反馈该问题：

* [提交 issue][issue]

* 添加 `opencurve_bot` 微信号，进入微信群



## 420006

**类别：** [通用逻辑][common-logic]、[CurveBS 客户端相关][curvebs-client-logic]

**描述：** 取消 CurveBS 卷映射失败

**如何解决：** 用户可通过错误线索和查看对应的日志文件来锁定具体的错误，并进行相应的排查

## 420007

**类别：** [通用逻辑][common-logic]、[CurveBS 客户端相关][curvebs-client-logic]

**描述：** 已存在的 target 守护进程异常

**如何解决：** 用户可将该 target 守护进程停止掉，再重新启动一个新的 target 守护进程

## 420008

**类别：** [通用逻辑][common-logic]、[CurveBS 客户端相关][curvebs-client-logic]

**描述：** target 守护进程异常

**如何解决：** 用户需要通过 `docker logs` 命令查看该主机上 target 守护进程容器的相关日志来排查问题，target 守护进程的容器名为 `curvebs-target-deamon`

## 430000

**类别：** [通用逻辑][common-logic]、[CurveFS 客户端相关][curvefs-client-logic]

**描述：** 主机上的路径已被挂载

**如何解决：** 用户需通过 `curveadm umount` 命令卸载该挂载点后，再进行重新挂载，详见[卸载文件系统][umount]

## 430001

**类别：** [通用逻辑][common-logic]、[CurveFS 客户端相关][curvefs-client-logic]

**描述：** 创建文件系统失败

**如何解决：** **该错误码暂未被使用**

## 430002

**类别：** [通用逻辑][common-logic]、[CurveFS 客户端相关][curvefs-client-logic]

**描述：** 挂载 CurveFS 文件系统失败

**如何解决：** 一般造成该错误的原因是用户填写的 S3 信息错误导致创建文件系统失败，从而导致挂载文件系统失败，用户可检查客户端配置文件中 S3 相关的配置项是否正确。关于这个错误码，我们后续会做进一步拆分，帮助用户更好定位问题

## 430003

**类别：** [通用逻辑][common-logic]、[CurveFS 客户端相关][curvefs-client-logic]

**描述：** 卸载 CurveFS 文件系统失败

**如何解决：** 用户需要根据错误线索和查看对应的错误日志来锁定具体的错误，并进行相应的排查

## 440000

**类别：** [通用逻辑][common-logic]、[PolarFS 相关][polarfs-logic]

**描述：** 获取系统发行版信息失败

**如何解决：** 目前我们 PolarFS 只支持安装在 `debian10` 系统上，CurveAdm 会提前去获取安装主机的系统发行版信息。用户可根据错误线索和查看相应的错误日志来锁定具体的错误，并进行排查

## 440001

**类别：** [通用逻辑][common-logic]、[PolarFS 相关][polarfs-logic]

**描述：** 当前系统不支持安装 PolarFS

**如何解决：** 目前我们 PolarFS 只支持安装在 `debian10` 系统上，请确保当前的操作系统满足这一要求

## 440002

**类别：** [通用逻辑][common-logic]、[PolarFS 相关][polarfs-logic]

**描述：** 安装 PolarFS 包失败

**如何解决：** 用户需要根据错误线索和查看对应的错误日志来锁定具体的错误，并进行相应的排查

## 450000

**类别：** [通用逻辑][common-logic]、[playground 相关][service-logic]

**描述：** 未找到指定的 playground

**如何解决：** 在删除 playground 时，需要指定 playground 的名称，用户可通过 `playground ls` 命令查看所有的 playground 列表，并指定相应的 playground 名称进行删除

```shell
$ curveadm playground ls
```

```shell
Get Playground Status: [OK]

Id  Name                           Create Time          Status
--  ----                           -----------          ------
1   playground-curvebs-1661257926  2022-08-23 20:32:09  Up 3 minutes
2   playground-curvebs-1661258009  2022-08-23 20:33:32  Up About a minute
```

示例：

```shell
$ curveadm playground rm playground-curvebs-1661257926
```

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
