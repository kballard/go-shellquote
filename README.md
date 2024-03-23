# go-shellquote 

This is a clone of the no longer maintaind project
[go-shellquote](https://github.com/kballard/go-shellquote), name is kept the
same, but import path is altered.

The main purpose is to have features like:
- go.mod
- release or versioned tags
- CI with lint and test
- local lint and test in containers

See: 
- https://github.com/kballard/go-shellquote/issues/10
- https://github.com/kballard/go-shellquote/issues/13
- https://github.com/kballard/go-shellquote/issues/5

## PACKAGE

package shellquote
    import "github.com/Hellseher/go-shellquote"

    Shellquote provides utilities for joining/splitting strings using sh's
    word-splitting rules.

## VARIABLES

var (
    UnterminatedSingleQuoteError = errors.New("Unterminated single-quoted string")
    UnterminatedDoubleQuoteError = errors.New("Unterminated double-quoted string")
    UnterminatedEscapeError      = errors.New("Unterminated backslash-escape")
)


## FUNCTIONS

func Join(args ...string) string
    Join quotes each argument and joins them with a space. If passed to
    /bin/sh, the resulting string will be split back into the original
    arguments.

func Split(input string) (words []string, err error)
    Split splits a string according to /bin/sh's word-splitting rules. It
    supports backslash-escapes, single-quotes, and double-quotes. Notably it
    does not support the $'' style of quoting. It also doesn't attempt to
    perform any other sort of expansion, including brace expansion, shell
    expansion, or pathname expansion.

    If the given input has an unterminated quoted string or ends in a
    backslash-escape, one of UnterminatedSingleQuoteError,
    UnterminatedDoubleQuoteError, or UnterminatedEscapeError is returned.


