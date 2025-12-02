---
id: 23
database_id: 1276052450
node_id: I_kwDODa89W85MDwPi
status: open
title: "build: compiled feature set"
labels: ["type: feature","scope: code","impact: medium","effort: medium"]
url: https://github.com/octomation/go-tool/issues/23
created_at: 2022-06-19T12:49:30Z
updated_at: 2023-04-29T12:18:34Z
---

# build: compiled feature set

**Motivation:** compile binary with predefined features.

See https://github.com/octomation/go-tool/blob/f59cf8b6efb783cc90ebf4752f3557afc76c4894/internal/config/features.go#L6-L11.

```bash
$ go build main.go -- --supported:feature=false
```

**Inspiration:** https://github.com/evanw/esbuild/releases/tag/v0.14.46, `--supported:feature=false`.
