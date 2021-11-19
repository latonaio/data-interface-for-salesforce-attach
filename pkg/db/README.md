# DB
MySQL のアクセスを担当します。

## 事前準備
接続先の MySQL に関する設定は 環境変数で行います。
下記の環境変数を設定してください。

| env_name | description |
| --- | --- |
| MYSQL_HOST | ホスト名 |
| MYSQL_PORT | ポート番号 |
| MYSQL_USER | ユーザー名 |
| MYSQL_PASSWORD | パスワード |
| MYSQL_DB_NAME | データベース名 |
| MAX_OPEN_CONNECTION | 最大コネクション数 |
| MAX_IDLE_CONNECTION | アイドル状態の最大コネクション数 |

## 使用方法
はじめに DB コネクションの確立を行います。
```
func main() {
	ctx := context.Background()
	if err := db.NewDBConPool(ctx); err != nil {
		panic(err)
	}
}
```
### レコードの登録
Register() メソッドを使用して登録処理を行います。

具体例: Contract の登録の例
```
func main() {
    ctx := context.Background()
    if err := db.NewDBConPool(ctx); err != nil {
    	panic(err)
    }

	cID := "contruct1"
	cName := "hoge"
	cStartDate := "2021-01-01"
	cEndDate := "2021-11-30"
 	customerID := "customer1"
	contract := models.Contract{
		CustomerID: &customerID,
		ContractID: &cID,
		Name: &cName,
		StartDate: &cStartDate,
		EndDate: &cEndDate,
	}
	if err := contract.Register(); err != nil {
		fmt.Printf("failed register: %v", err)
	}
}
```

### レコードの取得
下記の具体例を参考に各種レコードを取得します。

具体例 1: contract_id を指定して contract を取得する
```
func main() {
    ctx := context.Background()
    if err := db.NewDBConPool(ctx); err != nil {
    	panic(err)
    }

	contract, err := models.ContractByID("XYZ_11")
    if err != nil {
    	fmt.Printf("error: %v", err)
    }
    fmt.Printf("contract: %v\n", contract)
}
``` 

具体例 2: customer_id を指定して 複数の contract を取得する
```
func main() {
    ctx := context.Background()
    if err := db.NewDBConPool(ctx); err != nil {
    	panic(err)
    }

	contracts, err := models.ContractsByCustomerID("XYZ_11")
    if err != nil {
    	fmt.Printf("error: %v", err)
    }
    fmt.Printf("contracts: %v\n", contracts)
}
```