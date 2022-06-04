# Pseudo redis implementation in go

For practice, I implemented a pseudo redis in go.


## Brief description

You can set, get, and delete a key-value database on memory
<br />create new database (by default its name is "default")
<br />get all database
<br />query based on regex on keys
<br />dump database onto .csv file. It will create a directory ./dumps and then create a file if not exists based on input
<br />load .csv onto memory
<br />For the scalability section, I split every 100 data into a section


## Run

To run, you should run `go run main.go` in the terminal or run `go build` in the main directory and then run `./redisak` in terminal

###
after
for set run
`set city_temp 18`<br />
for get run
`get city`<br />
for del run
`del city_temp`<br />
for keys run
`keys city_*`<br />
for use/add new database run
`use new_database`<br />
for get all database run
`list`<br />
for dump run
`dump default ./dumps/default`<br />
for load run
`load ./dumps/default default`<br />
for exit run
`exit`<br />


## requirements
go version 1.18
