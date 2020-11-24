# weather

A command line app to know the weather from any city. ☀️

## Getting started

Make sure you have Go installed on your machine (1.14+).

```
go get github.com/tomsarry/weather
go install github.com/tomsarry/weather/cmd/weather
```

## How to use it

To know the weather of _city_, type the following into your terminal:

```
$ weather -c "city"
```

### Example

For instance if you want to understand how cold it can get in the magnificent city of Montreal, use:

```
$ weather -c "Montreal"
```

and you should see something like this printed on your screen:

```
Results for Montreal, CA:
Temperature: -45.6
Weather: Clouds - (overcast clouds)
```
