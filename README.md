# Phrase Generator

This package implements a high entropy, performant passphrase
generator in Go. It uses a mask to represent the type of item that is 
included, whether a word, special character, or any other modifier. 

## Driving Principles
The goal is to create a phrase generator that creates phrases with as 
much or more entropy than a Diceware generated phrase, along with 
additional configurability and customizability and the ability to 
calculate the strength of a generated phrase in real time.

Kerckhoff's Principle holds that even if all parts of a system are known, 
a secure system should still be secure. Anything less is Security through
Obscurity, which is no security at all. 

Discouraging discussion of weaknesses of the algorithm or any 
vulnerabilities is counterproductive; the Ganymede team takes any potential
security issues very seriously. Please submit a PR or Issue if you feel 
there is a potential vulnerability or weakness.

## Mask

Unlike other phrase generators, this generator doesn't rely on selecting
a specific number of words, rather, words are generated independently 
until the correct value is achieved. 

This is enabled by the generation of a "mask" specifying the unique item
types that go into a passphrase. Depending on which modifiers are specified
when the passphrase is generated, it ensures the exact number of correct
items are outputted when the phrase is generated. 

This also allows for even greater customizability of the generator, 
allowing it to generate a passphrase tailored to any modifier.

## Entropy
The entropy of a passphrase can be also be calculated using the mask. 
Each unique identifier has an entropy, the sum of each adds up to the
total entropy of a given phrase without any knowledge of the phrase itself.

Without any modifiers, the entropy of each word in the list is 13.2 bits,
compared to the 12.9 bits in the Diceware word list.

### Calculating Entropy
Calculating the entropy of a phrase is easy using just the phrase mask.

Given a mask M:

`M = [WORD][SEP][NUMBER][SPEC_CHAR][WORD][SEP][LAST_WORD]`

We can measure the entropy of the passphrase using the formula below, with 
H being the total Entropy. 

**Shannon Entropy Formula**

$$ H(x) = -\sum_{i=1}^n [P(x_i) * log_bP(x_i)] = \sum_{i=1}^n [P(x_i) * log_b(1 / P(x_i))] $$

Entropy is calculated using the Shannon entropy formula, and measures the
"unpredictability" of a password or passphrase. A higher entropy password
offers additional protection from those that may try to crack a passphrase.

A higher entropy is better, it's recommended to use a password with
a minimum of 25-30 bits of entropy for non-vital accounts and 60 bits
or more for important accounts.
