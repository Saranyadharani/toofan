// Topic: Advanced Types

// 1. Union Types
type ID = string | number;
const userId: ID = "user_123";
const productId: ID = 456;

// 2. Intersection Types
type Admin = User & { permissions: string[] };

// 3. Type Guards
function isString(value: unknown): value is string {
    return typeof value === "string";
}

// 4. Type Aliases
type ApiResponse<T> = { data: T; status: number };

// 5. Literal Types
type Status = "pending" | "approved" | "rejected";
const orderStatus: Status = "pending";