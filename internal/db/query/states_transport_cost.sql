-- name: BulkInsertStateTransportCosts :copyfrom
INSERT INTO states_transport_cost (
    State, 
    Air_Cost, 
    Bus_Cost_Inter, 
    Bus_Cost_Intra, 
    Motorcycle_Cost, 
    Water_Cost
) VALUES (
    $1, $2, $3, $4, $5, $6
);

-- name: GetStateCost :one
SELECT * FROM states_transport_cost
WHERE State = $1;

-- name: ListStatesCosts :many
SELECT * FROM states_transport_cost
ORDER BY State;


