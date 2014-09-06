# vagrant-global-status

A fast vagrant global-status command.

## Description

`vagrant-global-status` is fast `vagrant global-status` command.  
It parses directly vagrant machine-index file, not use Ruby VM.

```console
$ time vagrant global-status
vagrant global-status  0.97s user 0.09s system 99% cpu 1.072 total

$ time vagrant-global-status
vagrant-global-status  0.00s user 0.00s system 68% cpu 0.007 total # fast!!
```

## How to use

WIP...

## Installation

### Developer

```console
$ go get github.com/monochromegane/vagrant-global-status/...
```

### User

Download from the following url.

- [https://github.com/monochromegane/vagrant-global-status/releases](https://github.com/monochromegane/vagrant-global-status/releases)

## Usage

```console
$ vagrant-global-status
3369de2 default virtualbox poweroff /Users/name/vagrant-dir
```

## Code status

[![wercker status](https://app.wercker.com/status/8169cbc1fb1e057433f7a06f0dd0cf97/m/master "wercker status")](https://app.wercker.com/project/bykey/8169cbc1fb1e057433f7a06f0dd0cf97)

[![Build Status](https://travis-ci.org/monochromegane/vagrant-global-status.svg?branch=master)](https://travis-ci.org/monochromegane/vagrant-global-status)

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

