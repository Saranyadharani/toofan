// Topic: Records and Pattern Matching

record Point(int x, int y) {}
record Circle(Point center, int radius) {}

if (shape instanceof Circle(Point(var x, var y), var r)) {
    System.out.println("Circle at (" + x + "," + y + ")");
}