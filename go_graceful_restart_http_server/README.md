# [Golang] graceful restart  http server

## 檔案目錄

```
├── server.go # 測試檔案入口
├── go.mod
├── go.sum
├── .air.toml # $ air init 自動產生
└── tmp       # $ air init 自動產生
    └── main

```

- 1. 下載air套件 (go 1.16以上適用此指令)
```
go install github.com/cosmtrek/air@latest
```

- 2. 進入專案資料夾，初始化後目錄會自動產生 `.air.toml` 檔案
```
air init
```

- 3. 啟動專案
```
air
```

> Air 套件原址 https://github.com/cosmtrek/air