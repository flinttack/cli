# Flinttack CLI

Flinttack records the trust decision for each dependency in `flinttack.lock`, fails CI
when that state drifts, and analyzes your own source code for vulnerabilities, committed
secrets and unsafe infrastructure settings. Analysis runs locally; your source code is
never uploaded.

## Install

```sh
curl -fsSL https://raw.githubusercontent.com/flinttack/cli/main/scripts/install.sh | sh
```

The installer puts two executables next to each other:

- `flinttack`: the command-line interface (this repository).
- `flinttack-engine`: the source-code analysis engine that `flinttack` starts for
  `sast`, `security`, `chain` and related commands.

`flinttack` looks for the engine in the directory of its own executable, then on `PATH`.
Set `FLINTTACK_ENGINE` to use an engine at another location.

## Usage

```sh
flinttack scan .                 # dependency trust report
flinttack lock .                 # write flinttack.lock
flinttack verify .               # fail when the lock drifts
flinttack sast .                 # find vulnerabilities in your source
flinttack security . --format sarif
```

Run `flinttack --help` for all commands. Documentation is in [`docs/`](docs/).

## Development

```sh
go vet ./...
go test ./...
```

Tests that need the analysis engine are skipped unless `FLINTTACK_ENGINE` points to a
`flinttack-engine` executable.

## License

Apache-2.0. See [LICENSE](LICENSE).
