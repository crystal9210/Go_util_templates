package sample

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	mock "test_mod/defaultpack_mock/mock_sample"
)

func TestSample(t *testing.T) {
	// モックコントローラの生成とリソース解法の処理実装
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// sample.writerとio.Writerを共に実装
	w := mock.NewMockwriter(ctrl)

	gomock.InOrder(
		w.EXPECT().Write([]byte("hoge")).Return(4, nil),
		w.EXPECT().Write([]byte("fuga")).Return(4, nil),
	)

	// io.Writerとして渡す
	fmt.Fprintf(w, "hoge")
	fmt.Fprintf(w, "fuga")
}
