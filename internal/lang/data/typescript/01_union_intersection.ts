// Topic: Union and Intersection Types

type SuccessResponse = { status: "success"; data: unknown };
type ErrorResponse = { status: "error"; message: string };
type ApiResponse = SuccessResponse | ErrorResponse;

type Admin = User & { permissions: string[]; role: "admin" };