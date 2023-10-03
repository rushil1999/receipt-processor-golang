# receipt-processor-golang
* This is a web app developed completely in Golang. Gin web framework has been used for routing purposes, along with other packages for utilities.
* The structure of the project has been made such that it is production-ready and can be expanded to add more functionality.
* The folder has 3 sub-folders and a main.go file:
  * api: This folder contains the routes and the controller function
  * helpers: This folder has all the utilities functions that help in executing logic
  * models: This folder has structs used to represent real-life entities like receipts, items, etc
  * main.go: This file starts the server and imports routes

# How to start locally
* Install Golang into your system ( I have used version 1.21.1 )
* Go to root directory `cd receipt-processor`
* run `go run .`
* This will start the server locally on port 8000 (Make sure that port is not used by any other application)

# Routes
* GET `/receipts`: Returns all the receipts in indented JSON
* POST `/receipts/process`: Will take JSON body and return a JSON object with 1 key `id` depicting the id assigned to the receipt
* GET `receipts/{id}/process`: Will return the points assigned to receipt

# Key Features
* The project is well-organized and modular
* Error Handling is added, so invalid input will returned with an appropriate error message
* Points are stored with receipts. So once a receipt is calculated for points, it will NOT be re-calculated
   
