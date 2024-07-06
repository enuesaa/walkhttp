# designdoc

## TODO
- インタフェースの調整が必要
  - web or cli
  - サブコマンドを作るのか

## インタフェース
- サブコマンドなし

### コンソール
```console
$ walkhttp
> GET /
****
* GET /
> Accept application/json
***
* GET /
* Accept application/json
>
# execute GET here.
> ctl
# serve web console here.
> export
# export http request results as json file.
# also, this file can be reused when supplied to cli argument.
```
