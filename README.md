# cdx - Continuous Deployment Tooling

![Github Actions](https://github.com/azunymous/cdx/workflows/Build/badge.svg?branch=master)

## Useful utilities for continuous integration and deployment pipelines
- Semantically versioning via Git tags
- Promoting versions of applications or modules
- Getting the latest version or the latest promoted version from Git tags

## Install binary release (Mac or Linux)
Download the [latest release](https://github.com/azunymous/cdx/releases/latest) 
binary for your platform. Make it executable (`chmod +x <binary name>`) and
move it somewhere on your `PATH`

## Install with Go
```shell script
GO111MODULE=on go get github.com/azunymous/cdx/cmd/cdx
```

## `cdx tag` commands

`cdx tag` allows you to manage your tags for versioning easily. It can be used for
both manual and automated releasing of new versions.

This is designed for multiple versioned applications in a single repository. The format of 
the tag is:

```
<module/app name>-<semantic version>
```

e.g `my-app-1.0.0`

`cdx` can be used to mark successful test runs or production candidates. 
This results in a tag like the following: 

```
<module/app name>-<semantic version>+<promotion-stage>
```

e.g `my-app-1.0.0+passed-extended-tests`

These are currently lightweight tags.

### Release - `cdx tag release -n <app name>` 
The `release` command increments the version of the provided application.

`cdx tag release -n my-app` on a repository with previously tagged
 version `my-app-0.1.0` for example will bump the version 
 from `my-app-0.1.0` to `my-app-0.2.0` for the currently checked out 
 commit.

The semantic version field to be bumped can be configured with `--increment` 
e.g 
```
cdx tag release -n my-app --increment major
```

### Promote - `cdx tag promote -n <app name> <promotion stage>`

This will promote the current commit to the provided stage for the provided
application/module.

e.g if the current commit is tagged `my-app-0.1.0`.

```
cdx tag promote -n my-app production
```
Will tag the checked out commit with `my-app-0.1.0+production`

### Latest - `cdx tag latest -n <app name> [promotion stage]`

The `latest` command returns the highest version of the application/module. 
If you provide a promotion stage, cdx returns the highest version of that module which has been
 promoted to that stage instead.
 
To get only tags of the current commit you can use the `--head` flag.

## Build 

You can build `cdx` with `go build ./cmd/cdx` at the root of this repository. 
You need Go installed with support for Go modules. 

## Notes
- cdx uses lightweight tags for versioning. Currently `latest` can detect annotated 
tags, but cdx does not support versioning with them.
- The `--push` flag delegates to `git push`, other commands use the go-git library and have no
dependencies
- `cdx tag` does not enforce ordered tagging. If you run `cdx tag release` from a branch or on an
untagged commit, it will search for the highest version across all references. The exception to this
is if the `--push` flag is used, which makes sure the current commit is on `origin/master`.
- Pre-releases are currently ignored. `my-app-1.0.0-rc1+passed-extended-tests` is ignored and
the latest version for the `passed-extended-tests` stage is `my-app-0.999.0+passed-extended-tests`
for example. The same applies for finding the latest tag in general: `0.1.0` vs `0.2.0-RC1`, 
`0.1.0` would be returned as the latest version. 