# Phrase Generator

This package implements a high entropy, extremely fast passphrase
generator in Go. It uses a mask to represent the type of item that is 
included, whether a word, special character, or any other modifier. 

A common method of generating passphrases is the Diceware algorithm.

# Mask

Unlike other phrase generators, this generator doesn't rely on selecting
a specific number of words, rather, words are generated independently 
until the correct value is achieved. 

This is enabled by the generation of a "mask" specifying the unique item
types that go into a passphrase. Depending on which modifiers are specified
when the passphrase is generated, it ensures the exact number of correct
items are outputted when the phrase is generated. 

This also allows for even greater customizability of the generator, 
allowing it to generate a passphrase tailored to any modifier.