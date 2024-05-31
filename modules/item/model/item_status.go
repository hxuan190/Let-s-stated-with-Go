package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type ItemStatus int

const (
	ItemStatusDoing ItemStatus = iota
	ItemStatusDone
	ItemStatusDeleted
)

var AllItemStatuses = [3]string{"Doing", "Done", "Deleted"}

func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}

	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	stringData := strings.ReplaceAll(string(data), "\"", "")
	value, err := ParseStringToItemStatus(stringData)
	if err != nil {
		return errors.New(fmt.Sprintf(
			"Failed to scan data from mysql: %s",
			value,
		))
	}
	*item = value
	return nil
}

func (item *ItemStatus) String() string {
	return AllItemStatuses[*item]
}

func ParseStringToItemStatus(s string) (ItemStatus, error) {
	for item := range AllItemStatuses {
		if AllItemStatuses[item] == s {
			return ItemStatus(item), nil
		}
	}
	return ItemStatus(0), errors.New("invalid item string")
}

func (item *ItemStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("Failed to scan data from mysql: %s", value))
	}
	v, err := ParseStringToItemStatus(string(bytes))
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to scan data from mysql: %s", value))
	}
	*item = v

	return nil
}

func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}
	return item.String(), nil
}
