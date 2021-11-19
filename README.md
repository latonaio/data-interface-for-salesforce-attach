# data-interface-for-salesforce-attach
data-interface-for-salesforce-attach は、Salesforce オブジェクトに pdfファイルを添付するマイクロサービスです。

## 動作環境
data-interface-for-salesforce-attach は、aion-coreのプラットフォーム上での動作を前提としています。  
使用する際は、事前に下記の通りAIONの動作環境を用意してください。     

* OS: Linux OS   
* CPU: ARM/AMD/Intel  
* Kubernetes   
* [AION](https://github.com/latonaio/aion-core)のリソース    

## セットアップ
以下のコマンドを実行して、docker imageを作成してください。
```
$ cd /path/to/data-interface-for-salesforce-attach
$ make docker-build
```

## 起動方法
以下のコマンドを実行して、podを立ち上げてください。
```
$ cd /path/to/data-interface-for-salesforce-attach
$ kubectl apply -f data-interface-for-salesforce-pdf.yaml
```

## kanban との通信
### kanban から受信するデータ
kanban から受信する metadata に下記の情報を含む必要があります。

| key | value |
| --- | --- |
| method | 使用する HTTP メソッド |
| object_name | オブジェクト名 |
| object_id | オブジェクト id |
| connection_type | request |
| file_name | ファイル名 |
| path | ファイルパス |
| file_extension | ファイル拡張子 |

具体例: 
```example
# metadata (map[string]interface{}) の中身

"method": "put"
"object_name": "ContractPDF"
"object_id": "test_id"
"file_name": "xxxxx"
"path": "/"
"file_estension": "pdf"
"connection_type": "request"
```

### kanban に送信するデータ
kanban に送信する metadata は下記の情報を含みます。

| key | type | description |
| --- | --- | --- |
| method | string | 使用する HTTP メソッド |
| object | string | 文字列 "pdf" |
| query_params | map[string]string | クエリパラメータ |
| body | string | ファイルの中身(base64エンコードされている)|
| connection_key | string | pdf |
| is_body_base64 | bool | true |

具体例: 
```example
# metadata (map[string]interface{}) の中身

"method": "put"
"object": "ContractPDF"
"path_param": "xxxx"
"query_params": map[string]string{"pdfName": "xxxxx.pdf"}
"body": "yyyyyyyyyyy"
"is_body_base64": true
"connection_type": "request"
```

## kanban(salesforce-api-kube) から受信するデータ
kanban からの受信可能データは下記の形式です

| key | value |
| --- | --- |
| key | オブジェクト名 |
| content | ファイルの登録状態 |
| connection_type | 文字列 "response" |

具体例:
```example
# metadata (map[string]interface{}) の中身

"key": "ContractPDF"
"content": "{id: xxxxx, file_name: yyyyy}"
"connection_type": "response"
```