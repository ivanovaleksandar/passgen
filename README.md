# Password generator

A small utility tool for generating passwords

# Build
```
go build 
```

## Usage

Generate a 16 characters long password

```
passgen -s 16
```

Exclude usage of symbols
```
passgen -y=false
```

Include only a subset of letters
```
passgen -L="aeiou"
```

List all available options using `--help`