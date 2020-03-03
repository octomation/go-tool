> # 🧩 Tool
>
> Template for typical Go tool.

[![Build][build.icon]][build.page]
[![Template][template.icon]][template.page]

## 💡 Idea

...

Full description of the idea is available [here][design.page].

## 🏆 Motivation

...

## 🤼‍♂️ How to

...

## 🧩 Installation

### Homebrew

```bash
$ brew install :owner/tap/:binary
```

### Binary

```bash
$ curl -sSfL https://:install.sh | sh
# or
$ wget -qO-  https://:install.sh | sh
```

### Source

```bash
# use standard go tools
$ go get github.com/:owner/:repository@:version
# or use egg tool
$ egg tools add github.com/:owner/:repository@:version
```

> [egg][]<sup id="anchor-egg">[1](#egg)</sup> is an `extended go get`.

### Bash and Zsh completions

```bash
$ :binary completion bash > /path/to/bash_completion.d/:binary.sh
$ :binary completion zsh  > /path/to/zsh-completions/_:binary.zsh
```

## 🤲 Outcomes

...

<sup id="egg">1</sup> The project is still in prototyping.[↩](#anchor-egg)

---

made with ❤️ for everyone

[build.page]:       https://travis-ci.org/:owner/:repository
[build.icon]:       https://travis-ci.org/:owner/:repository.svg?branch=master
[design.page]:      https://www.notion.so/33715348cc114ea79dd350a25d16e0b0?r=0b753cbf767346f5a6fd51194829a2f3
[promo.page]:       https://github.com/:owner/:repository
[template.page]:    https://github.com/octomation/go-tool
[template.icon]:    https://img.shields.io/badge/template-go--tool-blue

[egg]:              https://github.com/kamilsk/egg
