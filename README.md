# helm-apply
A Helm Plugin to allow you to apply a specific file inside the chart over the top of the default values, allowing you to ship default configurations that are easily swappable in your chart


## Usage

```shell
helm apply upgrade mychart release-name ha.yaml -f local.yaml  --wait

helm apply template mychart release-name ha.yaml -f local.yaml --dry-run --debug
```
