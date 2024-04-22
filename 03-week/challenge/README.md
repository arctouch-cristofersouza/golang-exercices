
## Challenge

Create a struct with will contain props from a card (number and suit) and an implement `interface deck` which has two funcionts:
 - New: where you create the deck with 12 cards from each suit
 - Sort: You should sort from ace to king (1-12) in the following suit order Spades, Diamonds, Clubs and Hearts 
    (Ace - Spades ... King - Spades, Ace - Diamonds ... King - Diamonds, Ace - Clubs and so on..)
- Shuffle: shuffle deck

All should change struct and NOT return a new struct

hit: Rand.Shuffle

```
    a := []int{1, 2, 3, 4, 5, 6, 7, 8}
    rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

```

here's a playground link about shuffle https://go.dev/play/p/-sb5UyWAJLs