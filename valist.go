package valist

type TeamMeta struct {
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
	// Licenses contains a list of licenses for the release.
	Licenses []string `json:"licenses"`
}

type LicenseMeta struct {
	// Image is the URI of the license image
	Image string `json:"image"`
	// Name is the unique license name.
	Name string `json:"name"`
	// Description is a description of the license.
	Description string `json:"description"`
	// ExternalURL is a link to the license website.
	ExternalURL string `json:"external_url"`
}
