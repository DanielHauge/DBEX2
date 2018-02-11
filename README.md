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

## TL:DR version:
Put this command in linux terminal (**NOT RECOMMENDED**)
```
wget -O - https://raw.githubusercontent.com/Games-of-Threads/DBEX2-DFH/master/SetupDBEX2.sh | bash
```


---------

## How to run!
### Pre-requisite
- Have a running debian distribution. (Ubuntu 16 or < is recommended)
>- If you do not have access to a linux. Follow this [guide](https://github.com/datsoftlyngby/soft2018spring-databases-teaching-material)
- (**OPTIONAL BUT RECOMMENDED**): Use the VM (vagrantfile) in this repository. A vagrantfile is provided at [link](https://raw.githubusercontent.com/Games-of-Threads/DBEX2-DFH/master/Vagrantfile). If using this vagrantfile. It has private network, so you can access stuff from inside at all ports on IP: 192.168.33.10. If IP allready taken on hostmachine, you can change the IP in the vagrantfile.
-------------

### Setup
To setup this exercise. simply be in the shell linux terminal, and put in this command. Note: It will install latest version of: Docker, Docker compose, wget and git.
**__RECOMMENDED__**
```
wget -O - https://raw.githubusercontent.com/Games-of-Threads/DBEX2-DFH/master/SetupDBEX2.sh | bash
```
If you believe, you allready have the newest docker, docker-compose, wget, git version installed. You can also use the express script instead.
```
wget -O - https://raw.githubusercontent.com/Games-of-Threads/DBEX2-DFH/master/ExpressSetup.sh | bash
```
This will setup the database and application. It can take a minute or two. (Longer if slow computer).

------------------------
------------------------

### How to Use.
The golang application has 2 interfaces. A command line interface, and a API endpoint. So it is up to you how you prefer to interact with systems. When the setup is done. It will display likely IP's you will need to go to, to access it's API Endpoint. It will also display the command to attach into the container to use the command line interface. which is:
```
docker attach GoApp
```
Note: If you attach to the container and leave it without correct escaping (CTRL+P -> CTRL-Q). The container will be deleted. You can run another run with the following command if this were to happen, and you need to get one up again without needing to run the entire setup again.
```
sudo docker run --rm -it --link dbms:mongo --publish=9191:9191 --name GoApp dbex2:dfh
```

Both CLI and API will point to either commands or routes for picking a questions to be answered.

---------------------

##### CLI Commands

Question #     | Question 1 | Question 2 | Question 3 | Question 4 | Question 5a | Question 5b
-------------- | ---------- | ---------- | ---------- | ----------- | ---------- | --------------
CommandToWrite | userscount | mostlinks | mostmentions | mostactive | mostgrumpy | mosthappy

---------------------

##### API Routes

Question #     | Question 1 | Question 2 | Question 3 | Question 4 | Question 5a | Question 5b
-------------- | ---------- | ---------- | ---------- | ----------- | ---------- | --------------
Route to use   | http://ip:9191/usercount | http://ip:9191/mostlinks | http://ip:9191/mentioned | http://ip:9191/mostactive | http://ip:9191/mostgrumpy | http://ip:9191/mosthappy

Note: If using the vagrantfile provided. the IP will be: 192.168.33.10:9191/route

---------------------
---------------------

### Expected answers

##### Question 1.
Expected result: 659621

Note: However, Robo 3T with the same exact quiry document. Will bring this result: **659774**.

This is most likely due to a difference in how each (golang library driver and robo) handle distinction.

##### Question 2.
Expected result: 

```
Name: lost_dog - Value: 549
Name: tweetpet - Value: 310
Name: VioletsCRUK - Value: 251
Name: what_bugs_u - Value: 246
Name: tsarnick - Value: 245
Name: SallytheShizzle - Value: 229
Name: mcraddictal - Value: 217
Name: Karen230683 - Value: 216
Name: keza34 - Value: 211
```

##### Question 3. 
Expected result:

```
Name: @mileycyrus - Value: 3770
Name: @tommcfly - Value: 3614
Name: @ddlovato - Value: 2903
Name: I - Value: 2618
Name: @DavidArchie - Value: 1089
Name: @Jonasbrothers - Value: 1063
Name: @DonnieWahlberg - Value: 1017
Name: @jordanknight - Value: 1005
Name: i - Value: 978
```

##### Question 4.
Expected result:
```
Name: lost_dog - Value: 549
Name: webwoke - Value: 345
Name: tweetpet - Value: 310
Name: SallytheShizzle - Value: 281
Name: VioletsCRUK - Value: 279
Name: mcraddictal - Value: 276
Name: tsarnick - Value: 248
Name: what_bugs_u - Value: 246
Name: Karen230683 - Value: 238
```

##### Question 5.
Expected results.
- A: Grumpy
```
Name: stephyway - Value: 0
Name: bobzipp - Value: 0
Name: tattoodancer35 - Value: 0
Name: paligurl93 - Value: 0
Name: adbillingsley - Value: 0
Name: sdancingsteph - Value: 0
Name: bpbabe - Value: 0
Name: RobFoxKerr - Value: 0
```
- B: Happy
```
Name: stephyway - Value: 4
Name: bobzipp - Value: 4
Name: tattoodancer35 - Value: 4
Name: paligurl93 - Value: 4
Name: adbillingsley - Value: 4
Name: sdancingsteph - Value: 4
Name: bpbabe - Value: 4
Name: RobFoxKerr - Value: 4
```


