# gh-issue-projector

TODO

## Usage

    usage: gh-issue-projector TODO


## Example

    TODO

## Installation

intel-linux users:

    sudo /bin/sh -c ' wget https://github.com/alexcb/gh-issue-projector/releases/latest/download/gh-issue-projector-linux-amd64 -O /usr/local/bin/gh-issue-projector && chmod +x /usr/local/bin/gh-issue-projector'

raspberrypi-v4-linux users:

    sudo /bin/sh -c ' wget https://github.com/alexcb/gh-issue-projector/releases/latest/download/gh-issue-projector-linux-arm64 -O /usr/local/bin/gh-issue-projector && chmod +x /usr/local/bin/gh-issue-projector'

intel-mac users:

    sudo /bin/sh -c ' wget https://github.com/alexcb/gh-issue-projector/releases/latest/download/gh-issue-projector-darwin-amd64 -O /usr/local/bin/gh-issue-projector && chmod +x /usr/local/bin/gh-issue-projector'

m1/2-mac users:

    sudo /bin/sh -c ' wget https://github.com/alexcb/gh-issue-projector/releases/latest/download/gh-issue-projector-darwin-arm64 -O /usr/local/bin/gh-issue-projector && chmod +x /usr/local/bin/gh-issue-projector'

## Building

First download [earthly](https://github.com/earthly/earthly).

Then run:

    earthly +gh-issue-projector-all

builds are written to `build/<OS>/<arch>/gh-issue-projector` (where `OS` is either `linux` or `darwin` (MacOS), and `arch` is either `amd64` (intel-based) or `arm64` (M1, raspberry pi v4, etc))


## Licensing
gh-issue-projector is licensed under the Mozilla Public License Version 2.0. See [LICENSE](LICENSE).
