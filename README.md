# juvo-cnb

## How to build

### Builder

1. Install docker, kubectl, and [pack cli](https://buildpacks.io/docs/for-platform-operators/how-to/integrate-ci/pack/)
1. Run:
    ```bash
    cd cnb-components/builders/juvo-builder
    ./make-builder v{version}
    ```

### Buildpack

1. Install [go](https://go.dev/doc/install)
1. Run
   ```
   cd cnb-components/buildpacks/juvo-poetry-buildpack
   ./build.sh
   ```

## Maintainers

* Diego Balseiro <dbalseiro@stackbuilders.com>
* Brad James <bwjames@hollandhart.com>
