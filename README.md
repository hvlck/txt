# txt

`text utilities (golang)`

`txt` is mainly a lot of text normalization utilities, primarily for full-text search. This library is primarily intended for my own use, but pull requests are always welcome. The primary focus is on performance, for use in a full-text search engine I'm building.

## Utilities

+ fast levenshtein distance calculator (~1.3*10^12 ops/s)
+ an implementation of the Porter stemming algorithm
+ stopword remover
+ various reading functions
  + read time
  + word/char count
  + (not implemented) text difficulty (Flescher-Kincaid)
+ a generic [Trie](https://en.wikipedia.org/wiki/Trie) data structure
  + inserts: ~6.7*10^14 ops/s
  + exact matches: ~3.3*10^15 ops/s
  + partial matches: ~1.2*10^14 ops/s

## Roadmap

+ text normalisation
  + color names (words -> hex)
  + URL normalizations
  + remove fractional numbers to `nth` precision
  + synoyms/thesaurus
    + replace multiple words with one word that is a synoym of it
+ text difficulty
+ support for multiple languages

## Credits

+ [Stopwords](https://dev.mysql.com/doc/refman/8.0/en/fulltext-stopwords.html#fulltext-stopwords-stopwords-for-myisam-search-indexes)
+ [Words](http://www.gwicks.net/dictionaries.htm)
+ [Colors](https://www.colorcodehex.com/sort-by-hex-value.html)
+ [Word Frequency](https://www.kaggle.com/datasets/rtatman/english-word-frequency?resource=download)
