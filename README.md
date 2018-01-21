# N5 Web Application

N5 is a web application that displays the next five races that can be bet on, sorted by the closing date.
When a race is closed, you can no longer bet or view that race.

To view the details of the race, click on the race number and a new page will load displaying the race type, time, competitors and other details.

## Getting Started

Before running the N5 Application make sure you have installed the correct dependencies, programs and have correctly set up the database connection. If this is not done, erroneous results may occur.

#####You have been warned /s.


### Prerequisites

The main programs required to run this project are

- A Golang IDE (Goland is my fave)
- Docker
- A database visualiser tool (not necessary but makes life so much better)
- Some sort of command line tool (is it possible to not have one?)
- PostgreSQL

### Installing

Follow this guide to install and run the project.

1. Install IDE of choice
2. Set GOPATH and GOROOT variables
3. To install all required packages and dependencies that you do not have, on your command line tool navigate to the N5 Web App and run ``` go get -d ./... ```
4. Install Docker
5. Run docker and either ``` docker pull postgres ``` or open up your command line tool, navigate to the project and then run ``` docker-compose -f stacker.yml up ```
6. Once database container is running, open up database visualiser tool and connect
7. In your IDE of choice, make sure you have set up your database to postgreSQL and you connect to the IP and Port shown by the docker container
8. Start the database connection in your IDE
9. Open tables.sql in the sql folder of the project and run the database creation statements within
10. Check your database visualisation tool or run SQL commands to make sure the tables have been created correctly
11. Run main.go in your IDE
12. Open up a web browser and navigate to ```127.0.0.1:8080``` or ```localhost:8080```

###And there you have it! You are now successfully running the application. 

## Navigating the web page information

####Main Page
On the main page that is loaded on ```127.0.0.1:8080``` or ```localhost:8080``` you will find the next five races with the smallest time remaining til closing in comparison to your current time. 

It displays the race number, the race type, the race location, it's closing time, and the time remaining (in seconds) until the race is closed.

###Race Details Page

To view the details of each race, click on the race number and a new page will load displaying the race information.

The race details page displays the race location, race type, the competitors names, their starting positions, the closing time, and the time remaining until the race is closed (in seconds).

## Built With

* [Goland](https://www.jetbrains.com/go/) - Golang IDE used
* [Docker](https://www.docker.com/) - The software containerization platform
* [PostgreSQL](https://www.postgresql.org/) - Database used
* [SQLBoiler](https://github.com/volatiletech/sqlboiler) - Database ORM tool
* [Gin Gonic](https://github.com/gin-gonic/gin) - Web framework

## Authors

* **Sir Professor Lord David Gray** - *Only the best* - [Subrising](https://github.com/subrising)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Me for being awesome
* StackOverflow for helping me remember things I've forgotten
* LadBrokes for giving me this opportunity
* Google for obvious reasons
