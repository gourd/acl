#
# This file is only for test
#
# To build / install this library, please use command:
# go get github.com/gourd/acl
#

test: pretest
	go test -v

pretest:
	sqlite3 _test/acl_test.db < _test/schema.sqlite3.sql

clean:
	rm -f _test/acl_test.db

$PHONY: test pretest clean
