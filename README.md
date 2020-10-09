[![Go Report Card](https://goreportcard.com/badge/github.com/MadJlzz/gopypi)](https://goreportcard.com/report/github.com/MadJlzz/gopypi)
[![Release](https://img.shields.io/github/release/MadJlzz/gopypi.svg?style=flat-square)](https://github.com/MadJlzz/gopypi/releases/latest)

# gopypi

gopypi is an implementation of the [PEP 503](https://www.python.org/dev/peps/pep-0503/) also known
as the `Simple Repository API`.

Everybody knows the `PyPi` ([The Python Package Index](https://pypi.org/)) and since I needed to store **private packages**,
I decided to make my own implementation with Golang.

## Installation

As usual with go, just use the Good Ol' Classic:
```bash
go get github.com/MadJlzz/gopypi
```

Now simply install it: 
```bash
go install
```

## Usage

### Localhost backend

Simple static file server exposing packages locally.

```bash
$ ./gopypi-local -help
Usage of C:\GithubTech\go\src\github.com\MadJlzz\gopypi\gopypi-local:
  -package-location string
        Location from which we should load packages. (default "C:/DefaultStorage")
  -port string
        Port of the app (default "3000")
```

### Google Cloud Storage 

PyPi server implementation backed by Google Cloud Storage.

You should provide credentials with `storage.buckets.get` permission in order to access bucket
and generate signed urls.
(see [service accounts](https://cloud.google.com/iam/docs/service-accounts) for more details)

```bash
$ ./gopypi-gcs -help
Usage of C:\GithubTech\go\src\github.com\MadJlzz\gopypi\gopypi-gcs:
  -credentials string
        GCP JSON credentials file. (default "credentials/service-account-dev.json")
  -port string
        Port of the app (default "3000")
```

### Download a package

Just use `pip` as you'd expect to:
```bash
pip install --extra-index-url https://<gopypi server>/simple/ <private_package>
```

:warning: **If you're package has the same name with one in pypi.org**: Order your
indexes in a pip configuration file to search first in your private registry. 

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change. 

## License

This project is under license [GNU GPL V3](LICENSE).
