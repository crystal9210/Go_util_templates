package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

// smsServiceMock 構造体:MessageService インターフェースのモック実装;github.com/stretchr/testify/mock パッケージから Mock インターフェースを埋め込んでいる、下のTestChargeCustomer()で使用される
// これは↓のSendChargeNotificationメソッドを実装しているので、MyServiceインターフェースを実装している
type MessageServiceMock struct {
	mock.Mock
}

// SendChargeNotification:SendChargeNotification のモックメソッド
func (m *MessageServiceMock) SendChargeNotification(value int) error {
	fmt.Println("Mocked charge notification function")
	fmt.Printf("Value passed in: %d\n", value)
	args := m.Called(value)
	return args.Error(0)
}

// TestChargeCustomer:MyServiceのChargeCustomerメソッドをテストするメソッド、SMEServiceのモックを作成する
func TestChargeCustomer(t *testing.T) {
	smsService := new(MessageServiceMock)
	// smsServiceという名前のMessageServiceMock構造体のメソッドSendChargeNotificationが100という引数で呼び出されたときに、nilを返すようにモックを設定
	smsService.On("SendChargeNotification", 100).Return(nil)

	myService := MyService{smsService}
	// テスト対象のmyServiceのChargeCustomerメソッドを呼び出し、引数100を渡す、このメソッドは、料金の請求を行う
	// ☆ChargeCustomer内で引数がSendChargeNotificationに渡される、そして、この処理が行われたことがtにより報告される機構が内部的にある
	// 参照：
	// tはtesting.T型のオブジェクトです。このオブジェクトは、テスト中に発生したすべてのアサーションやエラーを追跡し、テストの結果を報告します。
	err := myService.ChargeCustomer(100)

	if err != nil {
		t.Errorf("Error charging customer: %v", err)
	}

	// tで追跡した内容を確認
	// 説明：モックが期待通りのメソッド呼び出しを受け取ったかどうかを検証します。成功した場合、何も出力されずにテストが続行されます。失敗した場合、t.Errorfを使用してテストに失敗したことを報告します。これにより、テストの結果が失敗としてマークされ、関連するエラーメッセージが表示されます。エラーメッセージには、期待されたメソッド呼び出しと実際の呼び出しが含まれ、どのように異なるかが説明されます。
	smsService.AssertExpectations(t)
}
