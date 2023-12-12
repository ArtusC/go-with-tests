* Learn Go with tests:
  * https://quii.gitbook.io/learn-go-with-tests/

  ---------------------------------------------------------------

* About go modules and start a GO aplication
   * https://github.com/golang/go/wiki/Modules#example

---------------------------------------------------------------

Now, we can run the tests separately, as described into makeflile:

* to run all tests, just run: 
   * ```make check```
*  to run a specific test file, example: 
   * ```make check testfile=./context/v3/```
*  to run a specific test name, example: 
   * ```make check testfile=./context/v3/ testname=TestServer```
* if you want to run these tests with a race flag, put -race=race, example: 
   * ```make check testfile=./context/v3/ testname=TestServer -race=race```
