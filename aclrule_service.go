// Generated by gourd (version 0.1dev)
// Generated at 2015/10/25 01:37:43 (+0800)
// Note: If you want to re-generate this file in the future,
//       do not change it.

package acl

import (
	"github.com/gourd/service"
	"github.com/gourd/service/upperio"
	"net/http"

	"log"
	"upper.io/db"
)

func init() {
	// define service provider with proxy
	service.Providers.DefineFunc("ACLRule", func(r *http.Request) (s service.Service, err error) {
		return GetACLRuleService(r)
	})
}

// GetACLRuleService provides raw ACLRuleService
func GetACLRuleService(r *http.Request) (s *ACLRuleService, err error) {

	// obtain database
	db, err := upperio.Open(r, "default")
	if err != nil {
		return
	}

	// define service and return
	s = &ACLRuleService{db}
	return
}

// ACLRuleService serves generic CURD for type ACLRule
// Generated by gourd CLI tool
type ACLRuleService struct {
	Db db.Database
}

// Create a ACLRule in the database, of the parent
func (s *ACLRuleService) Create(
	cond service.Conds, ep service.EntityPtr) (err error) {

	// get collection
	coll, err := s.Coll()
	if err != nil {
		return
	}

	// apply random uuid string to string id

	//TODO: convert cond into parentkey and
	//      enforce to the entity

	// add the entity to collection

	id, err := coll.Append(ep)

	if err != nil {
		log.Printf("Error creating ACLRule: %s", err.Error())
		err = service.ErrorInternal
		return
	}

	// apply the key to the entity
	e := ep.(*ACLRule)
	e.Id = int64(id.(int64))

	return
}

// Search a ACLRule by its condition(s)
func (s *ACLRuleService) Search(
	q service.Query) (result service.Result, err error) {

	// get collection
	coll, err := s.Coll()
	if err != nil {
		return
	}

	// retrieve entities by given query conditions
	var res db.Result
	conds := upperio.Conds(q.GetConds())
	if conds == nil {
		res = coll.Find()
	} else {
		res = coll.Find(conds)
	}

	// handle paging
	if q.GetOffset() != 0 {
		res = res.Skip(uint(q.GetOffset()))
	}
	if q.GetLimit() != 0 {
		res = res.Limit(uint(q.GetLimit()))
	}

	result = upperio.NewResult(res)
	return
}

// One returns the first ACLRule matches condition(s)
func (s *ACLRuleService) One(
	c service.Conds, ep service.EntityPtr) (err error) {

	// retrieve results from database
	l := &[]ACLRule{}
	q := service.NewQuery().SetConds(c)
	res, err := s.Search(q)
	if err != nil {
		return
	}

	// dump results into pointer of map / struct
	err = res.All(l)
	if err != nil {
		return
	}

	// if not found, report
	if len(*l) == 0 {
		err = service.ErrorNotFound
		return
	}

	// assign the value of given point
	// to the first retrieved value
	(*ep.(*ACLRule)) = (*l)[0]
	return nil
}

// Update ACLRule on condition(s)
func (s *ACLRuleService) Update(
	c service.Conds, ep service.EntityPtr) (err error) {

	// get collection
	coll, err := s.Coll()
	if err != nil {
		return
	}

	// get by condition and ignore the error
	cond, _ := c.GetMap()
	res := coll.Find(db.Cond(cond))

	// update the matched entities
	err = res.Update(ep)
	if err != nil {
		log.Printf("Error updating ACLRule: %s", err.Error())
		err = service.ErrorInternal
	}
	return
}

// Delete ACLRule on condition(s)
func (s *ACLRuleService) Delete(
	c service.Conds) (err error) {

	// get collection
	coll, err := s.Coll()
	if err != nil {
		return
	}

	// get by condition and ignore the error
	cond, _ := c.GetMap()
	res := coll.Find(db.Cond(cond))

	// remove the matched entities
	err = res.Remove()
	if err != nil {
		log.Printf("Error deleting ACLRule: %s", err.Error())
		err = service.ErrorInternal
	}
	return nil
}

// AllocEntity allocate memory for an entity
func (s *ACLRuleService) AllocEntity() service.EntityPtr {
	return &ACLRule{}
}

// AllocEntityList allocate memory for an entity list
func (s *ACLRuleService) AllocEntityList() service.EntityListPtr {
	return &[]ACLRule{}
}

// Len inspect the length of an entity list
func (s *ACLRuleService) Len(pl service.EntityListPtr) int64 {
	el := pl.(*[]ACLRule)
	return int64(len(*el))
}

// Coll return the raw upper.io collection
func (s *ACLRuleService) Coll() (coll db.Collection, err error) {
	// get raw collection
	coll, err = s.Db.Collection("acl_rules")
	if err != nil {
		log.Printf("Error connecting collection acl_rules: %s",
			err.Error())
		err = service.ErrorInternal
	}
	return
}

// Close the database session that ACLRule is using
func (s *ACLRuleService) Close() error {
	return s.Db.Close()
}
