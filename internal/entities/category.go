package entities

import (
	"fmt"
	"time"
)

type Category struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreateAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCategory(name string) (*Category, error) {
	category := &Category{
		Name:      name,
		CreateAt:  time.Now(),
		UpdatedAt: time.Now(),
	}
	err := category.IsValid()

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (c *Category) IsValid() error {
	if len(c.Name) < 5 {
		return fmt.Errorf("Name must be greater than5. got %d", len(c.Name))
	}

	return nil
}
