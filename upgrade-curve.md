升级服务
===

* [第 1 步：修改集群拓扑](#第-1-步修改集群拓扑)
* [第 2 步：提交集群拓扑](#第-2-步提交集群拓扑)
* [第 3 步：升级指定服务](#第-3-步升级指定服务)

第 1 步：修改集群拓扑
---

修改集群拓扑中的镜像名为需升级的镜像名：

```shell
$ vim topology.yaml
```

```yaml
kind: curvebs
global:
  container_image: opencurvedocker/curvebs:v1.2  # 修改镜像
  ...
```

第 2 步：提交集群拓扑
---

```shell
$ curveadm config commit topology.yaml
```

第 3 步：升级指定服务
---

```shell
$ curveadm upgrade
```

CurveAdm 默认升级集群中的所有服务，如需升级指定服务，可通过添加以下 3 个选项来实现：

* `--id`: 升级指定 `id` 的服务
* `--host`: 升级指定主机的所有服务
* `--role`: 升级指定角色的所有服务

以上 3 个选项可任意组合使用，服务对应的 `id`、`host`、`role` 可通过 [curveadm status](#查看集群状态) 来查看。

#### 示例 1：升级 id 为 `c9570c0d0252` 的服务

```shell
$ curveadm upgrade --id c9570c0d0252
```

#### 示例 2：升级 `10.0.1.1` 这台主机上的所有 `MDS` 服务
```shell
$ curveadm upgrade --host 10.0.1.1 --role mds
```

> :bulb: **提醒：**
>
> `upgrade` 默认会滚动升级指定的每一个服务，用户在升级每一个服务后
> 需进入容器内确定集群是否健康。若用户想一次性执行升级操作，可添加 `-f` 选项强制升级
