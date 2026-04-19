// Concept 1: Records (Java 14+)
record Person(String name, int age, String email) {}

// Concept 2: Switch Expressions (Java 14+)
enum Day { MON, TUE, WED, THU, FRI, SAT, SUN }
int getDayNumber(Day day) {
    return switch (day) {
        case MON -> 1;
        case TUE -> 2;
        case WED -> 3;
        case THU -> 4;
        case FRI -> 5;
        case SAT, SUN -> 0;
    };
}

// Concept 3: Text Blocks (Java 15+)
String json = """
    {
        "name": "John",
        "age": 30
    }
    """;

// Concept 4: Pattern Matching for instanceof (Java 16+)
Object obj = "Hello";
if (obj instanceof String str) {
    System.out.println(str.toUpperCase());
}

// Concept 5: Sealed Classes (Java 17+)
sealed interface Vehicle permits Car, Truck {}
record Car(String model) implements Vehicle {}
record Truck(int capacity) implements Vehicle {}

// Concept 6: CompletableFuture for Async
import java.util.concurrent.CompletableFuture;
CompletableFuture<String> future = CompletableFuture.supplyAsync(() -> "Result");