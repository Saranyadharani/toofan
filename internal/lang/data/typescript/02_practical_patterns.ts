// Concept 1: Factory Pattern with Generics
class Registry<T> {
    private items = new Map<string, T>();
    register(id: string, item: T): void { this.items.set(id, item); }
    get(id: string): T | undefined { return this.items.get(id); }
}

// Concept 2: Async/Await with Error Handling
async function fetchWithTimeout<T>(
    url: string,
    timeoutMs: number
): Promise<T> {
    const controller = new AbortController();	
    const timeoutId = setTimeout(() => controller.abort(), timeoutMs);
    try {
        const response = await fetch(url, { signal: controller.signal });
        return await response.json();
    } finally {
        clearTimeout(timeoutId);
    }
}

// Concept 3: Decorator Pattern
function loggedMethod<This, Args extends unknown[], Return>(
    target: (this: This, ...args: Args) => Return,
    context: ClassMethodDecoratorContext<This, (this: This, ...args: Args) => Return>
) {
    return function(this: This, ...args: Args): Return {
        console.log(`Calling ${String(context.name)} with:`, args);
        return target.call(this, args);
    };
}

// Concept 4: Utility Type Combinations
type Nullable<T> = { [P in keyof T]: T[P] | null };
type ApiEndpoint = `/api/${string}/${string}`;

// Concept 5: Promise.all with typed results
async function parallelRequests<T, U>(task1: Promise<T>, task2: Promise<U>): Promise<[T, U]> {
    return Promise.all([task1, task2]);
}

// Concept 6: Optional Chaining and Nullish Coalescing
function getDisplayName(user: { profile?: { name?: string | null } }): string {
    return user.profile?.name ?? "Anonymous";
}