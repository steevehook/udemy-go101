# Section 8 Assignment

## Instructions

1. Create `HTTPError` custom type which has the following fields:

```
StatusCode  - int
Code        - string
Context     - []ErrorField
```

2. Create `ErrorField` custom type which has the following fields:

```
Code        - int
Field       - string
Constraints - map[string]string
```

Note: `Code` field should have the format: `APP_NAME`.`ACTION_NAME`.`ERROR_TYPE`
for example: `notes.read.format_error`, `notes.write.validation_error`

`HTTPError` -> `Code` represents generic error code
and `ErrorField` -> `Code` represents field specific code

Note: `StatusCode` is for internal use only, thus should
not make it to the final JSON

3. Output couple of error to STDOUT in JSON format

## Expectations

```
./assignment

{
    "code": "notes.read.format_error",
    "context": [
        {
            "code": "notes.read.invalid_length"
            "field": "title",
            "constraints": {
                "min": "10",
                "max": "2000"
            }
        },
    ]
}
```

[Back](https://github.com/steevehook/udemy-go101)

---

Happy hacking gophers 🚀🚀🚀

<img src="https://github.com/steevehook/udemy-go101/raw/master/udemy-go101.svg?sanitize=true" width="150px"/>
