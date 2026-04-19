// Topic: Optional

import java.util.Optional;

// 1. Create Optional
Optional<String> optional = Optional.of("value");
Optional<String> empty = Optional.empty();

// 2. ifPresent
optional.ifPresent(v -> System.out.println(v));

// 3. orElse
String result = empty.orElse("default");

// 4. orElseThrow
String value = optional.orElseThrow(() -> new RuntimeException("Missing"));

// 5. map
Optional<Integer> length = optional.map(String::length);