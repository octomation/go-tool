> # 🧩 Go tool
>
> Template for typical tool written on Go.

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
$ curl -fsSL https://install.octolab.org/octomation/tool | sh
# or
$ wget -qO-  https://install.octolab.org/octomation/tool | sh
```

> Don't forget about [security](https://www.idontplaydarts.com/2016/04/detecting-curl-pipe-bash-server-side/).

### Source

```bash
# use standard go tools
$ go get go.octolab.org/template/tool@latest
# or use egg tool
$ egg tools add go.octolab.org/template/tool@latest
```

> [egg][] is the `extended go get`.

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
[docs.page]:        https://pkg.go.dev/go.octolab.org/template/tool
[docs.icon]:        https://img.shields.io/badge/docs-pkg.go.dev-blue
[mirror.page]:      https://bitbucket.org/kamilsk/go-tool
[mirror.icon]:      https://img.shields.io/badge/mirror-bitbucket-blue
[promo.page]:       https://github.com/octomation/go-tool
[quality.page]:     https://goreportcard.com/report/go.octolab.org/template/tool
[quality.icon]:     https://goreportcard.com/badge/go.octolab.org/template/tool
[template.page]:    https://github.com/octomation/go-tool
[template.icon]:    https://img.shields.io/badge/template-go--tool-blue

[egg]:              https://github.com/kamilsk/egg
