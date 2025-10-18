package stripeclient

import (
	"context"

	"github.com/stripe/stripe-go/v83"
)

// CreateCustomer creates a new Stripe customer for V1 Consortium users
func (sc *StripeClient) CreateCustomer(ctx context.Context, data *CustomerData) (*CustomerData, error) {
	params := &stripe.CustomerCreateParams{
		Email:       stripe.String(data.Email),
		Name:        stripe.String(data.Name),
		Description: stripe.String(data.Description),
		Phone:       stripe.String(data.Phone),
		Metadata:    data.Metadata,
	}

	if data.Address != nil {
		params.Address = &stripe.AddressParams{
			Line1:      stripe.String(data.Address.Line1),
			Line2:      stripe.String(data.Address.Line2),
			City:       stripe.String(data.Address.City),
			State:      stripe.String(data.Address.State),
			PostalCode: stripe.String(data.Address.PostalCode),
			Country:    stripe.String(data.Address.Country),
		}
	}

	customer, err := sc.client.V1Customers.Create(ctx, params)
	if err != nil {
		return nil, WrapStripeError(err)
	}

	return &CustomerData{
		ID:          customer.ID,
		Email:       customer.Email,
		Name:        customer.Name,
		Description: customer.Description,
		Phone:       customer.Phone,
		Metadata:    customer.Metadata,
		Address:     customer.Address,
	}, nil
}

// GetCustomer retrieves a customer by ID
func (sc *StripeClient) GetCustomer(ctx context.Context, customerID string) (*CustomerData, error) {
	customer, err := sc.client.V1Customers.Retrieve(ctx, customerID, nil)
	if err != nil {
		return nil, WrapStripeError(err)
	}

	return &CustomerData{
		ID:          customer.ID,
		Email:       customer.Email,
		Name:        customer.Name,
		Description: customer.Description,
		Phone:       customer.Phone,
		Metadata:    customer.Metadata,
		Address:     customer.Address,
	}, nil
}

// UpdateCustomer updates an existing customer
func (sc *StripeClient) UpdateCustomer(ctx context.Context, customerID string, data *CustomerData) (*CustomerData, error) {
	params := &stripe.CustomerUpdateParams{}

	if data.Email != "" {
		params.Email = stripe.String(data.Email)
	}
	if data.Name != "" {
		params.Name = stripe.String(data.Name)
	}
	if data.Description != "" {
		params.Description = stripe.String(data.Description)
	}
	if data.Phone != "" {
		params.Phone = stripe.String(data.Phone)
	}
	if data.Metadata != nil {
		params.Metadata = data.Metadata
	}
	if data.Address != nil {
		params.Address = &stripe.AddressParams{
			Line1:      stripe.String(data.Address.Line1),
			Line2:      stripe.String(data.Address.Line2),
			City:       stripe.String(data.Address.City),
			State:      stripe.String(data.Address.State),
			PostalCode: stripe.String(data.Address.PostalCode),
			Country:    stripe.String(data.Address.Country),
		}
	}

	customer, err := sc.client.V1Customers.Update(ctx, customerID, params)
	if err != nil {
		return nil, WrapStripeError(err)
	}

	return &CustomerData{
		ID:          customer.ID,
		Email:       customer.Email,
		Name:        customer.Name,
		Description: customer.Description,
		Phone:       customer.Phone,
		Metadata:    customer.Metadata,
		Address:     customer.Address,
	}, nil
}

// DeleteCustomer deletes a customer
func (sc *StripeClient) DeleteCustomer(ctx context.Context, customerID string) error {
	_, err := sc.client.V1Customers.Delete(ctx, customerID, nil)
	if err != nil {
		return WrapStripeError(err)
	}
	return nil
}
