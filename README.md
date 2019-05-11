# Optimization of Trafficlight-Circuits

Group: ateam

### Build status / Test-Coverage

develop:
[![Build Status](https://travis-ci.com/ob-algdatii-ss19/leistungsnachweis-ateam.svg?token=4zw9EzexndWUV9DTxZpz&branch=develop)](https://travis-ci.com/ob-algdatii-ss19/leistungsnachweis-ateam)

master:
[![Build Status](https://travis-ci.com/ob-algdatii-ss19/leistungsnachweis-ateam.svg?token=4zw9EzexndWUV9DTxZpz&branch=master)](https://travis-ci.com/ob-algdatii-ss19/leistungsnachweis-ateam)

Determine testcoverage on your local machine and display results in the browser:
```
go test ./... -v -coverpkg=./... -coverprofile=cover.out
go tool cover -html=cover.out
```

### Heroku

The develop- and master-branch will be automatically deployed to heroku.

[heroku-master](https://leistungsnachweis-ateam.herokuapp.com): ![Heroku](https://heroku-badge.herokuapp.com/?app=leistungsnachweis-ateam)

[heroku-develop](https://leistungsnachweis-ateam-dev.herokuapp.com): ![Heroku](https://heroku-badge.herokuapp.com/?app=leistungsnachweis-ateam-dev)

### Getting started

* download github-Projekt
* run main-funcition in main.go
* open webserver in your browser ```localhost:8080```

*Hint:* To see javascript logging, open the developertools in your browser.
In the console you can find the logging output of the js-scripts.


### Project structure

```
 |-- backend
 |   |-- adjGraph
 |   |-- algorithms
 |-- frontend (includes all the gui-things)
 |   |-- images  (svg image)
 |   |-- scripts (javascript files)
 |   |-- styles  (css files)
 |   |-- index.html
 |   |-- intersection.html
 |   |-- result.html
 |-- visualizeGraph (tool to display json graph in the browser)
 |-- main.go
```
