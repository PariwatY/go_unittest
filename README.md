# go_unittest


# Command for running test
1. Run all test
```
go test ./...
```
2. Run specific test
- Example 1
```
go test go_unittest/services -v -run="TestCheckGrade"
```
- Example 2
```
go test go_unittest/services -v -run="TestCheckGrade/should_success_when_inputGrade80_then_returnA"
```

3. Run check test coverage
```
go test go_unittest/services -cover
```


# Config VS Code to see your code coverage 
1. For Mac command+shift+p
2. find settings.json in vscode
3. Use below config put into settings.json
```
 "go.coverOnSave": true,
    "go.coverOnSingleTest": true,
    "go.coverageDecorator": {
        "type": "gutter",
        "coveredHighlightColor": "rgba(64,128,128,0.5)",
        "uncoveredHighlightColor": "rgba(128,64,64,0.25)",        
        "coveredGutterStyle": "blockgreen",
        "uncoveredGutterStyle": "blockred"
    }
```