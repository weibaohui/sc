[![GoDev](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)](https://pkg.go.dev/github.com/weibaohui/sc)
[![BuildStatus](https://github.com/weibaohui/sc/workflows/tests/badge.svg)](https://github.com/weibaohui/sc/actions?workflow=tests)
[![Go Report Card](https://goreportcard.com/badge/github.com/weibaohui/sc)](https://goreportcard.com/report/github.com/weibaohui/sc)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/weibaohui/sc)](https://www.tickgit.com/browse?repo=github.com/weibaohui/sc)
[![codecov](https://codecov.io/gh/weibaohui/sc/branch/master/graph/badge.svg)](https://codecov.io/gh/weibaohui/sc)

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

```
Usage:
sc [flags]

Flags:
-d, --debug         调试
-h, --help          help for sc
-p, --path string   扫描路径 (default ".")
```

# 输出值

```json
{
  "": {
    "Code": 10,
    "Blank": 4,
    "Comment": 0
  },
  ".go": {
    "Code": 341,
    "Blank": 52,
    "Comment": 0
  },
  ".md": {
    "Code": 15,
    "Blank": 6,
    "Comment": 0
  },
  ".mod": {
    "Code": 1,
    "Blank": 1,
    "Comment": 0
  },
  ".sum": {
    "Code": 286,
    "Blank": 1,
    "Comment": 0
  },
  "sum": {
    "Code": 653,
    "Blank": 64,
    "Comment": 0
  }
}

```

# 说明

默认排除了隐藏文件及文件夹 使用魔法数识别二进制文件并排除

# todo

- docker √
- git ing
