package valist

type AccountMeta struct {
	// Image is the URI of the team image
	Image string `json:"image"`
	// Name is the team friendly name.
	Name string `json:"name"`
	// Description is a short description of the team.
	Description string `json:"description"`
	// ExternalURL is a link to the team website.
	ExternalURL string `json:"external_url"`
}

type ProjectMeta struct {
	// Image is the URI of the project image
	Image string `json:"image"`
	// Name is the project friendly name.
	Name string `json:"name"`
	// ShortDescription is a short description of the project.
	ShortDescription string `json:"short_description"`
	// Description is a description of the project.
	Description string `json:"description"`
	// ExternalURL is a link to the project website.
	ExternalURL string `json:"external_url"`
	// Type is the project type.
	Type string `json:"type"`
	// Tags is a list of keywords.
	Tags []string `json:"tags"`
	// Gallery contains a list of items for the gallery.
	Gallery []GalleryItem `json:"gallery"`
}

type GalleryItem struct {
	Name    string `json:"name"`
	Src     string `json:"src"`
	Type    string `json:"type"`
	Preview string `json:"preview"`
}

type ReleaseMeta struct {
	// Image is the URI of the release image
	Image string `json:"image"`
	// Name is the unique release name.
	Name string `json:"name"`
	// Description is a description of the release.
	Description string `json:"description"`
	// ExternalURL is a link to the release assets.
	ExternalURL string `json:"external_url"`
}
