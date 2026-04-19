// Concept 1: Union Types for API responses
type ApiResponse<T> = { success: true; data: T } | { success: false; error: string };

// Concept 2: Type Guards
function isString(value: unknown): value is string {
    return typeof value === "string";
}

// Concept 3: Discriminated Unions
interface Circle { kind: "circle"; radius: number; }
interface Square { kind: "square"; sideLength: number; }
type Shape = Circle | Square;

function getArea(shape: Shape): number {
    switch (shape.kind) {
        case "circle": return Math.PI * shape.radius ** 2;
        case "square": return shape.sideLength ** 2;
    }
}

// Concept 4: Generic Constraints
function getProperty<T, K extends keyof T>(obj: T, key: K): T[K] {
    return obj[key];
}

// Concept 5: Mapped Types with Constraints
type ReadonlyDeep<T> = {
    readonly [P in keyof T]: T[P] extends object ? ReadonlyDeep<T[P]> : T[P];
};

// Concept 6: Template Literal Types
type EventHandlers<T extends string> = {
    [K in T as `on${Capitalize<K>}`]: () => void;
};