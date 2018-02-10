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

## How to run it!
