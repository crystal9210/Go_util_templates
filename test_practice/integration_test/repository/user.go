package repository

import (
	"integration/model"
)

// ユーザー情報を取得するためのインターフェース
type UserRepository interface {
	FindByID(id int) (*model.User, error)
}

// MockUserRepository はUserRepositoryのモック実装です。
type MockUserRepository struct{}

// モックのユーザー情報を返すFindByID関数
func (m *MockUserRepository) FindByID(id int) (*model.User, error) {
	// テストのため、ID に 1 を指定した場合のみ、固定のユーザー情報を返します。
	if id == 1 {
		return &model.User{ID: 1, Name: "John Doe"}, nil
	}
	// 存在しないIDの場合はnilを返します。
	return nil, nil
}
