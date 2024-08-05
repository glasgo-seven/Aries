# ANSI

##### Content:

* ansi_Commands.go
  ```go
  func ANSI_ClearScreen() { ... }
  
  // ANSI text coloring
  func ANSI_SetColor16(color string) { ... }
  func ANSI_SetColor16_string(color string) string { ... }
  
  func ANSI_ResetColor16() { ... }
  func ANSI_ResetColor16_string() string { ... }
  
  func ANSI_SetColor256_FG(color string) { ... }
  func ANSI_SetColor256_FG_string(color string) string { ... }
  
  func ANSI_SetColor256_BG(color string) { ... }
  func ANSI_SetColor256_BG_string(color string) string { ... }
  
  func ANSI_SetColorRGB_FG(r string, g string, b string) { ... }
  func ANSI_SetColorRGB_FG_string(r string, g string, b string) string { ... }
  
  func ANSI_SetColorRGB_BG(r string, g string, b string) { ... }
  func ANSI_SetColorRGB_BG_string(r string, g string, b string) string { ... }
  
  func ANSI_ResetColor() { ... }
  func ANSI_ResetColor_string() string { ... }
  
  
  func ANSI_FormatString(str string, format string, reset string) string { ... }
  func ANSI_CommandExecute(command string, x string) { ... }
  
  
  // ANSI text formatting for log messages
  func ANSI_AsAlert(message string) string { ... }
  func ANSI_Alert(message string) { ... }
  func ANSI_AlertLine(message string) { ... }
  
  func ANSI_AsError(message string) string { ... }
  func ANSI_Error(message string) { ... }
  func ANSI_ErrorLine(message string) { ... }
  
  
  // ANSI text styling
  func ANSI_Bold(str string) string { ... }
  func ANSI_Dim(str string) string { ... }
  func ANSI_Italic(str string) string { ... }
  func ANSI_UnderLine(str string) string { ... }
  func ANSI_Blink(str string) string { ... }
  func ANSI_InverseColor(str string) string { ... }
  func ANSI_Hidden(str string) string { ... }
  func ANSI_Strike(str string) string {} ... 
  ```

* ansi_Constants.go
  ```go
  ```

* ansi_Formatting.go
  ```go
  ```
