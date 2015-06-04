package acl

import (
	"testing"
)

func testACLRules() *ACLRules {
	return &ACLRules{
		ACLRule{
			Actor:  "actor:1",
			Target: "target:1",
			Scope:  "scope:1",
		},
		ACLRule{
			Actor:  "actor:2",
			Target: "target:2",
			Scope:  "scope:2",
		},
		ACLRule{
			Actor:  "actor:1",
			Target: "target:3",
			Scope:  "scope:1",
		},
		ACLRule{
			Actor:  "actor:2",
			Target: "target:4",
			Scope:  "scope:2",
		},
		ACLRule{
			Actor:  "actor:5",
			Target: "target:5",
			Scope:  "scope:5",
		},
	}
}

func TestACLRulesSubset_HasScope(t *testing.T) {
	results := testACLRules().Select().HasScope("scope:1", "scope:2").All()
	if len(results) != 4 {
		t.Errorf("Unexpected result set size %d: %#v", len(results), results)
	}
}

func TestACLRulesSubset_HasTarget(t *testing.T) {
	results := testACLRules().Select().HasTarget("target:1", "target:2").All()
	if len(results) != 2 {
		t.Errorf("Unexpected result set size %d: %#v", len(results), results)
	}
}

func TestACLRulesSubset_HasActor(t *testing.T) {
	results := testACLRules().Select().HasActor("actor:1", "actor:2").All()
	if len(results) != 4 {
		t.Errorf("Unexpected result set size %d: %#v", len(results), results)
	}
}
