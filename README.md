## Thinko 框架初始化模板

Thinko 是一个轻量级 GO WEB 应用框架，提供一套结构化、模块化的开发环境，为减少开发学习成本，提高团队开发效率。

## 安装

#### 方式一

通过命令行去初始化项目,先安装命令行工具

```
git clone https://github.com/watsonhaw5566/thinko.git && cd think-core/cmd/thinko && go install
```

然后就可以在全局通过 ``thinko`` 命令去创建项目

```
think init demo-app
```

#### 方式二

也或者可以直接克隆项目使用

```
git clone https://github.com/watsonhaw5566/thinko-template.git
```

安装依赖

```
go mod tidy
```

启动项目

```
go run main.go
```