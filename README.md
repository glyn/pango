# pango - Package Analysis for Go

Some ways of developing Go code, particularly TDD, result in many small types composed together. But when you need to grok the overall structure, it can be laborious to wade through the code.

This is where pango comes in - it analyses the packages of your code and their interdependencies and produces various metrics which can alert you to certain structural problems.