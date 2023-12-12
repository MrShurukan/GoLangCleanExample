package application

import (
	"encoding/json"
	"os"
)
import "testClean/domain"

// 1) Просмотреть товары -> Вернуть все товары из базы
func GetAllGoods() ([]domain.Good, error) {
	fileCont, err := os.ReadFile("storage/goods.json")
	if err != nil {
		return nil, err
	}

	var goods []domain.Good
	err = json.Unmarshal(fileCont, &goods)
	if err != nil {
		return nil, err
	}

	return goods, nil
}

// 2) Создать заказ
// 3) Добавить товар в заказ
// 4) Удалить товар из заказа
// 5) Финализировать заказ -> Пометить, что он готов, заблокировать от изменения
