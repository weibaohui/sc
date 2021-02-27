[![GoDev](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)](https://pkg.go.dev/github.com/weibaohui/sc)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/weibaohui/sc)
[![BuildStatus](https://github.com/weibaohui/sc/workflows/build/badge.svg)](https://github.com/weibaohui/sc/actions?workflow=build)
[![Go Report Card](https://goreportcard.com/badge/github.com/weibaohui/sc)](https://goreportcard.com/report/github.com/weibaohui/sc)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/weibaohui/sc)](https://www.tickgit.com/browse?repo=github.com/weibaohui/sc)
[![codecov](https://codecov.io/gh/weibaohui/sc/branch/master/graph/badge.svg)](https://codecov.io/gh/weibaohui/sc)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/weibaohui/sc?sort=semver)
[![Stargazers over time](https://starchart.cc/weibaohui/sc.svg)](https://starchart.cc/weibaohui/sc)

# 简介

统计源码行数

# 编译

```
go build 
```

# 安装

```
go get -u github.com/weibaohui/sc
```

# 基本用法

## binary user
```
Usage:
  sc [flags]

Flags:
  -d, --debug               调试
      --exclude string      跳过文件夹列表,使用逗号分割
      --force               使用自定义配置覆盖默认初始配置，否则合并
  -h, --help                help for sc
  -p, --path string         扫描路径 (default ".")
      --skipSuffix string   跳过文件后缀列表,使用逗号分割
```

排除x文件夹，跳过后缀为.x .y .z 三种后缀

```shell
sc -d --skipSuffix ".x,.y,.z" --exclude "x"
```

# docker use

docker -v 挂载待扫描目录到容器里面 sc -p 扫描指定目录

```docker
docker run -it --rm -v $(pwd):/code/  weibh/sc  -p /code/ 
```

# 输出值

包含了git的用量统计，代码行数的统计

```json
{
  "git": {
    "Branch": 1,
    "Tags": 7,
    "Commit": {
      "main": 104
    },
    "AuthorCounts": {
      "weibaohui@chinamobile.com": {
        "Email": "weibaohui@chinamobile.com",
        "Name": "weibh",
        "CommitCount": 61,
        "Addition": 1442,
        "Deletion": 559
      },
      "weibaohui@yeah.net": {
        "Email": "weibaohui@yeah.net",
        "Name": "weibaohui",
        "CommitCount": 246,
        "Addition": 11242,
        "Deletion": 8613
      }
    },
    "CurrentBranch": "main"
  },
  "source": {
    "FileTypeCounter": {
      "": {
        "Code": 28,
        "Blank": 7,
        "Comment": 0
      },
      ".go": {
        "Code": 2498,
        "Blank": 251,
        "Comment": 0
      },
      ".md": {
        "Code": 116,
        "Blank": 19,
        "Comment": 0
      },
      ".sum": {
        "Code": 366,
        "Blank": 1,
        "Comment": 0
      },
      "Sum": {
        "Code": 3008,
        "Blank": 278,
        "Comment": 0
      }
    }
  }
}

```

# 说明

默认排除了隐藏文件及文件夹 使用魔法数识别二进制文件并排除



## Thanks

感谢 [JetBrains 公司](https://www.jetbrains.com/?from=sc) 为本开源项目提供的免费正版 Intellij GoLand 的 License 支持。
