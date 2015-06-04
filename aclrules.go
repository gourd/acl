package acl

// ACLRules represents list of ACLRules
type ACLRules []ACLRule

func (rs *ACLRules) Select() (cout ACLRulesPipe) {
	ch := make(chan ACLRule)
	cout = ch
	go func() {
		for _, r := range *rs {
			ch <- r
		}
		close(ch)
	}()
	return cout
}

// ACLRulesPipe represet result filtered by user
type ACLRulesPipe <-chan ACLRule

// Filter filter out rules that doesn't pass the given test
func (rr ACLRulesPipe) Filter(t func(r ACLRule) bool) (cout ACLRulesPipe) {
	ch := make(chan ACLRule)
	cout = ch
	go func() {
		for r := range rr {
			if t(r) {
				ch <- r
			}
		}
		close(ch)
	}()
	return cout
}

// HasActor filter out rules that doesn't concern any of the given actors
func (rr ACLRulesPipe) HasActor(as ...ACLActor) ACLRulesPipe {
	return rr.Filter(func(r ACLRule) bool {
		for _, a := range as {
			if r.Actor == string(a) {
				return true
			}
		}
		return false
	})
}

// HasScope filter out rules that doesn't concern any of the given scopes
func (rr ACLRulesPipe) HasScope(ss ...ACLScope) ACLRulesPipe {
	return rr.Filter(func(r ACLRule) bool {
		for _, s := range ss {
			if r.Scope == string(s) {
				return true
			}
		}
		return false
	})
}

// HasTarget filter out rules that doesn't concern any of the given targets
func (rr ACLRulesPipe) HasTarget(ts ...ACLTarget) ACLRulesPipe {
	return rr.Filter(func(r ACLRule) bool {
		for _, t := range ts {
			if r.Target == string(t) {
				return true
			}
		}
		return false
	})
}

// Exists test if any result is coming out of the range
func (rr ACLRulesPipe) Exists() bool {
	t := false

	// find the first result, if there is any
	for _ = range rr {
		t = true
		break
	}

	// receive all results and release memory
	go func() {
		for _ = range rr {
		}
	}()

	return t
}

// All returns all piped rules as a new ACLRules
func (rr ACLRulesPipe) All() (rs ACLRules) {
	for r := range rr {
		rs = append(rs, r)
	}
	return
}
