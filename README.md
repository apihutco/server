<p align="center"><a href="https://apihut.co/" target="_blank"><img style="width: 250px;height: 250px" src="https://github.com/apihutco/docs/blob/main/docs/static/apihut.png?raw=true" alt="Logo"></a></p>


<p align="center">
<a href="https://github.com/apihutco/server" target="_blank"><img src="https://img.shields.io/badge/version-v2.0-brightgreen" alt="Version"></a>
<a href="https://github.com/apihutco/server" target="_blank"><img src="https://img.shields.io/github/go-mod/go-version/apihutco/server?style=flat&logo=go" alt="Go Version"></a>
<a href="https://github.com/apihutco/server/blob/main/LICENSE" target="_blank"><img src="https://img.shields.io/github/license/apihutco/server" alt="License"></a>
<a href="https://goreportcard.com/report/github.com/apihutco/server"><img src="https://goreportcard.com/badge/github.com/apihutco/server" alt="Go Report Card"></a>
<a href="https://docs.apihut.co" target="_blank"><img src="https://img.shields.io/badge/docs-current-4c6ef5?logo=vercel&logoWidth=10" alt="Docs"></a>
</p>

# APIHut

这是一个极简的接口聚合中心，旨在为个人项目提供一个通用的开放数据来源。

- [支持的接口](https://docs.apihut.co/)

<!--Start reconstruct: `2022.09.26`-->

## 构建

### Command

```bash
# 运行
make dev
# 构建
make build
```

### Docker

```bash
docker build -t apihutco/server:latest .
```

## CICD

`CICD` 目前使用 [Gitea](https://github.com/go-gitea/gitea) + [Argo Workflow](https://github.com/argoproj/argo-workflows/) + [Argo CD](https://github.com/argoproj/argo-cd) + [Kubernetes](https://github.com/kubernetes/kubernetes) 的工作流，通过 `webhooks` 与 模板 触发自动化的 GitOps 。相关的配置文件位于独立的 [deploy](https://github.com/apihut/deploy) 项目中。