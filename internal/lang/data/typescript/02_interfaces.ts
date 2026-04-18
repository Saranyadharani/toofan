// Topic: Working with Interfaces

interface User {
    id: number;
    name: string;
    email: string;
    isActive?: boolean;
}

interface Admin extends User {
    role: "admin" | "superadmin";
    permissions: string[];
}

const newUser: User = {
    id: 1,
    name: "Alice",
    email: "alice@example.com"
};