# Wildfire - The Task
The task was to create a "Production Ready" web service which utilizes two existing web services:\
[Random name generator](https://names.mcquay.me/api/v0/)\
[Random Chuck Norris joke generator](http://joke.loc8u.com:8888/joke?limitTo=nerdy&firstName=John&lastName=Doe)

Use the the results of each and return that result to the user.

## Requirements
Some of the basic requirements of this task were
- Take 2-4 hours on the task
- Written in Go
- Clearly written README.md that at *Minimum* provides clear instructions on using the service
- Service should support concurrent requests and remain responsive under load

## Instructions for Use

The easiest way for you to use this web service would be to clone this repository
```
git clone https://github.com/Conor-Fleming/Wildfire-Task
```

From here you can simply use `$ go run main.go` and it will start the server.

Once running, you can use curl to make requests to ```localhost:8080```
```
$ curl localhost:8080                                                                                   
Clive Moodie programs occupy 150% of CPU, even when they are not executing.
```

Alternatively once the server has been started you can navigate to (http://localhost:8080/) and use refresh for new content.

## Testing
To test this program I created a simple client that used a wait group and go routine that sent off 1000 requests to the service at the same time. 
This test returned all successful status codes of 200. (yay)
