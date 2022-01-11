快速上手 CurveAdm 开发
===

* [CurveAdm 整体设计](#curveadm-整体设计)
* [项目组织结构](#项目组织结构)
* [CurveAdm 的实现](#curveadm-的实现)
* [代码格式规范](#代码格式规范)

CurveAdm 整体设计
---

用户可通过以下文档了解 CurveAdm 的整体设计：

* [CurveAdm 简介](overview)

项目组织结构
---

CurveAdm 项目组织结构参照的是 [golang-standards/project-layout][project-layout]，各目录详情如下：

```shell
├── cli       # 命令行实现
│   ├── cli          # curveadm 对象，该对象贯穿整个命令执行
│   ├── command      # 各类命令的实现，如 deploy、start、stop 等
│   └── curveadm.go  # 命令执行入口
├── cmd       # 项目主干
│   └── curveadm     # main 函数
├── configs   # 配置文件模板 
│   ├── bs           # CurveBS 相关配置文件模板
│   └── fs           # CurveFS 相关配置文件模板
├── internal  # 私有库代码
│   ├── configure    # 各类配置文件的解析，如集群拓扑文件、客户端配置文件
│   ├── errors       # 错误码
│   ├── storage      # SQLite 数据库，如存储集群拓扑文件、服务对应容器 ID
│   ├── task         # 具体的任务实现，每一个任务包含多个步骤，内含一些共用的步骤抽象
│   ├── tools        # 各类工具实现，如更新 CurveAdm
│   ├── tui          # 终端输出实现
│   └── utils        # utils
├── pkg       # 公有库代码，可提供给外部程序使用 
│   ├── log          # 日志
│   ├── module       # SSH 客户端以及各类命令的封装抽象层，如 docker cli，mkdir, rmdir 等
│   └── variable     # 变量的实现
└── scripts   # 脚本
    └── install.sh   # 项目安装脚本
```

CurveAdm 的实现
---

CurveAdm 的实现主要分为以下 5 个部分：

| 模块      | 说明                                                         |
| :---      | :---                                                         |
| command   | 命令行的实现                                                 |
| configure | 解析配置文件                                                 |
| task      | 根据配置信息生成相应任务，每一个任务分为多个步骤             |
| tasks     | 执行任务，顺序执行任务每一步骤，其中一个步骤出错，则任务失败 |
| tui       | 根据任务结果输出相应信息到终端                               |

CurveAdm 所有命令实现都统一遵循以上 5 个部分流程， 用户可从一个简单命令实现切入，来了解整个执行流程，
我们推介从 [curveadm start][curveadm-start] 命令来了解。

代码格式规范
---

CurveAdm 代码遵守 golang 官方代码风格，使用 go 安装包中自带的 [go fmt][go-fmt] 来统一格式化，
希望用户提交代码时也遵守这一规范。

[go-fmt]: https://go.dev/blog/gofmt
[project-layout]: https://github.com/golang-standards/project-layout/blob/master/README_zh.md
[curveadm-start]: https://github.com/opencurve/curveadm/blob/master/cli/command/start.go#L41
