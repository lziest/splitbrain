# splitbrain
An example when Go's vendoring works against productity

## [mood](https://github.com/lziest/mood) 

package `mood` has a method `Show()` which prints out a exported variable `Status`

## [funkybrain](https://github.com/lziest/funkybrain)

package `funkybrain` appears to import `mood` and also vendor it.
Nothing bad happens, right?

## Now meet splitbrain
Now at this package, we decided to use `mood` and `funkybrain`.
But guess what, our `mood` may change to be "sad",
but our invisible `mood` for funkybrain is for ever "happy".
When we change our visible "mood", our funkybrain actually never cares.

## Installation

With a [correctly configured](https://golang.org/doc/code.html#GOPATH) Go installation:

```
go get -u github.com/lziest/splitbrain
```

## Play it

```
$ splitbrain --mood sad
mood.Status ptr = 0x19be30
mood.Status ptr = 0x19be40
It appears you can change your mood to what you want: I'm sad
Deep inside the funkybrain, it never changes: I'm happy

