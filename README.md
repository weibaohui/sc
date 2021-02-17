[![GoDev](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)](https://pkg.go.dev/github.com/weibaohui/sc)
[![BuildStatus](https://github.com/weibaohui/sc/workflows/build/badge.svg)](https://github.com/weibaohui/sc/actions?workflow=build)
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

## binary user
```
Usage:
  sc [flags]

Flags:
  -d, --debug         调试
  -h, --help          help for sc
  -p, --path string   扫描路径 (default ".")
  -s, --silent        静默执行
```

# docker use

```docker
docker run -it --rm -v $(pwd):/code/  weibh/sc  -p /code/ -s
```

# 输出值

包含了git的用量统计，代码行数的统计

```json
{
  "git": {
    "Branch": 2,
    "Tags": 0,
    "Commit": {
      "git": 56,
      "main": 62
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
        "CommitCount": 163,
        "Addition": 9638,
        "Deletion": 7235
      }
    }
  },
  "source": {
    "FileTypeCounter": {
      "": {
        "Code": 26,
        "Blank": 8,
        "Comment": 0
      },
      ".go": {
        "Code": 2166,
        "Blank": 338,
        "Comment": 0
      },
      ".log": {
        "Code": 1,
        "Blank": 1,
        "Comment": 0
      },
      ".md": {
        "Code": 64,
        "Blank": 17,
        "Comment": 0
      },
      ".mod": {
        "Code": 8,
        "Blank": 4,
        "Comment": 0
      },
      ".sum": {
        "Code": 421,
        "Blank": 1,
        "Comment": 0
      },
      "Sum": {
        "Code": 2686,
        "Blank": 369,
        "Comment": 0
      }
    }
  }
}
```

# 说明

默认排除了隐藏文件及文件夹 使用魔法数识别二进制文件并排除

# todo

- docker √
- git ing

## Thanks

感谢 [JetBrains 公司](https://www.jetbrains.com/?from=sc) 为本开源项目提供的免费正版 Intellij GoLand 的 License 支持。
