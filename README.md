[![Go Report Card](https://goreportcard.com/badge/github.com/MadJlzz/gopypi)](https://goreportcard.com/report/github.com/MadJlzz/gopypi)
[![Release](https://img.shields.io/github/release/MadJlzz/gopypi.svg?style=flat-square)](https://github.com/MadJlzz/gopypi/releases/latest)

# gopypi

gopypi is partial implementation of the [PEP 503](https://www.python.org/dev/peps/pep-0503/) also known as
the `Simple Repository API`.

Everybody knows `PyPi` ([The Python Package Index](https://pypi.org/)) when it comes to retrieve Python dependencies and
since I needed to store **private packages**, I decided to make my own implementation with Golang.

## Installation

### Google Cloud Platform

#### App Engine

At the moment, `gopypi` has been tested and deployed only within Google App Engine. It is handy as it comes with several
features like autoscaling or even TLS termination.

To deploy `gopypi` in App Engine, just update the `app.yaml` with the correct
`API_KEY` that you've generated. You can use this snippet to generate easily a token:

```go
package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	// Pure bytes generated key.
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	fmt.Printf("%x\n", b)
	// If you prefer the base64 encoded version.
	s := base64.StdEncoding.EncodeToString(b)
	fmt.Printf("%s", s)
}
```

Once you've update the `API_KEY` environment variable, use `gcloud` to deploy:

```bash
gcloud app deploy app.yaml
```

#### Secret Manager

`gopypi` needs a private key to sign URLs of packages stored in GCS to let `pip` download them.

To let the app retrieve this private key, you will need to create it first from the `service account` that runs the
AppEngine service.

Then, go to the `secret manager` and create a new key called `gopypi-sa-private-key`. Store the JSON service account
file there.

Don't forget to add a permission for the service account that runs `gopypi` to be able to retrieve secrets.
(`secretmanager.secrets.get`: Secret Manager Viewer `roles/secretmanager.viewer`)

You can change the name of the secret to use by changing the secret name of the `configs/gcp.yaml` file.

#### Cloud Storage

`gopypi` leverages GCS to store and serve packages.

Create a bucket with the name `gopypi` or the one of you choice. (check `configs/gcp.yaml` and update the bucket value!)

Like for retrieving the secrets, the `service account` running `gopypi` will need permissions. Go to the newly create
bucket and add `storage.buckets.get` for the `sa`.

That's it you should be good to go!

You can choose the directory structure of your choice for packages as long as you keep the root with packages name. (
e.g. `pydantic/...` or `pydantic/1.8.2/...`)

### Download a package

Just use `pip` as you'd expect to:

```bash
pip install --extra-index-url https://<gopypi server>/simple/ <private_package>
```

:warning: **If you're package has the same name with one in pypi.org**: Order your indexes in a pip configuration file
to search first in your private registry.

## Contributing

Pull requests are always welcome. Don't hesitate to open an issue for questions or changes.

## License

This project is under license [GNU GPL V3](LICENSE).
