//go:generate gourd gen service -type=ACLRule -coll=acl_rules $GOFILE
package acl

// ACLBucket represents object or container which
// is the receiver of a certain action
type ACLTarget string

// ACLUser represents certain type of user / program
// who is the actor of a certain action
type ACLActor string

const (
	AllUsers              ACLActor = "user:all"
	AllAuthenticatedUsers ACLActor = "user:authenticated"
)

// ACLScope represents the permission scope
// of the actor can act on target. (e.g. READER, WRITER, OWNER)
type ACLScope string

const (
	ScopeOwner  ACLScope = "OWNER"
	ScopeReader ACLScope = "READER"
	ScopeWriter ACLScope = "WRITER"
)

// ACLRule represents an access control list rule entry
// which describe the permission scope which certain
// actor can act on the target
type ACLRule struct {
	Id     int64  `db:"id,omitempty" json:"-"`
	Actor  string `db:"actor" json:"actor"`
	Target string `db:"target" json:"target"`
	Scope  string `db:"scope" json:"scope"`
}
