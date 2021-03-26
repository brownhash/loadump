# loadump

Load-Dump: Go on with load testing of your endpoints with full power

## Usage

### Config

```json
{
    "config": {
        "parallelism": 1000,
        "waitPeriod": 1,
        "executionMinutes": 5,
    },
    "http": [
        "configOverride": {
            "parallelism": 2500,
            "executionMinutes": 15,
        },
        "services": {
            "appName": {
                "addr": "http://127.0.0.1:8080/api/yourApi",
                "method": "GET",
                "header": {
                    "sampleHeader": "sampleHeaderValue",
                },
                "body": {
                    "sampleKey": "sampleValue"
                }
            },
        }
    ]
}
```
