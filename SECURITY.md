# Security Policy

## Reporting a vulnerability

Please report security issues privately through GitHub Security Advisories
on [flinttack/flinttack](https://github.com/flinttack/flinttack/security/advisories/new).

Do not open a public issue for an unreleased vulnerability.

We aim to acknowledge reports within 5 business days.

Flinttack may send public dependency names and versions to OSV and
the npm registry. Private registries are not queried in the default
`public-only` network mode. See [Privacy](docs/privacy.md).

## Supported versions

Only the latest tagged release is supported until v1.0.

## What Flinttack verifies about itself

Release artifacts are published with:

- `checksums.txt`
- SPDX SBOMs
- GitHub artifact attestations
- Immutable releases (tag and assets locked after publish)

Consumers can verify either the release archive or the extracted binary:

```bash
gh attestation verify flinttack_darwin_arm64.tar.gz --repo flinttack/flinttack
gh attestation verify ./flinttack --repo flinttack/flinttack
```

`scripts/install.sh` checks the published SHA-256 before install. If `gh`
is available, it also verifies the GitHub attestation.
