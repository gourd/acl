package acl

import (
	"github.com/gourd/service"
	"net/http"
)

// ACL returns the ACL entries concerning the named target
func ACL(r *http.Request, container, target ACLTarget) (rules ACLRules, err error) {

	// get service
	s := service.Providers.MustService(r, "ACLRule")

	// search container rules
	crules := []ACLRule{}
	cq := service.NewQuery().AddCond("target", container)
	result, err := s.Search(cq)
	if err != nil {
		return
	}

	err = result.All(&crules)
	if err != nil {
		return
	}
	for _, r := range crules {
		rules = append(rules, r)
	}

	// search target rules
	trules := []ACLRule{}
	tq := service.NewQuery().AddCond("target", target)
	result, err = s.Search(tq)
	if err != nil {
		return
	}

	err = result.All(&trules)
	if err != nil {
		return
	}

	for _, r := range trules {
		rules = append(rules, r)
	}

	return
}

// PutACLRule saves the named ACL entity with the provided actor, scope for the named target.
func PutACLRule(r *http.Request, target ACLTarget, actor ACLActor, scope ACLScope) (err error) {

	// get service
	s := service.Providers.MustService(r, "ACLRule")

	// search existing rule
	rule := &ACLRule{}
	conds := service.NewConds()
	conds.Add("target", string(target))
	conds.Add("actor", string(actor))
	err = s.One(conds, rule)

	// should test if err is Not Found
	// need to turn error into singleton
	if err != nil {
		err = s.Create(service.NewConds(), &ACLRule{
			Actor:  string(actor),
			Target: string(target),
			Scope:  string(scope),
		})
	} else {
		rule.Actor = string(actor)
		rule.Target = string(target)
		rule.Scope = string(scope)
		err = s.Update(conds, rule)
	}

	return
}

// DeleteACLRule deletes the named ACL entity for the named object.
func DeleteACLRule(r *http.Request, target ACLTarget, actor ACLActor) (err error) {

	// get service
	s := service.Providers.MustService(r, "ACLRule")

	// delete the rules
	conds := service.NewConds()
	conds.Add("target", string(target))
	conds.Add("actor", string(actor))
	s.Delete(conds)

	return
}
