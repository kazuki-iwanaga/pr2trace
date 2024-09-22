> [!CAUTION]
> This repository is not ready.

<div align="center">
    <picture>
        <source media="(prefers-color-scheme: dark)" srcset="./docs/logo-dark-mode.svg">
        <source media="(prefers-color-scheme: light)" srcset="./docs/logo-white-mode.svg">
        <img alt="pr2trace logo" src="./docs/logo-with-background.svg" width="256">
    </picture>
</div>

<div align="center">
    <h1>pr2trace</h1>
</div>

Convert a GitHub Pull Request to OpenTelemetry-compatible telemetry.

# Install

## homebrew
```
brew install kazuki-iwanaga/tap/pr2trace
```

## go install
```
go install github.com/kazuki-iwanaga/pr2trace@latest
```

## GitHub releases
```
# e.g. Install v0.0.5 for Linux(x86_64)
wget https://github.com/kazuki-iwanaga/pr2trace/releases/download/v0.0.5/pr2trace_Linux_x86_64.tar.gz
tar xvf pr2trace_Linux_x86_64.tar.gz pr2trace
# sudo mv pr2trace /usr/local/bin/pr2trace
```