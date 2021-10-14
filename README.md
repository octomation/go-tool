> # 🧩 Tool
>
> Template for typical Go tool.

[![Build][build.icon]][build.page]
[![Documentation][docs.icon]][docs.page]
[![Quality][quality.icon]][quality.page]
[![Template][template.icon]][template.page]
[![Coverage][coverage.icon]][coverage.page]
[![Mirror][mirror.icon]][mirror.page]

## 💡 Idea

```bash
$ tool do action
```

A full description of the idea is available [here][design.page].

## 🏆 Motivation

...

## 🤼‍♂️ How to

...

## 🧩 Installation

### Homebrew

```bash
$ brew install octolab/tap/tool
```

### Binary

```bash
$ curl -fsSL https://raw.githubusercontent.com/octomation/go-tool/main/bin/install | sh
# or
$ wget -qO-  https://raw.githubusercontent.com/octomation/go-tool/main/bin/install | sh
```

> Don't forget about [security](https://www.idontplaydarts.com/2016/04/detecting-curl-pipe-bash-server-side/).

### Source

```bash
# use standard go tools
$ go get github.com/octomation/go-tool@latest
# or use egg tool
$ egg tools add github.com/octomation/go-tool@latest
```

> [egg][] is an `extended go get`.

### Shell completions

```bash
$ tool completion > /path/to/completions/...
# or
$ source <(tool completion)
```

## 🤲 Outcomes

...

<p align="right">made with ❤️ for everyone</p>

[awesome.icon]:     https://awesome.re/mentioned-badge.svg
[build.page]:       https://github.com/octomation/go-tool/actions/workflows/ci.yml
[build.icon]:       https://github.com/octomation/go-tool/actions/workflows/ci.yml/badge.svg
[coverage.page]:    https://codeclimate.com/github/octomation/go-tool/test_coverage
[coverage.icon]:    https://api.codeclimate.com/v1/badges/8491ba0aada439d2df0c/test_coverage
[design.page]:      https://www.notion.so/33715348cc114ea79dd350a25d16e0b0
[docs.page]:        https://pkg.go.dev/go.octolab.org
[docs.icon]:        https://img.shields.io/badge/docs-pkg.go.dev-blue
[mirror.page]:      https://bitbucket.org/kamilsk/go-tool
[mirror.icon]:      https://img.shields.io/badge/mirror-bitbucket-blue
[promo.page]:       https://github.com/octomation/go-tool
[quality.page]:     https://goreportcard.com/report/go.octolab.org
[quality.icon]:     https://goreportcard.com/badge/go.octolab.org
[template.page]:    https://github.com/octomation/go-tool
[template.icon]:    https://img.shields.io/badge/template-go--tool-blue

[egg]:              https://github.com/kamilsk/egg
