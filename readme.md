# Semantic Version Tags

## About

`semvertag` can read and write GiT tags formatted as a semantic version (MAJOR.MINOR.POINT) and can compute the next version for use in build tools.

It's built with `gopkg.in/src-d/go-git.v4` as a monolithic binary and offers cross-platform and dependency-free GiT access. All you need to do is select a release binary and add it to your repo!

## Usage

```
> ./semvertag [get|set] [latest|next|nextminor|nextmajor]
```

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