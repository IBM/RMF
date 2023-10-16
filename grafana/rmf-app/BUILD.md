# How to build IBM RMF for z/OS Grafana Plugin

## Pre-requisites

- `NodeJS`: >=16
- `Yarn`: 1.x.x
- `Go`: >=1.21
- `GNU Make`: >=3.81
- `jq`: >=v1.6
- `zip`:  >=3.0
- `mdsum`: >=0.9

## Build steps

- Download or clone the latest plugin source code from GitHub url: https://github.com/IBM/RMF.
- Navigate to the directory `grafana/rmf-app`.
- Execute the command: `make all`.  
  This creates the `./build` directory and once successful you can find the
  `ibm-rmf-grafana-<VERSION>.zip` and `ibm-rmf-grafana-<VERSION>.zip.md5` files there.  
