# Create Viam Location from Golang

This example shows how to create a Location in Viam using the gRPC golang clients.

## Prerequisite Overview

1. Install golang (https://go.dev/doc/install)
2. Build Viam RDK CLI [See Below](#cli-install
3. Login to Viam on the CLI
4. Find your organization ID

### 2. CLI Install

Checkout the RDK repo.

1. `git clone https://github.com/viamrobotics/rdk`

2. `cd rdk`

Build The CLI

3. `go build -o ~/go/bin/viam cli/viam/main.go`


### 3. Login To Viam on the CLI

Run `~/go/bin/viam auth` and follow the steps.

### 4. Find your organization ID

`~/go/bin/viam organizations list`

Example:
```
Adam's Org (id: 1111111-2222-3333-ba2a-32b38905be3c)
```

## Run the sample code

Checkout this sample app: `git clone https://github.com/AdamMagaluk/viam-sample-api-access && cd viam-sample-api-access`

```
VIAM_ACCESS_TOKEN=`~/go/bin/viam auth print-access-token` go run main.go 1111111-2222-3333-ba2a-32b38905be3c "New Location"
```

Note: copy the organization ID above into the first argument.