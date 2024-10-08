# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and is generated by [Changie](https://github.com/miniscruff/changie).


## [September 03, 2024]((https://github.com/OpsLevel/opslevel-jq-parser/compare/v2024.8.19...v2024.9.3))
### Dependency
* Bump github.com/opslevel/opslevel-go/v2024 from 2024.8.16 to 2024.9.3

## [August 19, 2024]((https://github.com/OpsLevel/opslevel-jq-parser/compare/v2024.4.26...v2024.8.19))
### Dependency
* Bump github.com/opslevel/opslevel-go/v2024 from 2024.4.26 to 2024.5.13
* Bump github.com/opslevel/opslevel-go/v2024 from 2024.5.13 to 2024.5.31
* bump opslevel-go version to v2024.8.1
* bump opslevel-jq-parser module go version to 1.22
* bump opslevel-go to v2024.8.16

## [April 26, 2024]((https://github.com/OpsLevel/opslevel-jq-parser/compare/v2024.3.15...v2024.4.26))
### Dependency
* Bump golang.org/x/net from 0.22.0 to 0.23.0
* bump opslevel-go version

## [March 15, 2024]((https://github.com/OpsLevel/opslevel-jq-parser/compare/v2024.2.26...v2024.3.15))
### Feature
* Enforce uniqueness on all fields in ServiceRegistration
* Ensure repositories have an alias and they are unique
### Refactor
* Properties are now returned as []opslevel.PropertyInput
### Bugfix
* have parsed jq expr return blank string instead of null

## [February 26, 2024]((https://github.com/OpsLevel/opslevel-jq-parser/compare/v2024.1.13...v2024.2.26))
### Feature
* Add generic Deduplicated function for slices and test cases
* Bump actions/cache from 3 to 4
* Add support for parsing property assignments
### Dependency
* Bump codecov/codecov-action from 3 to 4
* Bump github.com/rs/zerolog from 1.31.0 to 1.32.0
* Bump arduino/setup-task from 1 to 2
### Bugfix
* Fix config to use empty string

## [January 13, 2024]((https://github.com/OpsLevel/opslevel-jq-parser/compare/v2023.12.11...v2024.1.13))

## [December 11, 2023]((https://github.com/OpsLevel/opslevel-jq-parser/compare/v2023.11.2...v2023.12.11))
### Feature
* Add ability to parse a system

## [November 02, 2023]((https://github.com/OpsLevel/opslevel-jq-parser/compare/v0.0.0...v2023.11.2))
### Feature
* Initial Release
