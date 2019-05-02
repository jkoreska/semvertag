# Semantic Version Tags

## About

`semvertag` can read and write GiT tags formatted as a semantic version (MAJOR.MINOR.PATCH) and can compute the next version for use in build tools.

It's built with `gopkg.in/src-d/go-git.v4` as a monolithic binary and offers cross-platform and dependency-free GiT access.

All you need to do is select a [release binary](releases) and add it to your repo or build system!

Currently, prefixes are not supported ([here's why](https://github.com/semver/semver/issues/1)).

## Usage

```
> ./semvertag [get|set] [latest|next|nextminor|nextmajor]
```

Run `semvertag get next` when starting a build to set the version, then `semvertag set next` and `git push --tags origin` if the build is successful.

Note: `semvertag` expects a `.git` directory in the current working directory.

### Examples

```
> ./semvertag get latest
0.1.0
```
```
> ./semvertag get next
0.1.1
```
```
> ./semvertag get nextmajor
1.0.0
```
```
> ./semvertag set next
0.1.1
> ./semvertag get latest
0.1.1
```