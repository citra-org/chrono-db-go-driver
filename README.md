# IntelliLog-GoLang-Driver
GoLang Driver for IntelliLog Database


# Docs

### Usage
Run
```
go get github.com/citra-org/chrono-db-go-driver
```

### FUNCTIONS

1. Connect <<**URI** *string*>>

This function will start a TCP connection with credentials provided, if correct keep the connection  remans open until disconnected.

2. Close

This function will diconnect from chrono db

3. CreateStream <<**CHRONO_NAME** *string* :: **STREAM_NAME** *string*>>

4. WriteEvent <<**CHRONO_NAME** *string* :: **STREAM_NAME** *string* :: **EVENTS** { **HEADER**, **BODY** *(string, string)* } >>

5. Read <<**CHRONO_NAME** *string* :: **STREAM_NAME** *string*>>

Please use the above functions only, the other necessary functions are still under dev

[Check this sample usage](https://github.com/citra-org/dosis/tree/main/chrono-db-go-driver#readme)

### Note

THIS CODEBASE IS STILL UNDER DEV, HAS LOT OF BUGS & NEED SEVERAL IMPROVEMENTS. PLEASE DONT USE IN PROD


### Contributors

**Creator**: Sugam Kuber ([Github](https://github.com/sugamkuber)) ([LinkedIn](https://linkedin.com/in/sugamkuber))
