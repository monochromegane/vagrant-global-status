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

## Usage

```console
$ vagrant-global-status
3369de2 default virtualbox poweroff /Users/name/vagrant-dir
```

### With vagrant-peco

If you are using [peco](https://github.com/peco/peco), you can select vagrant-global-status, and execute command !

See: [vagrant-peco](https://github.com/monochromegane/vagrant-peco)

![](https://raw.githubusercontent.com/monochromegane/vagrant-peco/master/images/vagrant-peco-up.gif)

## Installation

### Developer

```console
$ go get github.com/monochromegane/vagrant-global-status/...
```

### User

Download from the following url.

- [https://github.com/monochromegane/vagrant-global-status/releases](https://github.com/monochromegane/vagrant-global-status/releases)

### Note

`vagrant-global-status` depends on the following.

- [Vagrant 1.6 or higher (vagrant global-status)](http://www.vagrantup.com/blog/feature-preview-vagrant-1-6-global-status.html)

## Code status

[![wercker status](https://app.wercker.com/status/8169cbc1fb1e057433f7a06f0dd0cf97/m/master "wercker status")](https://app.wercker.com/project/bykey/8169cbc1fb1e057433f7a06f0dd0cf97)

[![Build Status](https://travis-ci.org/monochromegane/vagrant-global-status.svg?branch=master)](https://travis-ci.org/monochromegane/vagrant-global-status)

## GoDoc

- [GoDoc: vagrant-global-status](https://godoc.org/github.com/monochromegane/vagrant-global-status)

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

