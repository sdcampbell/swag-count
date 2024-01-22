# swag-count
Count the number of Swagger API endpoints for API pentest scoping.

Usage: `swag-count -u [URL to swagger.json]` or `swag-count -f [path to swagger.json]`

Examples:

```
swag-count -u https://petstore.swagger.io/v2/swagger.json
Number of API endpoints: 20
```

```
swag-count -f ~/Downloads/swagger.json                   
Number of API endpoints: 20
```