# golang-json

## 簡述

有關 `json` 的相關操作練習。

## 步驟及說明

* 設定書本的結構。
* 取得書本結構切片。
* 進行轉換，將結構轉換為 `json` 格式。
* 打印在終端機上查看結果。

## 安裝 docker 環境及執行程式

* clone GitHub repository
```
$ get clone https://github.com/yuyuancha/golang-json.git
```

* 透過 docker-compose 開啟 docker 容器
```
$ docker-compose up -d
```

* 執行 main.go
```
$ docker-compose exec app go run main.go
```

* 關閉 docker 容器
```
docker-compose down
```
