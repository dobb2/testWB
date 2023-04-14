package entities

import (
	"errors"
	"strconv"
	"strings"
)

type Item struct {
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

type ArrayItem []Item

func (it *ArrayItem) Scan(value interface{}) error {
	var valueStr string
	if value == nil {
		*it = make([]Item, 0)
		return nil
	}

	switch value.(type) {
	case string:
		valueStr = value.(string)
	default:
		return errors.New("Incompatible type for Items")
	}
	itemsStr := strings.Split(valueStr, "\",\"")
	for _, itemStr := range itemsStr {
		treatedStr := strings.Trim(itemStr, "{}()")
		itemElements := strings.Split(treatedStr, ",")
		itemElements[9] = strings.Trim(itemElements[9], "\"\\")

		ChartID, _ := strconv.Atoi(itemElements[0])
		Price, _ := strconv.Atoi(itemElements[2])
		Sale, _ := strconv.Atoi(itemElements[5])
		TotalPrice, _ := strconv.Atoi(itemElements[7])
		NmID, _ := strconv.Atoi(itemElements[8])
		Status, _ := strconv.Atoi(itemElements[10])
		elem := Item{
			ChrtID:      ChartID,
			TrackNumber: itemElements[1],
			Price:       Price,
			Rid:         itemElements[3],
			Name:        itemElements[4],
			Sale:        Sale,
			Size:        itemElements[6],
			TotalPrice:  TotalPrice,
			NmID:        NmID,
			Brand:       itemElements[9],
			Status:      Status,
		}
		*it = append(*it, elem)
	}

	return nil
}

/*
func (it ArrayItem) Value() (driver.Value, error) {
	itemsStr := "ARRAY["
	for i, item := range it {
		itemElemStr := "(" + strconv.FormatInt(int64(item.ChrtID), 10) + ", '" +
			item.TrackNumber + "', " +
			strconv.FormatInt(int64(item.Price), 10) + ", '" +
			item.Rid + "', '" +
			item.Name + "', " +
			strconv.FormatInt(int64(item.Sale), 10) + ", '" +
			item.Size + "', " +
			strconv.FormatInt(int64(item.TotalPrice), 10) + ", " +
			strconv.FormatInt(int64(item.NmID), 10) + ", '" +
			item.Brand + "', " +
			strconv.FormatInt(int64(item.Status), 10) + ")::Items"
		itemsStr += itemElemStr
		if i+1 != len(it) {
			itemsStr += ","
		}
	}
	itemsStr += "]"
	return itemsStr, nil
}

*/
