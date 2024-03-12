package sample

import (
	"testing"

	mock "test_mod/mock/mock_sample"

	"github.com/golang/mock/gomock"
)

func TestSample1(t *testing.T) {
	// 新しいモックコントローラの作成；テストのモックオブジェクトを管理するためのもの
	ctrl := gomock.NewController(t)
	// テストの最後にモックコントローラが解放されるようにします。これにより、テストが終了する際にモックの期待値が確認される
	defer ctrl.Finish()

	// mock.NewMockSample(ctrl) を使用して、Sample インターフェースのモックを作成
	mockSample := mock.NewMockSample(ctrl)
	// mockSample.EXPECT().Method("hoge").Return(1) を使用して、Method メソッドが引数 "hoge" を受け取った場合に1を返すことを期待
	// ☆mockSample.EXPECT().Method("hoge").Return(1) における Method は、MockSampleMockRecorder 構造体の Method メソッドが呼び出されています。
	// mockSample.EXPECT() は、MockSample 構造体の EXPECT メソッドを呼び出します。このメソッドは *MockSampleMockRecorder 型の値を返します。その後、Method("hoge") は、MockSampleMockRecorder 構造体の Method メソッドを呼び出します。
	mockSample.EXPECT().Method("hoge").Return(1).Times(2) // Times:呼び出される回数を指定

	// 何回でも良い場合は，AnyTimes
	// mockSample.EXPECT().Method("hoge").Return(1).AnyTimes()

	t.Log("result:", mockSample.Method("hoge"))
	t.Log("result:", mockSample.Method("hoge"))
	// t.Log("result:", mockSample.Method("hoge")) →呼び出される回数が違うとエラーが出る
}
