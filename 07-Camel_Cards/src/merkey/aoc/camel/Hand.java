package merkey.aoc.camel;

import java.util.HashMap;

public class Hand implements Comparable<Hand>{

    private final String cards;
    private final long bid;
    private final HandStrength strength;

    public Hand(String cards, long bid) {
        this.cards = cards;
        this.bid = bid;
        this.strength = evaluateStrength(cards);
    }

    HandStrength evaluateStrength(String cards) {
        HandStrength strength = HandStrength.HighCard;
        HashMap<Byte, Integer> cardCounts = new HashMap<>();
        for (byte c : cards.getBytes()) {
            int count = 1;
            if (cardCounts.containsKey(c)) {
                count = cardCounts.get(c) + 1;
            }
            cardCounts.put(c, count);
        }
        int pairs = 0;
        int threes = 0;
        int fours = 0;
        int fives = 0;
        for (Byte b : cardCounts.keySet()) {
            if (cardCounts.get(b) == 5) fives++;
            if (cardCounts.get(b) == 4) fours++;
            if (cardCounts.get(b) == 3) threes++;
            if (cardCounts.get(b) == 2) pairs++;
        }
        if (fives > 0) {
            strength = HandStrength.FiveOfAKind;
        } else if (fours > 0) {
            strength = HandStrength.FourOfAKind;
        } else if (threes > 0) {
            if (pairs > 0) {
                strength = HandStrength.FullHouse;
            } else {
                strength = HandStrength.ThreeOfAKind;
            }
        } else if (pairs == 1) {
            strength = HandStrength.Pair;
        } else if (pairs == 2) {
            strength = HandStrength.TwoPair;
        }

        return strength;
    }

    public String getCards() {
        return cards;
    }

    public long getBid() {
        return bid;
    }

    public HandStrength getStrength() {
        return strength;
    }

    @Override
    public String toString() {
        return "Hand{" +
                "cards='" + cards + '\'' +
                ", bid=" + bid +
                ", strength=" + strength +
                '}';
    }

    @Override
    public int compareTo(Hand other) {
        int compareValue = this.strength.compareTo(other.strength);
        if (compareValue == 0) {
            int i = 0;
            while ((compareValue == 0) && (i < 5)) {
                compareValue = Integer.compare(
                        cardValue(this.cards.charAt(i)),
                        cardValue(other.cards.charAt(i)));
                i++;
            }
        }
        return compareValue;
    }

    private int cardValue(char card) {
        int value = 0;
        switch (card) {
            case 'T':
                value = 10;
                break;
            case 'J':
                value = 11;
                break;
            case 'Q':
                value = 12;
                break;
            case 'K':
                value = 13;
                break;
            case 'A':
                value = 14;
                break;
            default:
                value = Integer.decode(String.valueOf(card));
        }
        return value;
    }
}
