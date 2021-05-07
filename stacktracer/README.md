# Purpose 
    To print customized echo logs
## Input
    String message to print
## Input example:
    logger.Echo("Hey error in main")
## Output 
    FileName : packageName/FunctionName: Line number: : ErrorMessage
## Output Example
    main.go : main.main : Line: 16 -> Hey error in main
    stackController.go : stack/controller.TestLogger : Line: 9 -> Hey error in controller