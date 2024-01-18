# Goal:
We will create an application which listens to HTTP requests and responds to them.

# Description:
You have been asked to create a web server where users can track how many games players have won.

* `GET /players/{name}` should return a number indicating the total number of wins

* `POST /players/{name}` should record a win for that name, incrementing for every subsequent POST


We will follow the TDD approach, getting working software as quickly as we can and then making small iterative improvements until we have the solution. 


By taking this approach we:

* Keep the problem space small at any given time

* Don't go down rabbit holes

* If we ever get stuck/lost, doing a revert wouldn't lose loads of work.

# Chapter:
https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server

# Running:

* Run this a few times, change the player names if you like:

	`curl -X POST http://localhost:5000/players/<SOME_NAME>`

* Check scores with curl:

	`curl -X GET http://localhost:5000/players/<SOME_NAME>`
