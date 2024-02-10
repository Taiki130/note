PRにラベルが含まれるかどうか

```yaml
- if: contains(fromJSON(${{ fromJSON(github.event.pull_request.labels.*.name) }}), 'help wanted')
  shell: bash
  run: |
    # do something..
```

shellでラベルを使った処理

```yaml
- shell: bash
  env:
    LABELS: ${{ fromJSON(github.event.pull_request.labels.*.name) }}
  run: |
    for i in $(seq 0 $(($(echo "$LABELS" | jq length) - 1)))
    do
      label=$(echo "$LABELS" | jq -r .[$i])
      # do something..
      echo "$label"
    done
```

shellで配列に特定の値が含まれているかチェック

ref: https://github.com/slackhq/vitess/blob/bebddba7bbdf675d3f01badd59947a37f5cad5d4/.github/workflows/endtoend.yml#L12
