package schema

import (
	"regexp"
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/edge"
)

// ImageRepo holds the schema definition for the ImageRepo entity.
type ImageRepo struct {
	ent.Schema
}

// Fields of the ImageRepo.
func (ImageRepo) Fields() []ent.Field {
	return []ent.Field{
        field.String("name").
            // Regexp validation for group name.
            Match(regexp.MustCompile("[a-zA-Z_]+$")),
    }
}

// Edges of the ImageRepo.
func (ImageRepo) Edges() []ent.Edge {
	return []ent.Edge{
        edge.To("images", Image.Type),
    }
}
