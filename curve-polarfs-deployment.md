部署 Curve 版 PolarFS
===

* [第 1 步：环境准备](#第-1-步环境准备)
* [第 2 步：导入主机](#第-2-步导入主机)
* [第 3 步：准备客户端配置文件](#第-3-步准备客户端配置文件)
* [第 4 步：安装 PolarFS](#第-4-步安装-polarfs)
* [第 5 步：创建 curve 卷](#第-5-步创建-curve-卷)
* [第 6 步：格式化 curve 卷](#第-6-步格式化-curve-卷)
* [第 7 步：启动 pfsd 守护进程](#第-7-步启动-pfsd-守护进程)
* [其他：卸载 PolarFS](#其他卸载-polarfs)

第 1 步：环境准备
---

* [安装 CurveAdm](install-curveadm#安装-curveadm)

> 📢 **注意：**
>
> 目前 CurveAdm 只适配 `Debian10` 操作系统，未测试其他操作系统

第 2 步：导入主机
---

用户需导入安装 PolarFS 的主机列表，详见[主机管理][hosts]。

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
  - host: polarfs-host
    hostname: 10.0.1.1
```

### 2. 导入主机列表
```shell
$ curveadm hosts commit hosts.yaml
```

第 3 步：准备客户端配置文件
---

PolarFS 通过 curve-sdk 的方式将数据写入后端的 curve 卷，我们在安装 PolarFS 的同时需要指定 CurveBS 客户端配置文件来控制 curve-sdk 的行为

```shell
$ vim client.yaml
```

```shell
kind: curvebs
mds.listen.addr: 10.0.1.1:6700,10.0.1.2:6700,10.0.1.3:6700
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

第 4 步：安装 PolarFS
---

```shell
$ curveadm polarfs install --host polarfs-host -c client.yaml
```

* `--host`: 将 polarfs 安装到指定主机，用户可自行选择，请确保该主机已被导入
* `-c`: 指定 CurveBS 客户端配置文件

第 5 步：创建 curve 卷
---

我们在主机上安装了 `curve` 工具，该工具可用于创建卷，用户需要使用该工具创建实际存储 PolarFS 数据的 curve 卷

```shell
$ curve create --filename /volume --user my --length 10 --stripeUnit 16384 --stripeCount 64
```

用户可通过 `curve create -h` 命令查看创建卷的详细说明。上面的列子中，我们创建了一个拥有以下属性的卷：
 * 卷名为 `/volume`
 * 所属用户为 `my`
 * 大小为 `10GB`
 * 条带大小为 `16KB`
 * 条带个数为 `64`

**特别需要注意的是**，在数据库场景下，我们强烈建议使用条带卷，只有这样才能充分发挥 Curve 的性能优势，而 `16384 * 64` 的条带设置是目前最优的条带设置。

第 6 步：格式化 curve 卷
---

```shell
$ sudo pfs -C curve mkfs pool@@volume_my_
```

与我们在本地挂载文件系统前要先在磁盘上格式化文件系统一样，我们也要在我们的 curve 卷上格式化 PolarFS 文件系统。**特别需要注意**的是，由于 PolarFS 解析的特殊性，我们将以 `pool@${volume}_${user}_` 的形式指定我们的 curve 卷，此外还需要将卷名中的 `/` 替换成 `@`

第 7 步：启动 pfsd 守护进程
---

```shell
$ sudo /usr/local/polarstore/pfsd/bin/start_pfsd.sh -p pool@@volume_user_
```

如果 `pfsd` 启动成功，那么至此 curve 版 PolarFS 已全部部署完成，用户只需要根据 PolarDB 官方文档再在其上部署 PolarDB 数据库即可，详见[PolarDB 部署][polardb-deployment]


其他：卸载 PolarFS
---

```shell
$ curveadm polarfs uninstall --host polarfs-host
```

[hosts]: https://github.com/opencurve/curveadm/wiki/hosts
[important-config]: https://github.com/opencurve/curveadm/wiki/topology#curvebs-重要配置项
[curvebs-client-conf]: https://github.com/opencurve/curve/blob/master/conf/client.conf
[nebd-design]: https://github.com/opencurve/curve/blob/master/docs/cn/nebd.md
[polardb-deployment]: https://apsaradb.github.io/PolarDB-for-PostgreSQL/zh/deploying/deploy.html
