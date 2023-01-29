扩容集群
===

* [第 1 步：提交主机列表](#第-1-步提交主机列表)
* [第 2 步：格式化磁盘（CurveBS）](#第-2-步格式化磁盘)
* [第 3 步：修改集群拓扑](#第-3-步修改集群拓扑)
* [第 4 步：扩容集群](#第-4-步扩容集群)

第 1 步：提交主机列表
---

### 1. 添加新增机器至主机列表

```shell
$ vim hosts.yaml
```

```yaml
global:
  user: curve
  ssh_port: 22
  private_key_file: /home/curve/.ssh/id_rsa

hosts:
  - host: server-host1
    hostname: 10.0.1.1
  - host: server-host2
    hostname: 10.0.1.2
  - host: server-host3
    hostname: 10.0.1.3
  - host: server-host4  # 新增机器
    hostname: 10.0.1.4
  - host: server-host5  # 新增机器
    hostname: 10.0.1.5
  - host: server-host6  # 新增机器
    hostname: 10.0.1.6
```

### 2. 提交主机列表

```shell
$ curveadm hosts commit hosts.yaml
```

第 2 步：格式化磁盘
---

> 此步骤只针对扩容 CurveBS 集群，若扩容的是 CurveFS 集群，请直接跳过此步骤。

### 1. 准备磁盘列表

```shell
$ vim format.yaml
```

```yaml
host:
  - server-host4
  - server-host5
  - server-host6
disk:
  - /dev/sda:/data/chunkserver0:90  # device:mount_path:format_percent
  - /dev/sdb:/data/chunkserver1:90
  - /dev/sdc:/data/chunkserver2:90
```

### 2. 开始格式化

```shell
$ curveadm format -f format.yaml
```

> :warning: **警告：**
>
> format.yaml 中只需要填写新增机器上的磁盘即可，**切勿**将已在集群中服务的磁盘列表填入其中，避免造成无法挽回的事故。

第 3 步：修改集群拓扑
---

将扩容的服务列表添加到拓扑文件中：


```shell
$ vim topology.yaml
```

```yaml
kind: curvebs
global:
  container_image: opencurvedocker/curvebs:v1.2
  variable:
    home_dir: /tmp
    machine1: server-001
    machine2: server-002
    machine3: server-003
    machine4: server-004  # 新增机器
    machine5: server-005  # 新增机器
    machine6: server-006  # 新增机器

chunkserver_services:
  config:
    listen.ip: ${service_host}
    listen.port: 82${format_replicas_sequence}
    data_dir: /data/chunkserver${service_replicas_sequence}
    copysets: 100
  deploy:
    - host: ${machine1}
      replicas: 3
    - host: ${machine2}
      replicas: 3
    - host: ${machine3}
      replicas: 3
    - host: ${machine4}  # 新增服务
      replicas: 3
    - host: ${machine5}  # 新增服务
      replicas: 3
    - host: ${machine6}  # 新增服务
      replicas: 3
```

> :warning: **警告：**
>
> * 每一次只能扩容同一种角色的服务
> * 对于 chunkserver/metaserver 服务来说，每次扩容都会新增一个逻辑池，新增的服务都位于该逻辑池中，请确保每次扩容至少增加 3 台主机


第 4 步：扩容集群
---

```shell
$ curveadm scale-out topology.yaml
```

> :bulb: **提醒：**
>
> 扩容操作属于幂等操作，用户在执行失败后可重复执行，不用担心服务残留问题

