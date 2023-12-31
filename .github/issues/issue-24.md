---
code: null
id: I_kwDODa89W85NT8-l
databaseId: 1297076133
number: 24
url: https://github.com/octomation/go-tool/issues/24
title: "Taskfile: forks: migration style and bulletproof"
issueType: null
labels: []
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2022-07-07T09:14:51Z
updatedAt: 2023-04-29T12:15:14Z
lastEditedAt: null
closedAt: 2023-04-29T12:15:14Z
---

# Taskfile: forks: migration style and bulletproof

**Motivation:** current script has a lot of disadvantages.

- If something went wrong, I have to find this place manually and start from scratch or continue by copy-paste commands.
  - solution is to make it transactionable: `run continue`
- I cannot replace published release, because it tries to rebase from that point
  - see https://github.com/octomation/go-tool/blob/f59cf8b6efb783cc90ebf4752f3557afc76c4894/.github/fork/Taskfile#L6-L7
  - solution is to remove this tag if it present on origin
  - good example is `godownloader` problem when I replace it time after time

**Forks**

- https://github.com/kamilsk/mergestat
  - kamilsk/mergestat/issues/11
  - kamilsk/mergestat/issues/6
- https://github.com/kamilsk/godownloader
  - kamilsk/godownloader/issues/5
- https://github.com/kamilsk/go-tools
  - kamilsk/go-tools/issues/15
- https://github.com/kamilsk/golangci-lint
  - kamilsk/golangci-lint/issues/57
  - kamilsk/golangci-lint/issues/56

<!-- 2023-04-29T12:15Z https://github.com/octomation/go-tool/issues/24#issuecomment-1528775071
implemented, https://github.com/octomation/taskfiles
-->
