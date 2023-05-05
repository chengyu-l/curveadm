# CurveFS 使用 memcached

- [CurveFS 使用 memcached](#curvefs-使用-memcached)
  - [使用 Curveadm playbook 部署 memcached 集群](#使用-curveadm-playbook-部署-memcached-集群)
    - [第 1 步：环境准备](#第-1-步环境准备)
    - [第 2 步：在中控机上安装 CurveAdm](#第-2-步在中控机上安装-curveadm)
    - [第 3 步：导入主机列表](#第-3-步导入主机列表)
      - [1. 准备主机列表](#1-准备主机列表)
        - [global](#global)
        - [hosts](#hosts)
        - [host 和 hostname](#host-和-hostname)
        - [lables](#lables)
        - [env](#env)
      - [2. 导入主机列表](#2-导入主机列表)
    - [第 4 步：部署 memcached 集群](#第-4-步部署-memcached-集群)
    - [第 5 步：检查 memcached 集群状态](#第-5-步检查-memcached-集群状态)
    - [其他运维操作](#其他运维操作)
      - [停止 memcached 集群](#停止-memcached-集群)
      - [清理 memcached 集群](#清理-memcached-集群)
- [添加 memcached 集群](#添加-memcached-集群)
- [CurveFS client 使用 memcached 集群](#curvefs-client-使用-memcached-集群)

`memcached` 可以为 `curvefs client` 提供缓存功能，解决 `client` 对接 s3 时性能不足的问题。
下面介绍如何使用 `curveadm playbook` 部署 `memcached` 。
所使用到的所有文件可以从 [memcached](https://github.com/opencurve/curveadm/tree/master/playbook/memcached) 中获取。

## 使用 Curveadm playbook 部署 memcached 集群

### 第 1 步：环境准备

* [软硬件环境需求](install-curveadm#软硬件环境需求)
* [安装依赖](install-curveadm#安装依赖)

### 第 2 步：在中控机上安装 CurveAdm

* [安装 CurveAdm](install-curveadm#安装-curveadm)

### 第 3 步：导入主机列表

用户需导入部署集群所需的机器列表和相关的信息，以便在之后的各类配置文件中填写部署服务的主机名，
请确保在之后各类配置文件出现的主机名都已导入，详见[主机管理][hosts]。

#### 1. 准备主机列表

```shell
vim hosts.yaml
```

```yaml
global:
  user: curve
  ssh_port: 22
  private_key_file: /home/curve/.ssh/id_rsa

hosts:
  - host: server-host1
    hostname: 10.0.1.1
    labels:
      - memcached
    envs:
      - IMAGE=memcached:1.6.17
      - LISTEN=10.0.1.1
      - PORT=11211
      - USER=root
      - MEMORY_LIMIT=32768 # item memory in megabytes
      - MAX_ITEM_SIZE=8m # adjusts max item size (default: 1m, min: 1k, max: 1024m)
      - EXT_PATH=/mnt/memcachefile/cachefile:1024G
      - EXT_WBUF_SIZE=8 # size in megabytes of page write buffers.
      - EXT_ITEM_AGE=1 # store items idle at least this long (seconds, default: no age limit)
  - host: server-host2
    hostname: 10.0.1.2
    labels:
      - memcached
    envs:
      - IMAGE=memcached:1.6.17
      - LISTEN=10.0.1.2
      - PORT=11211
      - USER=root
      - MEMORY_LIMIT=32768 # item memory in megabytes
      - MAX_ITEM_SIZE=8m # adjusts max item size (default: 1m, min: 1k, max: 1024m)
      - EXT_PATH=/mnt/memcachefile/cachefile:1024G
      - EXT_WBUF_SIZE=8 # size in megabytes of page write buffers.
      - EXT_ITEM_AGE=1 # store items idle at least this long (seconds, default: no age limit)
  - host: server-host3
    hostname: 10.0.1.3
    labels:
      - memcached
    envs:
      - IMAGE=memcached:1.6.17
      - LISTEN=10.0.1.3
      - PORT=11211
      - USER=root
      - MEMORY_LIMIT=32768 # item memory in megabytes
      - MAX_ITEM_SIZE=8m # adjusts max item size (default: 1m, min: 1k, max: 1024m)
      - EXT_PATH=/mnt/memcachefile/cachefile:1024G
      - EXT_WBUF_SIZE=8 # size in megabytes of page write buffers.
      - EXT_ITEM_AGE=1 # store items idle at least this long (seconds, default: no age limit)
```

下面介绍一下 `hosts.yaml` 中各个部分的作用。

##### global

其中 `global` 部分记录 `ssh` 使用的端口、用户、私钥等信息。

##### hosts

`hosts` 记录了要部署服务的机器的信息：

```yaml
  - host: server-host1
    hostname: 10.0.1.1
    labels:
      - memcached
    envs:
      - IMAGE=memcached:1.6.17
      - LISTEN=10.0.1.1
      - PORT=11211
      - USER=root
      - MEMORY_LIMIT=32768 # item memory in megabytes
      - MAX_ITEM_SIZE=8m # adjusts max item size (default: 1m, min: 1k, max: 1024m)
      - EXT_PATH=/mnt/memcachefile/cachefile:1024G
      - EXT_WBUF_SIZE=8 # size in megabytes of page write buffers.
      - EXT_ITEM_AGE=1 # store items idle at least this long (seconds, default: no age limit)
```

##### host 和 hostname

`hostname` 记录要部署机器的 `ip`，使用 `host` 作为机器的别名，一个 `host` 表示一台机器，是 `curveadm playbook` 执行的基本单位。
如果想在一台机器上部署多个服务，可以保持 `hostname` 一致，然后使用不同的 `host` 来区别。
以下方式表示在机器 `10.0.1.1` 上同时部署两个服务：

```yaml
- host: server-host1
    hostname: 10.0.1.1
...
- host: server-host2
    hostname: 10.0.1.1
```

##### lables

`labels` 表示机器的标签，允许有多个，拥有同一个 `label` 的机器表示归属于同一个组，`curveadm playbook` 管理的基本单位。
`curveadm playbook` 执行时使用 `-l` 或者 `--label` 指定标签来识别要执行命令的机器。
如以下命令就表示在所有拥有 label `memcached` 机器上执行 `xxx.sh` 脚本。

```bash
curveadm playbook --label memcached xxx.sh
```

##### env

`env` 里设置一些环境变量供后续脚本执行时使用，大部分为 `memcached` 的启动参数，这部分内容可以参考部署脚本 `deploy.sh` 中的 `init` 的内容。
在创建镜像时指定启动的参数，在 `init` 中根据需求配置对应的启动参数。
如需要添加新的参数请自行修改 `hosts.yaml`中的 `env` 和 `deploy.sh` 中的 `init()` 内容。

```shell
...
init() {
    ${g_docker_cmd} pull ${IMAGE} >& /dev/null
    if [ "${LISTEN}" ]; then
        g_start_args="${g_start_args} --listen=${LISTEN}"
    fi
    ...
    if [ "${EXT_PATH}" ]; then
        g_start_args="${g_start_args} --extended ext_path=/memcached/data${EXT_PATH}"
        volume_path=(${EXT_PATH//:/ })
        g_volume_bind="--volume ${volume_path}:/memcached/data${volume_path}"
    fi
    ...
}
...
```
#### 2. 导入主机列表

```bash
curveadm hosts commit hosts.yaml
```

### 第 4 步：部署 memcached 集群

以上述的 `hosts.yaml` 为例，部署一个 `memcached` 集群。
假设已经成功导入主机，执行一下命令部署 `memcached` 集群：

```bash
curveadm playbook -l memcached deploy.sh
```

成功部署的输出如下：


```shell
TOTAL: 3 hosts

--- server-host1 [SUCCESS]
[✔] create container [memcached-11211]
[✔] start container [memcached-11211]
[✔] wait 3 seconds, check container status...
CONTAINER ID        NAMES               STATUS
2836176cdfe8        memcached-11211     Up 3 seconds

--- server-host2 [SUCCESS]
[✔] create container [memcached-11211]
[✔] start container [memcached-11211]
[✔] wait 3 seconds, check container status...
CONTAINER ID        NAMES               STATUS
65308cc2f676        memcached-11211     Up 3 seconds

--- server-host3 [SUCCESS]
[✔] create container [memcached-11211]
[✔] start container [memcached-11211]
[✔] wait 3 seconds, check container status...
CONTAINER ID        NAMES               STATUS
9f5045852974        memcached-11211     Up 3 seconds
```

同时 `deploy.sh` 脚本还会检查容器名称 `memcached-${PORT}` 和端口号是否占用。
容器名称已占用认为已经创建会跳过创建返回仍是成功。
端口号通过 `lsof -i:${port}` 检查是否占用，如已占用则返回失败。

### 第 5 步：检查 memcached 集群状态

使用以下命令检查 `memcached` 的状态

```bash
curveadm playbook -l memcached status.sh
```

输出如下：

```bash
TOTAL: 3 hosts

--- server-host1 [SUCCESS]
CONTAINER ID        NAMES               STATUS
2836176cdfe8        memcached-11212     Up 2 minutes
memcached addr: 10.0.1.1:11211

--- server-host2 [SUCCESS]
CONTAINER ID        NAMES               STATUS
65308cc2f676        memcached-11211     Up 2 minutes
memcached addr: 10.0.1.2:11211

--- server-host3 [SUCCESS]
CONTAINER ID        NAMES               STATUS
9f5045852974        memcached-11211     Up 2 minutes
memcached addr: 10.0.1.3:11211
```

### 其他运维操作

#### 停止 memcached 集群

使用以下命令停止 `memcached` 集群:

```bash
curveadm playbook -l memcached stop.sh
```

输出如下：

```bash
TOTAL: 3 hosts

--- server-host1 [SUCCESS]
[✔] stop container[memcached-11211]

--- server-host2 [SUCCESS]
[✔] stop container[memcached-11211]

--- server-host3 [SUCCESS]
[✔] stop container[memcached-11211]
```

#### 清理 memcached 集群
---

使用以下命令停止 `memcached` 集群:

```bash
curveadm playbook -l memcached clean.sh
```

输出如下：

```bash
TOTAL: 3 hosts

--- server-host1 [SUCCESS]
[✔] rm container[memcached-11211]

--- server-host2 [SUCCESS]
[✔] rm container[memcached-11211]

--- server-host3 [SUCCESS]
[✔] rm container[memcached-11211]
```

# 添加 memcached 集群

在使用 memcached 集群前需要将memcached 集群添加到 mds，添加 memcached 集群的命令如下：

首先查看集群有哪些服务，服务状态是否正常：
```bash
curveadm status 

----------------------------------------------------------------
Get Service Status: [OK]  

cluster name      : **********
cluster kind      : curvefs
cluster mds addr  : **********,**********,**********
cluster mds leader: ********** / ddd28e44de66
……
```

然后进入某一个容器（请将后面的 id 更换为自己的Id）：
```bash
curveadm enter ddd28e44de66
```

在容器中执行以下命令，添加 memcached 集群到 mds（后面的地址根据自己的情况修改）：
```bash
curve fs create cachecluster  --servers 10.0.1.1:11211,10.0.1.2:11211,10.0.1.3:11211
```

添加成功后，可以通过以下命令查看 memcached 集群的信息：
```bash
curve fs list cachecluster
```

# CurveFS client 使用 memcached 集群

如果 CurveFS client 端需要使用 memcached 集群，需要在挂载时在 `client.yaml` 中指定配置 `fuseClient.supportKVcache: true`。
这样挂载时 client 会向 mds 申请分配一个 memcached 集群。
  
```yaml
# client.yaml
kind: curvefs
s3.ak: <>
s3.sk: <>
s3.endpoint: <>
s3.bucket_name: <>
container_image: opencurvedocker/curvefs:latest
mdsOpt.rpcRetryOpt.addrs: 10.0.1.1:6700,10.0.1.2:6700,10.0.1.3:6700
log_dir: /home/curve/curvebs/logs/client
data_dir: /data/curvefs
fuseClient.supportKVcache: true
```
