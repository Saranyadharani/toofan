// Topic: CompletableFuture

import java.util.concurrent.CompletableFuture;

CompletableFuture.supplyAsync(() -> fetchUser())
    .thenApply(User::orders)
    .thenAccept(System.out::println)
    .exceptionally(ex -> { System.err.println(ex); return null; });