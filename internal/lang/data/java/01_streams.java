// Topic: Streams

import java.util.*;
import java.util.stream.*;

// 1. Filter
List<String> names = Arrays.asList("Alice", "Bob", "Charlie");
names.stream().filter(n -> n.startsWith("A")).collect(Collectors.toList());

// 2. Map
names.stream().map(String::toUpperCase).collect(Collectors.toList());

// 3. Reduce
int sum = Stream.of(1, 2, 3, 4).reduce(0, (a, b) -> a + b);

// 4. Collect
List<String> filtered = names.stream().filter(n -> n.length() > 3).collect(Collectors.toList());

// 5. FindFirst
Optional<String> first = names.stream().filter(n -> n.length() > 3).findFirst();