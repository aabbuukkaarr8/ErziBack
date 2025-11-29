package product

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// PriceEntry представляет одну запись цены
type PriceEntry struct {
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

// Prices представляет массив цен (JSONB)
type Prices []PriceEntry

// Value реализует driver.Valuer для записи в БД
func (p Prices) Value() (driver.Value, error) {
	if p == nil {
		return []byte("[]"), nil
	}
	return json.Marshal(p)
}

// Scan реализует sql.Scanner для чтения из БД
func (p *Prices) Scan(value interface{}) error {
	if value == nil {
		*p = Prices{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, p)
}

type Model struct {
	ID          int
	Title       string
	Description string
	ImageURL    string
	IsActive    bool
	Category    string
	CreatedAt   time.Time
	Prices      Prices
}

type Attribute struct {
	ID        int
	ProductID int
	Key       string
	Value     string
}
