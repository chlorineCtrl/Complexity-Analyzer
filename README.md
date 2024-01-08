**INSTALLATION:**

```bash
go get github.com/chlorineCtrl/Complexity-Analyzer/cmd/complexity
go install github.com/chlorineCtrl/Complexity-Analyzer/cmd/complexity
```

**USAGE:**

```
go vet -vettool=$(which complexity) [Flags] [Directory/file]
```

**FLAGS:**

```
--cycloover: show functions with the Cyclomatic complexity > N (default: 10)

--maintunder: show functions with the Maintainability index < N (default: 20)
```

**UNINSTALLATION:**

```
go clean -i github.com/chlorineCtrl
/Complexity-Analyzer/cmd/complexity
```

**OUTPUT:**
![output1](https://github.com/chlorineCtrl/Complexity-Analyzer/blob/main/images/msin.png)
