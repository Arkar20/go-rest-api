package customer

func ToCustomerDTO(customer Customer) map[string]interface{} {
	customerMap := map[string]interface{}{
		"id":              customer.ID,
		"name":            customer.Name,
		"phnum1":          customer.PhNum1,
		"phnum2":          customer.PhNum2,
		"company":         customer.Company,
		"address":         customer.Address,
		"retail_customer": customer.RetailCustomer,
		"marketplace_id":  customer.MarketplaceID,
	}

	return customerMap
}

func ToCustomerDTOArray(customers []Customer) []map[string]interface{} {
	var customerDTOs []map[string]interface{}
	for _, customer := range customers {
		customerDTOs = append(customerDTOs, ToCustomerDTO(customer))
	}
	return customerDTOs
}
