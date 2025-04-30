-- Close All Requests: Building #11 is undergoing a major renovation. Implement a query to close all
-- requests from apartments in this building.

-- SOLUTION
-- UPDATE queries, like SELECT queries, can have WHERE clauses. To implement this query, we get a list of all
-- apartment IDs within building #11 and the list of update requests from those apartments.
UPDATE Requests
SET Status = 'Closed'
Where AptID IN (
    SELECT AptID FROM Apartments
    WHERE BuildingID = 11
)

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

-- to be ontinued at page 453