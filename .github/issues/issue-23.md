---
code: null
id: I_kwDODa89W85MDwPi
databaseId: 1276052450
number: 23
url: https://github.com/octomation/go-tool/issues/23
title: "build: compiled feature set"
issueType: null
labels:
  - "type: feature"
  - "scope: code"
  - "impact: medium"
  - "effort: medium"
assignees:
  - kamilsk
milestone: "[[milestone-2]]"
parent: null
issueFields: {}
state: OPEN
stateReason: null
duplicateOf: null
createdAt: 2022-06-19T12:49:30Z
updatedAt: 2023-04-29T12:18:34Z
lastEditedAt: null
closedAt: null
---

# build: compiled feature set

**Motivation:** compile binary with predefined features.

See https://github.com/octomation/go-tool/blob/f59cf8b6efb783cc90ebf4752f3557afc76c4894/internal/config/features.go#L6-L11.

```bash
$ go build main.go -- --supported:feature=false
```

**Inspiration:** https://github.com/evanw/esbuild/releases/tag/v0.14.46, `--supported:feature=false`.
