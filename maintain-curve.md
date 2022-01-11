CurveAdm 常见运维操作
===

* [查看集群列表](#查看集群列表)
* [切换集群](#切换集群)
* [删除集群](#删除集群)
* [导出集群](#导出集群)
* [导入集群](#导入集群)
* [查看集群拓扑](#查看集群拓扑)
* [修改集群拓扑](#修改集群拓扑)
* [对比集群拓扑](#对比集群拓扑)
* [查看集群状态](#查看集群状态)
* [启动服务](#启动服务)
* [停止服务](#停止服务)
* [重启服务](#重启服务)
* [修改服务配置](#修改服务配置)
* [进入服务容器](#进入服务容器)
* [清理集群](#清理集群)
* [获得 Curve 团队技术支持](#获得-curve-团队技术支持)

查看集群列表
---

```shell
$ curveadm cluster ls
```

若想显示集群详细信息，可添加 `-v` 选项，该选项将显示`集群 ID`、`集群 UUID`、`集群创建时间`、`集群描述`等相关信息：

```shell
$ curveadm cluster ls -v
```

切换集群
---

切换指定集群为当前管理集群。切换集群后，之后的操作都将作用于该集群。

```shell
$ curveadm cluster checkout <cluster-name>
```

切换集群后，再次查看集群列表时，当前集群名前会出现 `*` 样图标，我们以此来确定当前操作集群。

删除集群
---

```shell
$ curveadm cluster rm <cluster-name>
```

> :warning: **警告：**
> 
> 删除集群后，该集群相关信息将全部被清除，请谨慎操作。`curveadm` 支持同时管理多个集群，在不必要的情况下，请勿删除集群。

导出集群
---

我们可将当前集群信息保存为本地文件，通常我们在以下 2 种情况下需要导出集群：
  * 需将集群信息做定时备份，以防丢失
  * 与其他用户共享集群信息，以供其他用户操作集群

```shell
$ curveadm cluster export <cluster-name> [-o database-file-path]
``` 

> :bulb: **提示：**
> 
> 导出的集群文件以集群 `UUID` 为键，该 `UUID` 全局唯一，值则主要保存了以下 2 类信息：
> * 集群的服务配置，即集群拓扑 
> * 每个服务的相关信息，包括服务 ID、服务运行的容器 ID 等

导入集群
---

```shell
$ curveadm cluster import <cluster-name> [-f database-file-path]
```

> 📢 **注意：**
> * 导入集群时指定的集群名不可与已存在的集群名重复，否则将导致导入失败
> * 导入集群时会将集群拓扑原封不动导入，所以可能会导致某些配置项失效，如 SSH 的私钥路径配置项 `private_key_path`，
> 遇到此类情况，用户可修改集群拓扑并[提交](#提交集群拓扑)后即可正常操作集群

查看集群拓扑
---

```shell
$ curveadm config show
```

> :bulb: **提醒：**
>
> 当本地拓扑文件丢失时，我们可以通过保存当前的集群拓扑来恢复：
> ```shell
> $ curveadm config show > topology.yaml
> ```

修改集群拓扑
---

修改本地集群拓扑：

```shell
$ vim topology.yaml
```

当我们修改了本地的拓扑文件后，需向 CurveAdm 提交我们的修改，集群拓扑才能生效：

```shell
$ curveadm config commit <topology.yaml>
```

> :bulb: **提醒：**
> 
> 在提交本地集群拓扑时，终端会显示本地集群拓扑与当期集群拓扑之间的差异，请仔细对比，以防错误提交。

对比集群拓扑
---

我们可通过以下命令来查看本地集群拓扑文件与当前集群拓扑之间的差异：

```shell
$ curveadm config diff <topology.yaml>
```

查看集群状态
---

```shell
$ curveadm status
```

CurveAdm 默认会显示服务 ID、服务角色、主机地址、已部署的副本服务数量、容器 ID、运行状态：

```shell
Get Service Status: [OK]

cluster name    : my-cluster
cluster kind    : curvebs
cluster mds addr: 10.0.1.1:6666,10.0.1.2:6666,10.0.1.3:6666

Id            Role           Host      Replica  Container Id  Status
--            ----           ----      -------  ------------  ------
c9570c0d0252  etcd           10.0.1.1  1/1      ced84717bf4b  Up 45 hours
493b7831907c  etcd           10.0.1.2  1/1      907f8b84f527  Up 45 hours
8438cc5ecb52  etcd           10.0.1.3  1/1      44eca4798424  Up 45 hours
505da008b59c  mds            10.0.1.1  1/1      37c05bbb39af  Up 45 hours
e7bfb934182b  mds            10.0.1.2  1/1      044b56281928  Up 45 hours
1b322781339c  mds            10.0.1.3  1/1      b00481b9872d  Up 45 hours
<replica>     chunkserver    10.0.1.1  3/3      <replica>     RUNNING
<replica>     chunkserver    10.0.1.2  3/3      <replica>     RUNNING
<replica>     chunkserver    10.0.1.3  3/3      <replica>     RUNNING
2912bbdbcb48  snapshotclone  10.0.1.1  1/1      8b7a14b872ff  Up 45 hours
b862ef6720ed  snapshotclone  10.0.1.2  1/1      8e2a4b9e16b4  Up 45 hours
ed4533e903d9  snapshotclone  10.0.1.3  1/1      a35c30e3143d  Up 45 hours
```

* 若想查看其余信息，如日志目录、数据目录等，可添加 `-v` 参数
* 对于同一台主机上的[副本][replicas]服务来说，其状态默认是折叠的，可添加 `-s` 参数来显示每一个副本服务

启动服务
---

```shell
$ curveadm start
```

CurveAdm 默认启动集群中的所有服务，如需启动指定服务，可通过添加以下 3 个选项来实现：

* `--id`: 启动指定 `id` 的服务
* `--host`: 启动指定主机的所有服务
* `--role`: 启动指定角色的所有服务

以上 3 个选项可任意组合使用，服务对应的 `id`、`host`、`role` 可通过 [curveadm status](#查看集群状态) 来查看。

#### 示例 1： 启动 id 为 `c9570c0d0252` 的服务

```shell
$ curveadm start --id c9570c0d0252
```

#### 示例 2： 启动 `10.0.1.1` 这台主机上的所有 `MDS` 服务
```shell
$ curveadm start --host 10.0.1.1 --role mds 
```

停止服务
---

```shell
$ curveadm stop
```

CurveAdm 默认停止集群中的所有服务，如需停止指定服务，可通过添加以下 3 个选项来实现：

* `--id`: 停止指定 `id` 的服务
* `--host`: 停止指定主机的所有服务
* `--role`: 停止指定角色的所有服务

以上 3 个选项可任意组合使用，服务对应的 `id`、`host`、`role` 可通过 [curveadm status](#查看集群状态) 来查看。

#### 示例 1： 停止 id 为 `c9570c0d0252` 的服务

```shell
$ curveadm stop --id c9570c0d0252
```

#### 示例 2： 停止 `10.0.1.1` 这台主机上的所有 `MDS` 服务
```shell
$ curveadm stop --host 10.0.1.1 --role mds 
```

> :warning: **警告：**
> 
> 停止服务可能造成集群不健康，导致客户端 IO 失败，请谨慎操作。

重启服务
---

```shell
$ curveadm restart
```

CurveAdm 默认重启集群中的所有服务，如需重启指定服务，可通过添加以下 3 个选项来实现：

* `--id`: 重启指定 `id` 的服务
* `--host`: 重启指定主机的所有服务
* `--role`: 重启指定角色的所有服务

以上 3 个选项可任意组合使用，服务对应的 `id`、`host`、`role` 可通过 [curveadm status](#查看集群状态) 来查看。

#### 示例 1： 重启 id 为 `c9570c0d0252` 的服务

```shell
$ curveadm restart --id c9570c0d0252
```

#### 示例 2： 重启 `10.0.1.1` 这台主机上的所有 `MDS` 服务
```shell
$ curveadm restart --host 10.0.1.1 --role mds 
```

修改服务配置
---

在集群运行的过程中，我们可能需要修改服务的配置，并重启服务。具体操作步骤如下：

#### 第 1 步：编辑本地集群拓扑文件，修改对应服务的配置项

```shell
$ vim topology.yaml
```

> :bulb: **提醒：**
>
> 当本地拓扑文件丢失时，我们可以通过保存当前的集群拓扑来恢复：
> ```shell
> $ curveadm config show > topology.yaml
> ```
 
#### 第 2 步：提交修改
```shell
$ curveadm config commit topology.yaml
```

#### 第 3 步：重新加载服务
 
```shell
$ curveadm reload
```
 
CurveAdm 默认重新加载集群中的所有服务，如需重新加载指定服务，可通过添加以下 3 个选项来实现：

* `--id`: 重新加载指定 `id` 的服务
* `--host`: 重新加载指定主机的所有服务
* `--role`: 重新加载指定角色的所有服务

以上 3 个选项可任意组合使用，服务对应的 `id`、`host`、`role` 可通过 [curveadm status](#查看集群状态) 来查看。

#### 示例 1：重新加载 id 为 `c9570c0d0252` 的服务

```shell
$ curveadm reload --id c9570c0d0252
```

#### 示例 2：重新加载 `10.0.1.1` 这台主机上的所有 `MDS` 服务
```shell
$ curveadm reload --host 10.0.1.1 --role mds 
```

> :bulb: **提醒：**
> 
> 命令 [restart](#重启服务) 与 `reload` 的区别在于，
> `reload` 会根据当前集群拓扑的变更修改相应服务的配置，然后再重启服务，
> 而 `restart` 则只会简单的重启服务。

进入服务容器
---

我们可以远程进入服务容器内，查看服务进程、配置、日志、数据等信息:

```shell
$ curveadm enter <id>
```

服务对应的 `id` 可通过 [curveadm status](#查看集群状态) 来查看。

> :bulb: **提醒：**
> 
> CurveAdm 默认进入该服务的根目录，服务根目录包含了服务所需的所有文件，其目录结构如下：
> 
> ```shell
> /curvefs/mds  # 服务根目录
> |-- conf      # 配置目录
> |   `-- mds.conf
> |-- data      # 数据目录
> |-- logs      # 日志目录
> |   `-- curvefs-mds.log.INFO.20220120-142115.6
> `-- sbin      # 二进制目录
>     `-- curvefs-mds
> ```

清理集群
---

```shell
$ curveadm clean
```

CurveAdm 默认清理集群中所有服务的所有对象，如需清理指定服务或对象，可通过添加以下 4 个选项来实现：

* `--id`: 清理指定 `id` 所对应的服务
* `--host`: 清理指定主机所对应的服务
* `--role`: 清理指定角色所对应的服务
* `--only`: 清理指定对象 (`log`: 日志，`data`: 数据，`container`: 容器)

以上 4 个选项可任意组合使用，服务对应的 `id`、`host`、`role` 可通过 [curveadm status](#查看集群状态) 来查看。

#### 示例 1：清理 id 为 `c9570c0d0252` 服务的所有对象

```shell
$ curveadm clean --id c9570c0d0252
```

#### 示例 2：清理所有 MDS 服务的日志和数据

```shell
$ curveadm clean --role mds --only log,data
```

#### 示例 3：清理 `10.0.1.1` 这台主机上 `MDS` 服务的所有对象

```shell
$ curveadm clean --host 10.0.1.1 --role mds
```

> 📢 **注意：**
> 
> 当清理服务容器时，请确保相应服务已停止，你可以使用 [stop](#停止服务) 命令来停止指定服务。
 
获得 Curve 团队技术支持
---

当你的服务遇到异常，需要 Curve 团队的技术支持时，你可以通过 `support` 命令来将集群的一些必要信息，
如服务配置和日志进行打包压缩，并上传到 Curve 中心。
我们保证所有信息都将被加密，并只有你拥有秘钥：

```shell
$ curveadm support
```

> :bulb: **提醒：**
> 
> 当数据上传成功后，终端会显示相应的秘钥，请将该秘钥[告诉 Curve 团队成员][support]，以供他们分析及排查问题。

[support]: https://github.com/opencurve/curveadm/wiki/others#%E9%97%AE%E9%A2%98%E4%B8%8E%E5%8F%8D%E9%A6%88
[replicas]: https://github.com/opencurve/curveadm/wiki/topology#replica 