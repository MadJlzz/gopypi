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

This section will be updated later when the project will be more mature.

Also, as a reference, this solution will be made by keeping in mind that it will be deployed within
Google Cloud AppEngine for serving traffic and Google Cloud Storage for saving packages.

A local solution will also be available.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change. 

## License

This project is under license [GNU GPL V3](LICENSE).
