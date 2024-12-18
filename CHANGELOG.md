# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

#### binary-codec

- Updated `definitions`.
- New `DecodeLedgerData` function.
- `Quality` encoding/decoding functions.
- New `XChainBridge` and `Issue` types.

#### address-codec

- Address validation with `IsValidAddress`, `IsValidClassicAddress` and `IsValidXAddress`.
- Address conversion with `XAddressToClassicAddress` and `ClassicAddressToXAddress`.
- X-Address encoding/decoding with `EncodeXAddress` and `DecodeXAddress`.

#### keypairs

#### xrpl


### Changed


#### binary-codec

- Exported `FieldInstance` type.
- Updated `NewBinaryParser` constructor to accept `definitions.Definitions` as a parameter.
- Updated `NewSerializer` to `NewBinarySerializer` constructor.
- Refactored `FieldIDCodec` to be a struct with `Encode` and `Decode` methods.
- `FromJson` methods to `FromJSON`
- `ToJson` methods to `ToJSON`

#### address-codec



#### keypairs

#### xrpl

### Deprecated
- Method `oldFunction()` will be removed in next major version
- Configuration option `deprecatedSetting` is no longer recommended

### Removed
- Dropped support for legacy platform Y
- Eliminated unused utility functions

### Fixed
- Resolved critical bug in data processing module
- Corrected rendering issue on mobile devices
- Fixed memory leak in background service

### Security
- Patched vulnerability in authentication mechanism
- Updated encryption libraries to address potential security risks
