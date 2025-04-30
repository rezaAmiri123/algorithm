package main
// Apartment
// AptID		int
// UnitNumber	varchar(10)
// BuildingID	int

// AptTenants
// TenantID		int
// AptID			int

// Tenants
// TenantID   int
// TenantName varchar(100)


// Multiple Apartments: Write a SQL query to get a list of tenants who are renting more than one
// apartment.

// SELECT TenantName FROM Tenants
// INNER JOIN 
// 	(SELECT TenantID FROM AptTenants GROUP BY TenantID HAVING COUNT(*) > 1) C
// ON Tenants.TenantID = C.TenantID

// Whenever you write a GROUP BY clause in an interview (or in real life), make sure that anything in the
// SELECT clause is either an aggregate function or contained within the GROUP BY clause.
