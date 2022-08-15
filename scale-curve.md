扩容集群
===

* [第 1 步：修改集群拓扑](#第-1-步修改集群拓扑)
* [第 2 步：扩容集群](#第-2-步扩容集群)

第 1 步：修改集群拓扑
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
    machine1: debian10-001
    machine2: debian10-002
    machine3: debian10-003
    machine4: debian10-004  # 新增机器
    machine5: debian10-005  # 新增机器
    machine6: debian10-006  # 新增机器

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


第 2 步：扩容集群
---

```shell
$ curveadm scale-out topology.yaml
```

> :bulb: **提醒：**
>
> 扩容操作属于幂等操作，用户在执行失败后可重复执行，不用担心服务残留问题

