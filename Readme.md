## go-kit DDD
Domain Driven Design is prevelent and rising standard for organizing your microservice code. This design architecture emphasis on Code organization and separation of concern

The above code is organized in following packages
- `/api` - contains api url routes
- `/sessionmanager` - single module that handle single responsibility i.e manage sessions of app
- `/strconv` - single module that handle single responsibility i.e transform string [Follow the tutorial here](https://gokit.io/examples/stringsvc.html)


Middlewares