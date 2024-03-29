// utilities for managing stopwords
// stopword list courtesy of https://dev.mysql.com/doc/refman/8.0/en/fulltext-stopwords.html#fulltext-stopwords-stopwords-for-myisam-search-indexes
// our version is somewhat modified
package txt

import (
	"strings"
)

var STOPWORDS = map[string]bool{"a": true, "a's": true, "able": true, "about": true, "above": true, "according": true, "accordingly": true, "across": true, "actually": true, "after": true, "afterwards": true, "again": true, "against": true, "ain't": true, "all": true, "allow": true, "allows": true, "almost": true, "alone": true, "along": true, "already": true, "also": true, "although": true, "always": true, "am": true, "among": true, "amongst": true, "an": true, "and": true, "another": true, "any": true, "anybody": true, "anyhow": true, "anyone": true, "anything": true, "anyway": true, "anyways": true, "anywhere": true, "apart": true, "appear": true, "appreciate": true, "appropriate": true, "are": true, "aren't": true, "around": true, "as": true, "aside": true, "ask": true, "asking": true, "associated": true, "at": true, "available": true, "away": true, "awfully": true, "be": true, "became": true, "because": true, "become": true, "becomes": true, "becoming": true, "been": true, "before": true, "beforehand": true, "behind": true, "being": true, "believe": true, "below": true, "beside": true, "besides": true, "best": true, "better": true, "between": true, "beyond": true, "both": true, "brief": true, "but": true, "by": true, "c'mon": true, "c's": true, "came": true, "can": true, "can't": true, "cannot": true, "cant": true, "cause": true, "causes": true, "certain": true, "certainly": true, "changes": true, "clearly": true, "cocom": true, "come": true, "comes": true, "concerning": true, "consequently": true, "consider": true, "considering": true, "contain": true, "containing": true, "contains": true, "corresponding": true, "could": true, "couldn't": true, "course": true, "currently": true, "definitely": true, "described": true, "despite": true, "did": true, "didn't": true, "different": true, "do": true, "does": true, "doesn't": true, "doing": true, "don't": true, "done": true, "down": true, "downwards": true, "during": true, "each": true, "edu": true, "eg": true, "eight": true, "either": true, "else": true, "elsewhere": true, "enough": true, "entirely": true, "especially": true, "et": true, "etc": true, "even": true, "ever": true, "every": true, "everybody": true, "everyone": true, "everything": true, "everywhere": true, "ex": true, "exactly": true, "example": true, "except": true, "far": true, "few": true, "fifth": true, "first": true, "five": true, "followed": true, "following": true, "follows": true, "for": true, "former": true, "formerly": true, "forth": true, "four": true, "from": true, "further": true, "furthermore": true, "get": true, "gets": true, "getting": true, "given": true, "gives": true, "go": true, "goes": true, "going": true, "gone": true, "got": true, "gotten": true, "greetings": true, "had": true, "hadn't": true, "happens": true, "hardly": true, "has": true, "hasn't": true, "have": true, "haven't": true, "having": true, "he": true, "he's": true, "hello": true, "help": true, "hence": true, "her": true, "here": true, "here's": true, "hereafter": true, "hereby": true, "herein": true, "hereupon": true, "hers": true, "herself": true, "hi": true, "him": true, "himself": true, "his": true, "hither": true, "hopefully": true, "how": true, "howbeit": true, "however": true, "i'd": true, "i'll": true, "i'm": true, "i've": true, "ie": true, "if": true, "ignored": true, "immediate": true, "in": true, "inasmuch": true, "inc": true, "indeed": true, "indicate": true, "indicated": true, "indicates": true, "inner": true, "insofar": true, "instead": true, "into": true, "inward": true, "is": true, "isn't": true, "it": true, "it'd": true, "it'll": true, "it's": true, "its": true, "itself": true, "just": true, "keep": true, "keeps": true, "kept": true, "know": true, "known": true, "knows": true, "last": true, "lately": true, "later": true, "latter": true, "latterly": true, "least": true, "less": true, "lest": true, "let": true, "let's": true, "like": true, "liked": true, "likely": true, "little": true, "look": true, "looking": true, "looks": true, "ltd": true, "mainly": true, "many": true, "may": true, "maybe": true, "me": true, "mean": true, "meanwhile": true, "merely": true, "might": true, "more": true, "moreover": true, "most": true, "mostly": true, "much": true, "must": true, "my": true, "myself": true, "name": true, "namely": true, "nd": true, "near": true, "nearly": true, "necessary": true, "need": true, "needs": true, "neither": true, "never": true, "nevertheless": true, "new": true, "next": true, "nine": true, "no": true, "nobody": true, "non": true, "none": true, "noone": true, "nor": true, "normally": true, "not": true, "nothing": true, "novel": true, "now": true, "nowhere": true, "obviously": true, "of": true, "off": true, "often": true, "oh": true, "ok": true, "okay": true, "old": true, "on": true, "once": true, "one": true, "ones": true, "only": true, "onto": true, "or": true, "other": true, "others": true, "otherwise": true, "ought": true, "our": true, "ours": true, "ourselves": true, "out": true, "outside": true, "over": true, "overall": true, "own": true, "particular": true, "particularly": true, "per": true, "perhaps": true, "placed": true, "please": true, "plus": true, "possible": true, "presumably": true, "probably": true, "provides": true, "que": true, "quite": true, "qv": true, "rather": true, "rd": true, "re": true, "really": true, "reasonably": true, "regarding": true, "regardless": true, "regards": true, "relatively": true, "respectively": true, "right": true, "said": true, "same": true, "saw": true, "say": true, "saying": true, "says": true, "second": true, "secondly": true, "see": true, "seeing": true, "seem": true, "seemed": true, "seeming": true, "seems": true, "seen": true, "self": true, "selves": true, "sensible": true, "sent": true, "serious": true, "seriously": true, "seven": true, "several": true, "shall": true, "she": true, "should": true, "shouldn't": true, "since": true, "six": true, "so": true, "some": true, "somebody": true, "somehow": true, "someone": true, "something": true, "sometime": true, "sometimes": true, "somewhat": true, "somewhere": true, "soon": true, "sorry": true, "specified": true, "specify": true, "specifying": true, "still": true, "sub": true, "such": true, "sup": true, "sure": true, "t's": true, "take": true, "taken": true, "tell": true, "tends": true, "th": true, "than": true, "thank": true, "thanks": true, "thanx": true, "that": true, "that's": true, "thats": true, "the": true, "their": true, "theirs": true, "them": true, "themselves": true, "then": true, "thence": true, "there": true, "there's": true, "thereafter": true, "thereby": true, "therefore": true, "therein": true, "theres": true, "thereupon": true, "these": true, "they": true, "they'd": true, "they'll": true, "they're": true, "they've": true, "think": true, "third": true, "this": true, "thorough": true, "thoroughly": true, "those": true, "though": true, "three": true, "through": true, "throughout": true, "thru": true, "thus": true, "to": true, "together": true, "too": true, "took": true, "toward": true, "towards": true, "tried": true, "tries": true, "truly": true, "try": true, "trying": true, "twice": true, "two": true, "un": true, "under": true, "unfortunately": true, "unless": true, "unlikely": true, "until": true, "unto": true, "up": true, "upon": true, "us": true, "use": true, "used": true, "useful": true, "uses": true, "using": true, "usually": true, "value": true, "various": true, "very": true, "via": true, "viz": true, "vs": true, "want": true, "wants": true, "was": true, "wasn't": true, "way": true, "we": true, "we'd": true, "we'll": true, "we're": true, "we've": true, "welcome": true, "well": true, "went": true, "were": true, "weren't": true, "what": true, "what's": true, "whatever": true, "when": true, "whence": true, "whenever": true, "where": true, "where's": true, "whereafter": true, "whereas": true, "whereby": true, "wherein": true, "whereupon": true, "wherever": true, "whether": true, "which": true, "while": true, "whither": true, "who": true, "who's": true, "whoever": true, "whole": true, "whom": true, "whose": true, "why": true, "will": true, "willing": true, "wish": true, "with": true, "within": true, "without": true, "won't": true, "wonder": true, "would": true, "wouldn't": true, "yes": true, "yet": true, "you": true, "you'd": true, "you'll": true, "you're": true, "you've": true, "your": true, "yours": true, "yourself": true, "yourselves": true, "zero": true}

func ContainsStopwords(str string) bool {
	m := false
	w := strings.Split(str, " ")

	for _, v := range w {
		if _, ok := STOPWORDS[v]; ok {
			return true
		}
	}

	return m
}

func removeAtIndex(s []string, idx int) []string {
	return append(s[:idx], s[idx+1:]...)
}

func RemoveStopwords(str string) string {
	str = strings.ReplaceAll(strings.ToLower(str), "\n", "")

	w := strings.Split(str, " ")

	for i := 0; i < len(w); i++ {
		v := w[i]
		if STOPWORDS[v] {
			w = removeAtIndex(w, i)
			i--
		}
	}

	return strings.Join(w, " ")
}

func FilterStopwords(tokens []string) []string {
	withoutStops := []string{}
	for _, v := range tokens {
		v = strings.ToLower(v)
		if _, ok := STOPWORDS[v]; !ok {
			withoutStops = append(withoutStops, v)
		}
	}

	return withoutStops
}
