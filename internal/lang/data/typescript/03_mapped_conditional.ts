// Topic: Mapped and Conditional Types

type Nullable<T> = { [P in keyof T]: T[P] | null };

type IsArray<T> = T extends any[] ? true : false;