// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: copyfrom.go

package sqlc

import (
	"context"
)

// iteratorForBulkInsertStateFoodCosts implements pgx.CopyFromSource.
type iteratorForBulkInsertStateFoodCosts struct {
	rows                 []BulkInsertStateFoodCostsParams
	skippedFirstNextCall bool
}

func (r *iteratorForBulkInsertStateFoodCosts) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForBulkInsertStateFoodCosts) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].State,
		r.rows[0].Cost,
	}, nil
}

func (r iteratorForBulkInsertStateFoodCosts) Err() error {
	return nil
}

func (q *Queries) BulkInsertStateFoodCosts(ctx context.Context, arg []BulkInsertStateFoodCostsParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"states_food_costs"}, []string{"state", "cost"}, &iteratorForBulkInsertStateFoodCosts{rows: arg})
}

// iteratorForBulkInsertStateTransportCosts implements pgx.CopyFromSource.
type iteratorForBulkInsertStateTransportCosts struct {
	rows                 []BulkInsertStateTransportCostsParams
	skippedFirstNextCall bool
}

func (r *iteratorForBulkInsertStateTransportCosts) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForBulkInsertStateTransportCosts) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].State,
		r.rows[0].AirCost,
		r.rows[0].BusCostInter,
		r.rows[0].BusCostIntra,
		r.rows[0].MotorcycleCost,
		r.rows[0].WaterCost,
	}, nil
}

func (r iteratorForBulkInsertStateTransportCosts) Err() error {
	return nil
}

func (q *Queries) BulkInsertStateTransportCosts(ctx context.Context, arg []BulkInsertStateTransportCostsParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"states_transport_cost"}, []string{"state", "air_cost", "bus_cost_inter", "bus_cost_intra", "motorcycle_cost", "water_cost"}, &iteratorForBulkInsertStateTransportCosts{rows: arg})
}
