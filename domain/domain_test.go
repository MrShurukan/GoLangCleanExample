package domain

import "testing"

func TestOrder_AddGoodToEmpty(t *testing.T) {
	// 1) Подготовить фейковые данные
	var order = Order{}
	var good = Good{Name: "Тест", Price: 100}
	// 2) Запустить
	order.AddGood(good)
	// 3) Проверить
	if len(order.Positions) == 0 {
		t.Error("Товар не был добавлен в order")
		t.Fail()
	}
	if len(order.Positions) > 1 {
		t.Errorf("При добавлении одного товара появилось %d товаров", len(order.Positions))
		t.Fail()
	}
	if order.Positions[0].Good != good {
		t.Errorf("Товар %s не был добавлен корректно: %s", good.Name, order.Positions[0].Good.Name)
		t.Fail()
	}
}

func TestOrder_AddGoodDuplicate(t *testing.T) {
	// 1) Подготовить фейковые данные
	var order = Order{}
	var good = Good{Name: "Тест", Price: 100}
	var good2 = Good{Name: "Тест", Price: 100}
	// 2) Запустить
	order.AddGood(good)
	order.AddGood(good2)
	// 3) Проверить
	if len(order.Positions) == 0 {
		t.Error("Товар не был добавлен в order")
		return
	}
	if len(order.Positions) > 1 {
		t.Errorf("При добавлении двух дубликатов появилось %d товаров", len(order.Positions))
		return
	}
	if order.Positions[0].Good != good {
		t.Errorf("Товар %s не был добавлен корректно: %s", good.Name, order.Positions[0].Good.Name)
	}
	if order.Positions[0].Amount != 2 {
		t.Errorf("Количество товара не было обновлено: %d, ожидалось 2", order.Positions[0].Amount)
	}
}
