# Goal:
We will make our endpoints return JSON and explore how to do routing.

# Description:
Our product owner has a new requirement; to have a new endpoint called /league which returns a list of all players stored. She would like this to be returned as JSON.

# Chapter:
https://quii.gitbook.io/learn-go-with-tests/build-an-application/json

# Running:

* On the **json-routing-embbeding-version/v3** folder, build the application and run it:
	`go build . && ./json-routing-embedding-version`

* Run this a few times, change the player names if you like:

	`curl -X POST http://localhost:5000/players/<SOME_NAME>`

* Check scores with curl:

	`curl -X GET http://localhost:5000/players/<SOME_NAME>`

* Check the league form players with curl:

	`curl -X GET http://localhost:5000/league`