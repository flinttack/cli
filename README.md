# Flinttack

Flinttack records the trust decision for each dependency in `flinttack.lock`, fails CI
when that state drifts, and analyzes your own source code for vulnerabilities, committed
secrets and unsafe infrastructure settings. Analysis runs locally; your source code is
never uploaded.

- **Dependency trust**: `lock` records who and what you trust, `verify` fails when it drifts.
- **Vulnerabilities**: `audit` checks dependencies against OSV and ranks findings by severity.
- **Source code**: `sast` finds vulnerabilities, secrets and unsafe infrastructure settings.
- **One report**: `security` combines all of the above; output as text, JSON, SARIF, JUnit
  or Markdown.

## Install

```sh
curl -fsSL https://flinttack.sh/install | sh
# a specific version
curl -fsSL https://flinttack.sh/install | sh -s -- --version v0.1.1
```

The installer supports macOS and Linux (amd64 and arm64), checks the SHA-256 of the
download, verifies its GitHub attestation, and installs to `/usr/local/bin` or, if that is
not writable, `~/.local/bin`. Set `INSTALL_DIR` to choose another directory.

Attestation verification needs the [GitHub CLI](https://cli.github.com) (`gh`), signed in
with `gh auth login`; without it the installer stops. Pass `--skip-attestation` to install
without verifying. Windows builds are published as zip archives on the
[releases page](https://github.com/flinttack/flinttack/releases).

Two executables work together:

- `flinttack`: the command-line interface (this repository).
- `flinttack-engine`: the source-code analysis engine that `flinttack` starts for
  `sast`, `security`, `chain` and related commands. It is released separately, from
  [flinttack/engine-releases](https://github.com/flinttack/engine-releases).

The installer runs `flinttack engine install` for you (`--skip-engine` leaves it out). That
command downloads the newest engine that speaks this `flinttack`'s protocol, checks its
sha256 and its GitHub attestation, and keeps it in the user cache. `flinttack engine update`
installs a newer one and removes the old ones, `flinttack engine list` shows the published
releases that fit, and `flinttack engine status` shows which engine is in use. An installed
version is not downloaded again unless you pass `--force`. Set `FLINTTACK_ENGINE_VERSION` to pin
a version. `flinttack doctor` reports an engine that is missing or that this `flinttack` cannot
use. The GitHub Action installs the engine for `sast` and `security`; its `engine-version`
input pins a release. Releases before the engine
was split out shipped it next to `flinttack`, and that copy wins over installed ones;
`flinttack engine update` installs the newest engine and removes that copy.

`flinttack` looks for the engine in this order: `FLINTTACK_ENGINE`, the directory of its own
executable, the engines installed by `flinttack engine install`, then `PATH`. Commands that
do not analyze source code, such as `scan`, `lock` and `verify`, work without the engine.

## Usage

```sh
flinttack scan .                 # dependency trust report
flinttack lock .                 # write flinttack.lock
flinttack verify .               # fail when the lock drifts
flinttack audit .                # dependency vulnerabilities by severity
flinttack sast .                 # find vulnerabilities in your source
flinttack security . --format sarif
```

Run `flinttack --help` for all commands, and `--offline` to skip remote lookups.
Exit code 0 means success, 1 a trust violation or finding, 2 an operational error.
[`examples/`](examples/) walks the lock, update and verify loop for each package manager.

### In CI

```yaml
- uses: flinttack/flinttack@v0.1.0
  with:
    command: security
```

A [pre-commit](https://pre-commit.com) hook is provided in
[`.pre-commit-hooks.yaml`](.pre-commit-hooks.yaml).

## Uninstall

```sh
flinttack uninstall
```

This removes the `flinttack` executable, the `flinttack-engine` next to it, the engines
downloaded by `flinttack engine install` and the response cache. `--dry-run` lists what would go, `--yes` skips the prompt, `--keep-cache`
keeps the cache. Project files such as `flinttack.lock` are never touched.

## Documentation

- [Trust model](docs/trust-model.md) and [threat model](docs/threat-model.md)
- [Lockfile format](docs/lockfile.md) and [specification](docs/specification.md)
- [Privacy](docs/privacy.md): what leaves your machine, and what does not
- [Enterprise](docs/enterprise.md)

## Development

```sh
go vet ./...
go test ./...
```

Building this repository gives you the CLI only. The engine lives in a separate,
private repository; tests that need it are skipped unless `FLINTTACK_ENGINE` points to a
`flinttack-engine` executable; `flinttack engine install` fetches a released one.
`scripts/package.sh <goos> <goarch> <version>` builds the CLI release archive.

Read [CONTRIBUTING.md](CONTRIBUTING.md) before opening a pull request, and
[SECURITY.md](SECURITY.md) to report a vulnerability.

## License

Apache-2.0. See [LICENSE](LICENSE).
