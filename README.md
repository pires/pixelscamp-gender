# pixelscamp-gender

I wanted to know the gender ratio for PixelsCamp 2016. I started by finding lists of names on the web, and while it was easy
to find for UK and US, the same didn't apply for Portugal. Thanks to [Publico](https://publico.pt) and [Central de Dados](https://github.com/centraldedados/nomes_proprios)
for helping with the latter.

**Results as of Saturday, 24th of September, 2016**:
```
Pixels:
 > Total: 863
 > Female: 88 (10.20%)
 > Male: 775 (89.80%)
 > Unknown: 0 (-0.00%)
```

## Dependencies

You must have Go installed, obviously. Then:
```
go get -u github.com/jbrukh/bayesian
go get -u github.com/hstove/gender/classifier
```

## Run

```
go run main.go
```

## Train (optional)

```
cd classifier
go run classifier.go
```
