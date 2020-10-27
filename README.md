What is this project?
----
This was a take-home assessment for CogoLabs to build a url shortener. The language was go, which I had never worked with before. I also had never worked with docker. It was a learn-something-new project. 

While running the app, you can enter a valid URL into the prompt at localhost:8080. It'll return a shortened url, for example "url/1". Where from "localhost:8080/url/1" you'll be redirected to your url.

How to run
----
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

To run program, download the project. With docker running, inside the project (urlshortener folder), run: 

	docker build -t image-name .
	docker run --publish 8080:8080 --name test --rm image-name

The app can be viewed at "localhost:8080"
Once a shortenedURL is created, using "localhost:8080/url/<shortenedURL>" 
	will redirect to the requested URL
	
Then, to stop running the app, run from a new terminal:

	docker stop test

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Bugs and future work
---------

Bugs:

1. The app doesn't validate the URL entered to be shortened. If the user enters "test" to be shortened (not a valid url). "test" will be added to the database as a URL. Then, when the user calls localhost:8080/url/<shortenedURL>" an error occurs when trying to redirect to the faulty url "test".

2. The app doesn't check to see if a URL already exists in the database. 



Future work: 

1. Address the known bugs above (primarily, focus on validating urls that are enterred, both for format and for pre-existance)

2. Adding testing (learn how to write tests in Go)

3. Style -- was is convention for Go?
