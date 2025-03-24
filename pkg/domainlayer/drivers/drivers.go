package drivers

import "time"

type Drivers struct {
	id                    int
	make                  string
	model                 string
	registration_number   string
	seating_capacity      int
	capacity              int
	year_of_purchase      int
	total_kms_driven      int
	last_serviced_on      time.Time
	current_mileage       int
	insurance_valid_until time.Time
	current_status        string
	created_at            time.Time
}
