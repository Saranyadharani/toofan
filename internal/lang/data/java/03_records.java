// Topic: Records

// 1. Basic Record
record Point(int x, int y) {}

// 2. Record with validation
record Person(String name, int age) {
    Person {
        if (age < 0) throw new IllegalArgumentException("Age cannot be negative");
    }
}

// 3. Record with custom method
record Rectangle(int width, int height) {
    int area() { return width * height; }
}

// 4. Record with static field
record Config(String url) {
    static final int DEFAULT_PORT = 8080;
}

// 5. Record in collection
List<Point> points = List.of(new Point(1, 2), new Point(3, 4));