部署 CurveBS 客户端
===

* [第 1 步：环境准备](#第-1-步环境准备)
* [第 2 步：导入主机](#第-2-步导入主机)
* [第 3 步：准备客户端配置文件](#第-3-步准备客户端配置文件)
* [第 4 步：映射 CurveBS 卷](#第-4-步映射-curvebs-卷)
* [其他：取消 CurveBS 卷映射](#其他取消-curvebs-卷映射)

第 1 步：环境准备
---

* [软硬件环境需求](install-curveadm#软硬件环境需求)
* [安装依赖](install-curveadm#安装依赖)


第 2 步：导入主机
---

用户需导入客户端所需的主机列表，如果你在部署集群时已将客户端主机导入，可直接跳过此步骤。
请确保在之后映射/取消映射中指定的主机都已导入，详见[主机管理][hosts]。

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

第 4 步：映射 CurveBS 卷
---

```shell
$ curveadm map <volume-user>:<volume-name> --host <host> -c client.yaml --create --size 10GB
```

* `<volume-user>`: 该卷所属用户名，用户可自行定义
* `<volume-name>`: 卷名，用户可自行定义
* `--host`: 将卷挂载到指定主机，用户可自行选择，请确保该主机已被导入
* `--create`：当卷不存在时，则自行创建
* `--size`: 指定创建卷的大小，默认为 `10GB`

当用户映射卷成功后，在**相应的主机上**即能看到 CurveBS 卷对应的 nbd 设备：

```shell
$ lsblk | grep nbd
```

用户也可以在中控机上查看所有客户端的状态：

```shell
$ curveadm client status
```

```shell
Get Client Status: [OK]

Id            Kind     Host          Container Id  Status       Aux Info
--            ----     ----          ------------  ------       --------
362d538778ad  curvebs  server-host1  cfa00fd01ae8  Up 36 hours  {"user":"curve","volume":"/test1"}
b0d56cfaad14  curvebs  server-host2  c0301eff2af0  Up 36 hours  {"user":"curve","volume":"/test2"}
c700e1f6acab  curvebs  server-host3  52554173a54f  Up 36 hours  {"user":"curve","volume":"/test3"}
```

> 📢 **注意：**
>
> * 卷名必须以 `/` 为起始，如 `/test`、`/curve/test`
> * 目前卷大小只支持以 `10GB` 为最小单位，即创建的卷大小只能是 10GB、20GB、30GB...，依此类推
> * 我们以 `<用户名>:<卷名>` 作为建来存储卷的相关信息，请勿挂载相同的卷

### 示例
```shell
$ curveadm map curve:/test --host server-host1 -c client.yaml --create
```

其他：取消 CurveBS 卷映射
---

```shell
$ curveadm unmap <user>:<volume-name> --host server-host1
```

[hosts]: https://github.com/opencurve/curveadm/wiki/hosts
[important-config]: https://github.com/opencurve/curveadm/wiki/topology#curvebs-重要配置项
[curvebs-client-conf]: https://github.com/opencurve/curve/blob/master/conf/client.conf
[nebd-design]: https://github.com/opencurve/curve/blob/master/docs/cn/nebd.md
