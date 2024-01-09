package showcase

import uuid "github.com/satori/go.uuid"

type Showcase struct {
	ID    uuid.UUID
	Title string
	Body  string
	Image string
	Sort  int
}
