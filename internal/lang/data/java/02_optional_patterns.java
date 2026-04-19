// Topic: Optional Patterns

import java.util.Optional;

String value = Optional.ofNullable(getData())
    .filter(s -> !s.isEmpty())
    .map(String::trim)
    .orElseThrow(() -> new IllegalStateException("Invalid data"));