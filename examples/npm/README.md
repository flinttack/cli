# Flinttack npm example

Walk through the core loop on a tiny npm lockfile: lock, change a
dependency, inspect trust drift, then verify.

```bash
cd examples/npm
flinttack lock --offline --no-fail
flinttack verify --offline
```

`--offline` does not query OSV or the registry. The default policy
treats unchecked vulnerabilities and malicious reports as `unknown`,
so `lock` exits `1` unless `--no-fail` is set. `flinttack.lock` in this
directory is that unknown snapshot for `ms@2.1.3`.

Simulate a dependency change:

```bash
cp after/package-lock.json package-lock.json
flinttack diff --offline
flinttack verify --offline
```

`diff` reports the version and digest change. `verify` fails until you
accept the new tree:

```bash
flinttack lock --offline
flinttack verify --offline
```

Restore the original lockfile with `git checkout -- package-lock.json`.
