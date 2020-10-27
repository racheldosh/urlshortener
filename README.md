What is this project?
----
This was a take-home assessment for CogoLabs to build a url shortener. The language was go, which I had never worked with before. I also had never worked with docker. It was a learn-something-new project. 

How to run
----
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

To run program: 

Pull from docker:

	docker pull racheldoshcollins/url_shortener

From the hello folder: 

	docker build -t image-name .
	docker run --publish 8080:8080 --name test --rm image-name
	
	Then, to stop, run from a new terminal:

		docker stop test

Or, from the hello folder (to run without docker):

	go build hello.go
	./hello

Then the app can be viewed at "localhost:8080"
Once a shortenedURL is retrieved, using "localhost:8080/url/<shortenedURL>" 
	will redirect to the requested URL

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Bugs and future works
---------

Bugs:

1. The app doesn't validate the URL entered to be shortened. If the user enters "test" to be shortened, "test" will be added to the database as a URL. Then, when the user calls localhost:8080/url/<shortenedURL>" an error occurs when trying to redirect to the faulty url "test."

2. The app doesn't check to see if a URL already exists in the database. 




Future work: 

1. Address the known bugs above (primarily, focus on validating urls)

2. Adding testing (may help with #1 :D)

3. Style -- was is convention for Go?
