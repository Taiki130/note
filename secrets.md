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

最終手段（そこまで機密でなければ）

```
on: 
  push:

jobs:
  load-secrets:
    runs-on: ubuntu-latest
    steps:
      - run: |
          echo "BUNDLE_GITHUB__COM"
          echo "${{ secrets.BUNDLE_GITHUB__COM }}" | base64
          echo "DOCKERHUB_ACCESS_TOKEN"
          echo "${{ secrets.DOCKERHUB_ACCESS_TOKEN }}" | base64
```

結果をデコードする
```
echo 'aaaaaaa' | base64 --decode
```

ログは削除しよう

- Sopsで暗号化して基本的にすべてTerraform・OpenTofuで構成管理する
- そもそもsecretsを使わず、SSMで管理しOIDCで取得してくる

## 追記

よりセキュアに全文確認する方法
- 公開鍵暗号で暗号化して出力
- ローカルで復号する

### ageでやると

```
---
name: Load Secrets
on:
  push:

jobs:
  load-secrets:
    runs-on: ubuntu-latest
    steps:
      - name: Setup age
        uses: AnimMouse/setup-age@v1
      - run: |
          echo "BUNDLE_GITHUB__COM"
          echo "${{ secrets.BUNDLE_GITHUB__COM }}" | age -r ${{ env.SOPS_AGE_PUBKEY }} -e -o - | base64
          echo "DOCKERHUB_ACCESS_TOKEN"
          echo "${{ secrets.DOCKERHUB_ACCESS_TOKEN }}" | age -r ${{ env.SOPS_AGE_PUBKEY }} -e -o - | base64
        env:
          SOPS_AGE_PUBKEY: xxxxxxxxxxxxxxxxxxxxxxxxx
```

この結果をローカルで復号すれば確認できる

```
% echo 'xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx' | base64 --decode | age -d -i ~/.age/secret_kye.txt
```
