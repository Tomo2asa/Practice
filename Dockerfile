# 使うDockerイメージを指定
FROM golang:latest

# ディレクトリを移動(cdコマンドみたいな感じ)
WORKDIR /go

# Practice1の中身を仮想環境にコピー
ADD . /go

# 実行時のコマンド指定
CMD ["go", "run", "main.go"]