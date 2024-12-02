xgo - Extra Go
==============

Copyright (c) 2020, 2023, Geert JM Vanderkelen

Package `xgo` gathers extra, common functionality which gets reimplemented
in each project. It is organized in sub-packages which mimic the Go standard
library.

`xgo` is meanly used by other projects the author is maintaining and does nothing
extraordinary except taking away the tedious repeating.

The package grew from an old package, which was split into various separate
repositories within github.com/golistic. However, this is way too much overhead and
maintenance, so we decided to revert back to a single repository: `xgo`.

Requires Go 1.23 or later.

Index
-----

The following list shows sub-packages of `xgo`. Most have the same names as their
counterparts in the Go standard library, for example, `xos` and `os`. However, we
add some more like `xconv` and `xptr`.

* `xconv` - (basic) type conversions similar 
* `xnet` - from validating email addresses to finding te next free TCP port
* `xos` - wrapping around `os` with functions like `IsDir` or `IsRegularFile` and mapping environment
* `xptr` - getting pointer to value; probably the most reimplemented functionality 
* `xreflect` - handy tools doing reflection such as `PatchStruct`
* `xslice` - missing pieces of `slice`, with for example `AsAny` to return any slice as `[]any`
* `xsql` - extra functionality around SQL drivers including managing DSN (Data Source Name)
* `xstrings` - extends `strings` with useful helpers such as generic `Join` and `RepeatJoin`
* `xt` - basic wrappers around the `testing` standard package but with a short name
* `xtime` - helpers around `time.Time`

License
-------

Distributed under the MIT license. See `LICENSE.txt` for more information.

[1]: https://pkg.go.dev/std