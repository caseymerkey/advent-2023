package merkey.aoc.camel;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Collections;

public class Main {
    public static void main(String[] args) {

        String file = "input.txt";

        ArrayList<Hand> allHands = new ArrayList<>();
        String line = null;
        try (BufferedReader reader = new BufferedReader(new FileReader(file))) {
            while ((line = reader.readLine()) != null){
                String[] fields = line.split(" ");
                Hand hand = new Hand(fields[0], Integer.parseInt(fields[1]));
                allHands.add(hand);
            }
        } catch (IOException e) {
            System.out.println(e.getMessage());
        }

        Collections.sort(allHands);
        allHands.forEach(System.out::println);

        long total = 0L;
        for (int i=0; i<allHands.size(); i++) {
            total += allHands.get(i).getBid() * (i+1);
        }
        System.out.println("Total: " + total);
    }
}