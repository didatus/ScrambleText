# ScrambleText Module for Go

This small module scrambles a text in a reproducable way. Nothing magical. Just using the ASCII value of every character of the input string to reorder each character separately into a new scrambled output string.

#### Examples:

Scramble("Hello World")
> "WeHodo lllr"

Scramble("john.doe@company.test")
> "opon.ocdtehsa.t@mjeyn"

Scramble("你好，世界")
> "你，好世界"

#### Tests

Yeah, there is a test with the examples above. It's with a table data provider and runs the test in parallel. Very cool .. unnecessary and somewhat overkill but cool. I like Go

#### Why?

Why did i developed this module?
Well, it seems like a pointless module for learning Go Development. And yes, i new in Go and still learning, but i have a real use case for this module. A somewhat special use case, but i need this module in a project. 