部署 CurveBS 客户端
===

* [第 1 步：环境准备](#第-1-步环境准备)
* [第 2 步：准备客户端配置文件](#第-2-步准备客户端配置文件)
* [第 3 步：映射 CurveBS 卷](#第-3-步映射-curvebs-卷)
* [其他：取消 CurveBS 卷映射](#其他取消-curvebs-卷映射)
 
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
> cluster name    : my-cluster
> cluster kind    : curvebs
> cluster mds addr: 10.0.1.1:6700,10.0.1.2:6700,10.0.1.3:6700
> ...
> ```

第 3 步：映射 CurveBS 卷
---

```shell
$ curveadm map <user>:<volume-name> -c client.yaml --create --size 10GB
```

* `<user>`: 该卷所属用户名，用户可自行定义
* `<volume-name>`: 卷名，用户可自行定义
* `--create`：当卷不存在时，则自行创建
* `--size`: 指定创建卷的大小，默认为 `10GB`

当用户映射卷成功后，在相应的机器上即能看到 CurveBS 卷对应的 nbd 设备：

```shell
$ lsblk | grep nbd
```

> 📢 **注意：**
> 
> * 卷名必须以 `/` 为起始，如 `/test`、`/curve/test`
> * 目前卷大小只支持以 `10GB` 为最小单位，即创建的卷大小只能是 10GB、20GB、30GB...，依此类推
> * 我们以 `<用户名>:<卷名>` 作为建来存储卷的相关信息，请勿挂载相同的卷

> :bulb: **提醒：**
> 
> 目前映射 CurveBS 卷是通过 SSH 远程映射的。在相应的机器上，
> 我们会检查是否已存在 [NEBD-part2][nebd-design] 服务，如若没有，我们会自动创建容器并启动。
 
### 示例
```shell
$ curveadm map curve:/test -c client.yaml --create
```

其他：取消 CurveBS 卷映射 
---

```shell
$ curveadm unmap <user>:<volume-name> -c client.yaml
``` 

[important-config]: https://github.com/opencurve/curveadm/wiki/topology#curvebs-重要配置项
[curvebs-client-conf]: https://github.com/opencurve/curve/blob/master/conf/client.conf
[nebd-design]: https://github.com/opencurve/curve/blob/master/docs/cn/nebd.md