package infrastructure

import (
	"encoding/json"
	"os"
	"testClean/domain"
)

func GetAllGoodsFromStorage() ([]domain.Good, error) {
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
