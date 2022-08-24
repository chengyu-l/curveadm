# 301000

**类别：** [配置文件][configure]、[配置文件通用][configure-common]

**描述：** 不支持的配置值类型

**如何解决：** 配置文件中某些配置项对应的值只支持一些特定的类型，用户需要根据错误线索来锁定哪个配置项的值不符合要求，并进行相应的修正

# 301002

**类别：** [配置文件][configure]、[配置文件通用][configure-common]

**描述：** 配置项对应的值要求为布尔类型

**如何解决：** 用户需要根据错误线索来锁定哪个配置项对应的值要求为布尔类型，但实际却并不是，并将其根据需求修改成布尔类型（`true`、`false`）

# 301003

**类别：** [配置文件][configure]、[配置文件通用][configure-common]

**描述：** 配置项对应的值要求为整型

**如何解决：** 用户需要根据错误线索来锁定哪个配置项对应的值要求为整型，但实际却并不是，并将其根据需求修改成整型

# 301004

**类别：** [配置文件][configure]、[配置文件通用][configure-common]

**描述：** 配置项对应的值要求为非空字符串

**如何解决：** 用户需要根据错误线索来锁定哪个配置项对应的值要求为非空字符串，但实际却并不是，并将其根据需求修改成非空字符串

# 301005

**类别：** [配置文件][configure]、[配置文件通用][configure-common]

**描述：** 配置项对应的值要求为正整数

**如何解决：** 用户需要根据错误线索来锁定哪个配置项对应的值要求为正整数，但实际却并不是，并将其根据需求修改成正整数

# 301006

**类别：** [配置文件][configure]、[配置文件通用][configure-common]

**描述：** 配置项对应的值要求为字符串数组

**如何解决：** 用户需要根据错误线索来锁定哪个配置项对应的值要求为字符串数组，但实际却并不是，并将其根据需求修改成字符串数组

# 301100

**类别：** [配置文件][configure]、[配置文件通用][configure-common]

**描述：** 不支持的变量类型

**如何解决：** 目前配置文件中支持的变量类型为以下 3 类：`int`、`bool`、`string`，用户需要根据错误线索来锁定哪一个变量的值不符合要求，并进行相应的修正，详见[变量][variable]

示例：

```yaml
global:
  variable:
    name: curve         # 正确，变量类型为字符串
    age: 3              # 正确，变量类型为整数
    cloud-native: true  # 正确，变量类型为布尔
    labels:             # 错误，变量类型为数组，为不支持类型
      - storage
      - distruted system
  ...
```

# 301101

**类别：** [配置文件][configure]、[配置文件通用][configure-common]

**描述：** 无效的变量值

**如何解决：** 配置文件中变量对应的值不能为空，用户需要根据错误线索来锁定哪一个变量的值不符合要求，并进行相应的修正, 详见[变量][variable]



示例：

```yaml
global:
  variable:
    name: curve    # 正确
    age: 3         # 正确
    cloud-native:  # 错误，变量对应的值为空
```

## 310000

**类别：** [配置文件][configure]、[CurveAdm 配置文件][curveadm-configure]

**描述：** 解析 CurveAdm 配置文件失败

**如何解决：** 在修改 CurveAdm 配置文件时，有可能引入了一些不符合规范的字段，用户可根据错误线索及错误日志来锁定不符合规范的字段，并进行相应的修正，详见[CurveAdm 配置文件][curveadm-configure-file]

示例：
```ini
[defaults]
log_level = info
sudo_alias = "sudo"
timeout = 300
auto_upgrade = true

[ssh_connections]
retries = 3
timeout = 10
```

## 311000

**类别：** [配置文件][configure]、[CurveAdm 配置文件][curveadm-configure]

**描述：** CurveAdm 配置文件中的日志等级配置项无效

**如何解决：** 目前 CurveAdm 支持的日志等级为以下 4 个：`debug`、`info`、`warn`, `error`，用户可根据需求调整日志等级为对应的级别，详见[CurveAdm 配置文件][curveadm-configure-file]

示例：
```ini
[defaults]
log_level = info
sudo_alias = "sudo"
timeout = 300
auto_upgrade = true

[ssh_connections]
retries = 3
timeout = 10
```

## 311001

**类别：** [配置文件][configure]、[CurveAdm 配置文件][curveadm-configure]

**描述：** CurveAdm 配置文件中引入了不支持的配置项

**如何解决：** CurveAdm 配置文件中只可以填其支持的配置选项，当出现其他不支持的配置选项时，该错误码将会被报告，用户可根据错误线索去除那些不支持的配置项或进行修正，详见[CurveAdm 配置文件][curveadm-configure-file]

示例：
```ini
[defaults]
log_level = info
sudo_alias = "sudo"
timeout = 300
auto_upgrade = true

[ssh_connections]
retries = 3
timeout = 10
```

## 320000

**类别：** [配置文件][configure]、[主机配置文件][hosts-configure]

**描述：** 未找到主机配置文件

**如何解决：** 在导入主机列表时，请确保指定的主机配置文件存在，详见[主机管理][hosts]

示例：

```shell
$ curveadm hosts commit /path/to/hosts.yaml
```

## 320001

**类别：** [配置文件][configure]、[主机配置文件][hosts-configure]

**描述：** 读取主机配置文件失败

**如何解决：** 用户可根据错误线索或查看错误日志，来排除对应的问题，详见[主机管理][hosts]


## 320002

**类别：** [配置文件][configure]、[主机配置文件][hosts-configure]

**描述：** 主机配置文件不存在任何内容

**如何解决：** 用户需要将要导入的主机列表及其配置写入要提交的主机配置文件中，详见[主机管理][hosts]

示例：

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
    forward_agent: true
    become_user: admin
    labels:
      - client
```

## 320003

**类别：** [配置文件][configure]、[主机配置文件][hosts-configure]

**描述：** 解析主机配置文件失败

**如何解决：** 导入的主机配置文件有可能存在不符合规范的配置项，用户需要根据错误线索锁定那些配置项不符合规范并进行相应的修正，详见[主机管理][hosts]

示例：

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
    forward_agent: true
    become_user: admin
    labels:
      - client
```

## 321000

**类别：** [配置文件][configure]、[主机配置文件][hosts-configure]

**描述：** 主机配置文件中引入了不支持的配置项

**如何解决：** 主机配置文件中只可以填其支持的配置选项，当出现其他不支持的配置选项时，该错误码将会被报告，用户可根据错误线索去除那些不支持的配置项或进行修正，详见[主机管理][hosts]

## 321001

**类别：** [配置文件][configure]、[主机配置文件][hosts-configure]

**描述：** 主机缺少 `host` 配置项

**如何解决：** 主机列表中每一项中的 `host` 配置项是必须的，用户需要根据错误线索来锁定哪一个主机缺少 `host` 配置项并进行相应的修正，详见[主机管理][hosts]

示例：

```yaml
global:
  user: curve
  ssh_port: 22
  private_key_file: /home/curve/.ssh/id_rsa

hosts:
  - host: server-host1  # 正确，host 为 server-host1
    hostname: 10.0.1.1
  - hostname: 10.0.1.2  # 错误，缺少 host 配置项
```

## 321002

**类别：** [配置文件][configure]、[主机配置文件][hosts-configure]

**描述：** 主机缺少 `hostname` 配置项

**如何解决：** 主机列表中每一项中的 `hostname` 配置项是必须的，用户需要根据错误线索来锁定哪一个主机缺少 `hostname` 配置项并进行相应的修正，详见[主机管理][hosts]

示例：

```yaml
global:
  user: curve
  ssh_port: 22
  private_key_file: /home/curve/.ssh/id_rsa

hosts:
  - host: server-host1  # 正确，hostname 为 10.0.1.1
    hostname: 10.0.1.1
  - host: server-host2  # 错误，缺少 hostname 配置项
```

## 321003

**类别：** [配置文件][configure]、[主机配置文件][hosts-configure]

**描述：** 主机配置文件中的 SSH 端口超过最大端口限制

**如何解决：** 用户需要将主机配置文件中的 SSH 端口配置项 `ssh_port` 的值限制在 **[0, 65535]** 这个有效区间内，详见[主机管理][hosts]

正确示例：

```yaml
global:
  user: curve
  ssh_port: 22  # 正确
  ...
```

错误示例：

```yaml
global:
  user: curve
  ssh_port: 100000  # 错误，超过最大端口限制
  ...
```

## 321004

**类别：** [配置文件][configure]、[主机配置文件][hosts-configure]

**描述：** 主机配置文件中的 SSH 私钥路径必须为绝对路径

**如何解决：** 用户需要将主机配置文件中的私钥路径配置项 `private_key_file` 的值设为私钥的绝对路径，详见[主机管理][hosts]

正确示例：

```yaml
global:
  user: curve
  ssh_port: 22
  private_key_file: /home/curve/.ssh/id_rsa  # 正确
  ...
```

错误示例：

```yaml
global:
  user: curve
  ssh_port: 22
  private_key_file: ~/.ssh/id_rsa  # 错误，私钥路径需为绝对路径
  ...
```

## 321005

**类别：** [配置文件][configure]、[主机配置文件][hosts-configure]

**描述：** 主机配置文件中配置的 SSH 私钥文件不存在

**如何解决：** 用户需要确保主机配置文件中配置的私钥文件存在于本机。若用户以私钥的方式登录远程主机，则需要免密登录，可参考[<三步完成免密登录>](https://www.thegeekstuff.com/2008/11/3-steps-to-perform-ssh-login-without-password-using-ssh-keygen-ssh-copy-id/)，其他主机管理详见[主机管理][hosts]


## 321006

**类别：** [配置文件][configure]、[主机配置文件][hosts-configure]

**描述：** 主机配置文件中配置的 SSH 私钥文件的权限需为 `600`

**如何解决：** 用户需要确保主机配置文件中配置的私钥文件的权限为 `600`，否则有可能会出现登录失败的情况。当权限不符合要求时，用户可通过 `chmod` 命令修改其权限，详见[主机管理][hosts]

示例：

``` shell
$ chmod 600 ~/.ssh/id_rsa
$ ls -l ~/.ssh/id_rsa
-rw------- 1 root root 3243 8月   3 01:45 /root/.ssh/id_rsa
```

## 321007

**类别：** [配置文件][configure]、[主机配置文件][hosts-configure]

**描述：** 主机配置文件中存在重复的主机

**如何解决：** CurveAdm 以主机的 `host` 为键来查找主机的相应配置，用户需确保所有主机的 `host` 是独一无二的。当出现重复的主机 `host` 时，用户可根据错误线索来锁定哪些主机的 `host` 配置项重复了，并进行相应的修正，详见[主机管理][hosts]

示例：

```yaml
global:
  user: curve
  ssh_port: 22
  private_key_file: /home/curve/.ssh/id_rsa

hosts:
  - host: server-host1  # 正确
    hostname: 10.0.1.1
  - host: server-host1  # 错误，host 为 server-host1 的主机已被占用
    hostname: 10.0.1.1
```

## 321008

**类别：** [配置文件][configure]、[主机配置文件][hosts-configure]

**描述：** 主机配置文件中的 `hostname` 配置项要求为有效的 IP 地址

**如何解决：** 目前主机的 `hostname` 配置项被强制限制为主机的 IP 地址，用户需确保所有主机的 `hostname` 配置项都符合这一要求。后续我们有可能会开放域名的形式，详见[主机管理][hosts]

示例：

```yaml
global:
  user: curve
  ssh_port: 22
  private_key_file: /home/curve/.ssh/id_rsa

hosts:
  - host: server-host1
    hostname: 10.0.1.1         # 正确
  - host: server-host2
    hostname: curve.163.org    # 错误，hostname 不能为域名
  - host: server-host3
    hostname: 256.256.256.256  # 错误，该 IP 地址是无效的
```

## 330000

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 未找到用户指定的拓扑文件

**如何解决：** 用户需确保指定的集群拓扑文件存在于本机，详见[集群拓扑文件][topology]

## 330001

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 读取拓扑文件失败

**如何解决：** 用户可根据错误线索或查看错误日志来排除相应的问题，详见[集群拓扑文件][topology]

## 330002

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 集群拓扑文件中不存在任何内容

**如何解决：** 用户需将集群拓扑写入集群拓扑文件，并在命令行中指定该文件的绝对路径，详见[集群拓扑文件][topology]

示例：

```yaml
kind: curvefs
global:
  container_image: opencurvedocker/curvefs:latest
  data_dir: /home/curve/curvefs/data/${service_role}
  log_dir: /home/curve/curvefs/logs/${service_role}

etcd_services:
  config:
    listen.ip: ${service_host}
    listen.port: 2380
    listen.client_port: 2379

...
```

## 330003

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 解析集群拓扑失败

**如何解决：** 集群拓扑可能存在不符合规范的配置项，用户需要根据错误线索来锁定那些不符合规范的配置，并进行相应的修正，详见[集群拓扑文件][topology]

示例：

```yaml
kind: curvefs
global:
  container_image: opencurvedocker/curvefs:latest
  data_dir: /home/curve/curvefs/data/${service_role}
  log_dir: /home/curve/curvefs/logs/${service_role}

etcd_services:
  config:
    listen.ip: ${service_host}
    listen.port: 2380
    listen.client_port: 2379

...
```

## 330004

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 拓扑文件中的变量注册失败

**如何解决：** 变量的重复定义会导致变量注册失败，用户可根据错误线索或查看相应的错误日志来锁定具体的错误，并对其进行修正，详见[拓扑变量][variable]

示例：

```yaml
kind: curvefs
global:
  variable:
    name: curve    # 正确
    age: 18        # 正确
    name: curvefs  # 错误，变量重复定义
...
```

## 330005

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 拓扑文件中的变量解析失败

**如何解决：** 变量循环解析会导致变量解析失败，用户可根据错误线索或查看相应的错误日志来锁定具体的错误，并对其进行修正，详见[拓扑变量][variable]

示例：

```yaml
kind: curvefs
global:
  variable:
    from: ${to}
    to: ${from}
...
```

以上的变量块中，解析变量 `${from}` 依赖变量 `${to}` 的值，而解析变量 `${to}` 也要依赖变量 `${from}` 的值，这导致变量解析陷入死循环，无法成功解析所有变量

## 330006

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 设置变量失败

**如何解决：** *该错误码暂无使用场景*


## 330007

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 拓扑文件中的变量渲染失败

**如何解决：** 拓扑变量分为内建变量和自定义变量，用户需要确保在拓扑文件中填写的变量属于这 2 类，否则会导致变量渲染失败, 详见[拓扑变量][variable]

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
    listen.ip: ${service_host}  # 正确，${service_host} 为内建变量
    listen.port: 2380
    listen.client_port: 2379
  deploy:
    - host: ${machine1}  # 正确，${machine1} 为用户自定义变量
    - host: ${machine2}  # 正确，${machine2} 为用户自定义变量
    - host: ${machine3}  # 正确，${machine3} 为用户自定义变量
    - host: ${machine4}  # 错误，${machine4} 变量不存在
...
```

## 330008

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 计算集群拓扑创建的哈希值失败

**如何解决：** 这个错误属于代码逻辑错误，正常情况下不应该报告给用户，若用户收到该错误码，代表该版本 CurveAdm 代码逻辑存在问题，用户可通过以下 2 种方法反馈该问题：

* [提交 issue][issue]

* 添加 `opencurve_bot` 微信号，进入微信群




## 331000

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 不支持的集群类型

**如何解决：** 用户需在集群拓扑文件中指定该集群的正确类型。若部署 CurveBS 集群，则需指定 `kind` 配置项为 `curvebs`, 若部署 CurveFS 集群，则需指定 `kind` 配置项为 `curvefs`。详见[集群拓扑][topology]

CurveBS 集群拓扑示例：

```yaml
kind: curvebs
...
```

CurveFS 集群拓扑示例：

```yaml
kind: curvefs
...
```

## 331001

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 集群拓扑文件中不存在任何服务

**如何解决：** 集群拓扑文件是用来描述哪些服务部署在哪些机器上，用户需要根据集群的类型在集群拓扑文件中填写相应的服务列表，详见[集群拓扑][topology]

错误示例：

> 以下的集群拓扑文件中不存在任何服务

```yaml
kind: curvefs
global:
  container_image: opencurvedocker/curvefs:latest
  data_dir: /home/curve/curvefs/data/${service_role}
  log_dir: /home/curve/curvefs/logs/${service_role}
```

正确示例：

> 以下的集群拓扑文件中包含了 curvefs 所需的所有服务（etcd、mds、metaserver）

```yaml
kind: curvefs
global:
  container_image: opencurvedocker/curvefs:latest
  data_dir: /home/curve/curvefs/data/${service_role}
  log_dir: /home/curve/curvefs/logs/${service_role}

etcd_services:
  config:
    listen.ip: ${service_host}
    listen.port: 2380
    listen.client_port: 2379
  deploy:
    - host: server-host1
    - host: server-host2
    - host: server-host3

mds_services:
  config:
    ...
  deploy:
    ...

metaserver_services:
  config:
    ...
  deploy:
    ...
```

## 331002

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 集群拓扑文件中复制服务配置项要求为正整数

**如何解决：** 集群拓扑文件复制服务的配置项 `replicas` 的值要求为正整数，用户可根据错误线索来锁定哪一个台主机的 `replicas` 不符合要求，并进行相应的修正，详见[replicas][replicas]

示例：

```yaml
kind: curvebs
...
chunkserver_services:
  config:
    listen.port: 820${service_replicas_sequence}
  deploy:
    - host: server-host1
      replicas: 3         # 正确，replicas 为 3
    - host: server-host2
      replicas: -1        # 错误，replicas 需为正整数
    - host: server-host3
      replicas: 0         # 错误，replicas 需为正整数
```

## 331003

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 集群拓扑文件中的变量区块不符合规范

**如何解决：** 用户可在集群拓扑文件中的变量区块自定义变量，但要求该变量区块为字典类型，详见[拓扑变量][variable]

正确示例：

```yaml
kind: curvebs
global:
  variable:
    machine1: server-host1
    machine2: server-host2
    machine3: server-host3
```

错误示例：

```yaml
kind: curvebs
global:
  variable:
    - server-host1
    - server-host2
    - server-host3
```

## 331004

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 集群拓扑文件中存在重复的服务 ID

**如何解决：** 这个错误属于代码逻辑错误，正常情况下不应该报告给用户，若用户收到该错误码，代表该版本 CurveAdm 代码逻辑存在问题，用户可通过以下 2 种方法反馈该问题：

* [提交 issue][issue]

* 添加 `opencurve_bot` 微信号，进入微信群

## 332000

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 在提交拓扑时删除服务是被禁止的

**如何解决：** 为保证操作安全，在提交集群拓扑文件时，用户只能在拓扑文件中修改服务的配置项，不能对服务列表做任何增减操作

## 332001

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 在提交拓扑时增加服务是被禁止的

**如何解决：** 如果你需要新增服务，需要使用扩容命令，详见[扩容集群][scale-curve]

## 332002

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 在扩容集群时删除服务是被禁止的

**如何解决：** 为保证操作安全，扩容集群时，用户只能在拓扑文件中新增某一角色的服务列表，不能删除任何服务。用户可根据错误线索来在集群拓扑文件中恢复那些被删除的服务列表，详见[扩容集群][scale-curve]

## 332003

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 没有新增服务用于集群扩容

**如何解决：** 在扩容集群时，用户需要在拓扑文件中新增某一角色的服务列表，详见[扩容集群][scale-curve]

## 332004

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 每一次扩容操作要求新增服务列表为同一角色服务

**如何解决：** 在扩容集群时，用户需要在拓扑文件中新增扩容的服务列表，但要确保新增的服务列表都为同一角色，详见[扩容集群][scale-curve]

正确示例：

> 新增的 3 个服务都为 etcd 服务，符合要求

```yaml
kind: curvebs
global:
  container_image: opencurvedocker/curvebs:v1.2
  variable:
    machine1: debian10-001
    machine2: debian10-002
    machine3: debian10-003
    machine4: debian10-004  # 新增机器
    machine5: debian10-005  # 新增机器
    machine6: debian10-006  # 新增机器

etcd_services:
  config:
    listen.ip: ${service_host}
    listen.port: 2380
    listen.client_port: 2379
  deploy:
    - host: ${machine1}
    - host: ${machine2}
    - host: ${machine3}
    - host: ${machine4}  # 新增服务
    - host: ${machine5}  # 新增服务
    - host: ${machine6}  # 新增服务
...
```

错误示例：

> 新增的服务有 etcd 和 mds 服务，不符合要求

```yaml
kind: curvebs
global:
  container_image: opencurvedocker/curvebs:v1.2
  variable:
    machine1: debian10-001
    machine2: debian10-002
    machine3: debian10-003
    machine4: debian10-004  # 新增机器
    machine5: debian10-005  # 新增机器
    machine6: debian10-006  # 新增机器

etcd_services:
  config:
    listen.ip: ${service_host}
    listen.port: 2380
    listen.client_port: 2379
  deploy:
    - host: ${machine1}
    - host: ${machine2}
    - host: ${machine3}
    - host: ${machine4}  # 新增服务
    - host: ${machine5}  # 新增服务
    - host: ${machine6}  # 新增服务

mds_services:
  config:
    listen.ip: ${service_host}
    listen.port: 6700
    listen.dummy_port: 7700
  deploy:
    - host: ${machine1}
    - host: ${machine2}
    - host: ${machine3}
    - host: ${machine4}  # 新增服务
    - host: ${machine5}  # 新增服务
    - host: ${machine6}  # 新增服务
```

## 332005

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 每一次扩容 chunkserver 服务时需要至少新增 3 台主机来分布 zone

**如何解决：** 对于 chunkserver 服务来说，每次扩容都会新增一个逻辑池，新增的服务都位于该逻辑池中，请确保每次扩容至少增加 3 台以上的主机，否则 CurveBS 将没有足够的主机用来放置 3 副本。**特别需要注意**的是，3 台主机可以为同一主机，只要 `deploy` 中新增 3 个以上的 `host` 即可，即使他们实际指向同一台物理主机，详见[扩容集群][scale-curve]

正确示例：

> 新增 debian10-004、debian10-005、debian10-006 3 台主机

```yaml
kind: curvebs
global:
  container_image: opencurvedocker/curvebs:v1.2
  variable:
    machine1: debian10-001
    machine2: debian10-002
    machine3: debian10-003
    machine4: debian10-004  # 新增机器
    machine5: debian10-005  # 新增机器
    machine6: debian10-006  # 新增机器

chunkserver_services:
  config:
    listen.ip: ${service_host}
    listen.port: 82${format_replicas_sequence}
    data_dir: /data/chunkserver${service_replicas_sequence}
    copysets: 100
  deploy:
    - host: ${machine1}
      replicas: 3
    - host: ${machine2}
      replicas: 3
    - host: ${machine3}
      replicas: 3
    - host: ${machine4}  # 新增服务
      replicas: 3
    - host: ${machine5}  # 新增服务
      replicas: 3
    - host: ${machine6}  # 新增服务
      replicas: 3
...
```

错误示例：

> 只新增 debian10-004 这一台主机

```yaml
kind: curvebs
global:
  container_image: opencurvedocker/curvebs:v1.2
  variable:
    machine1: debian10-001
    machine2: debian10-002
    machine3: debian10-003
    machine4: debian10-004  # 新增机器
    machine5: debian10-005  # 新增机器
    machine6: debian10-006  # 新增机器

chunkserver_services:
  config:
    listen.ip: ${service_host}
    listen.port: 82${format_replicas_sequence}
    data_dir: /data/chunkserver${service_replicas_sequence}
    copysets: 100
  deploy:
    - host: ${machine1}
      replicas: 3
    - host: ${machine2}
      replicas: 3
    - host: ${machine3}
      replicas: 3
    - host: ${machine4}  # 新增服务
      replicas: 3
...
```

## 332006

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 每一次扩容 metaserver 服务时需要至少新增 3 台主机来分布 zone

**如何解决：** 对于 metaserver 服务来说，每次扩容都会新增一个逻辑池，新增的服务都位于该逻辑池中，请确保每次扩容至少增加 3 台以上的主机，否则 CurveFS 将没有足够的主机用来放置 3 副本。**特别需要注意**的是，3 台主机可以为同一主机，只要 `deploy` 中新增 3 个以上的 `host` 即可，即使他们实际指向同一台物理主机，详见[扩容集群][scale-curve]

正确示例：

> 新增 debian10-004、debian10-005、debian10-006 3 台主机

```yaml
kind: curvefs
global:
  container_image: opencurvedocker/curvefs:latest
  variable:
    machine1: debian10-001
    machine2: debian10-002
    machine3: debian10-003
    machine4: debian10-004  # 新增机器
    machine5: debian10-005  # 新增机器
    machine6: debian10-006  # 新增机器

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
    - host: ${machine4}  # 新增服务
    - host: ${machine5}  # 新增服务
    - host: ${machine6}  # 新增服务
...
```

错误示例：

> 只新增 debian10-004 这一台主机

```yaml
kind: curvefs
global:
  container_image: opencurvedocker/curvefs:latest
  variable:
    machine1: debian10-001
    machine2: debian10-002
    machine3: debian10-003
    machine4: debian10-004  # 新增机器
    machine5: debian10-005  # 新增机器
    machine6: debian10-006  # 新增机器

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
    - host: ${machine4}  # 新增服务
...
```

## 332007

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 在迁移服务的同时增加服务是被禁止的

**如何解决：** 为保证操作安全，在迁移服务的时候，用户只能在拓扑文件中修改同一服务角色下的一台主机名，不能对服务列表做任何增减操作，详见[迁移服务][migrate-curve]


## 332008

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 在迁移服务的同时删除服务是被禁止的

**如何解决：** 为保证操作安全，在迁移服务的时候，用户只能在拓扑文件中修改同一服务角色下的一台主机名，不能对服务列表做任何增减操作，详见[迁移服务][migrate-curve]

## 332009

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 没有服务用于迁移

**如何解决：** 在迁移服务时，用户需要在拓扑文件中修改同一服务角色下的一台主机名，从而产生迁移的服务列表，详见[迁移服务][migrate-curve]

## 332010

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 每一次迁移操作要求迁移的服务列表为同一角色服务

**如何解决：** 在迁移服务时，用户需确保迁移的服务列表都为同一角色服务，详见[迁移服务][migrate-curve]

正确示例：

> 迁移的服务都为 etcd 服务，符合要求

```yaml
kind: curvebs
global:
  container_image: opencurvedocker/curvebs:v1.2
  variable:
    machine1: debian10-001
    machine2: debian10-002
    machine3: debian10-003
    machine4: debian10-004  # 迁移目的机器

etcd_services:
  config:
    listen.ip: ${service_host}
    listen.port: 2380
    listen.client_port: 2379
  deploy:
    - host: ${machine1}
    - host: ${machine2}
    - host: ${machine4}  # 正确，原先服务部署于 ${machine3}，现在迁移到 ${machine4}
...
```

错误示例：

> 迁移的服务有 etcd 和 mds 服务，不符合要求

```yaml
kind: curvebs
global:
  container_image: opencurvedocker/curvebs:v1.2
  variable:
    machine1: debian10-001
    machine2: debian10-002
    machine3: debian10-003
    machine4: debian10-004  # 迁移目的机器

etcd_services:
  config:
    listen.ip: ${service_host}
    listen.port: 2380
    listen.client_port: 2379
  deploy:
    - host: ${machine1}
    - host: ${machine2}
    - host: ${machine4}  # 正确，原先服务部署于 ${machine3}，现在迁移到 ${machine4}

mds_services:
  config:
    listen.ip: ${service_host}
    listen.port: 6700
    listen.dummy_port: 7700
  deploy:
    - host: ${machine1}
    - host: ${machine2}
    - host: ${machine4}  # 错误，不能与 etcd 混在一起迁移
```

## 332011

**类别：** [配置文件][configure]、[拓扑配置文件][topology-configure]

**描述：** 每一次迁移操作要求迁移同一角色服务下某一台主机的所有服务实例

**如何解决：** 在迁移服务时，用户需确保同一角色服务下某一台主机的所有服务实例都被迁移，不能只迁移一部分，详见[迁移服务][migrate-curve]

正确示例：

```shell
chunkserver_services:
  config:
    listen.ip: ${service_host}
    listen.port: 82${format_replicas_sequence}
    data_dir: /data/chunkserver${service_replicas_sequence}
    copysets: 100
  deploy:
    - host: ${machine1}
      replicas: 3
    - host: ${machine2}
      replicas: 3
    - host: ${machine4}  # 将 ${machine3} 上的所有服务迁移到 ${machine4}
      replicas: 3
```

## 340000

**类别：** [配置文件][configure]、[格式化配置文件][format-configure]

**描述：** 未找到格式化配置文件

**如何解决：** 用户在执行格式化时需通过 `-f` 选项指定格式化配置文件，若用户没有指定，CurveAdm 默认使用当前目录下的 `format.yaml` 作为格式化配置文件。请确保对应的格式化配置文件存在，详见[格式化磁盘][format-disk]


## 340001

**类别：** [配置文件][configure]、[格式化配置文件][format-configure]

**描述：** 解析格式化配置文件失败

**如何解决：** 格式化配置文件中有可能出现了不符合规范的配置项，用户可根据错误线索锁定具体的错误，并对其进行修正，详见[格式化磁盘][format-disk]

示例：

```yaml
host:
  - server-host1
  - server-host2
  - server-host3
disk:
  - /dev/sda:/data/chunkserver0:90  # device:mount_path:format_percent
  - /dev/sdb:/data/chunkserver1:90
  - /dev/sdc:/data/chunkserver2:90
```

## 341000

**类别：** [配置文件][configure]、[格式化配置文件][format-configure]

**描述：** 格式化配置文件中的磁盘列表不符合规范

**如何解决：** 磁盘列表中的每一项由设备、挂载路径、格式化百分比这 3 部分组成，并由 `:` 符合间隔开，用户可根据错误线索锁定哪一项不符合规范并进行相应修正，详见[格式化磁盘][format-disk]

示例：
```yaml
host:
  - host1
  - host2
  - host3
disk:
  - /dev/sda:/data/chunkserver0      # 错误：缺少格式化百分比
  - /dev/sdb:/data/chunkserver1:     # 错误: 缺少格式化百分比
  - /dev/sdc:/data/chunkserver2:90:  # 错误: 多出一部分
  - /dev/sdd:/data/chunkserver3:90   # 正确
```

## 341001

**类别：** [配置文件][configure]、[格式化配置文件][format-configure]

**描述：** 格式化配置文件中存在无效的设备名

**如何解决：** 磁盘列表中的每一项由设备、挂载路径、格式化百分比这 3 部分组成，并由 `:` 符号间隔开，其中的设备名必须为绝对路径，用户可根据错误线索锁定哪一项不符合规范并进行相应修正

示例：
```yaml
host:
  - host1
  - host2
  - host3
disk:
  - sda:/data/chunkserver0:90       # 错误：磁盘 sda 的设备名应为 /dev/sda
  - /dev/sdb:/data/chunkserver1:90  # 正确
  - /dev/sdc:/data/chunkserver2:90  # 正确
```

## 341002

**类别：** [配置文件][configure]、[格式化配置文件][format-configure]

**描述：** 格式化配置文件中的挂载路径必须为绝对路径

**如何解决：** 磁盘列表中的每一项由设备、挂载路径、格式化百分比这 3 部分组成，并由 `:` 符合间隔开，其中的挂载路径必须为绝对路径，用户可根据错误线索锁定哪一项不符合规范并进行相应修正

示例：
```yaml
host:
  - host1
  - host2
  - host3
disk:
  - /dev/sda:chunkserver0:90        # 错误：挂载路径 chunkserver0 应为 /data/chunkserver0
  - /dev/sdb:/data/chunkserver1:90  # 正确
  - /dev/sdc:/data/chunkserver2:90  # 正确
```

## 341003

**类别：** [配置文件][configure]、[格式化配置文件][format-configure]

**描述：** 格式化配置文件中的格式化百分比需为整数

**如何解决：** 磁盘列表中的每一项由设备、挂载路径、格式化百分比这 3 部分组成，并由 `:` 符合间隔开，其中的格式化百分比需为整数，用户可根据错误线索锁定哪一项不符合规范并进行相应修正

示例：
```yaml
host:
  - host1
  - host2
  - host3
disk:
  - /dev/sda:/data/chunkserver0:10.5  # 错误：10.5 为浮点
  - /dev/sdb:/data/chunkserver1:xyz   # 错误：xyz 不是整数
  - /dev/sdc:/data/chunkserver2:90    # 正确
```

## 341004

**类别：** [配置文件][configure]、[格式化配置文件][format-configure]

**描述：** 格式化配置文件中的格式化百分比需在 **[1, 100]** 这个区间内

**如何解决：** 磁盘列表中的每一项由设备、挂载路径、格式化百分比这 3 部分组成，并由 `:` 符合间隔开，其中的格式化百分比需在 **[1, 100]** 这个区间内，用户可根据错误线索锁定哪一项不符合规范并进行相应修正

示例：
```yaml
host:
  - host1
  - host2
  - host3
disk:
  - /dev/sda:/data/chunkserver0:-1   # 错误：-1 不在 [1, 100] 区间内
  - /dev/sdb:/data/chunkserver1:101  # 错误：101 不在 [1, 100] 区间内
  - /dev/sdc:/data/chunkserver1:0    # 错误：0 不在 [1, 100] 区间内
  - /dev/sdd:/data/chunkserver2:1    # 正确
  - /dev/sde:/data/chunkserver3:100  # 正确
```

## 350000

**类别：** [配置文件][configure]、[客户端配置文件][client-configure]

**描述：** 解析客户端配置文件失败

**如何解决：** 客户端配置文件中有可能出现了不符合规范的配置项，用户可根据错误线索锁定具体的错误，并对其进行修正，详见[CurveBS 客户端配置文件][curvebs-client-configure]、[CurveFS 客户端配置文件][curvefs-client-configure]


## 350001

**类别：** [配置文件][configure]、[客户端配置文件][client-configure]

**描述：** 解析客户端配置文件中的变量出错

**如何解决：** 这个错误属于代码逻辑错误，正常情况下不应该报告给用户，若用户收到该错误码，代表该版本 CurveAdm 代码逻辑存在问题，用户可通过以下 2 种方法反馈该问题：

* [提交 issue][issue]

* 添加 `opencurve_bot` 微信号，进入微信群



## 351000

**类别：** [配置文件][configure]、[客户端配置文件][client-configure]

**描述：** 不支持的客户端配置文件类别

**如何解决：** 用户需要根据部署的客户端类型，在配置文件 `client.yaml` 中填写对应的 `kind` 配置项，来表明该配置文件的类型。如 CurveBS 的 `kind` 应为 `curvebs`，而 CurveFS 则为 `curvefs`

CurveBS 客户端配置文件示例：

```yaml
kind: curvebs
...
```

CurveFS 客户端配置文件示例：

```yaml
kind: curvefs
...
```

## 351001

**类别：** [配置文件][configure]、[客户端配置文件][client-configure]

**描述：** 客户端配置文件中配置项的值类型为无效类型

**如何解决：** 目前客户端配置文件配置项的值支持的类型为以下 3 个：`int`、`bool`、`string`，用户需要根据错误线索锁定哪些配置项不合符要求，并进行相应的修改

## 351002

**类别：** [配置文件][configure]、[客户端配置文件][client-configure]

**描述：** 进行 map 操作时需要指定 `curvebs` 类型的客户端配置文件

**如何解决：** 用户需要在客户端配置文件中指定 `kind` 配置项为 `curvebs`

client.yaml 示例：

```yaml
kind: curvebs
...
```

## 351003

**类别：** [配置文件][configure]、[客户端配置文件][client-configure]

**描述：** 进行 mount 操作时需要指定 `curvefs` 类型的客户端配置文件

**如何解决：** 用户需要在客户端配置文件中指定 `kind` 配置项为 `curvefs`

client.yaml 示例：

```yaml
kind: curvefs
...
```

## 351004

**类别：** [配置文件][configure]、[客户端配置文件][client-configure]

**描述：** 客户端配置文件中的 MDS 监听地址是无效的

**如何解决：** 用户可通过 `curveadm status` 命令查看集群 MDS 的监听地址，并将其填入客户端配置文件中

```shell
$ curveadm status
Get Service Status: [OK]

cluster name      : my-cluster
cluster kind      : curvebs
cluster mds addr  : 10.0.1.1:6700,10.0.1.2:6700,10.0.1.3:6700  # 此行为集群 MDS 监听地址
cluster mds leader: 10.0.1.1:6700 / 505da008b59c
```

CurveBS 客户端配置文件示例：

```yaml
kind: curvebs
mds.listen.addr: 10.0.1.1:6700,10.0.1.2:6700,10.0.1.3:6700
...
```

CurveFS 客户端配置文件示例：

```yaml
kind: curvefs
mdsOpt.rpcRetryOpt.addrs: 10.0.1.1:6700,10.0.1.2:6700,10.0.1.3:6700
...
```

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
