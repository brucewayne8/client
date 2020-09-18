# client
A client side application which sends requests through terminal interface, which is similar to curl and written in go language.


### Prerequisites:

go  - 1.13.8

### steps to follow for linux users:

- set GOBIN to pwd/bin by ```export GOBIN=path``` here path is path of directory where executable files needed to be stored and has permissions
- Make sure path variable can access GOBIN
- go install client.go

### steps to follow for windows users:

- Add GOBIN variable to system variables with path where exe files needed to be stored
- Add the same path to path variable in user variables
- go install client.go
### format for GET call :

    client GET {url}

    Example:  ```client GET "https://jsonplaceholder.typicode.com/posts/1"```

### format for POST call: 

    client POST {url} data

    format of data is "field1=value1&&field2=value2&&field3=value3"

    Example:   ```client POST "https://jsonplaceholder.typicode.com/posts/" "title=Brucewayne&&description=Batman&&Status=Single"```


