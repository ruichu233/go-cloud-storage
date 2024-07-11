package errno

var ErrUserAlreadyExist = &Errno{HTTP: 400, Code: "FailedOperation.UserAlreadyExist", Message: "User already exist."}
var ErrUserNotExist = &Errno{HTTP: 400, Code: "FailedOperation.UserNotExist", Message: "User not exist."}
var ErrPasswordIncorrect = &Errno{HTTP: 400, Code: "FailedOperation.PasswordIncorrect", Message: "Password incorrect."}
