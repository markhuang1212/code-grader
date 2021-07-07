# code-grader

This is a backend for grading C++ code in a secure and controlled manner.

## Usage

```shell
> docker-compose build
> docker-compose up
```

## API

### Grade Code Request

```bash
curl --location --request POST 'http://localhost:8080/api/v1/grade' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "TestCaseName": "example-1",
        "UserCode": "int main() { cout << \"Hello\"; }"
    }'
```

### Retrieve Grade Result

```bash
curl --location \
    --request GET 'http://localhost:8080/api/v1/result/cbaf98c458f3878df6f5'
```

## Author

Huang Meng \
2021