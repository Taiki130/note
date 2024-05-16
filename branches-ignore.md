下記のworkflowはrelease PR作成では動いてしまう
どうやらpull_request eventの場合 branches-ignoreの対象はbase branchになるらしい（push eventも？）

```
name: Auto Assign to PR

on:
  pull_request:
    types: [opened, reopened]
    branches-ignore:
      - 'release'
```

もし特定のhead branchを対象としてskipする場合は、下記のようにjobs.ifでskipするしかなさそう

```
name: Auto Assign to PR

on:
  pull_request:
    types: [opened, reopened]
    branches-ignore:
      - 'release'

jobs:
  if: ${{ github.head_ref != 'release' }}
```
