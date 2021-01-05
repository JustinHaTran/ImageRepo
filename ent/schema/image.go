package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/edge"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return []ent.Field{
        field.String("model"),
        field.Time("registered_at"),
    }
}

// Edges of the Image.
func (Image) Edges() []ent.Edge {
	return []ent.Edge{
        // Create an inverse-edge called "owner" of type `User`
        // and reference it to the "cars" edge (in User schema)
        // explicitly using the `Ref` method.
        edge.From("owner", User.Type).
            Ref("images").
            // setting the edge to unique, ensure
            // that a car can have only one owner.
			Unique(),
        edge.From("imagerepos", ImageRepo.Type).
            Ref("images"),
    }
}
