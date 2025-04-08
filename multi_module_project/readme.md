In case you are working with multiple modules you need to initiate workspace

```shell
go work init ./practice/ ./uitility/ # this enables working with the different modules and can import one anther modules packages

```


The import statement goes like
```go   
import IntHandler "utility/int"  // package_name "modulename/package_dir_name"
```

see the practice/main.py