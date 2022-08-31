## 500000

**类别：** [预检][precheck]、[拓扑文件预检项][topology-precheck]

**描述：** 拓扑文件中的 `s3.ak` 配置项是无效的

**如何解决：** 若用户想要部署快照克隆服务，需确保拓扑文件中的 `s3.ak`、`s3.sk`、`s3.nos_address`、`s3.snapshot_bucket_name` 这 4 个 S3 相关配置有效且正确。若用户不需要使用快照克隆服务，在部署的时候可通过添加 `--skip snapshotclone` 选项跳过部署快照克隆服务

拓扑文件示例：

```yaml
kind: curvebs
global:
  s3.ak: 0123456789abcdefghijklmnopqrstuv
  s3.sk: abcdefghijklmnopqrstuv0123456789
  s3.nos_address: nos-eastchina1.126.net
  s3.snapshot_bucket_name: curve
...
```
## 500001

**类别：** [预检][precheck]、[拓扑文件预检项][topology-precheck]

**描述：** 拓扑文件中的 `s3.sk` 配置项是无效的

**如何解决：** 若用户想要部署快照克隆服务，需确保拓扑文件中的 `s3.ak`、`s3.sk`、`s3.nos_address`、`s3.snapshot_bucket_name` 这 4 个 S3 相关配置有效且正确。若用户不需要使用快照克隆服务，在部署的时候可通过添加 `--skip snapshotclone` 选项跳过部署快照克隆服务

拓扑文件示例：

```yaml
kind: curvebs
global:
  s3.ak: 0123456789abcdefghijklmnopqrstuv
  s3.sk: abcdefghijklmnopqrstuv0123456789
  s3.nos_address: nos-eastchina1.126.net
  s3.snapshot_bucket_name: curve
...
```

## 500002

**类别：** [预检][precheck]、[拓扑文件预检项][topology-precheck]

**描述：** 拓扑文件中的 `s3.nos_addresss` 配置项是无效的

**如何解决：** 若用户想要部署快照克隆服务，需确保拓扑文件中的 `s3.ak`、`s3.sk`、`s3.nos_address`、`s3.snapshot_bucket_name` 这 4 个 S3 相关配置有效且正确。若用户不需要使用快照克隆服务，在部署的时候可通过添加 `--skip snapshotclone` 选项跳过部署快照克隆服务

拓扑文件示例：

```yaml
kind: curvebs
global:
  s3.ak: 0123456789abcdefghijklmnopqrstuv
  s3.sk: abcdefghijklmnopqrstuv0123456789
  s3.nos_address: nos-eastchina1.126.net
  s3.snapshot_bucket_name: curve
...
```
## 500003

**类别：** [预检][precheck]、[拓扑文件预检项][topology-precheck]

**描述：** 拓扑文件中的 `s3.snapshot_bucket_name` 配置项是无效的

**如何解决：** 若用户想要部署快照克隆服务，需确保拓扑文件中的 `s3.ak`、`s3.sk`、`s3.nos_address`、`s3.snapshot_bucket_name` 这 4 个 S3 相关配置有效且正确。若用户不需要使用快照克隆服务，在部署的时候可通过添加 `--skip snapshotclone` 选项跳过部署快照克隆服务

拓扑文件示例：

```yaml
kind: curvebs
global:
  s3.ak: 0123456789abcdefghijklmnopqrstuv
  s3.sk: abcdefghijklmnopqrstuv0123456789
  s3.nos_address: nos-eastchina1.126.net
  s3.snapshot_bucket_name: curve
...
```

## 501000

**类别：** [预检][precheck]、[拓扑文件预检项][topology-precheck]

**描述：** 拓扑文件中日志目录和数据目录必须为绝对路径

**如何解决：** 用户需确保拓扑文件中 `log_dir`、`data_dir` 这 2 个目录配置项的值为绝对路径。用户可根据错误线索锁定哪一个配置项不符合要求，并对其进行修改，详见[CurveBS 重要配置项][curvebs-important-configure]、[CurveFS 重要配置项][curvefs-important-configure]

示例：

```yaml
kind: curvefs
global:
  log_dir: /curvefs/logs/${service_role}  # 正确
  data_dir: ~/data/${service_role}        # 错误
...
```

## 501001

**类别：** [预检][precheck]、[拓扑文件预检项][topology-precheck]

**描述：** 拓扑文件中存在重复的数据目录

**如何解决：** 用户需确保拓扑文件中 `data_dir` 配置项不能使用主机上相同的目录。用户可根据错误线索锁定哪一个配置项不符合要求，并对其进行修改，用户也可以参考集群的[拓扑模版][config-template]对其进行修改

示例：

```yaml
kind: curvebs
global:
  variable:
    machine1: server-host1
    machine2: server-host2
    machine3: server-host3

etcd_services:
  config:
    listen.ip: ${service_host}
    listen.port: 2380
    listen.client_port: 2379
    data_dir: /tmp/${service_role}
  deploy:
    - host: ${machine1}  # 正确，该服务的数据目录为 server-host1 上的 /tmp/etcd
    - host: ${machine2}  # 正确，该服务的数据目录为 server-host2 上的 /tmp/etcd
    - host: ${machine3}  # 正确，该服务的数据目录为 server-host3 上的 /tmp/etcd

chunkserver_services:
  config:
    listen.ip: ${service_host}
    listen.port: 82${format_replica_sequence}  # 8200, 8201, 8202
    data_dir: /data/chunkserver
    copysets: 100
  deploy:
    - host: ${machine1}  # 错误，该主机上的 3 个服务数据目录都为 server-host1 上的 /data/chunkserver
      replica: 3
    - host: ${machine2}  # 错误，该主机上的 3 个服务数据目录都为 server-host2 上的 /data/chunkserver
      replica: 3
    - host: ${machine3}  # 错误，该主机上的 3 个服务数据目录都为 server-host3 上的 /data/chunkserver
      replica: 3
```

## 502000

**类别：** [预检][precheck]、[拓扑文件预检项][topology-precheck]

**描述：** 拓扑文件中存在重复的监听地址

**如何解决：** 用户需确保拓扑文件中不存在重复的监听地址，一个监听地址是由服务监听 IP（`listen.ip`）和监听端口（`listen.*_port`）组合而成的。用户可根据错误线索锁定哪一个配置项不符合要求，并对其进行修改，用户也可以参考集群的[拓扑模版][config-template]对其进行修改

示例：

```yaml
kind: curvebs
global:
  variable:
    machine1: server-host1
    machine2: server-host2
    machine3: server-host3

etcd_services:
  config:
    listen.ip: ${service_host}
    listen.port: 2380
    listen.client_port: 2379
  deploy:
    - host: ${machine1}  # 正确，该服务监听地址为 server-host1:{2379,2380}
    - host: ${machine2}  # 正确，该服务监听地址为 server-host2:{2379,2380}
    - host: ${machine3}  # 正确，该服务监听地址为 server-host3:{2379,2380}
...
chunkserver_services:
  config:
    listen.ip: ${service_host}
    listen.port: 8200
    ...
  deploy:
    - host: ${machine1}  # 错误，该主机上的 3 个服务的监听地址都为 server-host1:8200，存在冲突
      replica: 3
    ...
...
```

## 503000

**类别：** [预检][precheck]、[拓扑文件预检项][topology-precheck]

**描述：** 拓扑文件中 etcd 服务至少需要 3 个服务实例

**如何解决：** 用户需修改集群拓扑文件，确保拓扑文件中 etcd 的服务实例为 3 个以上，用户可参考集群的[拓扑模版][config-template]对其进行修改

正确示例：

> 该拓扑文件中存在 3 个 etcd 服务实例

```yaml
kind: curvebs
global:
  variable:
    machine1: server-host1
    machine2: server-host2
    machine3: server-host3

etcd_services:
  config:
    listen.ip: ${service_host}
    listen.port: 2380
    listen.client_port: 2379
  deploy:
    - host: ${machine1}
    - host: ${machine2}
    - host: ${machine3}
...
```

错误示例：

> 该拓扑文件中只存在 2 个 etcd 服务实例

```yaml
kind: curvebs
global:
  variable:
    machine1: server-host1
    machine2: server-host2
    machine3: server-host3

etcd_services:
  config:
    listen.ip: ${service_host}
    listen.port: 2380
    listen.client_port: 2379
  deploy:
    - host: ${machine1}
    - host: ${machine2}
...
```

## 503001

**类别：** [预检][precheck]、[拓扑文件预检项][topology-precheck]

**描述：** 拓扑文件中 mds 服务至少需要 3 个服务实例

**如何解决：** 用户需修改集群拓扑文件，确保拓扑文件中 mds 的服务实例为 3 个以上，用户可参考集群的[拓扑模版][config-template]对其进行修改

正确示例：

> 该拓扑文件中存在 3 个 mds 服务

```yaml
kind: curvebs
global:
  variable:
    machine1: server-host1
    machine2: server-host2
    machine3: server-host3

mds_services:
  config:
    listen.ip: ${service_host}
    listen.port: 6700
    listen.dummy_port: 7700
  deploy:
    - host: ${machine1}
    - host: ${machine2}
    - host: ${machine3}
...
```

错误示例：

> 该拓扑文件中只存在 2 个 mds 服务

```yaml
kind: curvebs
global:
  variable:
    machine1: server-host1
    machine2: server-host2
    machine3: server-host3

mds_services:
  config:
    listen.ip: ${service_host}
    listen.port: 6700
    listen.dummy_port: 7700
  deploy:
    - host: ${machine1}
    - host: ${machine2}
...
```

## 503002

**类别：** [预检][precheck]、[拓扑文件预检项][topology-precheck]

**描述：** 拓扑文件中 chunkserver 服务至少需要 3 个服务实例

**如何解决：** 用户需修改集群拓扑文件，确保拓扑文件中 chunkserver 的服务实例为 3 个以上，用户可参考集群的[拓扑模版][config-template]对其进行修改

正确示例：

> 该拓扑文件中存在 9 个 chunkserver 服务实例

```yaml
kind: curvebs
global:
  variable:
    machine1: server-host1
    machine2: server-host2
    machine3: server-host3

chunkserver_services:
  config:
    listen.ip: ${service_host}
    listen.port: 82${format_replica_sequence}  # 8200, 8201, 8202
    data_dir: /data/chunkserver${service_replicas_sequence}  # /data/chunkserver0, /data/chunksever1, /data/chunkserver2
    copysets: 100
  deploy:
    - host: ${machine1}
      replica: 3
    - host: ${machine2}
      replica: 3
    - host: ${machine3}
      replica: 3
...
```

错误示例：

> 该拓扑文件中只存在 2 个 chunkserver 服务实例

```yaml
kind: curvebs
global:
  variable:
    machine1: server-host1
    machine2: server-host2
    machine3: server-host3

chunkserver_services:
  config:
    listen.ip: ${service_host}
    listen.port: 82${format_replica_sequence}  # 8200, 8201, 8202
    data_dir: /data/chunkserver${service_replicas_sequence}  # /data/chunkserver0, /data/chunksever1, /data/chunkserver2
    copysets: 100
  deploy:
    - host: ${machine1}
      replica: 2
...
```

## 503003

**类别：** [预检][precheck]、[拓扑文件预检项][topology-precheck]

**描述：** 拓扑文件中 snapshotclone 服务至少需要 3 个服务实例

**如何解决：** 用户需修改集群拓扑文件，确保拓扑文件中 snapshotclone 的服务实例为 3 个以上，用户可参考集群的[拓扑模版][config-template]对其进行修改。若永不需要部署快照克隆服务，可以直接将整个 `snapshotclone_services` 区块删除，也可通过该预检

正确示例：

> 该拓扑文件中存在 3 个 snapshoclone 服务实例

```yaml
kind: curvebs
global:
  variable:
    machine1: server-host1
    machine2: server-host2
    machine3: server-host3

snapshotclone_services:
  config:
    listen.ip: ${service_host}
    listen.port: 5555
    listen.dummy_port: 8081
    listen.proxy_port: 8080
  deploy:
    - host: ${machine1}
    - host: ${machine2}
    - host: ${machine3}
...
```

错误示例：

> 该拓扑文件中只存在 2 个 snapshotclone 服务实例

```yaml
kind: curvebs
global:
  variable:
    machine1: server-host1
    machine2: server-host2
    machine3: server-host3

snapshotclone_services:
  config:
    listen.ip: ${service_host}
    listen.port: 5555
    listen.dummy_port: 8081
    listen.proxy_port: 8080
  deploy:
    - host: ${machine1}
    - host: ${machine2}
...
```

## 503004

**类别：** [预检][precheck]、[拓扑文件预检项][topology-precheck]

**描述：** 拓扑文件中 metaserver 服务至少需要 3 个服务实例

**如何解决：** 用户需修改集群拓扑文件，确保拓扑文件中 metaserver 的服务实例为 3 个以上，用户可参考集群的[拓扑模版][config-template]对其进行修改

正确示例：

> 该拓扑文件中存在 3 个 metaserver 服务实例

```yaml
kind: curvefs
global:
  variable:
    machine1: server-host1
    machine2: server-host2
    machine3: server-host3

metaserver_services:
  config:
    listen.ip: ${service_host}
    listen.port: 6800
    listen.external_port: 7800
    metaserver.loglevel: 0
  deploy:
    - host: ${machine1}
    - host: ${machine2}
    - host: ${machine3}
...
```

错误示例：

> 该拓扑文件中只存在 2 个 metaserver 服务实例

```yaml
kind: curvefs
global:
  variable:
    machine1: server-host1
    machine2: server-host2
    machine3: server-host3

metaserver_services:
  config:
    listen.ip: ${service_host}
    listen.port: 6800
    listen.external_port: 7800
    metaserver.loglevel: 0
  deploy:
    - host: ${machine1}
    - host: ${machine2}
...
```

## 503005

**类别：** [预检][precheck]、[拓扑文件预检项][topology-precheck]

**描述：** 拓扑文件中 etcd 服务至少需要 3 台主机

**如何解决：** 为保证服务的高可用，以及隔离故障域，拓扑文件至少需要 3 台主机用于部署 etcd 服务，**特别需要注意**的是，用户只需要在 `deploy` 列表中增加 3 台主机即可，即使它们实际指向的是同一台物理主机

正确示例：

> 该拓扑文件中存在 3 个 etcd 服务实例，并且拥有 server-host{1,2,3} 3 台主机

```yaml
kind: curvebs
global:
  variable:
    machine1: server-host1
    machine2: server-host2
    machine3: server-host3

etcd_services:
  config:
    listen.ip: ${service_host}
    listen.port: 2380
    listen.client_port: 2379
  deploy:
    - host: ${machine1}
    - host: ${machine2}
    - host: ${machine3}
...
```

错误示例：

> 该拓扑文件中拥有 3 个 etcd 服务实例，但是其都位于 server-host1 这 1 台主机上

```yaml
kind: curvebs
global:
  variable:
    machine1: server-host1

etcd_services:
  config:
    listen.ip: ${service_host}
    listen.port: 2380
    listen.client_port: 2379
  deploy:
    - host: ${machine1}
      replicas:
...
```

## 503006

**类别：** [预检][precheck]、[拓扑文件预检项][topology-precheck]

**描述：** 拓扑文件中 mds 服务至少需要 3 台主机

**如何解决：** 为保证服务的高可用，以及隔离故障域，拓扑文件至少需要 3 台主机用于部署 mds 服务，**特别需要注意**的是，用户只需要在 `deploy` 列表中增加 3 台主机即可，即使它们实际指向的是同一台物理主机

正确示例：

> 该拓扑文件中存在 3 个 mds 服务实例，并且拥有 server-host{1,2,3} 3 台主机

```yaml
kind: curvebs
global:
  variable:
    machine1: server-host1
    machine2: server-host2
    machine3: server-host3

mds_services:
  config:
    listen.ip: ${service_host}
    listen.port: 6700
    listen.dummy_port: 7700
  deploy:
    - host: ${machine1}
    - host: ${machine2}
    - host: ${machine3}
...
```

错误示例：

> 该拓扑文件中拥有 3 个 mds 服务实例，但是其都位于 server-host1 这 1 台主机上

```yaml
kind: curvebs
global:
  variable:
    machine1: server-host1

mds_services:
  config:
    listen.ip: ${service_host}
    listen.port: 6700
    listen.dummy_port: 7700
  deploy:
    - host: ${machine1}
      replicas: 3

...
```

## 503007

**类别：** [预检][precheck]、[拓扑文件预检项][topology-precheck]

**描述：** 拓扑文件中 chunkserver 服务至少需要 3 台主机

**如何解决：** 为保证服务的高可用，以及隔离故障域，拓扑文件至少需要 3 台主机用于部署 chunkserver 服务，**特别需要注意**的是，用户只需要在 `deploy` 列表中增加 3 台主机即可，即使它们实际指向的是同一台物理主机

正确示例：

> 该拓扑文件中存在 3 个 chunkserver 服务实例，并且拥有 server-host{1,2,3} 3 台主机

```yaml
kind: curvebs
global:
  variable:
    machine1: server-host1
    machine2: server-host2
    machine3: server-host3

chunkserver_services:
  config:
    listen.ip: ${service_host}
    listen.port: 82${format_replica_sequence}  # 8200, 8201, 8202
    data_dir: /data/chunkserver${service_replica_sequence}  # /data/chunkserver0, /data/chunksever1, /data/chunkserver2
    copysets: 100
  deploy:
    - host: ${machine1}
    - host: ${machine2}
    - host: ${machine3}
...
```

错误示例：

> 该拓扑文件中拥有 3 个 chunkserver 服务实例，但是其都位于 server-host1 这 1 台主机上

```yaml
kind: curvebs
global:
  variable:
    machine1: server-host1
    machine2: server-host2
    machine3: server-host3

chunkserver_services:
  config:
    listen.ip: ${service_host}
    listen.port: 82${format_replica_sequence}  # 8200, 8201, 8202
    data_dir: /data/chunkserver${service_replica_sequence}  # /data/chunkserver0, /data/chunksever1, /data/chunkserver2
    copysets: 100
  deploy:
    - host: ${machine1}
      replicas: 3
...
```

## 503008

**类别：** [预检][precheck]、[拓扑文件预检项][topology-precheck]

**描述：** 拓扑文件中 snapshotclone 服务至少需要 3 台主机

**如何解决：** 为保证服务的高可用，以及隔离故障域，拓扑文件至少需要 3 台主机用于部署 snapshotclone 服务，**特别需要注意**的是，用户只需要在 `deploy` 列表中增加 3 台主机即可，即使它们实际指向的是同一台物理主机。若用户不需要部署快照克隆部署，可将整个 `snapshotclone_services` 区块删除即可

正确示例：

> 该拓扑文件中存在 3 个 snapshoeclone 服务实例，并且拥有 server-host{1,2,3} 3 台主机

```yaml
kind: curvebs
global:
  variable:
    machine1: server-host1
    machine2: server-host2
    machine3: server-host3

snapshotclone_services:
  config:
    listen.ip: ${service_host}
    listen.port: 5555
    listen.dummy_port: 8081
    listen.proxy_port: 8080
  deploy:
    - host: ${machine1}
    - host: ${machine2}
    - host: ${machine3}
...
```

错误示例：

> 该拓扑文件中拥有 3 个 snapshotclone 服务实例，但是其都位于 server-host1 这 1 台主机上

```yaml
kind: curvebs
global:
  variable:
    machine1: server-host1
    machine2: server-host2
    machine3: server-host3

snapshotclone_services:
  config:
    listen.ip: ${service_host}
    listen.port: 5555
    listen.dummy_port: 8081
    listen.proxy_port: 8080
  deploy:
    - host: ${machine1}
      replicas: 3
...
```

## 503009

**类别：** [预检][precheck]、[拓扑文件预检项][topology-precheck]

**描述：** 拓扑文件中 metaserver 服务至少需要 3 台主机

**如何解决：** 为保证服务的高可用，以及隔离故障域，拓扑文件至少需要 3 台主机用于部署 metaserver 服务，**特别需要注意**的是，用户只需要在 `deploy` 列表中增加 3 台主机即可，即使它们实际指向的是同一台物理主机。

正确示例：

> 该拓扑文件中存在 3 个 metaserver 服务实例，并且拥有 server-host{1,2,3} 3 台主机

```yaml
kind: curvefs
global:
  variable:
    machine1: server-host1
    machine2: server-host2
    machine3: server-host3

metaserver_services:
  config:
    listen.ip: ${service_host}
    listen.port: 6800
    listen.external_port: 7800
    metaserver.loglevel: 0
  deploy:
    - host: ${machine1}
    - host: ${machine2}
    - host: ${machine3}
...
```

错误示例：

> 该拓扑文件中拥有 3 个 metaserver 服务实例，但是其都位于 server-host1 这 1 台主机上

```yaml
kind: curvebs
global:
  variable:
    machine1: server-host1
    machine2: server-host2
    machine3: server-host3

metaserver_services:
  config:
    listen.ip: ${service_host}
    listen.port: 6800
    listen.external_port: 7800
    metaserver.loglevel: 0
  deploy:
    - host: ${machine1}
      replicas: 3
...
```

## 510000

**类别：** [预检][precheck]、[SSH 预检项][ssh-precheck]

**描述：** SSH 连接失败

**如何解决：** 当出现 SSH 连接问题时，你可以根据主机的相应配置， 在本地手动模拟 curveadm 的连接操作，以此来排查相应问题，详见[主机管理][hosts]

```shell
$ ssh <user>@<hostname> -p <ssh_port> -i <private_key_file>
```

## 520000

**类别：** [预检][precheck]、[权限预检项][permission-precheck]

**描述：** 当前的执行用户不存在

**如何解决：** 请确保主机列表中的执行用户存在于相应的主机，详见[主机管理][hosts]

## 520001

**类别：** [预检][precheck]、[权限预检项][permission-precheck]

**描述：** 主机上的 hostname 无法被解析

**如何解决：** CurveAdm 以以下命令探测当前物理主机的 hostname 能否被解析，若主机名能被解析，增该 `ping` 执行成功

```shell
$ ping $(hostname)
```

若当前机器的主机名无法被解析，用户可手动在 `/etc/hosts` 文件添加一条记录

```shell
$ vim /etc/hosts
127.0.0.1  curve.163.org  # 我们假设当前机器的 hostname 为 curve.163.org
```

## 520002

**类别：** [预检][precheck]、[权限预检项][permission-precheck]

**描述：** 创建目录没有权限

**如何解决：** CurveAdm 会尝试在部署服务的主机上以执行用户去创建 `log_dir`、`data_dir`，关于执行用户详见[主机管理][hosts]。当出现这类情况下，用户需要根据错误线索和查看对应的错误日志锁定具体拥有权限的日志目录，并手动更改其权限

## 520003

**类别：** [预检][precheck]、[权限预检项][permission-precheck]

**描述：** 没有权限执行 docker 命令

**如何解决：** CurveAdm 会尝试在部署服务的主机上以执行用户去运行 `sudo docker info`， 关于执行用户详见[主机管理][hosts]。 当出现这类情况下，用户需要根据错误线索和查看对应的错误日志锁定具体权限问题，并手动修改用户的执行权限或修改执行用户。此外，如果用户想去除 `sudo`，可修改 CurveAdm 配置文件中的 `sudo_alias` 配置项，将其设置为空即可，详见[CurveAdm 配置文件][curveadm-configure-file]

示例：

```ini
[defaults]
log_level = info
sudo_alias = ""  # 将其置为空即可
timeout = 300
auto_upgrade = true

[ssh_connections]
retries = 3
timeout = 10
```

# 530000

**类别：** [预检][precheck]、[内核预检项][kernel-precheck]

**描述：** 当前主机的内核版本无法被识别

**如何解决：** 用户可根据错误线索和查看相应的错误日志来确认对应的主机内核版本是否不符合要求，并根据规范进行修改

# 530001

**类别：** [预检][precheck]、[内核预检项][kernel-precheck]

**描述：** 当前内核版本不支持 `renameat()` 接口

**如何解决：** 用户可将 chunkserver 中的 `fs.enable_renameat2` 配置项设置成 `false`，提交拓扑文件后重新执行预检即可

示例：

```yaml
chunkserver_services:
  config:
    ...
    fs.enable_renameat2: false
```

# 530002

**类别：** [预检][precheck]、[内核预检项][kernel-precheck]

**描述：** 未加载内核 nbd 模块

**如何解决：** CurveBS 客户端依赖内核 `nbd` 模块，若当前操作系统没有该模块，用户需自行编译并将其导入

# 530003

**类别：** [预检][precheck]、[内核预检项][kernel-precheck]

**描述：** 未加载内核 fuse 模块

**如何解决：** CurveFS 客户端依赖内核 `fuse` 模块，若当前操作系统没有该模块，用户需自行编译并将其导入

# 540000

**类别：** [预检][precheck]、[网络预检项][network-precheck]

**描述：** 相应主机上端口已被占用

**如何解决：** 用户可根据错误线索快速锁定哪台主机的哪台端口已被占用，可登录主机杀掉占用此端口的服务或修改拓扑文件中服务的端口，使其与此端口不冲突

示例：

> 改错误线索表明 server-host1 这台主机的 7700 端口已被占用

```yaml
---
Error-Code: 540000
Error-Description: port is already in use
Error-Clue: host=server-host1, port=7700
How to Solve:
  * Website: https://github.com/opencurve/curveadm/wiki/errno#540000
  * Log: /home/nbs/.curveadm/logs/curveadm-2022-08-25_10-19-39.log
  * WeChat: opencurve_bot
```

# 540001

**类别：** [预检][precheck]、[网络预检项][network-precheck]

**描述：** 对应主机无法到达

**如何解决：** CurveAdm 在需要连通的服务之间会通过 `ping` 来探测其网络连通性，当出现网络不联通时，用户需根据错误线索和查看相应的日志锁定哪些主机之间的网络是不联通的，并解决之

# 540002

**类别：** [预检][precheck]、[网络预检项][network-precheck]

**描述：** 无法连接虚拟服务的地址

**如何解决：** CurveAdm 为检测防火墙是否禁止拓扑文件中服务监听的地址访问，会预先在各主机上启动一批相同监听地址的虚拟服务，并在需要互相访问的主机之间利用 `telnet` 来检测其连通性。当用户收到该错误码时，用户需要根据错误线索快速锁定哪台主机的哪些服务端口被禁止访问，并调整相应的防火墙规则

## 550000

**类别：** [预检][precheck]、[时间预检项][date-precheck]

**描述：** 无效的时间

**如何解决：** 我们在每台部署服务的主机上运行 `data +"%s"` 命令来获取当前主机的时间戳，在我们测试的几个发行版中该命令返回的都是以下格式的时间戳。

```shell
$ date +"%s"
1661171289
```

## 550001

**类别：** [预检][precheck]、[时间预检][date-precheck]

**描述：** 部署服务的主机之间的最大时间差超过 30 秒

**如何解决：** 用户得保证所有服务的主机之间的时差最多不超过 30 秒，不然可能导致服务的异常。当收到该错误码时，用户需调整对应主机的时间，使其时差保持在 30 秒之内

## 560000

**类别：** [预检][precheck]、[服务预检][service-precheck]

**描述：** 在 chunkserver 配置的数据目录（`data_dir`）中未发现 chunkfile pool

**如何解决：** 为减少写 IO 放大，我们需要提前生成一批固定大小并预写过的 chunk 文件，详见 [chunkfile pool 设计][chunkfilepool]。用户可根据使用需求选择以下两种解决方案的任意一种：

##### 方案一： 关闭 chunkfile pool

对于没有性能要求的测试或体验用户，可以直接关闭 chunkfile pool，用户只需要在 chunkserver 服务的配置项中设置 `chunkfilepool.enable_get_chunk_from_pool` 为 false，并提交对应的集群拓扑即可：

```shell
$ vim topology.yaml
```

```yaml
chunkserver_services:
  config:
    data_dir: /data/chunkserver${service_replicas_sequence}
    chunkfilepool.enable_get_chunk_from_pool: false  # 关闭 chunkfile pool
    ...
  deploy:
    - host: ${machine1}
      replicas: 3
    ...
```

```shell
$ curveadm commit topology.yaml
```

##### 方案二： 格式化磁盘以生成 chunkfile pool

详见[格式化磁盘][format-disk]。**特别需要注意**的是，`format` 命令只是启动后台格式化进程，并不意味着格式化已完成，用户需要通过 `curveadm format --status` 命令查看所有的磁盘都已格式化完成，再进行部署。

## 570000

**类别：** [预检][precheck]、[客户端预检][others-precheck]

**描述：** CurveFS 客户端配置文件中的 `s3.ak` 配置项是无效的

**如何解决：** 如果你以 S3 作为 CurveFS 的数据存储后端，需要在客户端的配置文件中指定对应 S3 的相关配置，请确保 `s3.ak`、`s3.sk`、`s3.endpoint`、`s3.bucket_name` 这 4 个配置项有效且正确

client.yaml 示例：

```yaml
kind: curvefs
s3.ak: 0123456789abcdefghijklmnopqrstuv
s3.sk: abcdefghijklmnopqrstuv0123456789
s3.endpoint: nos-eastchina1.126.net
s3.bucket_name: curve
...
```

## 570001

**类别：** [预检][precheck]、[客户端预检][others-precheck]

**描述：** CurveFS 客户端配置文件中的 `s3.sk` 配置项是无效的

**如何解决：** 如果你以 S3 作为 CurveFS 的数据存储后端，需要在客户端的配置文件中指定对应 S3 的相关配置，请确保 `s3.ak`、`s3.sk`、`s3.endpoint`、`s3.bucket_name` 这 4 个配置项有效且正确

client.yaml 示例：

```yaml
kind: curvefs
s3.ak: 0123456789abcdefghijklmnopqrstuv
s3.sk: abcdefghijklmnopqrstuv0123456789
s3.endpoint: nos-eastchina1.126.net
s3.bucket_name: curve
...
```

## 570002

**类别：** [预检][precheck]、[客户端预检][others-precheck]

**描述：** CurveFS 客户端配置文件中的 `s3.endpoint` 配置项是无效的

**如何解决：** 如果你以 S3 作为 CurveFS 的数据存储后端，需要在客户端的配置文件中指定对应 S3 的相关配置，请确保 `s3.ak`、`s3.sk`、`s3.endpoint`、`s3.bucket_name` 这 4 个配置项有效且正确

client.yaml 示例：

```yaml
kind: curvefs
s3.ak: 0123456789abcdefghijklmnopqrstuv
s3.sk: abcdefghijklmnopqrstuv0123456789
s3.endpoint: nos-eastchina1.126.net
s3.bucket_name: curve
...
```

## 570003

**类别：** [预检][precheck]、[客户端预检][others-precheck]

**描述：** CurveFS 客户端配置文件中的 `s3.bucket_name` 配置项是无效的

**如何解决：** 如果你以 S3 作为 CurveFS 的数据存储后端，需要在客户端的配置文件中指定对应 S3 的相关配置，请确保 `s3.ak`、`s3.sk`、`s3.endpoint`、`s3.bucket_name` 这 4 个配置项有效且正确

client.yaml 示例：

```yaml
kind: curvefs
s3.ak: 0123456789abcdefghijklmnopqrstuv
s3.sk: abcdefghijklmnopqrstuv0123456789
s3.endpoint: nos-eastchina1.126.net
s3.bucket_name: curve
...
```

## 590000

**类别：** [预检][precheck]、[其他预检项][others-precheck]

**描述：** 相关主机上未安装 Docker

**如何解决：** 用户可根据错误线索来确定哪台主机未安装 Docker，并登录相应主机安装 Docker，详见[安装依赖](install-requirements)

## 590001

**类别：** [预检][precheck]、[其他预检项][others-precheck]

**描述：** Docker 守护进程未运行

**如何解决：** 安装 Docker 后需确保 Docker 服务端的守护进程处于运行状态，你可以通过以下命令来启动 Docker:

```shell
$ sudo systemctl enable docker
$ sudo systemctl start docker
```

详见[启动 Docker][start-docker-daemon]

## 590002

**类别：** [预检][precheck]、[其他预检项][others-precheck]

**描述：** 磁盘空间已满

**如何解决：** 用户可根据错误线索来确定哪台主机哪块磁盘没有剩余空间，登录相应主机删除不必要的文件以腾出空间

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
[others-precheck]: https://github.com/opencurve/curveadm/wiki/errno#子类
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
