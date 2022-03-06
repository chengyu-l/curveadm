使用 CurveAdm 部署 CurveBS 集群
===

* [第 1 步：环境准备](#第-1-步环境准备)
* [第 2 步：在中控机上安装 CurveAdm](#第-2-步在中控机上安装-curveadm)
* [第 3 步：格式化磁盘](#第-3-步格式化磁盘)
* [第 4 步：准备集群拓扑文件](#第-4-步准备集群拓扑文件)
* [第 5 步：添加集群并切换集群](#第-5-步添加集群并切换集群)
* [第 6 步：部署集群](#第-6-步部署集群)
* [第 7 步：查看集群运行情况](#第-7-步查看集群运行情况)
* [第 8 步：验证集群健康状态](#第-8-步验证集群健康状态)

第 1 步：环境准备
---

* [软硬件环境需求](install-curveadm#软硬件环境需求)
* [安装依赖](install-curveadm#安装依赖)

第 2 步：在中控机上安装 CurveAdm
---

* [安装 CurveAdm](install-curveadm#安装-curveadm)

第 3 步：格式化磁盘
---

为减少写 IO 放大，我们需要提前生成一批固定大小并预写过的 `chunk` 文件，详见 [chunkfile pool 设计][chunkfilepool]。

### 1. 准备磁盘列表

```shell
$ vim format.yaml
```

```yaml
user: curve
ssh_port: 22
private_key_file: /home/curve/.ssh/publish_rsa
host:
  - 10.0.1.1
  - 10.0.1.2
  - 10.0.1.3
disk:
  - /dev/sda:/data/chunkserver0:90  # device:mount_path:format_percent
  - /dev/sdb:/data/chunkserver1:90
  - /dev/sdc:/data/chunkserver2:90
```

> :bulb: **提醒：**
> 
> `disk` 数组中的每一项由 3 部分组成，分别为设备、挂载路径、格式化百分比。其中的挂载路径将作为 chunkserver 服务的数据目录使用，
> 所以我们建议挂载路径名的后缀为从 0 开始的连续递增数字，以精简 CurveBS 拓扑的配置。
> 
> 配置文件中的其余配置项可参考[重要配置项][important-config]。

> :warning: **警告：**
> 
> 请确保以上磁盘列表中的磁盘只用于 chunkserver 服务，针对每块磁盘，我们都将重新格式化成 `ext4` 文件系统，盘上的数据将全部被擦除。

### 2. 开始格式化

```shell
$ curveadm format -f format.yaml
```

> :bulb: **提醒：**
> 
> 考虑到格式化整个过程耗时较长，`curveadm format` 命令对每块磁盘成功启动一个格式化进程容器后即返回，
> 所以该命令成功返回并不意味着格式化已完成。
> 
> 用户可通过 `curveadm format --status` 命令查看格式化进度，当 `Status` 显示为 `Idea` 状态，
> 并且 `Formatted` 显示已格式化百分比大于设定百分比时，表示该磁盘格式化已完成。


第 4 步：准备集群拓扑文件
---

我们根据常见的场景，给用户准备了不同的拓扑文件模板，用户可根据需求自行选择，并进行编辑调整：

* [单机部署][curvebs-stand-alone-topology]

  所有服务都运行在一台主机上，一般用于体验或测试

* [多机部署][curvebs-cluster-topology]
 
  通用的多机部署模板，可用于生产环境或测试

关于拓扑文件中的各配置项，请参考 [CurveBS 集群拓扑][curvebs-topology]。

```shell
$ vim topology.yaml
```

```yaml
kind: curvebs
global:
  user: curve
  ssh_port: 22
  private_key_file: /home/curve/.ssh/publish_rsa
  container_image: opencurvedocker/curvebs:v1.2
  log_dir: /home/${user}/logs/${service_role}${service_replica_sequence}
  data_dir: /home/${user}/data/${service_role}${service_replica_sequence}
  s3.ak: <>
  s3.sk: <>
  s3.nos_address: <>
  s3.snapshot_bucket_name: <>
  variable:
    machine1: 10.0.1.1
    machine2: 10.0.1.2
    machine3: 10.0.1.3

etcd_services:
  config:
    listen.ip: ${service_host}
    listen.port: 2380
    listen.client_port: 2379
  deploy:
    - host: ${machine1}
    - host: ${machine2}
    - host: ${machine3}

mds_services:
  config:
    listen.ip: ${service_host}
    listen.port: 6700
    listen.dummy_port: 7700
  deploy:
    - host: ${machine1}
    - host: ${machine2}
    - host: ${machine3}

chunkserver_services:
  config:
    listen.ip: ${service_host}
    listen.port: 82${format_replica_sequence}  # 8200, 8201, 8202
    data_dir: /data/chunkserver${service_replica_sequence}  # /data/chunkserver0, /data/chunksever1, /data/chunkserver2
    copysets: 100
  deploy:
    - host: ${machine1}
      replica: 3
    - host: ${machine2}
      replica: 3
    - host: ${machine3}
      replica: 3

snapshotclone_services:
  config:
    listen.ip: ${service_host}
    listen.port: 5555
    listen.dummy_port: 8081
    listen.proxy_port: 8080
  deploy:
    - host: ${machine1}
    - host: ${machine2}
    - host: ${machine3}
```

第 5 步：添加集群并切换集群
---

#### 1. 添加 'my-cluster' 集群，并指定集群拓扑文件

```shell
$ curveadm cluster add my-cluster -f topology.yaml
```

#### 2. 切换 'my-cluster' 集群为当前管理集群

```shell
$ curveadm cluster checkout my-cluster 
```

第 6 步：部署集群
---

```shell
$ curveadm deploy
```

如果部署成功，将会输出类似 `Cluster 'my-cluster' successfully deployed ^_^.` 的字样。

第 7 步：查看集群运行情况
---

```shell
$ curveadm status
```

CurveAdm 默认会显示服务 ID、服务角色、主机地址、已部署的副本服务数量、容器 ID、运行状态：

```shell
Get Service Status: [OK]

cluster name    : my-cluster
cluster kind    : curvebs
cluster mds addr: 10.0.1.1:6700,10.0.1.2:6700,10.0.1.3:6700

Id            Role           Host      Replica  Container Id  Status
--            ----           ----      -------  ------------  ------
c9570c0d0252  etcd           10.0.1.1  1/1      ced84717bf4b  Up 45 hours
493b7831907c  etcd           10.0.1.2  1/1      907f8b84f527  Up 45 hours
8438cc5ecb52  etcd           10.0.1.3  1/1      44eca4798424  Up 45 hours
505da008b59c  mds            10.0.1.1  1/1      37c05bbb39af  Up 45 hours
e7bfb934182b  mds            10.0.1.2  1/1      044b56281928  Up 45 hours
1b322781339c  mds            10.0.1.3  1/1      b00481b9872d  Up 45 hours
<replica>     chunkserver    10.0.1.1  3/3      <replica>     RUNNING
<replica>     chunkserver    10.0.1.2  3/3      <replica>     RUNNING
<replica>     chunkserver    10.0.1.3  3/3      <replica>     RUNNING
2912bbdbcb48  snapshotclone  10.0.1.1  1/1      8b7a14b872ff  Up 45 hours
b862ef6720ed  snapshotclone  10.0.1.2  1/1      8e2a4b9e16b4  Up 45 hours
ed4533e903d9  snapshotclone  10.0.1.3  1/1      a35c30e3143d  Up 45 hours
```

* 若想查看其余信息，如日志目录、数据目录等，可添加 `-v` 参数
* 对于同一台主机上的[副本][replicas]服务来说，其状态默认是折叠的，可添加 `-s` 参数来显示每一个副本服务 

第 8 步：验证集群健康状态
---

集群服务正常运行，并不意味着集群的健康，所以我们在每一个容器内内置了 `curve_ops_tools` 工具。
该工具不仅可以查询集群的健康状态，还提供了许多其他特性，如显示各组件详细状态、集群容量、卷的管理、打快照等。

首先，我们需要进入任意一个服务容器内（服务 ID 可通过 `curveadm status` 查看）：

```shell
$ curveadm enter <Id>
```

在该容器内执行以下命令查看：
```shell
$ curve_ops_tool status
```

如果集群健康，在输出的开头会出现 `cluster is healthy` 的字样。

[important-config]: https://github.com/opencurve/curveadm/wiki/topology#curvebs-重要配置项
[chunkfilepool]: https://github.com/opencurve/curve/blob/master/docs/cn/chunkserver_design.md#224-chunkfilepool
[curvebs-stand-alone-topology]: https://github.com/opencurve/curveadm/blob/master/configs/bs/stand-alone/topology.yaml
[curvebs-cluster-topology]: https://github.com/opencurve/curveadm/blob/master/configs/bs/cluster/topology.yaml
[curvebs-topology]: https://github.com/opencurve/curveadm/wiki/topology#curvebs-集群拓扑
[replicas]: https://github.com/opencurve/curveadm/wiki/topology#replica 