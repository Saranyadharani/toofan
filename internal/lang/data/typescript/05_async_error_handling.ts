// Topic: Async Error Handling

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