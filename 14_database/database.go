package main

// Apartment
// AptID		int
// UnitNumber	varchar(10)
// BuildingID	int

// Complexes
// ComplexesID		int
// ComplexesName	varchar(100)

// Building
// BuildingID 		int
// ComplexesID 	int
// BuildingName 	varchar(100)
// Address 		varchar(500)

// AptTenants
// TenantID		int
// AptID			int

// Requests
// RequestID 	int
// Status 		varchar(100)
// AptID 		int
// Description varchar(500)

// Tenants
// TenantID   int
// TenantName varchar(100)

// Note that each apartment can have multiple tenants, and each tenant can have multiple apartments. Each
// apartment belongs to one building, and each building belongs to one complex
