# envfile
A little helper written in go to create a config file from environment variables

 ## Usage

You can copy the envfile binary located at https://github.com/Angelinsky7/envfile/releases/download/v0.0.1/envfile or build it from the source into your container/machine/project.
The current binary has been built with `CGO_ENABLED=0 GOOS=linux $(GO) build -a -installsuffix cgo -o $(OUT)/envfile` like the Makefile.

if you put the binary here : `/usr/bin/envfile` you can use like that :

| Command     | Type    | Caption | Values |
|-------------|---------|---------|--------|
| -formatter  | string  | The output formatter (default "json")        | json, yaml  |
| -help       | bool    | Print the help                               |   |
| -prefix     | string  | The prefix of the environment variables      |   |
| -r          | bool    | Is the prefix removed from the variable name |   |
| -separator  | string  | The key separator (default "__")             |   |
| -v          | bool    | Verbose.                                     |   |
| -version    | bool    | Print the version                            |   |

# Examples

```bash
/usr/bin/envfile -prefix NG_ -r > /usr/share/nginx/html/assets/config.json
```

This will crate a file containing all environment variables starting with NG_.
If the environment variables are 

```bash
NG_TEST1_VALUE1_SUB1=V1
NG_TEST1_VALUE1_SUB2=V2
NG_TEST1_VALUE2_SUB1=V3
NG_TEST1_VALUE3=V4
NG_TEST2=V5
```

the result file will be

```json
{
  "TEST1": {
    "VALUE1": {
      "SUB1": "V1",
      "SUB2": "V2"
    },
    "VALUE2": {
      "SUB1": "V3"
    },
    "VALUE3" : "V4"
  },
  "TEST2": "V5"
}
```

and if you use the yaml formatter :

```yaml
TEST1: 
  VALUE1: 
    SUB1: "V1"
    SUB2: "V2"
  VALUE2:
    SUB1: "V3"
  VALUE3: "V4"
TEST2: "V5"
```
