# Flinttack Cargo example

Walk through the core loop on a tiny Cargo lockfile: lock, change a
dependency, inspect trust drift, then verify.

```bash
cd examples/cargo
flinttack lock --offline --no-fail
flinttack verify --offline
```

`--offline` does not query OSV. The default policy treats unchecked
vulnerabilities and malicious reports as `unknown`, so `lock` exits
`1` unless `--no-fail` is set. `flinttack.lock` in this directory is
that unknown snapshot for `cfg-if@1.0.0`.

Simulate a dependency change:

```bash
cp after/Cargo.lock Cargo.lock
flinttack diff --offline
flinttack verify --offline
```

`diff` reports the version and digest change. `verify` fails until you
accept the new tree:

```bash
flinttack lock --offline
flinttack verify --offline
```

Restore the original lockfile with `git checkout -- Cargo.lock`.
