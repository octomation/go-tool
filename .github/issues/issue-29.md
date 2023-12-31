---
code: null
id: I_kwDODa89W85PXkzo
databaseId: 1331580136
number: 29
url: https://github.com/octomation/go-tool/issues/29
title: "go: extended support of go version since 1.18"
issueType: null
labels:
  - "type: feature"
  - "scope: code"
  - "impact: medium"
  - "effort: easy"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2022-08-08T09:20:38Z
updatedAt: 2023-09-13T10:25:41Z
lastEditedAt: null
closedAt: 2023-09-13T10:25:40Z
---

# go: extended support of go version since 1.18

**Motivation:** extract more information for "versioning".

https://github.com/octomation/go-tool/blob/87401abb06ada0ee8b388b98945f1e88f35c74bf/main.go#L30-L35

See https://go.dev/doc/go1.18#go-command,

> ### go version
>
> The go command now embeds version control information in binaries. It includes the currently checked-out revision, commit time, and a flag indicating whether edited or untracked files are present. Version control information is embedded if the go command is invoked in a directory within a Git, Mercurial, Fossil, or Bazaar repository, and the main package and its containing main module are in the same repository. This information may be omitted using the flag -buildvcs=false.
>
> Additionally, the go command embeds information about the build, including build and tool tags (set with -tags), compiler, assembler, and linker flags (like -gcflags), whether cgo was enabled, and if it was, the values of the cgo environment variables (like CGO_CFLAGS). Both VCS and build information may be read together with module information using go version -m file or runtime/debug.ReadBuildInfo (for the currently running binary) or the new [debug/buildinfo](https://go.dev/doc/go1.18#debug/buildinfo) package.
>
> The underlying data format of the embedded build information can change with new go releases, so an older version of go may not handle the build information produced with a newer version of go. To read the version information from a binary built with go 1.18, use the go version command and the debug/buildinfo package from go 1.18+.
