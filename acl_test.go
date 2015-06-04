package acl

import (
	"github.com/gourd/service/upperio"
	"net/http"
	"testing"
	"upper.io/db/sqlite"
)

func TestACLRulePut(t *testing.T) {

	// define db
	upperio.Define("default", sqlite.Adapter, sqlite.ConnectionURL{
		Database: `./_test/acl_test.db`,
	})

	var err error
	r := &http.Request{}

	// test parameters
	var tcontainer ACLTarget = "node_type:test_put"
	var ttarget ACLTarget = "node:test_put:1234"
	var tactor ACLActor = "role:admin"
	var tscope1 ACLScope = "OWNER1"
	var tscope2 ACLScope = "OWNER2"

	// create first dummy rule
	err = PutACLRule(r, ttarget, tactor, tscope1)
	if err != nil {
		t.Errorf(err.Error())
	}

	// test retrieve dummu rule
	rules, err := ACL(r, tcontainer, ttarget)
	if err != nil {
		t.Errorf(err.Error())
	} else if !rules.Select().HasActor(tactor).HasScope(tscope1).Exists() {
		t.Errorf("Inserted rule not found in %#v", rules)
	}

	// update dummy rule
	err = PutACLRule(r, ttarget, tactor, tscope2)
	if err != nil {
		t.Errorf(err.Error())
	}

	// test again to retrieve dummy rule
	rules, err = ACL(r, tcontainer, ttarget)
	if err != nil {
		t.Errorf(err.Error())
	} else if rules.Select().HasActor(tactor).HasScope(tscope1).Exists() {
		t.Errorf("Modified rule still found in %#v", rules)
	} else if !rules.Select().HasActor(tactor).HasScope(tscope2).Exists() {
		t.Errorf("Inserted rule not found in %#v", rules)
	}

}

func TestACLRuleDelete(t *testing.T) {
	// define db
	upperio.Define("default", sqlite.Adapter, sqlite.ConnectionURL{
		Database: `./_test/acl_test.db`,
	})

	var err error
	r := &http.Request{}

	// test parameters
	var tcontainer ACLTarget = "node_type:test_delete"
	var ttarget ACLTarget = "node:test_delete:1234"
	var tactor ACLActor = "role:admin"
	var tscope ACLScope = "OWNER"

	// create dummy rule
	err = PutACLRule(r, ttarget, tactor, tscope)
	if err != nil {
		t.Errorf(err.Error())
	}

	// test retrieve dummu rule
	rules, err := ACL(r, tcontainer, ttarget)
	if err != nil {
		t.Errorf(err.Error())
	} else if !rules.Select().HasTarget(ttarget).Exists() {
		t.Errorf("Inserted rule not found in %#v", rules)
	}

	// delete the dummy rule
	err = DeleteACLRule(r, ttarget, tactor)

	// test retrieve dummu rule
	rules, err = ACL(r, tcontainer, ttarget)
	if err != nil {
		t.Errorf(err.Error())
	} else if rules.Select().HasTarget(ttarget).Exists() {
		t.Errorf("Failed to delete dummy rule in %#v", rules)
	}

}
