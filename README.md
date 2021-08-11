# getip

[![Build Status](https://travis-ci.com/OpenIoTHub/getip.svg?branch=master)](https://travis-ci.com/OpenIoTHub/getip)

[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-white.svg)](https://snapcraft.io/publicip)

get my public ipv4 and ipv6 address by cli
```
getip
```
You can install the pre-compiled binary (in several different ways),
use Docker.

Here are the steps for each of them:

## Install the pre-compiled binary

**homebrew tap** :

```sh
$ brew install OpenIoTHub/tap/getip
```

**homebrew** (may not be the latest version):

```sh
$ brew install getip
```

**snapcraft**:

```sh
$ sudo snap install publicip
```

**scoop**:

```sh
$ scoop bucket add OpenIoTHub https://github.com/OpenIoTHub/scoop-bucket.git
$ scoop install getip
```

**deb/rpm**:

Download the `.deb` or `.rpm` from the [releases page][releases] and
install with `dpkg -i` and `rpm -i` respectively.

```sh
getip
```

**Shell script**:

```sh
$ curl -sfL https://install.goreleaser.com/github.com/OpenIoTHub/getip.sh | sh
```

**manually**:

Download the pre-compiled binaries from the [releases page][releases] and
copy to the desired location.

## Running with Docker

You can also use it within a Docker container. To do that, you'll need to
execute something more-or-less like the following:

```sh
$ docker run  -it openiothub/getip:latest
```

Note that the image will almost always have the last stable Go version.

[releases]: https://github.com/OpenIoTHub/getip/releases

