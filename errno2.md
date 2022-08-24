## 210000

**类别：** [命令行选项][command-options]、[服务选项][cluster-options]

**描述：** 未找到 ID 对应的服务

**如何解决：** 服务 ID 可通过 `curveadm status` 命令来查看，该命令显示的第一列即为服务 ID：

```shell
$ curveadm status
```

```shell
Get Service Status: [OK]

cluster name      : my-cluster
cluster kind      : curvefs
cluster mds addr  : 10.0.1.1:6700,10.0.1.2:6700,10.0.1.3:6700
cluster mds leader: 10.0.1.1:6700 / 505da008b59c

Id            Role        Host          Replicas  Container Id  Status
--            ----        ----          -------   ------------  ------
c9570c0d0252  etcd        server-host1  1/1       ced84717bf4b  Up 45 hours
493b7831907c  etcd        server-host2  1/1       907f8b84f527  Up 45 hours
8438cc5ecb52  etcd        server-host3  1/1       44eca4798424  Up 45 hours
505da008b59c  mds         server-host1  1/1       37c05bbb39af  Up 45 hours
e7bfb934182b  mds         server-host2  1/1       044b56281928  Up 45 hours
1b322781339c  mds         server-host3  1/1       b00481b9872d  Up 45 hours
2912bbdbcb48  metaserver  server-host1  1/1       8b7a14b872ff  Up 45 hours
b862ef6720ed  metaserver  server-host2  1/1       8e2a4b9e16b4  Up 45 hours
ed4533e903d9  metaserver  server-host3  1/1       a35c30e3143d  Up 45 hours
```

## 210001

**类别：** [命令行选项][command-options]、[服务选项][cluster-options]

**描述：** 不支持的 CurveBS 服务角色

**如何解决：** CurveBS 目前支持的服务角色仅为以下 4 个：`etcd`、`mds`、`metaserver`、`snapshotclone`。
请确保在需要输入服务角色的命令选项中填写上述服务角色中的一个，用户也可以使用 `help` 命令来查看相应命令的帮助提示信息。

## 210002

**类别：** [命令行选项][command-options]、[服务选项][cluster-options]

**描述：** 不支持的 CurveFS 服务角色

**如何解决：** CurveFS 目前支持的服务角色为以下 3 个：`etcd`、`mds`、`metaserver`。
请确保在需要输入服务角色的命令选项中填写上述服务角色中的一个，用户也可以使用 `help` 命令来查看相应命令的帮助提示信息。

## 210003

**类别：** [命令行选项][command-options]、[服务选项][cluster-options]

**描述：** 不支持跳过的服务角色

**如何解决：** 目前支持跳过的服务（即不部署该服务）仅为快照克隆服务，因为其他服务都为必须服务。当用户不需要使用快照克隆服务时，在部署时可添加该选项来跳过对其的部署：

示例：

```shell
$ curveadm deploy --skip snapshotclone
```

## 210004

**类别：** [命令行选项][command-options]、[服务选项][cluster-options]

**描述：** 不支持跳过的预检项

**如何解决：** 目前支持的预检项为以下 7 个，用户可选择相应的预检项跳过：

| 检查项 | 跳过选项   | 说明                                          |
| :---   | :---       | :---                                          |
| 拓扑   | topology   | 检查集群拓扑的合法性                          |
| SSH    | ssh        | 检查 SSH 的连通性                             |
| 权限   | permission | 检查当前用户执行 docker、创建目录等权限       |
| 内核   | kernel     | 检查内核版本、内核模块是否满足要求            |
| 网络   | network    | 检查网络连通性、防火墙等                      |
| 时间   | date       | 检查主机之间的时间差是否过大                  |
| 服务   | service    | 检查服务数量、chunkfile pool、S3 配置有效性等 |

详见[部署预检][precheck]

示例：

```shell
$ curveadm precheck --skip topology
```

## 210005

**类别：** [命令行选项][command-options]、[服务选项][cluster-options]

**描述：** 不支持的清理对象

**如何解决：** 目前支持的清理对象为以下 3 个：

| 清理对象 | 说明 |
| :--- | :--- |
| log | 日志目录 |
| data | 数据目录 |
| container | 容器 |

详见[清理集群][clean-cluster]

示例：

```shell
$ cuvreadm clean --only log,data
```

## 210006

**类别：** [命令行选项][command-options]、[服务选项][cluster-options]

**描述：** 没有匹配的服务

**如何解决：** 在用户指定服务 ID、服务主机、服务角色等限定条件后，CurveAdm 未找到任何一个匹配的服务。用户可通过 `curveadm status` 命令来查看各服务的属性，并指定正确的限制属性：

```shell
$ curveadm status
```

```shell
Get Service Status: [OK]

cluster name      : my-cluster
cluster kind      : curvefs
cluster mds addr  : 10.0.1.1:6700,10.0.1.2:6700,10.0.1.3:6700
cluster mds leader: 10.0.1.1:6700 / 505da008b59c

Id            Role        Host          Replicas  Container Id  Status
--            ----        ----          -------   ------------  ------
c9570c0d0252  etcd        server-host1  1/1       ced84717bf4b  Up 45 hours
493b7831907c  etcd        server-host2  1/1       907f8b84f527  Up 45 hours
8438cc5ecb52  etcd        server-host3  1/1       44eca4798424  Up 45 hours
505da008b59c  mds         server-host1  1/1       37c05bbb39af  Up 45 hours
e7bfb934182b  mds         server-host2  1/1       044b56281928  Up 45 hours
1b322781339c  mds         server-host3  1/1       b00481b9872d  Up 45 hours
2912bbdbcb48  metaserver  server-host1  1/1       8b7a14b872ff  Up 45 hours
b862ef6720ed  metaserver  server-host2  1/1       8e2a4b9e16b4  Up 45 hours
ed4533e903d9  metaserver  server-host3  1/1       a35c30e3143d  Up 45 hours
```

错误示例:

```shell
$ curveadm stop --id c9570c0d0252 --role mds
```

用户在 `stop` 服务时指定了服务的 ID 为 `c9570c0d0252`，同时又指定了服务的角色为 `mds`，而我们从以上集群状态显示来看，ID 为 `c9570c0d0252` 服务对应的角色为 `etcd`，所以 CurveAdm 无法找到同时满足以上 2 个条件的服务，遂报告了该错误码。

## 220000

**类别：** [命令行选项][command-options]、[客户端通用选项][client-common-options]

**描述：** 不支持的客户端类型

**如何解决：** 目前支持的客户端类型为以下 2 个：`curvebs`、`curvefs`

## 221000

**类别：** [命令行选项][command-options]、[CurveBS 客户端选项][curvebs-client-options]

**描述：** 无效的卷格式

**如何解决：** 一个有效的卷由卷所属用户和卷名两部分组成，并以 `:` 作为间隔符分隔开, 详见[映射 CurveBS 卷][map]

示例：

```shell
$ curveadm map user:/volume --host client-host --create  # 正确，卷为 user:/volume
$ curveadm map user: --host client-host --create         # 错误，缺少卷名
$ curveadm map :/volume --host client-host --create      # 错误，缺少卷所属用户
```

## 221001

**类别：** [命令行选项][command-options]、[CurveBS 客户端选项][curvebs-client-options]

**描述：** `root` 不能作为卷的所属用户

**如何解决：** `root` 用户作为特殊用户已被 CurveBS 使用，所以用户在挂载卷的时候不能以 `root` 用户作为卷的所属用户。用户可指定任意用户作为卷的所属用户，**特别需要注意**的是，该用户不需要在当前主机上真实存在，卷的所属用户是 CurveBS 的一个逻辑概念，表明该卷由谁创建，详见[映射 CurveBS 卷][map]


示例：

```shell
$ curveadm map root:/volume --host client-host --create   # 错误，root 不能作为卷的所属用户
$ curveadm map curve:/volume --host client-host --create  # 正确
$ curveadm map test:/volume --host client-host --create   # 正确
```


## 221002

**类别：** [命令行选项][command-options]、[CurveBS 客户端选项][curvebs-client-options]

**描述：** 卷名必须以 `/` 为起始

**如何解决：** 目前 CurveBS 的卷名必须以 `/` 为起始，如 `/test`、`/volume`，详见[映射 CurveBS 卷][map]

示例：

```shell
$ curveadm map curve:/volume --host client-host --create  # 正确，/volume 可以作为卷名
$ curveadm map curve:volume --host client-host --create   # 错误，volume 不能作为卷名
```

## 221003

**类别：** [命令行选项][command-options]、[CurveBS 客户端选项][curvebs-client-options]

**描述：** 卷大小必须以 `GB` 为结尾

**如何解决：** 目前 CurveBS 的卷大小必须以 `GB` 为结尾，如 `10GB`、`1024GB`，详见[映射 CurveBS 卷][map]

示例：

```shell
$ curveadm map curve:/volume --host client-host --create --size 10GB  # 正确，10GB 为有效卷大小
$ curveadm map curve:/volume --host client-host --create --size 10    # 错误，10 为无效卷大小
```

## 221004

**类别：** [命令行选项][command-options]、[CurveBS 客户端选项][curvebs-client-options]

**描述：** 卷大小必须为正整数

**如何解决：** 请确保以正整数来指定卷的大小，详见[映射 CurveBS 卷][map]

示例：

```shell
$ curveadm map curve:/volume --host client-host --create --size 10GB   # 正确，10GB 为有效卷大小
$ curveadm map curve:/volume --host client-host --create --size -10GB  # 错误，-10GB 为无效卷大小
```

## 221005

**类别：** [命令行选项][command-options]、[CurveBS 客户端选项][curvebs-client-options]

**描述：** 卷大小必须 10GB 的整数倍

**如何解决：** 目前卷大小只支持以 10GB 为最小单位，即创建的卷大小只能是 10GB、20GB、30GB...，依此类推，详见[映射 CurveBS 卷][map]

示例：

```shell
$ curveadm map curve:/volume --host client-host --create --size 10GB    # 正确，10GB 为有效卷大小
$ curveadm map curve:/volume --host client-host --create --size 1024GB  # 正确，1024GB 为有效卷大小
$ curveadm map curve:/volume --host client-host --create --size 15GB    # 错误，15GB 为无效卷大小
$ curveadm map curve:/volume --host client-host --create --size 5GB     # 错误，5GB 为无效卷大小
```

## 221006

**类别：** [命令行选项][command-options]、[CurveBS 客户端选项][curvebs-client-options]

**描述：**  客户端配置文件不存在

**如何解决：** 映射 CurveBS 卷时需要指定客户端配置文件，用户可通过 `-c` 选项指定配置文件，若用户未指定，CurveAdm 默认会以当前目录下的 `client.yaml` 作为客户端配置文件。

示例：

```shell
$  curveadm map user:/volume --host client-host -c /path/to/client.yaml
```

## 221007

**类别：** [命令行选项][command-options]、[CurveBS 客户端选项][curvebs-client-options]

**描述：**  未找到 ID 对应的客户端

**如何解决：** 客户端 ID 可通过 `curveadm client status` 命令来查看，该命令显示的第一列即为客户端 ID：

```shell
$ curveadm client status
```

```shell
Get Client Status: [OK]

Id            Kind     Host          Container Id  Status       Aux Info
--            ----     ----          ------------  ------       --------
362d538778ad  curvebs  client-host1  cfa00fd01ae8  Up 36 hours  {"user":"curve","volume":"/test1"}
b0d56cfaad14  curvebs  client-host2  c0301eff2af0  Up 36 hours  {"user":"curve","volume":"/test2"}
c700e1f6acab  curvebs  client-host3  52554173a54f  Up 36 hours  {"user":"curve","volume":"/test3"}
```

示例：

```shell
$ curveadm client enter 362d538778ad
```

## 220000

**类别：** [命令行选项][command-options]、[CurveFS 客户端选项][curvefs-client-options]

**描述：**  CurveFS 文件系统的挂载点必须为绝对路径

**如何解决：**  在挂载 CurveFS 文件系统时，指定的挂载点必须要求为绝对路径，详见[挂载 CurveFS 文件系统][mount]

示例：

```shell
$ curveadm curveadm mount /fs1 /mnt/test1 --host client-host  # 正确，/mnt/test1 为绝对路径
$ curveadm curveadm mount /fs1 test1 --host client-host       # 错误，test1 为相对路径
```

## 230000

**类别：** [命令行选项][command-options]、[playground 选项][playground-options]

**描述：** 不支持的 playground 类型

**如何解决：** 目前支持的 playground 类型为以下 2 个：`curvebs`、`curvefs`，用户需要在运行 playground 时通过 `--kind` 选项指定对应的类型

示例：
```shell
$ curveadm playground run --kind curvebs                              # 正确，运行一个 CurveBS 的 playground
$ curveadm playground run --kind curvefs --mountpoint /path/to/mount  # 正确，运行一个 CurveFS 的 playground，并指定对应的挂载点
$ curveadm playground run                                             # 错误，未指定 playground 类型
```

## 230001

**类别：** [命令行选项][command-options]、[playground 选项][playground-options]

**描述：** 对于 CurveFS 类型的 playground，必须指定对应的挂载点

**如何解决：** 当运行 CurveFS 类型的 playground 时，用户需要通过 `--mountpoint` 选项来指定文件系统的挂载点

示例：
```shell
$ curveadm playground run --kind curvefs --mountpoint /path/to/mount  # 正确，/path/to/mount 为挂载点
$ curveadm playground run --kind curvefs                              # 错误，未指定相应的挂载点
```

## 230002

**类别：** [命令行选项][command-options]、[playground 选项][playground-options]

**描述：**  playground 的挂载点要求为绝对路径

**如何解决：** 当运行 CurveFS 类型的 playground 时，用户通过 `--mountpoint` 选项指定的文件系统挂载点必须为绝对路径

示例：
```shell
$ curveadm playground run --kind curvefs --mountpoint /mnt/test  # 正确，/mnt/test 为有效挂载点
$ curveadm playground run --kind curvefs --mountpoint test       # 错误，test 为无效挂载点
```

## 230003

**类别：** [命令行选项][command-options]、[playground 选项][playground-options]

**描述：**  playground 指定的挂载点不存在

**如何解决：** 当运行 CurveFS 类型的 playground 时，用户通过 `--mountpoint` 选项指定的文件系统挂载点必须要求已存在于本机，若指定的挂载点未创建，用户需要在本机手动创建对应的挂载路径

[init]: https://github.com/opencurve/curveadm/wiki/errno#主类
[database]: https://github.com/opencurve/curveadm/wiki/errno#主类
[command-options]: https://github.com/opencurve/curveadm/wiki/errno#主类
[configure]: https://github.com/opencurve/curveadm/wiki/errno#主类
[common-logic]: https://github.com/opencurve/curveadm/wiki/errno#主类
[execute-command]:  https://github.com/opencurve/curveadm/wiki/errno#主类
[cluster-options]: https://github.com/opencurve/curveadm/wiki/errno#子类
[client-common-options]:  https://github.com/opencurve/curveadm/wiki/errno#子类
[curvebs-client-options]:  https://github.com/opencurve/curveadm/wiki/errno#子类
[curvefs-client-options]:  https://github.com/opencurve/curveadm/wiki/errno#子类
[playground-options]: https://github.com/opencurve/curveadm/wiki/errno#子类
[configure-common]: https://github.com/opencurve/curveadm/wiki/errno#子类
[curveadm-configure]: https://github.com/opencurve/curveadm/wiki/errno#子类
[hosts-configure]: https://github.com/opencurve/curveadm/wiki/errno#子类
[topology-configure]: https://github.com/opencurve/curveadm/wiki/errno#子类
[format-configure]: https://github.com/opencurve/curveadm/wiki/errno#子类
[client-configure]: https://github.com/opencurve/curveadm/wiki/errno#子类
[hosts-logic]:  https://github.com/opencurve/curveadm/wiki/errno#子类
[service-logic]: https://github.com/opencurve/curveadm/wiki/errno#子类
[curvefs-client-logic]: https://github.com/opencurve/curveadm/wiki/errno#子类
[curvebs-client-logic]: https://github.com/opencurve/curveadm/wiki/errno#子类
[polarfs-logic]: https://github.com/opencurve/curveadm/wiki/errno#子类
[playground-logic]: https://github.com/opencurve/curveadm/wiki/errno#子类
[topology-precheck]: https://github.com/opencurve/curveadm/wiki/errno#子类
[ssh-precheck]: https://github.com/opencurve/curveadm/wiki/errno#子类
[permission-precheck]: https://github.com/opencurve/curveadm/wiki/errno#子类
[kernel-precheck]: https://github.com/opencurve/curveadm/wiki/errno#子类
[network-precheck]: https://github.com/opencurve/curveadm/wiki/errno#子类
[date-precheck]: https://github.com/opencurve/curveadm/wiki/errno#子类
[service-precheck]: https://github.com/opencurve/curveadm/wiki/errno#子类
[command-common]: https://github.com/opencurve/curveadm/wiki/errno#子类
[others-command]: https://github.com/opencurve/curveadm/wiki/errno#子类
[docker-command]:  https://github.com/opencurve/curveadm/wiki/errno#子类
[ssh-connection]:  https://github.com/opencurve/curveadm/wiki/errno#子类
[shell-command]:  https://github.com/opencurve/curveadm/wiki/errno#子类
[database-common-solution]: https://github.com/opencurve/curveadm/wiki/errno#数据库类错误码
[insatall-directory]: https://github.com/opencurve/curveadm/wiki/overview#安装目录
[precheck]: https://github.com/opencurve/curveadm/wiki/precheck
[clean-cluster]: https://github.com/opencurve/curveadm/wiki/maintain-curve#清理集群
[map]: https://github.com/opencurve/curveadm/wiki/curvebs-client-deployment#第-4-步映射-curvebs-卷
[mount]: https://github.com/opencurve/curveadm/wiki/curvefs-client-deployment#第-4-步挂载-curvefs-文件系统
[variable]: https://github.com/opencurve/curveadm/wiki/topology#变量
[curveadm-configure-file]: https://github.com/opencurve/curveadm/wiki/install-curveadm#curveadm-配置文件
[hosts]: https://github.com/opencurve/curveadm/wiki/hosts
[topology]: https://github.com/opencurve/curveadm/wiki/topology
[replicas]: https://github.com/opencurve/curveadm/wiki/topology#replicas
[scale-curve]: https://github.com/opencurve/curveadm/wiki/scale-curve
[migrate-curve]: https://github.com/opencurve/curveadm/wiki/migrate-curve
[format-disk]: https://github.com/opencurve/curveadm/wiki/curvebs-cluster-deployment#第-4-步格式化磁盘
[curvebs-client-configure]: https://github.com/opencurve/curveadm/wiki/curvebs-client-deployment#第-3-步准备客户端配置文件
[curvefs-client-configure]: https://github.com/opencurve/curveadm/wiki/curvefs-client-deployment#第-3-步准备客户端配置文件
[add-and-checkout]: https://github.com/opencurve/curveadm/wiki/curvebs-cluster-deployment#第-6-步添加集群并切换集群
[umap]: https://github.com/opencurve/curveadm/wiki/curvebs-client-deployment#其他取消-curvebs-卷映射
[umount]: https://github.com/opencurve/curveadm/wiki/curvefs-client-deployment#其他卸载文件系统
[chunkfilepool]: https://github.com/opencurve/curve/blob/master/docs/cn/chunkserver_design.md#224-chunkfilepool
[config-template]: https://github.com/opencurve/curveadm/tree/master/configs
[execute-command-common-solution]: https://github.com/opencurve/curveadm/wiki/errno#命令执行类错误码
[start-docker-daemon]: https://yeasy.gitbook.io/docker_practice/install/ubuntu#qi-dong-docker
[issue]: https://github.com/opencurve/curveadm/issues
[topology]: https://github.com/opencurve/curveadm/wiki/topology
[install-requirements]: https://github.com/opencurve/curveadm/wiki/install-curveadm#安装依赖
[important-configure]: https://github.com/opencurve/curveadm/wiki/topology#curvebs-重要配置项
[curvebs-important-configure]: https://github.com/opencurve/curveadm/wiki/topology#curvebs-重要配置项
[curvefs-important-configure]: https://github.com/opencurve/curveadm/wiki/topology#curvefs-重要配置项
