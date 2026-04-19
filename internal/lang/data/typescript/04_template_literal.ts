// Topic: Template Literal Types

type EventName<T extends string> = `on${Capitalize<T>}`;
type HttpMethod = "get" | "post" | "put" | "delete";
type ApiEndpoint = `/api/${HttpMethod}/${string}`;