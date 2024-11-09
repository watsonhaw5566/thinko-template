<p align="center">
  <img width="300px" src="https://www.think-ts.cn/icon.png">
</p>

<p align="center">
  <a href="https://www.think-go.com.cn">
    <img src="https://img.shields.io/github/release/think-go/tg.svg?style=flat-square">
  </a>
  <a href="https://www.think-go.com.cn">
    <img src="https://pkg.go.dev/badge/github.com/think-go/tg?status.svg">
  </a>
  <a href="https://www.think-go.com.cn">
    <img src="https://codecov.io/gh/think-go/tg/branch/master/graph/badge.svg"/>
  </a>
  <a href="https://www.think-go.com.cn">
    <img src="https://img.shields.io/badge/%E4%BD%9C%E8%80%85-zhangyu-7AD6FD.svg"/>
  </a>
  <br>
</p>

<p align="center" style="font-weight:bold;font-size:20px;padding-top:5px;">一个轻量级的GO WEB应用框架</p>

- 💪 ORM思想链式操作CRUD
- 🔥 应用级提炼封装更贴近业务场景
- 🚀 高效路由管理，支持灵活的URL映射
- 🛠️ 自动化的代码生成工具，快速搭建项目基础结构

## ThinkGO框架

[ThinkGO](https://www.think-go.com.cn) 是一个轻量级的GO WEB应用框架，提供一套结构化、模块化的开发环境，为减少开发人员的学习成本，提高团队的开发效率而生。

## 目录结构

```
think-go
├── api                    // 请求返回及验证结构体定义
├── app
│   ├── controller         // 控制器调用service将结构呈现
│   │   └── hello.go
│   ├── dao                // 和数据库操作的逻辑都写这里
│   ├── entity             // 数据库表结构体及自定义数据库相关结构体
│   └── service            // 服务层编辑具体的逻辑代码，调用dao去和数据库交互
├── config                 // 框架配置文件
│   └── config.yaml
├── middleware             // 中间件目录
│   └── middleware.go
├── public                 // 静态资源目录
│   └── img
│       └── favicon.ico
├── router                 // 路由目录
│   └── router.go
├── view                   // 视图目录
│   ├── footer.html
│   └── index.html
├── main.go                // 入口文件
├── go.mod
├── go.sum
└── README.md
```

## 安装

#### 方式一

通过命令行去初始化项目,先安装命令行工具

```
git clone https://github.com/think-go/tg.git && cd tg/cmd/tg && go install
```

然后就可以在全局通过 ``tg`` 命令去创建项目

```
tg init demoApp
```

<p align="center">
  <img src="https://think-go.com.cn/tg.png">
</p>

#### 方式二

也或者可以直接克隆项目使用

```
git clone https://github.com/think-go/think-go.git
```

安装依赖

```
go mod tidy
```

启动项目

```
go run main.go
```

<p align="center">
  <img src="https://think-go.com.cn/think-go.png">
</p>


## 说明

``think-go`` 是基于 ``tg`` 核心包构建的基础工程项目，旨在为开发者提供一套结构化、模块化的开发环境。``think-go`` 精心设计了路由管理、中间件配置以及控制器实现等核心组件的组织方式与实现路径，确保了代码的高可读性和维护性。通过明确规定各功能模块的存放目录及实现方法，不仅简化了项目的搭建过程，还极大地方便了后续的迭代与扩展，使团队协作更加高效顺畅。无论是初学者还是有经验的开发者，都能在 ``think-go`` 的帮助下快速上手，专注于业务逻辑的实现，而无需从零开始搭建项目架构。
