# Orgmanager

## Usage

```sh
go run main.go -t "GITHUB TOKEN" -f file.json
```

The configuration file :
```json
{
	"labels": [{
		"name": "",
		"color": "",
		"repositories": [
			""
		]
	}],
	"organisation": ""
}

```

|Field|Type|Definition|
|---|---|---|
|name|String|Label name.|
|color|String|Label color in hexadecimal.
|repositories|Array of String|Regular expressions for specifying what repository will be affected by this label.|
