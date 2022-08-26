部署网络高性能版本 tgt
===

* [第 1 步：环境准备](#第-1-步环境准备)
* [第 2 步：导入主机](#第-2-步导入主机)
* [第 3 步：准备客户端配置文件](#第-3-步准备客户端配置文件)
* [第 4 步：启动 tgtd 守护进程](#第-4-步启动-tgtd-守护进程)
* [第 5 步：添加 target](#第-5-步添加-target)
* [第 6 步：显示所有 target](#第-6-步显示所有-target)
* [其他：删除 target ](#其他删除-target)
* [其他：停止 tgtd 守护进程](#其他停止-tgtd-守护进程)

第 1 步：环境准备
---

* [软硬件环境需求](install-curveadm#软硬件环境需求)
* [安装依赖](install-curveadm#安装依赖)

第 2 步：导入主机
---

用户需导入部署 target 的主机列表，详见[主机管理][hosts]。

### 1. 准备主机列表

```shell
$ vim hosts.yaml
```

```yaml
global:
  user: curve
  ssh_port: 22
  private_key_file: /home/curve/.ssh/id_rsa

hosts:
  - host: target-host
    hostname: 10.0.1.1
```

### 2. 导入主机列表
```shell
$ curveadm hosts commit hosts.yaml
```

第 3 步：准备客户端配置文件
---

由于 [curve-tgt][curve-tgt] 对接的是 CurveBS 客户端中的 [NEBD-Part2][nebd-design] 服务，
所以我们需要一份客户端配置文件来启动 NEBD-Part2 服务：

```shell
$ vim client.yaml
```

```shell
kind: curvebs
container_image: opencurvedocker/curvebs:v1.2
mds.listen.addr: 10.0.1.1:6700,10.0.1.2:6700,10.0.1.3:6700
log_dir: /home/curve/curvebs/logs/client
```

客户端配置文件中的配置项含义等同于集群拓扑文件中的配置项，详见 [CurveBS 重要配置项][important-config]。

所有未在客户端配置文件上出现的配置项，我们都将使用默认配置值，
你可以通过点击 [client 配置文件][curvebs-client-conf]来查看各配置项及相关默认值。

> :bulb: 关于 `mds.listen.addr` 配置项
>
> 由于所有的路由信息都存在于 MDS 服务中，客户端只需知晓集群中 MDS 服务地址即可正常进行 IO 正常。
>
> 配置文件中的 `mds.listen.addr` 配置项需填写集群中 MDS 服务地址，用户在部署好 CurveBS 集群后，
> 可通过 `curveadm status` 查看集群 MDS 服务地址：
>
> ```shell
> $ curveadm status
> Get Service Status: [OK]
>
> cluster name      : my-cluster
> cluster kind      : curvebs
> cluster mds addr  : 10.0.1.1:6700,10.0.1.2:6700,10.0.1.3:6700
> cluster mds leader: 10.0.1.1:6700 / 505da008b59c
> ...
> ```

第 4 步：启动 tgtd 守护进程
---

```shell
$ curveadm target start --host target-host -c client.yaml
```

* `--host`: 在指定主机上启动 `tgtd` 守护进程
* `-c`: 指定 CurveBS 客户端配置文件

第 5 步：添加 target
---

```shell
$ curveadm target add <user>:<volume-name> --host target-host --create --size 10GB
```

* `<user>`: 该卷所属用户名，用户可自行定义
* `<volume-name>`: 卷名，用户可自行定义
* `--host`: 将 target 添加到对应主机上的 tgtd 中
* `--create`：当卷不存在时，则自行创建
* `--size`: 指定创建卷的大小，默认为 `10GB`

> 📢 **注意：**
>
> * 卷名必须以 `/` 为起始，如 `/test`、`/curve/test`
> * 目前卷大小只支持以 `10GB` 为最小单位，即创建的卷大小只能是 10GB、20GB、30GB...，依此类推
> * 我们以 `<用户名>:<卷名>` 作为建来存储卷的相关信息，请勿挂载相同的卷
> * 目前我们一个 target 实例管理一个 CurveBS 卷

> :bulb: **提醒：**
>
> 目前 target 是通过 SSH 远程添加的。在相应的机器上，

### 示例
```shell
$ curveadm target add curve:/test -c client.yaml --create
```

第 6 步：显示所有 target
---

```shell
$ curveadm target list --host target-host
```

CurveAdm 会显示 target id、target name、对应的 CurveBS 卷、portal 等信息：

```shell
List Targets: [OK]

Tid  Target Name                                                       Store                   Portal
---  -----------                                                       -----                   ------
1    iqn.2022-02.com.opencurve:curve.d97372e236558863e4ffc63dcfa7eb27  cbd:pool//test1_curve_  10.182.2.43:3260
2    iqn.2022-02.com.opencurve:curve.41f7dfa48ebe871dd74b2db0d88f1243  cbd:pool//test2_curve_  10.182.2.43:3260
```

> :bulb: **提醒：**
>
> CurveAdm 会显示当前主机上全部已成功添加的 target 的相关信息，
> 这些将是之后客户端连接 target 的必要信息。

其他：删除 target
---

```shell
$ curveadm target rm TID [OPTION] --host target-host
```

> :bulb: **提醒：**
>
> TID 可以通过 `curveadm target list` 命令查看。

其他：停止 tgtd 守护进程
---

```shell
$ curveadm target stop --host target-host
```

[hosts]: https://github.com/opencurve/curveadm/wiki/hosts
[curve-tgt]: https://github.com/opencurve/curve-tgt
[nebd-design]: https://github.com/opencurve/curve/blob/master/docs/cn/nebd.md
[important-config]: https://github.com/opencurve/curveadm/wiki/topology#curvebs-重要配置项
[curvebs-client-conf]: https://github.com/opencurve/curve/blob/master/conf/client.conf
