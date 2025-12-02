---
id: 45
database_id: 1481251566
node_id: I_kwDODa89W85YShru
status: open
title: "dist: publish download script together with docs"
labels: ["type: feature","scope: inventory","impact: low","effort: easy"]
url: https://github.com/octomation/go-tool/issues/45
created_at: 2022-12-07T07:48:24Z
updated_at: 2023-04-29T12:20:41Z
---

# dist: publish download script together with docs

**Motivation:** simplicity.

Details are here https://github.com/evanw/esbuild/releases/tag/v0.16.0.

> Publish a shell script that downloads esbuild directly
>
> In addition to all of the existing ways to install esbuild, you can now also download esbuild directly like this:
>
> `curl -fsSL https://esbuild.github.io/dl/latest | sh`

See https://github.com/esbuild/esbuild.github.io/tree/gh-pages as an example.

**Result**

```bash
$ curl -fsSL https://octomation.github.io/go-tool/dl/latest | sh
# instead of
$ curl -fsSL https://raw.githubusercontent.com/octomation/go-tool/main/bin/install | sh
```
