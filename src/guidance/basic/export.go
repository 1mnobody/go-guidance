package basic

// 大写字母开头的变量是导出的，即在其他地方可以使用这个变量
var ExportedVar = 1

// 小写字母开头的则是未导出的，未导出的变量在包外是无法访问的
var unExportedVar = 2
