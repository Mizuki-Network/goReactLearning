package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	//Load()の前にGO_ENVの判定を定義しているので"$GO_ENV=dev go run migrate/migrate.go"で実行前に定義が必要となる
	if os.Getenv("GO_ENV") == "dev" { // もしGO_ENVが"dev"ならば
		err := godotenv.Load() // ルートディレクトリの.envファイルを読み込む
		if err != nil {
			log.Fatalln(err)
		}
	}
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	//gorm.Open()でDB接続を開くのに使用
	//postgres.Open()によりDB:postgresへ接続
	//&gorm.Config{}で接続やログ出力の設定を行えるが、空白{}によりデフォルトの値で起動
	//dbに*gorm.DB型が返り値として格納され、このオブジェクトを用いてCURD操作を行う
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected")
	return db
}

func CloseDB(db *gorm.DB) {
	// GORMのDBオブジェクトから、内部のsql.DBオブジェクトを取得(sql.DBは標準パッケージ)
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	// sql.DBオブジェクトのClose()メソッドを呼び出して接続を閉じる
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
