# Go Library Management API

## Tech Stack - Go

## How to Run the server

 Run ***"go run main.go"*** in the directory where main.go is located
 
### Endpoints

* Visit  ***" localhost:8000/api/ "***  for Total Library Details

* Visit ***" localhost:8000/api/booksAvailable "*** for all the books that are available
* Visit ***" localhost:8000/api/bookAvailable?book=[Book name] "*** to check if book is available
  for example localhost:8000/api/bookAvailable?book=Moby+Dick 

* Visit ***" localhost:8000/api/MostIssued "*** to get the most issued book

* Visit ***" localhost:8000/api/IssuedTo?book=[Book name] "*** to Check the user to which the book is issued

* Visit ***" localhost:8000/api/TopTrending "*** to get the most trending book
