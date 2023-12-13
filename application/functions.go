package application

import (
	"testClean/infrastructure"
)
import "testClean/domain"

// 1) Просмотреть товары -> Вернуть все товары из базы
func GetAllGoods() ([]domain.Good, error) {
	goods, err := infrastructure.GetAllGoodsFromStorage()
	if err != nil {
		return nil, err
	}
	return goods, nil
}

// 2) Создать заказ
// 3) Добавить товар в заказ
// 4) Удалить товар из заказа
// 5) Финализировать заказ -> Пометить, что он готов, заблокировать от изменения
