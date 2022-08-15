主机管理
===

模块简介
---

主机模块用来统一管理用户主机，以减少用户在各配置文件中重复填写主机 `SSH` 连接相关配置。

主机配置
---

### 示例

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

### 主机配置项

| 配置项           | 是否必填 | 默认值                  | 说明                                                                                                                                                                                           |
| :---             | :---     | :---                    | :---                                                                                                                                                                                           |
| host             | Y        |                         | 主机名                                                                                                                                                                                         |
| hostname         | Y        |                         | 主机地址                                                                                                                                                                                       |
| user             |          | $USER                   | 连接远端主机 SSH 服务的用户。若用户没有指定 `become_user`，则以 `user` 指定用户为执行用户， 用于执行部署操作。请确保执行用户拥有 sudo 权限，因为其将用于挂卸载文件系统、操作 docker cli 等操作 |
| ssh_port         |          | 22                      | 远端主机 SSH 服务端口                                                                                                                                                                          |
| private_key_file |          | /home/$USER/.ssh/id_rsa | 用于连接远端主机 SSH 服务的私钥路径                                                                                                                                                            |
| forward_agent    |          | false                   | 是否以 SSH 代理转发方式登录远程主机                                                                                                                                                            |
| become_user      |          |                         | 执行部署操作的用户。若用户没有指定 `become_user`，则以 `user` 指定用户为执行用户， 用于执行部署操作。请确保执行用户拥有 sudo 权限，因为其将用于挂卸载文件系统、操作 docker cli 等操作          |
| labels           |          |                         | 主机标签。单个主机可指定多个标签                                                                                                                                                               |

管理主机
---

* [导入主机列表](#导入主机列表)
* [查看主机列表](#查看主机列表)
* [显示主机配置](#显示主机配置)
* [登录主机](#登录主机)

### 导入主机列表

#### 第 1 步：准备主机列表

```
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
    forward_agent: true
    become_user: admin
    labels:
      - client
```

#### 第 2 步：导入主机列表

```shell
$ curveadm hosts commit hosts.yaml
```

### 查看主机列表

```shell
$ curveadm hosts ls
```

CurveAdm 默认会显示所有主机：

```
Host          Hostname  User   Port  Private Key File         Forward Agent  Become User  Labels
----          --------  ----   ----  ----------------         -------------  -----------  ------
server-host1  10.0.1.1  curve  22    /home/curve/.ssh/id_rsa  N              -            -
server-host2  10.0.1.2  curve  22    /home/curve/.ssh/id_rsa  N              -            -
server-host3  10.0.1.3  curve  22    /home/curve/.ssh/id_rsa  N              -            -
client-host   10.0.1.4  curve  22    /home/curve/.ssh/id_rsa  Y              admin        client
```

* 若想查看指定标签的主机列表，可指定 `-l` 参数。

### 显示主机配置

```shell
$ curveadm hosts show
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
    forward_agent: true
    become_user: admin
    labels:
      - client
```

### 登录主机

```shell
$ curveadm ssh <host>
```

#### 示例：登录 `server1` 主机

```
$ curveadm ssh server1

```

> :bulb: **提醒：**
>
> 当出现 SSH 连接问题时，你可以根据主机的相应配置，
> 在本地手动模拟 curveadm 的连接操作，以此来排查相应问题：
> ```shell
> $ ssh <user>@<hostname> -p <ssh_port> -i <private_key_file>
> ```
