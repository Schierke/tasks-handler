## GENERAL

The general apps following some dumbed-down hexagonal architecture with 4 main parts:

- Handler: the inbound, entry point of app (controller). Only one inbound HTTP is using here for now
- Service: Contains logics of our application.
- Repository: Outbound, here our only outbound is the Database (mongoDB) we connecting with.
- Models: contains entities, data models that we interact with


## Technical choices

- MongoDB: reason is because the dataset is really large and I don't see any many-many relationship so it's a good choice. Also you guys are using MongoDB so I want to try it to see.

- I use chi for HTTP because of its simplicity. The ease of adding middleware is also nice. For example:

```r.Use(handler.auth.Authenticator)``` before route to running authentication middleware.




## Improvments

We can improve many aspects of this little application:

- Test: No test (sorry :( ). I'm not too familliar with MongoDB so taking a lot of times looking into MongoDB docs. I include some mocks for fun but don't have the time to implement it.

- My MongoDB queries can be better.

- Authen (& author): the authorization can be simple bcrypt stocking on a DB. For authentification I prefer using go-jwt to create an JWT token (access & refresh). We can employ a redis cache system for refresh token. 

- Separating the usecase by folder should be better when project becomes large

## Scalability

- Sharding the database: we can have database for some of the most frequent "type" of job.

