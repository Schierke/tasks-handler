## Getting started
Before you start, make sure you have golang and docker installed with lastest version

For local execution, I recommend using task.

## INSTALL

### Using task

First installation: the easiest way is to do ```task local``` to launching mongoDB only. Then ```task migrate``` for creating collections. 

(Or if mongo is already launch, just task migrate is enough)

For second, third time, just ```task develop``` for launching everything 


### Using nothing (but still have docker)
First installation: ```docker-compose -f docker-compose.local.yml up --build``` for launching mongoDB only. Then build the app
Then ```go run main.go migrate``` for migrating the dataset

```docker-compose -f docker-compose.dev.yml up --build``` for running everything after that.

### Testing

You can import the postman file inside ```docs/postman``` for testing purpose.