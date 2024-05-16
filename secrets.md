terraformでimportしたり、GitHub APIでsecretsの中身を確認することはできなさそう。

secretsの内容をチェックする現状思いつく方法は、以下のようにsecretsの一部をログに出すぐらい

```
name: Load Secrets
on: 
  push:

jobs:
  load-secrets:
    runs-on: ubuntu-latest
    steps:
      - run: |
          echo "BUNDLE_GITHUB__COM"
          echo "${{ secrets.BUNDLE_GITHUB__COM }}" | cut -c -30
          echo "DOCKERHUB_ACCESS_TOKEN"
          echo "${{ secrets.DOCKERHUB_ACCESS_TOKEN }}"| cut -c -15
          echo "DOCKERHUB_USERNAME"
          echo "${{ secrets.DOCKERHUB_USERNAME }}" | cut -c -5
          echo "SLACK_WEBHOOK_URL"
          echo "${{ secrets.SLACK_WEBHOOK_URL }}" | cut -c -20
```

右から数文字のほうがチェックしやすいかも？

```
name: Load Secrets
on: 
  push:

jobs:
  load-secrets:
    runs-on: ubuntu-latest
    steps:
      - run: |
          echo "BUNDLE_GITHUB__COM"
          echo "${{ secrets.BUNDLE_GITHUB__COM }}" | tail -c 2
          echo "DOCKERHUB_ACCESS_TOKEN"
          echo "${{ secrets.DOCKERHUB_ACCESS_TOKEN }}"| tail -c 2
          echo "DOCKERHUB_USERNAME"
          echo "${{ secrets.DOCKERHUB_USERNAME }}" | tail -c 2
          echo "SLACK_WEBHOOK_URL"
          echo "${{ secrets.SLACK_WEBHOOK_URL }}" | tail -c 2
```

ログは削除しよう

- Sopsで暗号化して基本的にすべてTerraform・OpenTofuで構成管理する
- そもそもsecretsを使わず、SSMで管理しOIDCで取得してくる
