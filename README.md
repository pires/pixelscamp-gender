# pixels-gender

I wanted to know the gender ratio for PixelsCamp 2016. I started by finding lists of names on the web, and while it was easy
to find for UK and US, the same didn't apply for Portugal. Thanks to [Publico](https://publico.pt) and [Central de Dados](https://github.com/centraldedados/nomes_proprios)
for helping with the latter.

## Dependencies

You must have Go installed, obviously. Then:
```
go get -u github.com/jbrukh/bayesian
go get -u github.com/hstove/gender/classifier
```

## Run

```
curl "https://api.pixels.camp/users/?count=1000" > users.json
go run main.go
```

## Train (optional)

```
cd classifier
go run classifier.go
```

## TODO

* Retrieve directly from API instead of JSON file
