【how to use mockgen(gomock)】

1.Go,mockgen をインストール 2.インターフェースのモックを作る(spec.txt で参照先としてきじゅつしている記事と一緒) 3. mkdir -p mock_sample
4.mockgen -source sample.go -destination mock_sample/mock_sample.go 5.モックを使ったテストを書く → 実行
