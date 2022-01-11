部署 CurveFS 客户端
===

* [第 1 步：环境准备](#第-1-步环境准备)
* [第 2 步：准备客户端配置文件](#第-2-步准备客户端配置文件)
* [第 3 步：挂载 CurveFS 文件系统](#第-3-步挂载-curvefs-文件系统)
* [其他：卸载文件系统](#其他卸载文件系统)

第 1 步：环境准备
---

* [软硬件环境需求](install-curveadm#软硬件环境需求)
* [安装依赖](install-curveadm#安装依赖)
 
第 2 步：准备客户端配置文件
---

```shell
$ vim client.yaml
```

```shell
user: curve
host: 10.0.1.1
ssh_port: 22
private_key_file: /home/curve/.ssh/id_rsa
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
> ...

> 📢 **注意：**
> 
> 用户如需开启本地磁盘缓存，请务必配置 data_dir 配置项。

第 3 步：挂载 CurveFS 文件系统
---

```shell
$ sudo curveadm mount <curvefs-name> <mount-point> -c client.yaml
```

* `<curvefs-name>`: 文件系统名，用户可自行定义
* `<mount-point>`: 挂载路径，用户可自行定义，但必须为**绝对路径**

如果文件系统挂载成功，在相应的机器上即能查询到 CurveFS 文件系统对应的挂载项：

```shell
$ mount | grep <mount-point>
```

> 📢 **注意：**
> 
> 目前 CurveAdm 只支持本机挂载文件系统， 之后将支持通过 SSH 远程挂载。
> 在本机挂载文件系统时，请确保挂载用户有 `sudo` 权限。

### 示例
```shell
$ sudo curveadm mount /test /mnt/curve -c client.yaml
```

其他：卸载文件系统
---

```shell
$ sudo curveadm umount <mount-point>
``` 

[important-config]: https://github.com/opencurve/curveadm/wiki/topology#curvefs-重要配置项
[curvefs-client-conf]: https://github.com/opencurve/curve/blob/master/curvefs/conf/client.conf