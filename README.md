# Reproduction

see https://github.com/golangci/golangci-lint/issues/2752

```
bash -c 'cd hello && golangci-lint cache clean && golangci-lint run --config ../.golangci-lint.yml
```

```
$ golangci-lint version
golangci-lint has version 1.45.2 built from 8bdc4d3 on 2022-03-24T11:49:36Z
```
