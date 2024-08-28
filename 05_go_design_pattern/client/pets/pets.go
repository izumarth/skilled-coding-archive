package pets

type Pet struct {
	Species          string `json:"species"`
	Breed            string `json:"breed"`
	MinWeight        int    `json:"min_weight,omiteempty"`
	MaxWeight        int    `json:"max_weight,omiteempty"`
	AverageWeight    int    `json:"average_weight,omiteempty"`
	Weight           int    `json:"weight,omiteempty"`
	Description      string `json:"description,omiteempty"`
	Lifespan         int    `json:"lifespan,omiteempty"`
	GeographicOrigin string `json:"geographic_origin,omiteempty"`
	Color            string `json:"colort,omiteempty"`
	Age              int    `json:"age,omiteempty"`
	AgeEstimated     bool   `json:"age_estimated,omiteempty"`
}
