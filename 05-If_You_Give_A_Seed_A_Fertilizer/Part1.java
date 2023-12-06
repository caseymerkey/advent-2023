import java.io.*;
import java.util.*;
import java.util.regex.Matcher;
import java.util.regex.Pattern;
import java.util.stream.Collectors;

public class Part1 {

    public static long BAD_NUMBER = -9999;
    static class MegaMap {
        private final List<Map> listOfMaps = new ArrayList<>();
        private boolean sorted = false;

        public void addMap(Map map) {
            this.listOfMaps.add(map);
        }

        public long map(long input) {
            if (!sorted) {
                listOfMaps.sort((o1, o2) -> Long.compare(o1.sourceRangeStart, o2.sourceRangeStart));
                sorted = true;
            }

            // find the appropriate map
            for (Map m : listOfMaps) {
                if (m.isInputInRange(input)) {
                    return m.map(input);
                }
            }
            return input;
        }
    }

    static class Map {
        private final long destinationRangeStart;
        private final long sourceRangeStart;
        private final long rangeLength;

        public Map(long destinationRangeStart, long sourceRangeStart, long rangeLength) {
            this.destinationRangeStart = destinationRangeStart;
            this.sourceRangeStart = sourceRangeStart;
            this.rangeLength = rangeLength;
        }

        public long getDestinationRangeStart() {
            return destinationRangeStart;
        }

        public long getSourceRangeStart() {
            return sourceRangeStart;
        }

        public long getRangeLength() {
            return rangeLength;
        }

        public long map(long input) {
            if (isInputInRange(input)) {
                return destinationRangeStart + (input - sourceRangeStart);
            } else {
                return Part1.BAD_NUMBER;
            }
        }

        public boolean isInputInRange(long input) {
            return (input >= sourceRangeStart && input < sourceRangeStart + rangeLength);
        }
    }

    private final String[] paths = {"soil", "fertilizer", "water", "light", "temperature", "humidity", "location"};
    private List<Long> seeds = null;
    private final HashMap<String, MegaMap> mapOfMaps = new HashMap<>();

    private final Pattern mapLabelPattern = Pattern.compile("([a-z]+)-to-([a-z]+) map:");
    public static void main(String[] args) {
        File file = new File("input.txt");
        Part1 part1 = new Part1();
        part1.readFile(file);
        part1.evaluate();
    }

    private void evaluate() {
        long lowestLocation = Long.MAX_VALUE;
        for (long seed : seeds) {
            long nextValue = seed;
            for (String nextMap : paths) {
                MegaMap megaMap = mapOfMaps.get(nextMap);
                nextValue = megaMap.map(nextValue);
            }
            System.out.println(String.format("Seed %d ==> location %d", seed, nextValue));
            if (nextValue < lowestLocation) {
                lowestLocation = nextValue;
            }
        }
        System.out.println(String.format("Lowest location value: %d", lowestLocation));
    }

    void readFile(File file) {

        long lineNumber = 0;
        try (BufferedReader reader = new BufferedReader(new FileReader(file))) {
            String line = reader.readLine();
            // first line is seeds
            String[] seedStrings = line.substring(7).split(" ");
            seeds = Arrays.stream(seedStrings)
                    .map(Long::parseLong)
                    .collect(Collectors.toList());

            String mapName = null;
            MegaMap arrayOfMaps = new MegaMap();
            while ((line = reader.readLine()) != null) {
                if (!line.isBlank()) {
                    Matcher matcher = mapLabelPattern.matcher(line);
                    if (matcher.matches()) {
                        if (mapName != null) {
                            this.mapOfMaps.put(mapName, arrayOfMaps);
                            arrayOfMaps = new MegaMap();
                        }
                        mapName = matcher.group(2);
                    } else {
                        String[] mapString = line.split(" ");
                        arrayOfMaps.addMap(new Map(Long.parseLong(mapString[0]),
                                Long.parseLong(mapString[1]),
                                Long.parseLong(mapString[2])));
                    }
                }
            }
            this.mapOfMaps.put(mapName, arrayOfMaps);

        } catch (IOException e) {
            System.out.println(e.getMessage());
        }
    }
}
