# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html)

## [1.15.0] - 2024-12-08

The `xmaps` package is a new extension of the Go standard library's `maps` package. It
was based on earlier work excluding features now available in Go's native implementation.

Introduced the `changelog` command to automate the creation of ChangeLog entries. This command
leverages conventional commit conventions from conventionalcommits.org to categorize changes by
type and scope into the changelog sections following the standards of keepachangelog.com.

We revert back to a more human-readable version of the ChangeLog; removing `CHANGELOG.yaml`.

### Added

- **cmd**: add changelog command to generate changelog entries
- **xmaps**:
    - add OrderedMap (from previous implementation)
    - encode/decode OrderedMap as JSON
    - explicitly set OrderedMap value type

### Changed

- **xt**: equality check function and error messages

## [1.14.0] - 2024-12-02

### Added

- **xslice**: Add `Exclude`

### Build

- Upgrade to Go 1.23, and update direct dependencies.

## [1.13.0] - 2024-07-02

### Added

- **xtime**: `xtime.FirstBeforeSecond` returns true if the first time is before the second 
  time; if either of the times is nil or zero value, it returns true.

### Fixed

- **xt**: Modify `ErrorIs` so that it also can deal with non-wrapped errors.

## [1.12.0] - 2024-06-25

### Added

- **xslice**:
    - `RemoveFirst` removes the first element from the given slice and returns the resulting
       slice; if the slice is empty, it returns an empty slice.
    - `RemoveFirstN` removes the first "n" elements from the given slice and returns the resulting slice.
      If the slice is empty or "n" is larger than the length of the slice, it returns an empty slice.

## [1.11.0] - 2024-06-17

### Added

- **xt**:
    - Add `xt.ErrorIs()` to check a wrapped error.
    - Add the possibility to turn off colors in output of `xt` using the environmental
      variable `XT_NO_COLORS`. This is mostly to make testing easier, but can be useful in
      other situations as well.

## [1.10.0] - 2024-04-26

### Added

- **xfs**: Add `xfs.IsDir` for checking whether the argument is a directory within `fs.FS`.
- **xptr**: Add `xptr.ValueOr` which returns the first non-nil value. If the first argument is 
  not nil, it is returned; otherwise, the second argument is returned.

## [1.9.0] - 2024-02-09

### Added

- **xstring**:
    - `ScanTokens` scans a string and returns a slice of tokens. Tokens are separated by
      whitespace, unless they are within quotes. If a token is within quotes, the whitespace
      is preserved within the token. Any character that is not whitespace or a quote is part
      of a token.

## [1.8.0] - 2024-01-24

### Added

- **xtime**:
    - Add `MiddayForDate` and `UTCMiddayForDate`. Both functions perform similar operations
      as `Midday` and `UTCMidday` respectively, but for a specific date. The date is passed
      as year, month, day, like Go's `time.Date()`.

## [1.7.1] - 2024-01-16

### Fixed

- **xt**: Make `ok()` a Go test helper function.

## [1.7.0] - 2023-12-11

### Added

- **xtime**:
    - Add `Midday` and `UTCMidday`. `Midday` returns the current local time adjusted to
      midday (12:00 PM), while `UTCMidday` returns the current time in the UTC timezone,
      adjusted to midday. Useful for creating a standard timestamp where the specific time
      is less significant than the date.

## [1.6.1] - 2023-11-24

### Fixed

- **xgrpc**: Trim prefixes from gRPC errors and use quoted messages.

## [1.6] - 2023-11-15

### Added

- **xgrpc**: Add helper functionalities around gRPC with functions `xgrpc.CheckServiceAvailability`
  and `xgrpc.ErrorFromRPC`.

## [1.5.1]

### Fixed

- **xt**: Return all entries of `LogAgg`.

## [1.5] - 2023-09-27

### Added

- **xt**: Add log aggregation helper.

## [1.4] - 2023-09-26

### Added

- Add package `xrand` for functionality returning random data.

## [1.3.1] - 2023-09-24

### Fixed

- **xreflect**: Implement patching using pointer values.

## [1.3] - 2023-09-20

### Added

- Add package `xreflect` offering handy tools for reflection such as `PatchStruct`.

## [1.2] - 2023-09-17

### Added

- Add package `xslice` extending the slice package with the function `AsAny`.

## [1.1] - 2023-09-16

### Added

- Add package `xptr` with helper function `Of` to create pointers from values.

## [1.0] - 2023-08-26

- Initial release.

### Earlier Versions

- v0.9 - 2023-08-14
