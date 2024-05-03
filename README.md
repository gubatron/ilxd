[![Go](https://github.com/project-illium/ilxd/actions/workflows/go.yml/badge.svg)](https://github.com/project-illium/ilxd/actions/workflows/go.yml)
[![golangci-lint](https://github.com/project-illium/ilxd/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/project-illium/ilxd/actions/workflows/golangci-lint.yml)

<h1 align="center">
<img src="https://raw.githubusercontent.com/project-illium/faucet/master/static/logo-white.png" alt="Illium logo" title="Illium logo">
</h1>

# ilxd
illium full node implementation written in Go

This is an alpha version of the illium full node software. It very recently has become feature complete and can now
create and verify zk-snark proofs. We plan more testing with this version before moving to a formal beta. Also note the 
current performance (specifically the proving time) is not likely representative of the final performance.

If you want to test this alpha version you can download the binaries from the github releases page and run the node with
the `--alpha` flag.

```go
$ ilxd --alpha
```

### Install
Head over to the [releases](https://github.com/project-illium/ilxd/releases) page and download the lastest release for
your operating system. 

The release contains two binaries: `ilxd` and `ilxcli`. `ilxd` is the illium full node application and `ilxcli` is a 
command line application that is used to control and interact with a running node.

### Build From Source
Please note that the master branch is considered under active development and may contain bugs. If you are running in
a production environment please checkout a release tag.

Make sure you have the required dependencies:
```go
$ apt-get install build-essential pkg-config libssl-dev
```

Ilxd requires both Go and Rust to be installed on your system. You'll need to use the makefile to install it as it will
compile both the Go code and Rust bindings.

```
$ git clone https://github.com/project-illium/ilxd.git
$ cd ilxd
$ make install
```
This command builds both `ilxd` and `ilxcli`. The binaries will be put in `$GOPATH/bin`.

To put this directory in your path add this line to your `/etc/profile` (for a system-wide installation) or `$HOME/.profile`:

```
export PATH=$PATH:$GOPATH/bin
```

### GPU Acceleration
If you have an Nvidia GPU with the `nvidia-cuda-toolkit` installed you can make use of GPU acceleration to speed up the proving
and verifying time.

```
$ CUDA=1 make install
```

### Running with Docker
Using docker run:
```
$ docker run -d -p 9002:9002 -p 5001:5001 projectillium/ilxd:v0.0.11-alpha --alpha
```

Using docker-compose:
```
$ docker-compose up
```

### Usage
Vist [docs.illium.org](https://docs.illium.org/node) for a comprehensive guide to running a node.

### Contributing
We'd love your help! See the [contributing guidlines](https://github.com/project-illium/ilxd/blob/master/CONTRIBUTING.md) before submitting your first PR.
