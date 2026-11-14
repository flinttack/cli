# Flinttack Cabal example

Walk through the core loop on a tiny `cabal.project.freeze`: lock, change a
dependency, inspect trust drift, then verify.

```bash
cd examples/cabal
flinttack lock --offline --no-fail
flinttack verify --offline
```

`--offline` does not query OSV. A freeze file carries no hashes, so no
digest is recorded. The default policy treats unchecked vulnerabilities and
malicious reports as `unknown`, so `lock` exits `1` unless `--no-fail` is set.
`flinttack.lock` in this directory is that unknown snapshot for
`aeson@2.2.1.0`.

Simulate a dependency change:

```bash
cp after/cabal.project.freeze cabal.project.freeze
flinttack diff --offline
flinttack verify --offline
```

`diff` reports the version change. `verify` fails until you
accept the new tree:

```bash
flinttack lock --offline
flinttack verify --offline
```

Restore the original lockfile with `git checkout -- cabal.project.freeze`.
