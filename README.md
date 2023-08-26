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

Index
-----

* `xconv` provides functionality to convert things to other things (inspired by 
  standard library `strconv` but not limited to strings)
* `xnet` provides extra functionality around network I/O
* `xos` offers extra functionality around files and anything operating system related
* `xsql` provides extra functionality around SQL drivers including managing DSN (Data Source Name)
* `xstrings` extends `strings` with useful helpers
* `xt` offers wrappers around the `testing` standard package; the name was kept
  short because it used a lot
* `xtime` provides helpers around `time.Time`

License
-------

Distributed under the MIT license. See `LICENSE.txt` for more information.

[1]: https://pkg.go.dev/std