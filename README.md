# Database exercise 2 (Document Oriented Database)
This is a mini project for the database course. The excersise is based off the description in the buttom of this [resource](https://github.com/datsoftlyngby/soft2018spring-databases-teaching-material/blob/master/lecture_notes/02-Intro_to_MongoDB.ipynb)

## Description
This exercise goal is to build a simple application that can interact with a Document Oriented Database. The exercise involves having a database with preexisting data (from csv file) where the goal is to be able to answer 5 questions which the application shall have to answer. The questions are as follows.
- How many Twitter users are in the database?
- Which Twitter users link the most to other Twitter users? (Provide the top ten.)
- Who is are the most mentioned Twitter users? (Provide the top five.)
- Who are the most active Twitter users (top ten)?
- Who are the five most grumpy (most negative tweets) and the most happy (most positive tweets)? (Provide five users for each group)


The result is a One-line script setup, that be run on a linux machine. It will setup a MongoDB database and import the data to it and host it in a docker container, it also will build and run a golang application and run it in a container. 

-----------

## How to run!
### Pre-requisite
- Have a running debian distribution. (Ubuntu 16 or < is recommended)
>- If you do not have access to a linux. Follow this [guide](https://github.com/datsoftlyngby/soft2018spring-databases-teaching-material) to get familiar with vagrant. A vagrantfile is provided at [link], it is in the repository aswell.

### Setup
To setup this exercise. simply be in the shell linux terminal, and put in this command.
**__RECOMMENDED__**
```
wget -O - https://raw.githubusercontent.com/Games-of-Threads/DBEX2-DFH/master/SetupDBEX2.sh | bash
```
If you believe, you allready have the newest docker, docker-compose, wget, git version installed. You can also use the express script instead.
```
wget -O - https://raw.githubusercontent.com/Games-of-Threads/DBEX2-DFH/master/ExpressSetup.sh | bash
```
This will setup the database and application. It can take a minute or two. (Longer if slow computer).

### How to Use.
The golang application has 2 interfaces. A command line interface, and a API endpoint. So it is up to you how you prefer to interact with systems. When the setup is done. It will display likely IP's you will need to go to, to access it's API Endpoint. It will also display the command to attach into the container to use the command line interface. which is:
```
docker attach GoApp
```
Both CLI and API will point to either commands or routes for picking a questions to be answered.

##### CLI Commands

Question # | Question 1 | Question 2 | Question 3 | Question 4 | Question 5a | Question 5b
---------- | ---------- | ---------- | ---------- | ----------- | --------
CommandToWrite | userscount | mostlinks | mostmentions | mostactive | mostgrumpy | mosthappy

