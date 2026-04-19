// Topic: Async Patterns

// 1. Promise Creation
const fetchData = (): Promise<string> => {
    return new Promise((resolve) => resolve("data"));
};

// 2. Promise Chaining
fetchData()
    .then(data => console.log(data))
    .catch(error => console.error(error));

// 3. Async/Await
async function getData(): Promise<void> {
    const result = await fetchData();
    console.log(result);
}

// 4. Parallel Execution
const [user, posts] = await Promise.all([
    Promise.resolve("user"),
    Promise.resolve("posts")
]);

// 5. Error Handling
try {
    await fetchData();
} catch (error) {
    console.error("Failed:", error);
}