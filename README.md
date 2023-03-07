# Recipe API

## Create recipe payload

```json
{
	"name": "My Recipe",
	"keywords": ["test-key1", "test-key2"],
	"instructions": ["boil water", "turn on cooker"],
	"ingredients": ["water", "oil", "aggression"],
	"chefId": "random_id"
}
```

## Create chef payload

```json
{
	"name": "Fuad",
	"country": "Nigeria",
	"yearsOfExperience": 3
}
```

## Get chefs endpoint variants

This will return all chefs only.
```text
GET /chefs
```

This will include all chefs recipes in the response.
```text
GET /chefs?populate=true 
```

