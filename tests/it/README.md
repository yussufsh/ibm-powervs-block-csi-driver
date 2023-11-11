# Integration Testing

## Introduction

The integration test will create a PowerVS instance and run the CSI driver tests on it. The driver will be built locally for ppc64le and copied over to the pvm instance. The driver will run on the instance and a SSH tunnel will be created. A grpc client will connect to the driver via a tcp port and run the Node and Controller tests.

## Environments
The following environment variables will be required for running the integration test.
```
export IBMCLOUD_API_KEY=<api key>             # IBM Cloud API Key

```
Below environment variables are optional and is needed if you want to use custom values.
```
```

## Run the test
To run the test use the following command.
```
make test-integration
```
