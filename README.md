# Rohit-Sharma-Coding-Challenge

This is web API interface from sochain to get Block details and Transaction details

Getting Started
---------------

1. Install these dependencies.

    - Go 1.17.x
    - GNU make
    - Docker and Docker Compose

Make Tasks
----------

    build                # go build
    down                 # brings down dependencies
    up                   # bring up dependencies in a compose stack

# How to use 
To run the application in a container just run `make up` and to bring down the container just run `make down`.

There are 2 endpoints exposed in this application which are as follow

**To get Block details**

    /currencies/{crypto}/block-details/{blockHash}

* **GET**

* **URL Params**

  **Required:**

  `crypto=[string]`

   `blockHash=[string]`

**To get Transaction details**

    /currencies/{crypto}/tx-details/{transactionId}

* **GET**

* **URL Params**

  **Required:**

  `crypto=[string]`

  `transactionId=[string]`



## Logging
As of now we are using the gommon logger but we can create a wrapper around the logger to standardize the logging. 

## Caching
We can cache some information that doesn't change frequently for eg: Transaction details for 
tx Id: dbaf14e1c476e76ea05a8b71921a46d6b06f0a950f17c5f9f1a03b8fae467f10 remains same always so we can use go-redis to save that.

## Error handling
We can create custom error types that are specific to our application.
