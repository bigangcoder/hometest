#Human oriented README explaining your project's purpose, how to set it up, run tests,
#code linting, start contributing


# QUICK ONLINE API DOCS (can access directly on your web browser)
https://documenter.getpostman.com/view/13738369/UVsFzovo


# IMAGE LOCAL TEST (need to download image from docker hub)
## download image:(centos amd64x  version)
$ docker pull bigangcoder/hometest:0.91
## run image
$ docker run -p 9091:8080 bigangcoder/hometest:0.91

# SOURCE CODE LOCAL TEST(nedd to download source code from github)
https://github.com/bigangcoder/hometest



## project purpose  :three user stories that should be covered:

I want to be able to send a payment from one account to another (same currency)
I want to be able to see all payments
I want to be able to see available accounts

## settup database connection config file
<(project root path)> /hometest.globalgroup.com/backend/config/base_config.json

## local api docs 

<(project root path)>/docs/api.md

## Postman online api docs (for local runnable version   ex: 127.0.0.1 )

https://documenter.getpostman.com/view/13738369/UVsFy8J2

## database init sql file 

<(project root path)>/docs/init.sql

## how to run it up (in project root path)

go run main.go 

## how to run tests (in project root path)

go test -v 






