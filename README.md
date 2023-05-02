In an attempt to level up our Go skills Ed had a great idea about taking something on to learn together, and we wanted to extend an invitation to the team. We decided on a project-based approach to creating REST servers in Go using multiple methods.

The project aims to implement the following RESTful routes below using different approaches. 
```
POST   /task/              :  create a task, returns ID
GET    /task/<taskid>      :  returns a single task by ID
GET    /task/              :  returns all tasks
DELETE /task/<taskid>      :  delete a task by ID
GET    /tag/<tagname>      :  returns list of tasks with this tag
GET    /due/<yy>/<mm>/<dd> :  returns list of tasks due by this date
```

We will loosely follow [this post](https://eli.thegreenplace.net/2021/rest-servers-in-go-part-1-standard-library/), splitting our studies into 7 parts:
1. Only using the standard library
2. Using the router package
3. Using a web framework
4. Using OpenAPI/Swagger (We might skip this since we don't use either)
5. Using middleware 
6. Authentication
7. Using GraphQL

**The idea**
1. Anyone wanting to participate will do as much or as little to address part 1.
2. We will all hop on a call Friday(time TBD to accommodate the most people) to discuss the current part, issues, share, and generally talk about Go.
3. Rinse and repeat!
