# 課題１
課題１はシンプルなWebAPI一個の実装だからそんなに手間ではない。  
基礎を抑えるために用語について少し解説する。
## 用語とツールの解説
### Docker
Dockerは、コンテナ型の仮想化ツールである。詳しくは[ここ](https://qiita.com/gold-kou/items/44860fbda1a34a001fc1)   
*仮想化*とは、簡単に言ってしまえばコンピュータの中にもう一つコンピュータを作る技術のこと。これができると、WindowsやMacOSでLinuxが動かせたりするので、実際に使うサーバと同じ環境で開発ができるようになる。  
Dockerの特徴は、

1. 他の仮想化ツールより軽い(コンテナ型仮想化のメリット)
2. Dockerfileという簡単なファイルから仮想環境を作成できる。
3. 色々な人がDockerfileを共有してくれている

　の三点である。

## Dockerfile
Dockerで作成する仮想環境の作成手順を書いた手順書みたいなもの。  
文法とかルールについては[ここ](http://docs.docker.jp/engine/articles/dockerfile_best-practice.html)に書いてある。


### Go言語(Golang)
Googleの作った言語で、簡単で動作が速い。

### WebAPI
超簡単にいうと、URLにアクセスするとJSON形式やXML形式でデータを返してくれる仕組みのこと。詳しい解説は[ここ](https://qiita.com/NagaokaKenichi/items/df4c8455ab527aeacf02)  
REST APIと呼ばれる形式のものが多い
今回の課題では、`http://localhost:9090/`にアクセスすると

```json
{
    "message": "Hello, world"
}
```

が返ってくるWebAPIの作成が課題である

### JSON
JSONは、HTMLやYAMLとかの仲間で、文章のフォーマットみたいなもの。  
`{}`で囲われていて`key:value`の形であればJSONとよべる。  

## コードと動かし方
本来、GolangでWebサーバの開発を行う時は`gorilla/mux`や`julienschmidt/httprouter`を使って開発をする(その方が書くコードが少ないから)。  
しかし、簡単なWebAPIを作るだけなら標準ライブラリだけでいける。　　
サーバの中身については`main.go`を参照。
### 起動方法
```bash
docker build . -t practice1
docker run --rm -d  -p 9090:9090 practice1:latest
```
これで、`localhost:9090`にアクセスするとJSONが返ってくるはず。

dockerコマンドの`-p`オプションは、自分のPCの`9090`ポートへのアクセスをdockerで作成した環境の`9090`ポートに流し込むためのもの(ポートフォワーディング)。  

`docker build`は、DockerImageを作成するためのコマンド。
このDockerfileだと、コードを変更するたびにDockerImageを作り直す必要がある。
　


## 参考文献
* https://qiita.com/katekichi/items/d94e078b376151858ca4
* http://sgykfjsm.github.io/blog/2016/03/13/golang-json-api-tutorial/)
* http://otiai10.hatenablog.com/entry/2014/10/06/195442
* https://qiita.com/yasuno0327/items/be7fb992054f40b491cc