CREATE TABLE IF NOT EXISTS states_transport_cost (
    State VARCHAR(50) PRIMARY KEY,
    Air_Cost FLOAT,
    Bus_Cost_Inter FLOAT,
    Bus_Cost_Intra FLOAT,
    Motorcycle_Cost FLOAT,
    Water_Cost FLOAT
);


