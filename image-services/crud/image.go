package crud

// Image is ..
type Image struct {
	// Primary Key
	ImageID string `json:"imageID"`

	// Properties
	Description string `json:"description"`
	Key         string `json:"key"`

	// MetaData
	CreatedAt          string `json:"createdAt"`
	UpdatedAt          string `json:"updatedAt"`
	CreatedAtTimestamp int64  `json:"createdAtTimestamp"`
	UpdatedAtTimestamp int64  `json:"updatedAtTimestamp"`
}

// PartialImage is ..
type PartialImage struct {

	// Properties
	Description string `json:"description"`
	Key         string `json:"key"`
}

// NewImage is ..
type NewImage struct {
	// Properties
	Description string `json:"description"`
	Key         string `json:"key"`
}

// PatchImage is ..
type PatchImage struct {
	Description string `json:"description"`
}
