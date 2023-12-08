package merkey.aoc.camel;

import java.util.HashMap;

public class HandPart2 implements Comparable<HandPart2>{

    private final String cards;
    private final long bid;
    private final HandStrength strength;

    public HandPart2(String cards, long bid) {
        this.cards = cards;
        this.bid = bid;
        this.strength = evaluateStrength(cards);
    }

    HandStrength evaluateStrength(String cards) {
        HandStrength strength = HandStrength.HighCard;
        HashMap<Byte, Integer> cardCounts = new HashMap<>();
        int numberOfJokers = 0;
        for (byte c : cards.getBytes()) {
            if ((char)c != 'J') {
                int count = 1;
                if (cardCounts.containsKey(c)) {
                    count = cardCounts.get(c) + 1;
                }
                cardCounts.put(c, count);
            } else {
                numberOfJokers++;
            }
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

        strength = switch (numberOfJokers) {
            case 1 -> switch (strength) {
                case HighCard -> HandStrength.Pair;
                case Pair -> HandStrength.ThreeOfAKind;
                case TwoPair -> HandStrength.FullHouse;
                case ThreeOfAKind -> HandStrength.FourOfAKind;
                case FourOfAKind -> HandStrength.FiveOfAKind;
                default -> strength;
            };
            case 2 -> switch (strength) {
                case HighCard -> HandStrength.ThreeOfAKind;
                case Pair -> HandStrength.FourOfAKind;
                case ThreeOfAKind -> HandStrength.FiveOfAKind;
                default -> strength;
            };
            case 3 -> switch (strength) {
                case HighCard -> HandStrength.FourOfAKind;
                case Pair -> HandStrength.FiveOfAKind;
                default -> strength;
            };
            case 4, 5 -> HandStrength.FiveOfAKind;
            default -> strength;
        };

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
    public int compareTo(HandPart2 other) {
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
                value = 0;
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
