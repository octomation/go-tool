---
code: null
id: I_kwDODa89W85YShru
databaseId: 1481251566
number: 45
url: https://github.com/octomation/go-tool/issues/45
title: "dist: publish download script together with docs"
issueType: null
labels:
  - "type: feature"
  - "scope: inventory"
  - "impact: low"
  - "effort: easy"
assignees:
  - kamilsk
milestone: "[[milestone-2]]"
parent: null
issueFields: {}
state: OPEN
stateReason: null
duplicateOf: null
createdAt: 2022-12-07T07:48:24Z
updatedAt: 2023-04-29T12:20:41Z
lastEditedAt: null
closedAt: null
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

<!-- 2023-01-15T08:59Z https://github.com/octomation/go-tool/issues/45#issuecomment-1383093979
The correct way is to publish it to https://github.com/octomation/install.

Combine with
- octolab/homebrew-tap/issues/5
- kamilsk/homebrew-tap/issues/10
- https://github.com/tact-app/homebrew-tap
-->
