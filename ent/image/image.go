// Code generated by entc, DO NOT EDIT.

package image

const (
	// Label holds the string label denoting the image type in the database.
	Label = "image"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldFileLocation holds the string denoting the filelocation field in the database.
	FieldFileLocation = "file_location"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldPublic holds the string denoting the public field in the database.
	FieldPublic = "public"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeImagerepos holds the string denoting the imagerepos edge name in mutations.
	EdgeImagerepos = "imagerepos"

	// Table holds the table name of the image in the database.
	Table = "images"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "images"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_images"
	// ImagereposTable is the table the holds the imagerepos relation/edge. The primary key declared below.
	ImagereposTable = "image_repo_images"
	// ImagereposInverseTable is the table name for the ImageRepo entity.
	// It exists in this package in order to avoid circular dependency with the "imagerepo" package.
	ImagereposInverseTable = "image_repos"
)

// Columns holds all SQL columns for image fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldFileLocation,
	FieldDescription,
	FieldPrice,
	FieldPublic,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Image type.
var ForeignKeys = []string{
	"user_images",
}

var (
	// ImagereposPrimaryKey and ImagereposColumn2 are the table columns denoting the
	// primary key for the imagerepos relation (M2M).
	ImagereposPrimaryKey = []string{"image_repo_id", "image_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
