## GENERAL

The general apps following some dumbed-down hexagonal architecture with 4 main parts:

- Handler: the inbound, entry point of app
- Service: Contains logics of our application.
- Repository: Outbound, here our only outbound is the Database (mongoDB) we connecting with.
- Models: contains entities, data models that we interact with


## Technical choices

- MongoDB: reason is because the dataset is really large and I don't see any many-many relationship so it's a good choice. Also you guys are using MongoDB so I want to try it to see.

- I use chi for HTTP because of its simplicity.


## Improvments

We can improve many aspects of this little application:

- Test: No test (sorry :( ). I'm not too familliar with MongoDB so taking a lot of times looking into MongoDB docs. I include some mocks for fun but don't have the time to implement it.

- Maybe my MongoDB can be better.

## Scalability

- Sharding the database: we can have database for some of the most frequent "type" of job.

