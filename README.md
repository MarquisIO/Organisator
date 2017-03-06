# Organisator

A tool to help manage organisation-wide GitHub issue labels.

## Usage

```sh
go run main.go -t "GITHUB PERSONAL TOKEN" -f file.json
```

### Access token

Create a [Personal Access tokens](https://github.com/settings/tokens) and give a ```Full control of private repositories```.

### The configuration file :
```json
{
	"organisation": "MarquisIO",
	"labels": [{
		"name": "LABEL_NAME",
		"color": "123456",
		"repositories": [
			""
		]
	}]
}

```

|Field|Type|Definition|
|---|---|---|
|organisation|String|Which organisation the command will affect.|
|name|String|Label name.|
|color|String|Label color in hexadecimal.
|repositories|Array of String|Regular expressions for specifying what repository will be affected by this label.|
