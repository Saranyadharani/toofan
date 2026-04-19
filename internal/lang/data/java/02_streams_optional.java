// Concept 1: Stream Operations Pipeline
import java.util.*;
import java.util.stream.*;

List<String> result = Stream.of("apple", "banana", "cherry", "date")
    .filter(s -> s.length() > 4)
    .map(String::toUpperCase)
    .sorted(Comparator.reverseOrder())
    .limit(2)
    .collect(Collectors.toList());

// Concept 2: Grouping with Collectors
Map<Integer, List<String>> byLength = Stream.of("cat", "dog", "elephant")
    .collect(Collectors.groupingBy(String::length));

// Concept 3: Optional Patterns
Optional<String> optional = Optional.ofNullable(getValue());
String result2 = optional
    .filter(s -> !s.isEmpty())
    .map(String::trim)
    .orElseThrow(() -> new IllegalStateException("Invalid value"));

// Concept 4: Primitive Streams
IntStream.rangeClosed(1, 100)
    .filter(n -> n % 3 == 0)
    .average()
    .ifPresent(System.out::println);

// Concept 5: FlatMap for Nested Structures
List<List<Integer>> nested = Arrays.asList(
    Arrays.asList(1, 2),
    Arrays.asList(3, 4)
);
List<Integer> flattened = nested.stream()
    .flatMap(Collection::stream)
    .collect(Collectors.toList());

// Concept 6: Reduce for Aggregation
int sum = Stream.of(10, 20, 30, 40)
    .reduce(0, (a, b) -> a + b);

// Helper
String getValue() { return "test"; }