-- name: BulkInsertStateFoodCosts :copyfrom
INSERT INTO states_food_costs (
    State, 
   Cost
) VALUES (
    $1, $2
);