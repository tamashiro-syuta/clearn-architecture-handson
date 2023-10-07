package query_service

import (
	"testing"

	"gopkg.in/testfixtures.v2"

	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/db_test"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
)

var (
	query    *dbgen.Queries
	fixtures *testfixtures.Context
)

func TestMain(m *testing.M) {
	var err error

	// DBの立ち上げ
	resource, pool := dbTest.CreateContainer()
	defer dbTest.CloseContainer(resource, pool)

	// DBへ接続する
	db := dbTest.ConnectDB(resource, pool)
	defer db.Close()

	// テスト用DBをセットアップ
	dbTest.SetupTestDB()

	// テストデータの準備
	fixturePath := "../../fixtures"
	fixtures, err = testfixtures.NewFolder(db, &testfixtures.MySQL{}, fixturePath)
	if err != nil {
		panic(err)
	}

	query = dbgen.New(db)

	// テスト実行
	m.Run()
}

func resetTestData(t *testing.T) {
	t.Helper()
	if err := fixtures.Load(); err != nil {
		t.Fatal(err)
	}
}
