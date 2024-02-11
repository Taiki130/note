## actにわたすpayloadを調べるには

```
# change this to the event type you want to get the data for
on:
  pull_request:
    types: [opened, closed, reopened]

jobs:
  printJob:    
    name: Print event
    runs-on: ubuntu-latest
    steps:
    - name: Dump GitHub context
      env:
        GITHUB_CONTEXT: ${{ toJson(github) }}
      run: |
        echo "$GITHUB_CONTEXT"
```

要検証
ref: https://stackoverflow.com/questions/63803136/how-to-get-my-own-github-events-payload-json-for-testing-github-actions-locally
