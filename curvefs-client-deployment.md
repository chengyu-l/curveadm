部署 CurveFS 客户端
===

* [第 1 步：环境准备](#第-1-步环境准备)
* [第 2 步：导入主机](#第-2-步导入主机)
* [第 3 步：准备客户端配置文件](#第-3-步准备客户端配置文件)
* [第 4 步：挂载 CurveFS 文件系统](#第-4-步挂载-curvefs-文件系统)
* [其他：卸载文件系统](#其他卸载文件系统)

第 1 步：环境准备
---

* [软硬件环境需求](install-curveadm#软硬件环境需求)
* [安装依赖](install-curveadm#安装依赖)

第 2 步：导入主机
---

用户需导入客户端所需的主机列表，如果你在部署集群时已将客户端主机导入，可直接跳过此步骤。
请确保在之后挂载/卸载中指定的主机都已导入，详见[主机管理][hosts]。

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
  - host: server-host1
    hostname: 10.0.1.1
  - host: server-host2
    hostname: 10.0.1.2
  - host: server-host3
    hostname: 10.0.1.3
  - host: client-host
    hostname: 10.0.1.4
```

### 2. 导入主机列表
```shell
$ curveadm hosts commit hosts.yaml
```

第 3 步：准备客户端配置文件
---

```shell
$ vim client.yaml
```

```shell
kind: curvefs
s3.ak: <>
s3.sk: <>
s3.endpoint: <>
s3.bucket_name: <>
container_image: opencurvedocker/curvefs:latest
mdsOpt.rpcRetryOpt.addrs: 10.0.1.1:6700,10.0.1.2:6700,10.0.1.3:6700
log_dir: /home/curve/curvebs/logs/client
data_dir: /data/curvefs
```

客户端配置文件中的配置项含义等同于集群拓扑文件中的配置项，详见 [CurveFS 重要配置项][important-config]。

所有未在客户端配置文件上出现的配置项，我们都将使用默认配置值，
你可以通过点击 [client 配置文件][curvefs-client-conf]来查看各配置项及相关默认值。

> :bulb: 关于 `mdsOpt.rpcRetryOpt.addrs` 配置项
>
> 由于所有的路由信息都存在于 MDS 服务中，客户端只需知晓集群中 MDS 服务地址即可正常进行 IO 读写。
>
> 配置文件中的 `mdsOpt.rpcRetryOpt.addrs` 配置项需填写集群中 MDS 服务地址，用户在部署好 CurveFS 集群后，
> 可通过 `curveadm status` 查看集群 MDS 服务地址：
>
> ```shell
> $ curveadm status
> Get Service Status: [OK]
>
> cluster name    : my-cluster
> cluster kind    : curvefs
> cluster mds addr: 10.0.1.1:6700,10.0.1.2:6700,10.0.1.3:6700
> cluster mds leader: 10.0.1.1:6700 / 505da008b59c
> ...

> 📢 **注意：**
>
> 用户如需开启本地磁盘缓存，请务必配置 data_dir 配置项。

> 📢 **注意：**
>
> curve 自版本 release2.1.0 后支持多 s3，一个 fs 对应一个 s3 后端。
> 所有的 s3 信息（ak、sk、endpoint 和 bucket_name）保存在 mds 中，其他组件从 mds 中获取。
> 而 mds 中的 s3 信息在创建 fs 时指定。
> 因此 client.yaml 中的 s3 信息为必填项，相关信息会同步到创建 fs 工具的配置文件中。
> 若 fs 已经创建请保持一致，否则挂载失败；
> 若 fs 尚未创建则需保证 s3 信息可用，否则挂载失败。

第 4 步：挂载 CurveFS 文件系统
---

```shell
$ curveadm mount <curvefs-name> <mount-point> --host <host> -c client.yaml
```

* `<curvefs-name>`: 文件系统名，用户可自行定义
* `<mount-point>`: 挂载路径，用户可自行定义，但必须为**绝对路径**
* `--host`: 将卷挂载到指定主机，用户可自行选择，请确保该主机已被导入

如果文件系统挂载成功，在**相应的主机上**即能查询到 CurveFS 文件系统对应的挂载项：

```shell
$ mount | grep <mount-point>
```

用户也可以在中控机上查看所有客户端的状态：

```shell
$ curveadm client status
```

```
Get Client Status: [OK]

Id            Kind     Host          Container Id  Status       Aux Info
--            ----     ----          ------------  ------       --------
462d538778ad  curvefs  client-host1  dfa00fd01ae8  Up 36 hours  {"fsname":"/test1","mount_point":"/mnt/test1"}
c0d56cfaad14  curvefs  client-host2  c1301eff2af0  Up 36 hours  {"fsname":"/test2","mount_point":"/mnt/test2"}
d700e1f6acab  curvefs  client-host3  62554173a54f  Up 36 hours  {"fsname":"/test3","mount_point":"/mnt/test3"}
```

> 📢 **注意：**
>
> 若 curve release2.1 以上版本需要使用多 s3 功能，针对不同的 fs 修改 client.yaml 中的 s3 信息的相关配置即可。

### 示例
```shell
$ curveadm mount /test /mnt/test --host client-host -c client.yaml
```

其他：卸载文件系统
---

```shell
$ curveadm umount <mount-point> --host client-host
```

[hosts]: https://github.com/opencurve/curveadm/wiki/hosts
[important-config]: https://github.com/opencurve/curveadm/wiki/topology#curvefs-重要配置项
[curvefs-client-conf]: https://github.com/opencurve/curve/blob/master/curvefs/conf/client.conf
