pushイベントトリガーで実行されるworkflowはデフォルトブランチになくても、実行される。

しかし実行されない場合はsyntax errorになっている可能性がある。

`gh workflow view <workflow name>`

これで最近のrunが見れる

`gh run view 7665236335 -w`

ブラウザでみる

workflow_dispatchのeventを確認するにはすでにデフォルトブランチにあるworkflowを修正して確認する

https://qiita.com/trackiss/items/02eefc2ab8ccfd41768c
