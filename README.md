
# Building and Running
## After installing go:
  `go run main.go`

## Docker

```bash
make docker_build
make docker_run
```

# Implementation Notes

If I did this sort of dynamic attribute typing in Java, I'd use a class hierarchy and abstract functions. Since 
Go doesn't implement inheritance, I'm using interfaces, generics, and composition.

For the command line/Main I just wrote a main where you configure the products in line. 
Other options  :

- Take a 2JSON file specifying:
  - a list of products 
  - a list of rules

This would all 2 command line args and a json.unmarshall to the command


# Assumptions

## Spec

- For item B (Filter). I interpreted "threshold" as "% of Max Possible Score.":
  - if 2 rules each score 100, then 50% means > 100

## Rules

- Each condition is ANDED. We could define a struct that speficies a mix of And and Or groupings, but I'm deferring that for time

## Existence of Attribute Values

Each product will have all the attributes. If the attribute isn't specified, it's assumed to have the "0" value

For example, if the rule is "ImportedFrom = CN" and the country isn't specified, the rule caculates as if it were "" 
though in this case, a better attribute would be country of origin.

Another option would be to have each attribute have a default value

## Attributes

Since Price and Quantity are required Product will have a separate getter for price and quantity. 
If they are not present they will be treated has having a zero value.

For this we will asssume that there is a "registry" of valid attributes

