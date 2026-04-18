// Topic: Working with Collections

import java.util.*;

public class CollectionsDemo {
    public static void main(String[] args) {
        List<String> names = new ArrayList<>();
        names.add("Alice");
        names.add("Bob");
        
        Map<Integer, String> users = new HashMap<>();
        users.put(1, "John");
        users.put(2, "Jane");
        
        for (String name : names) {
            System.out.println(name);
        }
    }
}