# GO RELOADED

Go reloaded is a file formatter based on a given certain rules
It reads the contents of a given file the corrects it appropriately

### Rules

>
    1. Every instance of (hex) should replace the word before with the decimal version of the word (in this case the word will always be a hexadecimal number). (Ex: "1E (hex) files were added" -> "30 files were added")
    Every instance of (bin) should replace the word before with the decimal version of the word (in this case the word will always be a binary number). (Ex: "It has been 10 (bin) years" -> "It has been 2 years")

    2. Every instance of (up) converts the word before with the Uppercase version of it. (Ex: "Ready, set, go (up) !" -> "Ready, set, GO !")

    3. Every instance of (low) converts the word before with the Lowercase version of it. (Ex: "I should stop SHOUTING (low)" -> "I should stop shouting")

    4. Every instance of (cap) converts the word before with the capitalized version of it. (Ex: "Welcome to the Brooklyn bridge (cap)" -> "Welcome to the Brooklyn Bridge")
        For (low), (up), (cap) if a number appears next to it, like so: (low, <number>) it turns the previously specified number of words in lowercase, uppercase or capitalized accordingly. (Ex: "This is so exciting (up, 2)" -> "This is SO EXCITING")

    5. Every instance of the punctuations ., ,, !, ?, : and ; should be close to the previous word and with space apart from the next one. (Ex: "I was sitting over there ,and then BAMM !!" -> "I was sitting over there, and then BAMM!!").
        Except if there are groups of punctuation like: ... or !?. In this case the program should format the text as in the following example: "I was thinking ... You were right" -> "I was thinking... You were right".

    6. The punctuation mark ' will always be found with another instance of it and they should be placed to the right and left of the word in the middle of them, without any spaces. (Ex: "I am exactly how they describe me: ' awesome '" -> "I am exactly how they describe me: 'awesome'")
        If there are more than one word between the two ' ' marks, the program should place the marks next to the corresponding words (Ex: "As Elton John said: ' I am the most well-known homosexual in the world '" -> "As Elton John said: 'I am the most well-known homosexual in the world'")

    7. Every instance of a should be turned into an if the next word begins with a vowel (a, e, i, o, u) or a h. (Ex: "There it was. A amazing rock!" -> "There it was. An amazing rock!").


## Usage 
Download the repository 

```git
    git clone https://learn.zone01kisumu.ke/git/aaochieng/go-reloaded.git
```
Within the directory create```sample.txt``` file with malformed text i.e ```Simply add 42 (hex) and 10 (bin) and you will see the result is 68.```
Then run the below command in the root directory of the ownloaded repository

```golang
    go run . sample.txt result.txt
```

A new file ```result.txt``` will be created automatically. Based on the given example, the ```result.txt``` file should contain
the following text inside it ```Simply add 66 and 2 and you will see the result is 68.```

Based on the malformed sentence, it is supposed to convert the previous words before hex and bin and convert then to decimal then finally to string.

The code :

``` go

    // Converts a given string from binary to int64
    func Bindec(s string) int64 {
        value, err := strconv.ParseInt(s, 2, 64)
        if err != nil {
            fmt.Println(err)
        }
        return value
    }

    // Converts a given string from hexadecimal to int64
    func Hexdec(s string) int64 {
        value, err := strconv.ParseInt(s, 16, 64)
        if err != nil {
            fmt.Println(err)
        }
        return value
    }
```

## Code Testing 

To further test the code, add your tests based on the rules provided above in `goreloaded_test.go` file in the folder.

In `goreloaded_test.go` file modify the code section 

```go
    var addTests = []addTest{
	// Add your argument and expected result 
    // ...previous code
	{"Punctuation tests are ... kinda boring ,don't you think !?", "Punctuation tests are... kinda boring, don't you think!?"},
    }
```
Then run 

```go
    go test -v
```