# Changelog

All notable changes to **ValueStringBuilder** will be documented in this file. The project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

<!-- The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/) -->

## [Unreleased]

### Added

-   `SetRuneAt` is added to the string builder
-   `Skip` is added to the string builder

## [0.10.0] - 2023-08-31

### Added

-   `Reverse` is added to the string builder
-   `Substring` is added to the string builder

## [0.9.0] - 2023-08-26

### Added

-   `Trim`, `TrimStart`, `TrimEnd` are added to the string builder

## [0.8.0] - 2023-08-24

### Added

-   The string builder can now be used to append integers, boolean & collection of strings. `sb.AppendInt(1).AppendBool(true).AppendList([]string{"a", "b", "c"})`

## [0.7.0] - 2023-08-21

\### Added

-   The string builder can be used in chaining operations. `sb.Append("Hello").Append(" ").Append("World")`

## [0.6.1] - 2022-12-28

### Fixed

-   `Write` did not return the correct amount of added characters

## [0.6.0] - 2022-12-28

### Added

-   Implemented the `io.Writer` interface so it can be used with the `fmt` package

## [0.5.3] - 2022-12-28

\### Changed

-   Smaller internal refactorings

## [0.5.2] - 2022-12-27

### Fixed

-   Fixed a bug where `Append` and `AppendLine` could not handle umlauts

## [0.5.1] - 2022-12-27

### Fixed

-   Fixed a bug where find and friends and Replace did not work properly with umlauts or 2 (or more) byte character 

## [0.5.0] - 2022-12-26

### Added

-   Added the following methods: `ReplaceRune`, `Replace`

### Changed

-   Renamed `AsRune` to `AsRuneSlice` and included some more documentation to highlight that this is a shared memory block.

## [0.4.0] - 2022-12-26

### Added

-   Added the following methods: `FindFirst`, `FindLast`, `FindAll`

## [0.3.0] - 2022-12-25

### Added

-   Added the following methods: `RuneAt`, `AsRune`

## [0.2.0] - 2022-12-23

### Added

-   `Clear` methods

### Changed

-   Changed `NewFromString` to `NewStringBuilderFromString` to stay closer to standards

## [0.1.0] - 2022-12-21

This is the initial release for the `StringBuilder`.

### Added

-   `StringBuilder` with some major methods like (`Insert`, `Remove`, `Append`)

[Unreleased]: https://github.com/linkdotnet/golang-stringbuilder/compare/0.10.0...HEAD

[0.10.0]: https://github.com/linkdotnet/golang-stringbuilder/compare/0.9.0...0.10.0

[0.9.0]: https://github.com/linkdotnet/golang-stringbuilder/compare/0.8.0...0.9.0

[0.8.0]: https://github.com/linkdotnet/golang-stringbuilder/compare/0.7.0...0.8.0

[0.7.0]: https://github.com/linkdotnet/golang-stringbuilder/compare/0.6.1...0.7.0

[0.6.1]: https://github.com/linkdotnet/golang-stringbuilder/compare/0.6.0...0.6.1

[0.6.0]: https://github.com/linkdotnet/golang-stringbuilder/compare/0.5.3...0.6.0

[0.5.3]: https://github.com/linkdotnet/golang-stringbuilder/compare/0.5.2...0.5.3

[0.5.2]: https://github.com/linkdotnet/golang-stringbuilder/compare/0.5.1...0.5.2

[0.5.1]: https://github.com/linkdotnet/golang-stringbuilder/compare/0.5.0...0.5.1

[0.5.0]: https://github.com/linkdotnet/golang-stringbuilder/compare/0.4.0...0.5.0

[0.4.0]: https://github.com/linkdotnet/golang-stringbuilder/compare/0.3.0...0.4.0

[0.3.0]: https://github.com/linkdotnet/golang-stringbuilder/compare/0.2.0...0.3.0

[0.2.0]: https://github.com/linkdotnet/golang-stringbuilder/compare/0.1.0...0.2.0

[0.1.0]: https://github.com/linkdotnet/golang-stringbuilder/compare/12f8f67fb593ebe76a9794ea4a5362f6a8ae50d2...0.1.0
