race ?=
testfile ?=
testname ?=

git:
	git add .
	git commit -m "$m"
	git push

# run tests
# to run all tests, just run: make check
# to run a specific test file, example: make check testfile=./context/v3/
# to run a specific test name, example: make check testfile=./context/v3/ testname=TestServer
# if you want to run these tests with a race flag, put -race=race, example: make check testfile=./context/v3/ testname=TestServer -race=race
check:
	./hack/builder-check.sh testfile=$(testfile) testname=$(testname) race=$(race)