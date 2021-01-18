## Background
ispecはゼロイチでの開発が多いため、パッケージ構成やアーキテクチャを実装する機会も多い。

## Objective
パッケージ構成をコマンドラインから対話的に生成できるようにする。


## Usage

simple example
```
$ cookiecutter git@github.com:ispec-inc/going-to-go.git
component_id []: sample
owner [yamad07]:
organization_name [ispec-inc]:
db_name [sample]:
repository_name [sample]:
$ cd sample
$ cat README.md
```

monolithic repository
```
$ cd path/to/monorepo
$ cd ..
$ cookiecutter git@github.com:ispec-inc/going-to-go.git
component_id []: sample
owner [yamad07]:
organization_name [ispec-inc]:
db_name [sample]:
repository_name [sample]: monorepo/sample
```

## Future Work
- [ ] FirebaseやJWTなどの認証をいれるかどうかを選択でき、選択に応じてモジュールの追加や削除を行う。
