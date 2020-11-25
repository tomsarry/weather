# weather

A command line app to know the weather from any city. ☀️

## Getting started

Make sure you have Go installed on your machine (1.14+), and that $PATH leads to your binary go files.

```
go get github.com/tomsarry/weather
go install github.com/tomsarry/weather
```

## How to use it

* To know the weather of _city_, type the following into your terminal:

```
$ weather -c "city"
```

* You can also add a favorite location _city_ with the following (persistant through sessions):

```
$ weather -f "city"
```

Now, to quickly find the weather of your saved location you can simply do:

```
$ weather
```

### Example

For instance if you want to understand how cold it can get in the magnificent city of Montreal, use:

```
$ weather -c "Montreal"
Results for Montreal, CA:
Temperature: -5.4°C
Weather: Clouds - (overcast clouds)
```

or:

```
$ weather -f "Montreal"
New prefered city is Montreal.

$ weather
Using prefered city: Montreal.
Results for Montreal, CA:
Temperature: -5.4°C
Weather: Clouds - (overcast clouds)
```
