// Topic: Generics

// 1. Generic Function
function identity<T>(value: T): T {
    return value;
}

// 2. Generic Interface
interface Box<T> {
    content: T;
}

// 3. Generic Constraints
function getProperty<T, K extends keyof T>(obj: T, key: K): T[K] {
    return obj[key];
}

// 4. Generic Class
class Stack<T> {
    private items: T[] = [];
    push(item: T): void { this.items.push(item); }
    pop(): T | undefined { return this.items.pop(); }
}

// 5. Utility Types
type PartialUser = Partial<User>;