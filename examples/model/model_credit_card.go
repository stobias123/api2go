package model

type CreditCard struct {
	ID string
}

func (c CreditCard) GetName() string  {
	return "credit_cards"
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c CreditCard) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *CreditCard) SetID(id string) error {
	c.ID = id
	return nil
}

