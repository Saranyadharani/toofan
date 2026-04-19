// Concept 1: Virtual Threads (Java 21+)
Thread thread = Thread.startVirtualThread(() -> {
    System.out.println("Running in virtual thread");
});

// Concept 2: Structured Concurrency
try (var scope = new StructuredTaskScope.ShutdownOnFailure()) {
    Future<String> user = scope.fork(() -> fetchUser());
    Future<Integer> orders = scope.fork(() -> fetchOrderCount());
    scope.join();
    scope.throwIfFailed();
    String result = "User: " + user.resultNow() + ", Orders: " + orders.resultNow();
}

// Concept 3: Concurrent Collections
import java.util.concurrent.*;
ConcurrentHashMap<String, Integer> cache = new ConcurrentHashMap<>();
cache.computeIfAbsent("key", k -> expensiveLookup(k));

// Concept 4: Atomic Variables
import java.util.concurrent.atomic.*;
AtomicLong counter = new AtomicLong(0);
long value = counter.incrementAndGet();

// Concept 5: CompletableFuture Composition
CompletableFuture.supplyAsync(() -> "Hello")
    .thenApply(String::toUpperCase)
    .thenAccept(System.out::println)
    .exceptionally(ex -> {
        System.err.println("Error: " + ex);
        return null;
    });

// Concept 6: Rate Limiting with Semaphore
Semaphore semaphore = new Semaphore(10);
if (semaphore.tryAcquire()) {
    try {
        // Critical section
    } finally {
        semaphore.release();
    }
}

// Helpers
String fetchUser() { return "User"; }
int fetchOrderCount() { return 5; }
int expensiveLookup(String key) { return 42; }