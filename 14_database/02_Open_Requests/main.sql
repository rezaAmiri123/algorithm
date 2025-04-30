-- Open Requests: Write a SQL query to get a list of all buildings and the number of open requests
-- (Requests in which status equals'Open').

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

SELECT BuildingName,ISNULL(count,0) AS Count
FROM Building LEFT JOIN
    (SELECT Apartment.BuildingID, count(*) AS Count
    FROM Requests INNER JOIN Apartment
    ON Apartment.AptID = Requests.AptID
    WHERE Requests.Status = 'Open'
    GROUP BY Apartment.BuildingID) AS REqCounts
ON Requests.BuildingID = Building.BuildingID

-- Queries like this that utilize sub-queries should be thoroughly tested, even when coding by hand. It may be
-- useful to test the inner part of the query first, and then test the outer part.
