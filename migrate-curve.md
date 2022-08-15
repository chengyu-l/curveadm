迁移服务
===

* [第 1 步：修改集群拓扑](#第-1-步修改集群拓扑)
* [第 2 步：迁移服务](#第-2-步迁移服务)

第 1 步：修改集群拓扑
---

修改拓扑文件中需要迁移的主机：

```shell
$ vim topology.yaml
```

```yaml
kind: curvebs
global:
  variable:
    machine1: debian10-001
    machine2: debian10-002
    machine3: debian10-003
    machine4: debian10-004  # 新增机器

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
    - host: ${machine4}  # 将 ${machine3} 修改为 ${machine4}
      replicas: 3
```

> :warning: **警告：**
>
> * 每一次只能迁移同一种角色的服务
> * 每一次只能迁移同一台主机的服务

第 2 步：迁移服务
---

```shell
$ curveadm migrate topology.yaml
```

> :bulb: **提醒：**
>
> 迁移操作属于幂等操作，用户在执行失败后可重复执行，不用担心服务残留问题

