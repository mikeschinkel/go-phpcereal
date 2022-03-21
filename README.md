# go-phpcereal
Golang package for loading and writing serialized PHP with arbitrary schema


## Usage

```
p := phpcereal.NewParser(t.Value)
root, err := p.Parse()
```


## Requires

- go 1.18 - Generics

## License

- GNU Affero GPL v3.0