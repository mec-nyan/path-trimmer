# Path trimmer

![Example](./example.png)

Selectively reduce the length of the path string.

## TL;DR

Sometimes you work on deeply nested directories and the path to your _pwd_ becomes really long... This command will "trim" it a little bit to make it more suitable for display i.e. in your command line prompt.

## Get it

```sh
$ go get github.com/mec-nyan/path-trimmer@latest
```

## Use it

The command syntax is simple: `path-trimmer path [n]` where `path` is the path to trim and `[n]` is an optional argument indicating how many directory names (from the current and backwards) to keep intact (or _not to trim_, the default is 1 (don't shorten the name of the current working directory)).

For the rest of the path, if the names a longer than 3 chars it will shorten them and append an ellipsis to them (Unicode ellipsis, not three dots).

For example:

```sh
$ path-trimmer foo/bar
foo/bar
$ path-trimmer some_long_name/my_dir
som…/my_dir
$ path-trimmer a/deeply/nested/directory
a/dee…/nes…/directory
```

You can use it on your **Bash** prompt instead of `\w` or `\W`:

```sh
export PS1=`\u@\h $(path-trimmer $(pwd)) \$ `
```
